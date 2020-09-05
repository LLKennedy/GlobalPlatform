package apdu

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommand_ToBytes(t *testing.T) {
	type fields struct {
		Class                  Class
		Instruction            byte
		P1                     byte
		P2                     byte
		Data                   []byte
		ExpectResponseData     bool
		ExpectedResponseLength uint16
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{
			name: "case 1",
			fields: fields{
				Class:       byteClass(0x12),
				Instruction: 0x34,
				P1:          0x56,
				P2:          0x78,
			},
			want: []byte{0x12, 0x34, 0x56, 0x78},
		},
		{
			name: "case 2s.1",
			fields: fields{
				Class:                  byteClass(0x12),
				Instruction:            0x34,
				P1:                     0x56,
				P2:                     0x78,
				ExpectedResponseLength: 128,
			},
			want: []byte{0x12, 0x34, 0x56, 0x78, 0x80},
		},
		{
			name: "case 2s.2",
			fields: fields{
				Class:                  byteClass(0x12),
				Instruction:            0x34,
				P1:                     0x56,
				P2:                     0x78,
				ExpectedResponseLength: 256,
			},
			want: []byte{0x12, 0x34, 0x56, 0x78, 0x00},
		},
		{
			name: "case 2e.1",
			fields: fields{
				Class:                  byteClass(0x12),
				Instruction:            0x34,
				P1:                     0x56,
				P2:                     0x78,
				ExpectedResponseLength: 300,
			},
			want: []byte{0x12, 0x34, 0x56, 0x78, 0x01, 0x2C},
		},
		{
			name: "case 2e.2",
			fields: fields{
				Class:                  byteClass(0x12),
				Instruction:            0x34,
				P1:                     0x56,
				P2:                     0x78,
				ExpectResponseData:     true,
				ExpectedResponseLength: 0,
			},
			want: []byte{0x12, 0x34, 0x56, 0x78, 0x00, 0x00},
		},
		{
			name: "case 3s.1",
			fields: fields{
				Class:       byteClass(0x12),
				Instruction: 0x34,
				P1:          0x56,
				P2:          0x78,
				Data:        []byte{0x10, 0x20, 0x30},
			},
			want: []byte{0x12, 0x34, 0x56, 0x78, 0x03, 0x10, 0x20, 0x30},
		},
		{
			name: "case 3s.2",
			fields: fields{
				Class:       byteClass(0x12),
				Instruction: 0x34,
				P1:          0x56,
				P2:          0x78,
				Data: func() []byte {
					data := make([]byte, 256)
					return data
				}(),
			},
			want: func() []byte {
				data := []byte{0x12, 0x34, 0x56, 0x78, 0x00}
				data = append(data, make([]byte, 256)...)
				return data
			}(),
		},
		{
			name: "case 3e.1",
			fields: fields{
				Class:       byteClass(0x12),
				Instruction: 0x34,
				P1:          0x56,
				P2:          0x78,
				Data: func() []byte {
					data := make([]byte, 300)
					return data
				}(),
			},
			want: func() []byte {
				data := []byte{0x12, 0x34, 0x56, 0x78, 0x01, 0x2C}
				data = append(data, make([]byte, 300)...)
				return data
			}(),
		},
		{
			name: "case 3e.2",
			fields: fields{
				Class:       byteClass(0x12),
				Instruction: 0x34,
				P1:          0x56,
				P2:          0x78,
				Data: func() []byte {
					data := make([]byte, 65536)
					return data
				}(),
			},
			want: func() []byte {
				data := []byte{0x12, 0x34, 0x56, 0x78, 0x00, 0x00}
				data = append(data, make([]byte, 65536)...)
				return data
			}(),
		},
		{
			name: "case 4s.1",
			fields: fields{
				Class:                  byteClass(0x12),
				Instruction:            0x34,
				P1:                     0x56,
				P2:                     0x78,
				Data:                   []byte{0x10, 0x20, 0x30},
				ExpectedResponseLength: 128,
			},
			want: []byte{0x12, 0x34, 0x56, 0x78, 0x03, 0x10, 0x20, 0x30, 0x80},
		},
		{
			name: "case 4s.2",
			fields: fields{
				Class:                  byteClass(0x12),
				Instruction:            0x34,
				P1:                     0x56,
				P2:                     0x78,
				Data:                   []byte{0x10, 0x20, 0x30},
				ExpectedResponseLength: 256,
			},
			want: []byte{0x12, 0x34, 0x56, 0x78, 0x03, 0x10, 0x20, 0x30, 0x00},
		},
		{
			name: "case 4s.3",
			fields: fields{
				Class:                  byteClass(0x12),
				Instruction:            0x34,
				P1:                     0x56,
				P2:                     0x78,
				Data:                   make([]byte, 256),
				ExpectedResponseLength: 0x57,
			},
			want: func() []byte {
				data := []byte{0x12, 0x34, 0x56, 0x78, 0x00}
				data = append(data, make([]byte, 256)...)
				data = append(data, 0x57)
				return data
			}(),
		},
		{
			name: "case 4s.4",
			fields: fields{
				Class:                  byteClass(0x12),
				Instruction:            0x34,
				P1:                     0x56,
				P2:                     0x78,
				Data:                   make([]byte, 256),
				ExpectedResponseLength: 256,
			},
			want: func() []byte {
				data := []byte{0x12, 0x34, 0x56, 0x78, 0x00}
				data = append(data, make([]byte, 256)...)
				data = append(data, 0x00)
				return data
			}(),
		},
		{
			name: "case 4e.1",
			fields: fields{
				Class:                  byteClass(0x12),
				Instruction:            0x34,
				P1:                     0x56,
				P2:                     0x78,
				Data:                   make([]byte, 65536),
				ExpectedResponseLength: 128,
			},
			want: func() []byte {
				data := []byte{0x12, 0x34, 0x56, 0x78, 0x00, 0x00}
				data = append(data, make([]byte, 65536)...)
				data = append(data, 0x00, 0x80)
				return data
			}(),
		},
		{
			name: "case 4s.2",
			fields: fields{
				Class:                  byteClass(0x12),
				Instruction:            0x34,
				P1:                     0x56,
				P2:                     0x78,
				Data:                   make([]byte, 65536),
				ExpectedResponseLength: 0,
				ExpectResponseData:     true,
			},
			want: func() []byte {
				data := []byte{0x12, 0x34, 0x56, 0x78, 0x00, 0x00}
				data = append(data, make([]byte, 65536)...)
				data = append(data, 0x00, 0x00)
				return data
			}(),
		},
		{
			name: "case 4s.3",
			fields: fields{
				Class:                  byteClass(0x12),
				Instruction:            0x34,
				P1:                     0x56,
				P2:                     0x78,
				Data:                   make([]byte, 12),
				ExpectedResponseLength: 10000,
			},
			want: func() []byte {
				data := []byte{0x12, 0x34, 0x56, 0x78, 0x00, 0x0C}
				data = append(data, make([]byte, 12)...)
				data = append(data, 0x27, 0x10)
				return data
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Command{
				Class:                  tt.fields.Class,
				Instruction:            tt.fields.Instruction,
				P1:                     tt.fields.P1,
				P2:                     tt.fields.P2,
				Data:                   tt.fields.Data,
				ExpectResponseData:     tt.fields.ExpectResponseData,
				ExpectedResponseLength: tt.fields.ExpectedResponseLength,
			}
			assert.Equal(t, tt.want, c.ToBytes())
		})
	}
}

func TestCommand_Send(t *testing.T) {
	type fields struct {
		Class                  Class
		Instruction            byte
		P1                     byte
		P2                     byte
		Data                   []byte
		ExpectResponseData     bool
		ExpectedResponseLength uint16
	}
	type args struct {
		t Transport
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      Response
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Command{
				Class:                  tt.fields.Class,
				Instruction:            tt.fields.Instruction,
				P1:                     tt.fields.P1,
				P2:                     tt.fields.P2,
				Data:                   tt.fields.Data,
				ExpectResponseData:     tt.fields.ExpectResponseData,
				ExpectedResponseLength: tt.fields.ExpectedResponseLength,
			}
			got, err := c.Send(tt.args.t)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestResponseFromBytes(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name      string
		args      args
		want      Response
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ResponseFromBytes(tt.args.data)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

type byteClass byte

func (b byteClass) ToClassByte() byte {
	return byte(b)
}

func (b byteClass) ToInterindustry() (InterindustryClass, error) {
	return InterindustryClass{}, fmt.Errorf("invalid")
}
