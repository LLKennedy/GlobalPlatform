package gpapdu

import (
	"testing"

	"github.com/llkennedy/globalplatform/goimpl/apdu"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	type args struct {
		transport apdu.Transport
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewClient(tt.args.transport))
		})
	}
}

func TestClient_Delete(t *testing.T) {
	type fields struct {
		transport apdu.Transport
	}
	type args struct {
		deleteRelatedObjects bool
		cmd                  DeleteCommand
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantConfirmation []byte
		assertion        assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				transport: tt.fields.transport,
			}
			gotConfirmation, err := c.Delete(tt.args.deleteRelatedObjects, tt.args.cmd)
			tt.assertion(t, err)
			assert.Equal(t, tt.wantConfirmation, gotConfirmation)
		})
	}
}
