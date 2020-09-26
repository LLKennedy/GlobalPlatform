package bertlv

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthFromReader(t *testing.T) {
	type args struct {
		data io.Reader
	}
	tests := []struct {
		name          string
		args          args
		wantBytesRead int
		wantLength    uint64
		assertion     assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBytesRead, gotLength, err := LengthFromReader(tt.args.data)
			tt.assertion(t, err)
			assert.Equal(t, tt.wantBytesRead, gotBytesRead)
			assert.Equal(t, tt.wantLength, gotLength)
		})
	}
}

func TestLengthToBytes(t *testing.T) {
	type args struct {
		length uint64
	}
	tests := []struct {
		name     string
		args     args
		wantData []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantData, LengthToBytes(tt.args.length))
		})
	}
}
