package gpapdu

import (
	"testing"

	"github.com/llkennedy/globalplatform/goimpl/apdu"
	"github.com/stretchr/testify/assert"
)

func TestSendOnTransport(t *testing.T) {
	type args struct {
		t   apdu.Transport
		cmd Command
	}
	tests := []struct {
		name      string
		args      args
		want      apdu.Response
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SendOnTransport(tt.args.t, tt.args.cmd)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
