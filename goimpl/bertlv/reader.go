package bertlv

import (
	"fmt"
	"io"
)

// Reader reads BER-TLV data one object at a time from a data stream
type Reader struct {
	data io.Reader
}

// NewReader creates a new Reader
func NewReader(data io.Reader) (r Reader, err error) {
	if data == nil {
		err = fmt.Errorf("invalid data, must not be nil")
	} else {
		r = Reader{data}
	}
	return
}

// Read reads the next object from data
func (r Reader) Read() (bytesRead int, object Object, err error) {
	tagBytes, tag, tagErr := TagFromReader(r.data)
	bytesRead = tagBytes
	if tagErr != nil {
		err = fmt.Errorf("error getting tag: %w", tagErr)
		return
	}
	object.Tag = tag
	lengthBytes, length, lengthErr := LengthFromReader(r.data)
	bytesRead += lengthBytes
	if lengthErr != nil {
		err = fmt.Errorf("error reading length: %w", lengthErr)
		return
	}
	object.Length = length
	object.Value = make([]byte, length)
	valueBytes, valueErr := r.data.Read(object.Value)
	bytesRead += valueBytes
	if valueErr != nil {
		err = fmt.Errorf("error reading value: %w", valueErr)
	}
	return
}
