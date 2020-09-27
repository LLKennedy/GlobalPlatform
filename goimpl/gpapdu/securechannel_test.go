package gpapdu

import (
	"io"
	"testing"

	"github.com/llkennedy/globalplatform/goimpl/apdu"
	"github.com/stretchr/testify/assert"
)

func TestNewSecureChannelSession(t *testing.T) {
	type args struct {
		context          *Context
		s8mode           bool
		channelNumber    uint8
		keyVersionNumber uint8
	}
	tests := []struct {
		name      string
		args      args
		want      *SecureChannelSession
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSecureChannelSession(tt.args.context, tt.args.s8mode, tt.args.channelNumber, tt.args.keyVersionNumber)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestContext_InitializeUpdate(t *testing.T) {
	type fields struct {
		transport *apdu.TransportWrapper
		s8mode    bool
	}
	type args struct {
		channelNumber    uint8
		keyVersionNumber uint8
		randR            io.Reader
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      *SecureChannelSession
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Context{
				transport: tt.fields.transport,
				s8mode:    tt.fields.s8mode,
			}
			got, err := c.InitializeUpdate(tt.args.channelNumber, tt.args.keyVersionNumber, tt.args.randR)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSecureChannelSession_getRandR(t *testing.T) {
	type fields struct {
		randR io.Reader
	}
	tests := []struct {
		name   string
		fields fields
		want   io.Reader
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SecureChannelSession{
				randR: tt.fields.randR,
			}
			assert.Equal(t, tt.want, s.getRandR())
		})
	}
}
