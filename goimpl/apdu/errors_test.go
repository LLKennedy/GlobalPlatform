package apdu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError_Error(t *testing.T) {
	type fields struct {
		Raw uint16
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "empty defaults",
			want: "Unknown APDU Error: 0000",
		},
		{
			name: "invalid raw error",
			fields: fields{
				Raw: 0xFFFF,
			},
			want: "Unknown APDU Error: FFFF",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Error{
				Raw: tt.fields.Raw,
			}
			assert.Equal(t, tt.want, e.Error())
		})
	}
}

func TestError_String(t *testing.T) {
	type fields struct {
		Raw uint16
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Error{
				Raw: tt.fields.Raw,
			}
			assert.Equal(t, tt.want, e.String())
		})
	}
}
