package apdu

import (
	"encoding/binary"
	"fmt"
)

// Command is an APDU command as specified in ISO/IEC 7816. Lc may not be specified, since it is calculated from the provided data.
type Command struct {
	Class                  Class  // CLA
	Instruction            byte   // INS
	P1, P2                 byte   // P1, P2 = Parameter fields
	Data                   []byte // Command data field
	ExpectResponseData     bool   // Overridden by ExpectedResponseLength > 0, only serves to differentiate 0 = 0 and 0 = 65536
	ExpectedResponseLength uint16 // Le = expected number of bytes returned, 0 is interpreted as max (65536) only if ExpectResponseData is true
}

// ToBytes converts the command to bytes
func (c Command) ToBytes() []byte {
	var classByte byte
	if c.Class == nil {
		classByte = 0xFF // Deliberately invalid, since the input is invalid
	} else {
		classByte = c.Class.ToClassByte()
	}
	// Define command header
	header := []byte{classByte, c.Instruction, c.P1, c.P2}
	if len(c.Data) == 0 && !c.ExpectResponseData && c.ExpectedResponseLength == 0 {
		// No data in or out, only header + SW1SW2
		return header
	}
	var le []byte
	var lc []byte
	useExtendedLengths := c.ExpectedResponseLength > 256 || (c.ExpectResponseData && c.ExpectedResponseLength == 0) || len(c.Data) > 256
	if c.ExpectResponseData || c.ExpectedResponseLength > 0 {
		if useExtendedLengths {
			le = make([]byte, 2)
			binary.BigEndian.PutUint16(le, c.ExpectedResponseLength)
			if len(c.Data) == 0 {
				le = append([]byte{0}, le...)
			}
		} else if c.ExpectedResponseLength == 256 {
			le = []byte{0}
		} else {
			// We've confirmed 0 < le <= 255 with the checks above, this will never overflow
			le = []byte{byte(c.ExpectedResponseLength)}
		}
	}
	dataLen := len(c.Data)
	if dataLen != 0 {
		if dataLen > 65536 {
			panic("globalplatform/apdu: cannot send more than 65536 bytes of command data in one command")
		} else if useExtendedLengths {
			lc = make([]byte, 2)
			if dataLen < 65536 {
				binary.BigEndian.PutUint16(lc, uint16(dataLen))
				lc = append([]byte{0}, lc...)
			}
		} else {
			lc = make([]byte, 1)
			if dataLen < 256 {
				lc[0] = byte(dataLen)
			}
		}
	}
	out := make([]byte, len(header)+len(lc)+dataLen+len(le))
	buildOut := out[0:0]
	buildOut = append(buildOut, header...)
	buildOut = append(buildOut, lc...)
	buildOut = append(buildOut, c.Data...)
	buildOut = append(buildOut, le...)
	return out
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
