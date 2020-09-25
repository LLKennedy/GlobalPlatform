package bertlv

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTag_ToBytes(t *testing.T) {
	tests := []struct {
		name      string
		tr        Tag
		want      []byte
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "low number",
			tr: Tag{
				Class:               TagClassContextSpecific,
				ConstructedEncoding: true,
				Number:              27,
			},
			assertion: assert.NoError,
			want:      []byte{0b10111011},
		},
		{
			name: "class error",
			tr: Tag{
				Class:               5,
				ConstructedEncoding: true,
				Number:              27,
			},
			assertion: assert.Error,
		},
		{
			name: "30 edge",
			tr: Tag{
				Class:               TagClassApplication,
				ConstructedEncoding: true,
				Number:              30,
			},
			assertion: assert.NoError,
			want:      []byte{0b01111110},
		},
		{
			name: "31 edge",
			tr: Tag{
				Class:               TagClassApplication,
				ConstructedEncoding: true,
				Number:              31,
			},
			assertion: assert.NoError,
			want:      []byte{0b01111111, 0b00011111},
		},
		{
			name: "127 edge",
			tr: Tag{
				Class:               TagClassPrivate,
				ConstructedEncoding: false,
				Number:              127,
			},
			assertion: assert.NoError,
			want:      []byte{0b11011111, 0b01111111},
		},
		{
			name: "128 edge",
			tr: Tag{
				Class:               TagClassUniversal,
				ConstructedEncoding: false,
				Number:              128,
			},
			assertion: assert.NoError,
			want:      []byte{0b00011111, 0b10000001, 0b00000000},
		},
		{
			name: "1,234,567,890 example of large scale",
			tr: Tag{
				Class:               TagClassUniversal,
				ConstructedEncoding: false,
				Number:              1234567890,
			},
			assertion: assert.NoError,
			want:      []byte{0b00011111, 0b10000100, 0b11001100, 0b11011000, 0b10000101, 0b01010010},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.tr.ToBytes()
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTag_Write(t *testing.T) {
	tests := []struct {
		name      string
		tr        Tag
		wantW     string
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			tt.assertion(t, tt.tr.Write(w))
			assert.Equal(t, tt.wantW, w.String())
		})
	}
}

func TestTagFromBytes(t *testing.T) {
	type args struct {
		data io.Reader
	}
	tests := []struct {
		name         string
		args         args
		wantTag      Tag
		wantOverread bool
		assertion    assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTag, gotOverread, err := TagFromBytes(tt.args.data)
			tt.assertion(t, err)
			assert.Equal(t, tt.wantTag, gotTag)
			assert.Equal(t, tt.wantOverread, gotOverread)
		})
	}
}

func Test_reverseBitReader_ReadBit(t *testing.T) {
	tests := []struct {
		name string
		b    *reverseBitReader
		want byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.b.ReadBit())
		})
	}
}
