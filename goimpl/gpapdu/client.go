package gpapdu

import (
	"fmt"

	"github.com/llkennedy/globalplatform/goimpl/apdu"
)

// Client is a GP client wrapping APDU commands
type Client struct {
	transport apdu.Transport
}

// NewClient creates a new Client on the provided Transport
func NewClient(transport apdu.Transport) *Client {
	return &Client{transport}
}

// Delete deletes keys or card contents
func (c *Client) Delete(deleteRelatedObjects bool, cmd DeleteCommand) (confirmation []byte, err error) {
	if cmd == nil {
		return nil, fmt.Errorf("must supply a non-nil command")
	}
	fullData := cmd.unimplementableDeleteCommandToBytes()
	command := Command{
		Class: Class{
			IsGPCommand: true,
		},
		Instruction:        apdu.InstructionDeleteFile,
		ExpectResponseData: true,
	}
	if deleteRelatedObjects {
		command.P2 = b8
	}
	if len(fullData) > 255 {
		command.P1 = b8
		command.Data = fullData[:255]
	} else {
		command.Data = fullData
	}
	// res, sendErr := SendOnTransport(c.transport, command)
	// if sendErr != nil {
	// 	return nil, fmt.Errorf("sending command: %w", sendErr)
	// }
	// FIXME: multiple sends, multiple responds to map
	// r := bertlv.NewBytesReader(res.Data)
	// _, obj, readErr := r.ReadWithoutTag()
	// if readErr != nil {
	// 	return nil, fmt.Errorf("reading response: %w", readErr)
	// }
	return //obj.Value, nil
}
