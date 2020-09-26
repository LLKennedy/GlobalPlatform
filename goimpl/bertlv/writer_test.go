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
		args             args
		wantBytesWritten int
		assertion        assert.ErrorAssertionFunc
		wantOutput       []byte
	}{
		{
			name:      "example bitstring",
			assertion: assert.NoError,
			args: args{
				obj: Object{
					Tag: Tag{
						Class:               TagClassUniversal,
						ConstructedEncoding: false,
						Number:              3,
					},
					Length: 7000,
					Value:  []byte{0x04, 0x0A, 0x3B, 0x5F, 0x29, 0x1C, 0xD0},
				},
			},
			wantBytesWritten: 9,
			wantOutput:       []byte{0x03, 0x07, 0x04, 0x0A, 0x3B, 0x5F, 0x29, 0x1C, 0xD0},
		},
		{
			name:      "no value",
			assertion: assert.NoError,
			args: args{
				obj: Object{
					Tag: Tag{
						Class:               TagClassUniversal,
						ConstructedEncoding: false,
						Number:              3,
					},
					Length: 7000,
				},
			},
			wantBytesWritten: 2,
			wantOutput:       []byte{0x03, 0x00},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			w := Writer{
				output: buf,
			}
			gotBytesWritten, err := w.Write(tt.args.obj)
			tt.assertion(t, err)
			assert.Equal(t, tt.wantBytesWritten, gotBytesWritten)
			assert.Equal(t, tt.wantOutput, buf.Bytes())
		})
	}
}
