package gpapdu

import (
	"fmt"

	"github.com/llkennedy/globalplatform/goimpl/apdu"
)

// Command is the GP subset of apdu.Command, with fields altered or removed to help avoid breaking restrictions set by the GP spec
type Command struct {
	Class              Class  // CLA
	Instruction        byte   // INS
	P1, P2             byte   // P1, P2 = Parameter fields
	Data               []byte // Command data field
	ExpectResponseData bool   // Toggles whether any data is expected, expected length cannot be set
}

// ToAPDU returns the command as an apdu.Command, ready for byte conversion and transmission
func (c Command) ToAPDU() (apdu.Command, error) {
	if len(c.Data) > 255 {
		return apdu.Command{}, fmt.Errorf("too much data, one GP APDU command may only contain 255 bytes")
	}
	expectedLength := uint16(0)
	if c.ExpectResponseData {
		expectedLength = 256
	}
	cmd := apdu.Command{
		Class:                  c.Class,
		Instruction:            c.Instruction,
		P1:                     c.P1,
		P2:                     c.P2,
		Data:                   c.Data,
		ExpectedResponseLength: expectedLength,
	}
	return cmd, nil
}
