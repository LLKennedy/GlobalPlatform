package apdu

import (
	"encoding/binary"
	"fmt"
)

// Command is an APDU command as specified in ISO/IEC 7816. Lc may not be specified, since it is calculated from the provided data.
type Command struct {
	Class                  Class  // CLA
	Instruction            byte   // INS
	P1P2                   uint16 // P1, P2 = Parameter fields
	Data                   []byte // Command data field
	ExpectResponseData     bool   // Overridden by ExpectedResponseLength
	ExpectedResponseLength uint16 // Le = expected number of bytes returned, 0 is interpreted as max (65536)
}

// ToBytes converts the command to bytes
func (c Command) ToBytes() []byte {
	return nil // FIXME
}

// Send is the same as calling t.Send(c), but it also nil-checks t for you
func (c Command) Send(t Transport) (Response, error) {
	if t == nil {
		return Response{}, fmt.Errorf("cannot send APDU command on nil transport interface")
	}
	return t.Send(c)
}

// Response is an APDU response as specified in ISO/IEC 7816.
type Response struct {
	Data   []byte // Response data field
	SW1SW2 uint16
}

// ResponseFromBytes generates a Response object from raw bytes
func ResponseFromBytes(data []byte) (Response, error) {
	total := len(data)
	if total < 2 {
		return Response{}, fmt.Errorf("APDU response must be at least 2 bytes long")
	}
	res := Response{
		SW1SW2: binary.BigEndian.Uint16(data[total-2:]),
	}
	if total > 2 {
		res.Data = make([]byte, total-2)
		copy(res.Data, data[:total-2])
	}
	return res, nil
}
