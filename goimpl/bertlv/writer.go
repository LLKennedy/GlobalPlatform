package bertlv

import (
	"fmt"
	"io"
)

// Writer writes BER-TLV data one object at a time to a data stream
type Writer struct {
	output io.Writer
}

// NewWriter creates a new Writer
func NewWriter(output io.Writer) (w Writer, err error) {
	if output == nil {
		err = fmt.Errorf("invalid output, must not be nil")
	} else {
		w = Writer{output}
	}
	return
}

func (w Writer) Write(obj Object) (bytesWritten int, err error) {
	tagBytes, tagErr := obj.Tag.ToBytes()
	if tagErr != nil {
		err = fmt.Errorf("failed to convert tag to bytes: %w", err)
		return
	}
	bytesWritten, err = w.output.Write(tagBytes)
	if err != nil {
		err = fmt.Errorf("writing tag bytes: %w", err)
		return
	}
	// Override length value, we don't respect the incoming data
	obj.Length = uint64(len(obj.Value))
	lengthWritten, lengthErr := w.output.Write(LengthToBytes(obj.Length))
	bytesWritten += lengthWritten
	if lengthErr != nil {
		err = fmt.Errorf("writing length bytes: %w", err)
		return
	}
	valueWritten, valueErr := w.output.Write(obj.Value)
	bytesWritten += valueWritten
	if valueErr != nil {
		err = fmt.Errorf("writing value bytes: %w", err)
	}
	return
}
