package gpapdu

import (
	"testing"

	"github.com/llkennedy/globalplatform/goimpl/apdu"
	"github.com/stretchr/testify/assert"
)

func TestConfirmationFromResponse(t *testing.T) {
	type args struct {
		in apdu.Response
	}
	tests := []struct {
		name      string
		args      args
		wantRes   ResponseConfirmation
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := ConfirmationFromResponse(tt.args.in)
			tt.assertion(t, err)
			assert.Equal(t, tt.wantRes, gotRes)
		})
	}
}
