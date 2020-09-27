package gpapdu

import (
	"testing"

	"github.com/llkennedy/globalplatform/goimpl/apdu"
	"github.com/stretchr/testify/assert"
)

func TestClassFromByte(t *testing.T) {
	type args struct {
		in byte
	}
	tests := []struct {
		name      string
		args      args
		wantClass Class
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotClass, err := ClassFromByte(tt.args.in)
			tt.assertion(t, err)
			assert.Equal(t, tt.wantClass, gotClass)
		})
	}
}

func TestClass_ToClassByte(t *testing.T) {
	type fields struct {
		InterindustryClass apdu.InterindustryClass
		IsGPCommand        bool
	}
	tests := []struct {
		name   string
		fields fields
		want   byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Class{
				InterindustryClass: tt.fields.InterindustryClass,
				IsGPCommand:        tt.fields.IsGPCommand,
			}
			assert.Equal(t, tt.want, c.ToClassByte())
		})
	}
}
