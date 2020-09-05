package apdu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterindustryClass_ToClassByte(t *testing.T) {
	type fields struct {
		NotLastCommandOfChain bool
		SecureMessaging       CLASecureMessaging
		LogicalChannelNumber  uint8
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
			c := InterindustryClass{
				NotLastCommandOfChain: tt.fields.NotLastCommandOfChain,
				SecureMessaging:       tt.fields.SecureMessaging,
				LogicalChannelNumber:  tt.fields.LogicalChannelNumber,
			}
			assert.Equal(t, tt.want, c.ToClassByte())
		})
	}
}

func TestInterindustryClass_IsLastCommand(t *testing.T) {
	type fields struct {
		NotLastCommandOfChain bool
		SecureMessaging       CLASecureMessaging
		LogicalChannelNumber  uint8
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := InterindustryClass{
				NotLastCommandOfChain: tt.fields.NotLastCommandOfChain,
				SecureMessaging:       tt.fields.SecureMessaging,
				LogicalChannelNumber:  tt.fields.LogicalChannelNumber,
			}
			assert.Equal(t, tt.want, c.IsLastCommand())
		})
	}
}

func TestInterindustryClass_GetSMIndication(t *testing.T) {
	type fields struct {
		NotLastCommandOfChain bool
		SecureMessaging       CLASecureMessaging
		LogicalChannelNumber  uint8
	}
	tests := []struct {
		name   string
		fields fields
		want   CLASecureMessaging
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := InterindustryClass{
				NotLastCommandOfChain: tt.fields.NotLastCommandOfChain,
				SecureMessaging:       tt.fields.SecureMessaging,
				LogicalChannelNumber:  tt.fields.LogicalChannelNumber,
			}
			assert.Equal(t, tt.want, c.GetSMIndication())
		})
	}
}

func TestInterindustryClass_GetLogicalChannel(t *testing.T) {
	type fields struct {
		NotLastCommandOfChain bool
		SecureMessaging       CLASecureMessaging
		LogicalChannelNumber  uint8
	}
	tests := []struct {
		name   string
		fields fields
		want   uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := InterindustryClass{
				NotLastCommandOfChain: tt.fields.NotLastCommandOfChain,
				SecureMessaging:       tt.fields.SecureMessaging,
				LogicalChannelNumber:  tt.fields.LogicalChannelNumber,
			}
			assert.Equal(t, tt.want, c.GetLogicalChannel())
		})
	}
}
