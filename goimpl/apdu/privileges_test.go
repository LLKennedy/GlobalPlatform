package apdu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrivileges_ToBytes(t *testing.T) {
	type fields struct {
		SecurityDomain            bool
		DAPVerification           bool
		DelegatedManagement       bool
		CardLock                  bool
		CardTerminate             bool
		CardReset                 bool
		CVMManagement             bool
		MandatedDAPVerification   bool
		TrustedPath               bool
		AuthorizedManagement      bool
		TokenManagement           bool
		GlobalDelete              bool
		GlobalLock                bool
		GlobalRegistry            bool
		FinalApplication          bool
		GlobalService             bool
		ReceiptGeneration         bool
		CipheredLoadFileDataBlock bool
		ContactlessActivation     bool
		ContactlessSelfActivation bool
		reserved1                 bool
		reserved2                 bool
		reserved3                 bool
		reserved4                 bool
	}
	tests := []struct {
		name   string
		fields fields
		want   [3]byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Privileges{
				SecurityDomain:            tt.fields.SecurityDomain,
				DAPVerification:           tt.fields.DAPVerification,
				DelegatedManagement:       tt.fields.DelegatedManagement,
				CardLock:                  tt.fields.CardLock,
				CardTerminate:             tt.fields.CardTerminate,
				CardReset:                 tt.fields.CardReset,
				CVMManagement:             tt.fields.CVMManagement,
				MandatedDAPVerification:   tt.fields.MandatedDAPVerification,
				TrustedPath:               tt.fields.TrustedPath,
				AuthorizedManagement:      tt.fields.AuthorizedManagement,
				TokenManagement:           tt.fields.TokenManagement,
				GlobalDelete:              tt.fields.GlobalDelete,
				GlobalLock:                tt.fields.GlobalLock,
				GlobalRegistry:            tt.fields.GlobalRegistry,
				FinalApplication:          tt.fields.FinalApplication,
				GlobalService:             tt.fields.GlobalService,
				ReceiptGeneration:         tt.fields.ReceiptGeneration,
				CipheredLoadFileDataBlock: tt.fields.CipheredLoadFileDataBlock,
				ContactlessActivation:     tt.fields.ContactlessActivation,
				ContactlessSelfActivation: tt.fields.ContactlessSelfActivation,
				reserved1:                 tt.fields.reserved1,
				reserved2:                 tt.fields.reserved2,
				reserved3:                 tt.fields.reserved3,
				reserved4:                 tt.fields.reserved4,
			}
			assert.Equal(t, tt.want, p.ToBytes())
		})
	}
}

func TestPrivilegesFromBytes(t *testing.T) {
	type args struct {
		in [3]byte
	}
	tests := []struct {
		name string
		args args
		want Privileges
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, PrivilegesFromBytes(tt.args.in))
		})
	}
}
