package sp800108

import (
	"hash"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPRFHMAC(t *testing.T) {
	type args struct {
		hashGen func() hash.Hash
	}
	tests := []struct {
		name string
		args args
		want *PRFHMAC
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewPRFHMAC(tt.args.hashGen))
		})
	}
}

func TestPRFHMAC_OutputSizeBytes(t *testing.T) {
	type fields struct {
		hashGen func() hash.Hash
	}
	tests := []struct {
		name   string
		fields fields
		want   uint
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PRFHMAC{
				hashGen: tt.fields.hashGen,
			}
			assert.Equal(t, tt.want, p.OutputSizeBytes())
		})
	}
}

func TestPRFHMAC_Compute(t *testing.T) {
	type fields struct {
		hashGen func() hash.Hash
	}
	type args struct {
		key  []byte
		data []byte
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      []byte
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PRFHMAC{
				hashGen: tt.fields.hashGen,
			}
			got, err := p.Compute(tt.args.key, tt.args.data)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPRFHMAC_unimplementableNISTSP800108PRF(t *testing.T) {
	type fields struct {
		hashGen func() hash.Hash
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PRFHMAC{
				hashGen: tt.fields.hashGen,
			}
			p.unimplementableNISTSP800108PRF()
		})
	}
}
