package gpapdu

import (
	"fmt"

	"github.com/llkennedy/globalplatform/goimpl/apdu"
)

// Class extends the Interindustry class using b8 as a flag for GP-specific commands, no other changes
type Class struct {
	apdu.InterindustryClass
	IsGPCommand bool // Sets b8 depending on whether or not this is a proprietary GP command
}

// ClassFromByte converts a byte to a Class
func ClassFromByte(in byte) (class Class, err error) {
	b8Set := (in & b8) == b8
	baseClass, baseErr := apdu.InterindustryClassFromByte(in & 0x7F)
	if baseErr != nil {
		err = fmt.Errorf("could not convert byte to GP command class: %w", baseErr)
		return
	}
	class.IsGPCommand = b8Set
	class.InterindustryClass = baseClass
	return
}

// ToClassByte returns the formatted class byte
func (c Class) ToClassByte() byte {
	out := c.InterindustryClass.ToClassByte()
	if c.IsGPCommand {
		out = out | b8
	}
	return out
}
