package bertlv

import (
	"encoding/binary"
	"fmt"
	"io"
)

// LengthFromReader reads a length value from the reader
func LengthFromReader(data io.Reader) (bytesRead int, length uint64, err error) {
	dst := make([]byte, 1)
	bytesRead, err = data.Read(dst)
	if err != nil {
		return
	}
	firstByte := dst[0]
	if firstByte == 0x80 {
		// Indefinite isn't supported by ISO7816 and GP complies with that, so for the purposes of this library we don't need it
		err = fmt.Errorf("indefinite TLV lengths are not supported")
		return
	}
	short := (firstByte & b8) == 0
	if short {
		length = uint64(firstByte & 0x7F)
		return
	}
	// If the length is >127, the first 7 bits encodes the number of byte the rest of the length will be encoded with.
	// We only support <=64 bits (<=8 bytes)
	lengthLength := firstByte & 0x7F
	if lengthLength > 8 {
		err = fmt.Errorf("TLV lengths requiring more than 8 bytes of length data are not supported, %d bytes of length data were indicated", lengthLength)
		return
	}
	lengthBytes := make([]byte, lengthLength)
	read, readErr := data.Read(lengthBytes)
	bytesRead += read
	if readErr != nil {
		err = fmt.Errorf("could not read length data: %w", err)
	} else if read < int(lengthLength) {
		err = fmt.Errorf("%d length bytes were required, only %d could be read", lengthLength, read)
	} else {
		length = binary.BigEndian.Uint64(lengthBytes)
	}
	return
}

// LengthToBytes converts a uint64 length to to properly encoded length bytes
func LengthToBytes(length uint64) (data []byte) {
	if length < 128 {
		// It's that simple
		return []byte{byte(length)}
	}
	raw := make([]byte, 8)
	binary.BigEndian.PutUint64(raw, length)
	// Trim leading zero bytes to be kind to the receiver
	for i, next := range raw {
		if next != 0 {
			raw = raw[i:]
			break
		}
	}
	return append([]byte{byte(len(raw)) | b8}, raw...)
}
