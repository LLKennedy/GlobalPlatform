package bertlv

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWriter(t *testing.T) {
	tests := []struct {
		name       string
		wantW      Writer
		wantOutput string
		assertion  assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := &bytes.Buffer{}
			gotW, err := NewWriter(output)
			tt.assertion(t, err)
			assert.Equal(t, tt.wantW, gotW)
			assert.Equal(t, tt.wantOutput, output.String())
		})
	}
}

func TestWriter_Write(t *testing.T) {
	type fields struct {
		output io.Writer
	}
	type args struct {
		obj Object
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantBytesWritten int
		assertion        assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := Writer{
				output: tt.fields.output,
			}
			gotBytesWritten, err := w.Write(tt.args.obj)
			tt.assertion(t, err)
			assert.Equal(t, tt.wantBytesWritten, gotBytesWritten)
		})
	}
}
