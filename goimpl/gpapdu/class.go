package gpapdu

import "github.com/llkennedy/globalplatform/goimpl/apdu"

// Class extends the Interindustry class using b8 as a flag for GP-specific commands, no other changes
type Class struct {
	apdu.InterindustryClass
	IsGPCommand bool // Sets b8 depending on whether or not this is a proprietary GP command

}

// ToClassByte returns the formatted class byte
func (c Class) ToClassByte() byte {
	out := c.InterindustryClass.ToClassByte()
	if c.IsGPCommand {
		out = out | b8
	}
	return out
}
