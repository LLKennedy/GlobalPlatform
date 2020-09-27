package gpapdu

import (
	"encoding/asn1"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteCardContent_unimplementableDeleteCommandToBytes(t *testing.T) {
	type fields struct {
		ELFileOrAppID []byte
		CRTFDS        *ControlReferenceTemplateForDigitalSignature
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := DeleteCardContent{
				ELFileOrAppID: tt.fields.ELFileOrAppID,
				CRTFDS:        tt.fields.CRTFDS,
			}
			assert.Equal(t, tt.want, d.unimplementableDeleteCommandToBytes())
		})
	}
}

func TestDeleteKey_unimplementableDeleteCommandToBytes(t *testing.T) {
	type fields struct {
		IncludeKeyIdentifer     bool
		KeyIdentifier           byte
		IncludeKeyVersionNumber bool
		KeyVersionNumber        byte
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := DeleteKey{
				IncludeKeyIdentifer:     tt.fields.IncludeKeyIdentifer,
				KeyIdentifier:           tt.fields.KeyIdentifier,
				IncludeKeyVersionNumber: tt.fields.IncludeKeyVersionNumber,
				KeyVersionNumber:        tt.fields.KeyVersionNumber,
			}
			assert.Equal(t, tt.want, d.unimplementableDeleteCommandToBytes())
		})
	}
}

func TestASN1(t *testing.T) {
	c := ControlReferenceTemplateForDigitalSignature{
		ApplicationProviderID:     []byte{1},
		SecurityDomainID:          []byte{2},
		SecurityDomainImageNumber: []byte{3},
		TokenID:                   []byte{4},
	}
	data, err := asn1.MarshalWithParams(c, "tag:B6")
	assert.NoError(t, err)
	assert.Equal(t, []byte{}, data)
}
