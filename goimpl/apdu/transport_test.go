package apdu

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransportWrapper_Send(t *testing.T) {
	type fields struct {
		Transport Transport
	}
	type args struct {
		cmd Command
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
			T := &TransportWrapper{
				Mutex:     sync.Mutex{},
				Transport: tt.fields.Transport,
			}
			got, err := T.Send(tt.args.cmd)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
