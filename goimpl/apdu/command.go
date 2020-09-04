package apdu

// Command is an APDU command as specified in ISO/IEC 7816. Lc may not be specified, since it is calculated from the provided data.
type Command struct {
	Class                  Class  // CLA
	Instruction            byte   // INS
	P1P2                   uint16 // P1, P2 = Parameter fields
	Data                   []byte // Command data field
	ExpectedResponseLength uint16 // Le = expected number of bytes returned, 0 is interpreted as max (65536)
}

// ToBytes converts the command to bytes
func (c Command) ToBytes() []byte {
	return nil // FIXME
}
