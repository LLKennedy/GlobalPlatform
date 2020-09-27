package gpapdu

import (
	"testing"

	"github.com/llkennedy/globalplatform/goimpl/apdu"
	"github.com/stretchr/testify/assert"
)

func TestCommand_ToAPDU(t *testing.T) {
	type fields struct {
		Class              Class
		Instruction        apdu.Instruction
		P1                 byte
		P2                 byte
		Data               []byte
		ExpectResponseData bool
	}
	tests := []struct {
		name      string
		fields    fields
		want      apdu.Command
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Command{
				Class:              tt.fields.Class,
				Instruction:        tt.fields.Instruction,
				P1:                 tt.fields.P1,
				P2:                 tt.fields.P2,
				Data:               tt.fields.Data,
				ExpectResponseData: tt.fields.ExpectResponseData,
			}
			got, err := c.ToAPDU()
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
