package gpapdu

import (
	"crypto/rand"
	"fmt"
	"io"

	"github.com/llkennedy/globalplatform/goimpl/apdu"
)

const (
	// InstructionInitializeUpdate is the InitializeUpdate instruction
	InstructionInitializeUpdate apdu.Instruction = 0x50
)

// SecureChannelSession is a Secure Channel Session
type SecureChannelSession struct {
	randR io.Reader
}

// NewSecureChannelSession creates a new Secure Channel Session
func NewSecureChannelSession(context *Context, s8mode bool, channelNumber uint8, keyVersionNumber uint8) (*SecureChannelSession, error) {
	return context.InitializeUpdate(channelNumber, keyVersionNumber, nil)
}

// InitializeUpdate initiates a new Secure Channel Session
func (c *Context) InitializeUpdate(channelNumber uint8, keyVersionNumber uint8, randR io.Reader) (*SecureChannelSession, error) {
	s := &SecureChannelSession{
		randR: randR,
	}
	if c == nil || c.transport == nil {
		return nil, fmt.Errorf("cannot run InitializeUpdate command with nil transport")
	}
	class := Class{
		IsGPCommand: true,
		InterindustryClass: apdu.InterindustryClass{
			NotLastCommandOfChain: false,
			LogicalChannelNumber:  channelNumber,
			SecureMessaging:       apdu.CLASMNone,
		},
	}
	var randomHostChallenge []byte
	if c.s8mode {
		randomHostChallenge = make([]byte, 8)
	} else {
		randomHostChallenge = make([]byte, 16)
	}
	n, err := s.getRandR().Read(randomHostChallenge)
	if err != nil {
		return nil, fmt.Errorf("failed to generate host challenge: %v", err)
	}
	if (c.s8mode && n != 8) || (!c.s8mode && n != 16) {
		return nil, fmt.Errorf("failed to generate sufficient data for host challenge: got %d bytes, s8mode = %v", n, c.s8mode)
	}
	cmd := Command{
		Class:              class,
		Instruction:        InstructionInitializeUpdate,
		P1:                 keyVersionNumber,
		P2:                 0x00, // Required to always be 0x00
		Data:               randomHostChallenge,
		ExpectResponseData: true,
	}
	res, err := SendOnTransport(c.transport, cmd)
	if err != nil {
		return nil, err
	}
	if err = res.GetStatus().Error(); err != nil {
		return nil, err
	}
	mandatoryDataLength := 13
	if c.s8mode {
		mandatoryDataLength += 16
	} else {
		mandatoryDataLength += 32
	}
	dataLength := len(res.Data)
	if dataLength < mandatoryDataLength || dataLength > mandatoryDataLength+3 {
		return nil, fmt.Errorf("incorrect data returned from card: expected at least %d bytes and at most %d bytes, got %d", mandatoryDataLength, mandatoryDataLength+3, dataLength)
	}

	return nil, fmt.Errorf("not implemented")
}

// Inaccessible, only exists for testing purposes
func (s *SecureChannelSession) getRandR() io.Reader {
	if s == nil || s.randR == nil {
		return rand.Reader
	}
	if rand.Reader == nil {
		panic("crypto/rand.Reader set to nil")
	}
	return s.randR
}
