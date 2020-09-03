package apdu

import (
	"encoding/binary"
	"fmt"
)

const (
	errorNonSpecific                           = 0x6400
	errorWrongLengthLc                         = 0x6700
	errorLogicalChannelNotSupportedOrInUse     = 0x6881
	errorSecurityStatusNotSatisfied            = 0x6982
	errorConditionsOfUseNotSatisfied           = 0x6985
	errorIncorrectP1P2                         = 0x6A86
	errorInvalidInstruction                    = 0x6D00
	errorInvalidClass                          = 0x6E00
	errStringPrefix                            = "APDU General Error: "
	errStringNonSpecific                       = "Non specific diagnosis"
	errStringWrongLengthLc                     = "Wrong length in Lc"
	errStringLogicalChannelNotSupportedOrInUse = "Logical channel not supported or is not active"
	errStringSecurityStatusNotSatisfied        = "Security status not satisfied"
	errStringConditionsOfUseNotSatisfied       = "Conditions of use not satisfied"
	errStringIncorrectP1P2                     = "Incorrect P1 P2"
	errStringInvalidInstruction                = "Invalid instruction"
	errStringInvalidClass                      = "Invalid class"
)

// Error is a general error which may be returned by any command
type Error struct {
	// Raw are the raw bytes of the error
	Raw uint16
}

// NewError allows construction of an Error from bytes instead of raw uint16 data
func NewError(err [2]byte) Error {
	return Error{
		Raw: binary.BigEndian.Uint16(err[:]),
	}
}

// Error complies with the error interface
func (e Error) Error() string {
	return e.String()
}

// String complies with the stringer interface
func (e Error) String() string {
	out := "APDU error handling error" // shouldn't be possible to return this
	switch e.Raw {
	case errorNonSpecific:
		out = formatErrorStringGeneral(errStringNonSpecific)
	case errorWrongLengthLc:
		out = formatErrorStringGeneral(errStringWrongLengthLc)
	case errorLogicalChannelNotSupportedOrInUse:
		out = formatErrorStringGeneral(errStringLogicalChannelNotSupportedOrInUse)
	case errorSecurityStatusNotSatisfied:
		out = formatErrorStringGeneral(errStringSecurityStatusNotSatisfied)
	case errorConditionsOfUseNotSatisfied:
		out = formatErrorStringGeneral(errStringConditionsOfUseNotSatisfied)
	case errorIncorrectP1P2:
		out = formatErrorStringGeneral(errStringIncorrectP1P2)
	case errorInvalidInstruction:
		out = formatErrorStringGeneral(errStringInvalidInstruction)
	case errorInvalidClass:
		out = formatErrorStringGeneral(errStringInvalidClass)
	default:
		out = fmt.Sprintf("Unknown APDU Error: %04X", e.Raw)
	}
	return out
}

func formatErrorStringGeneral(errString string) string {
	return fmt.Sprintf("%s%s", errStringPrefix, errString)
}
