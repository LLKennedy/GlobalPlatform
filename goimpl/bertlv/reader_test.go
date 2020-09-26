package bertlv

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewReader(t *testing.T) {
	type args struct {
		data io.Reader
	}
	tests := []struct {
		name      string
		args      args
		wantR     Reader
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotR, err := NewReader(tt.args.data)
			tt.assertion(t, err)
			assert.Equal(t, tt.wantR, gotR)
		})
	}
}

func TestReader_Read(t *testing.T) {
	type fields struct {
		data io.Reader
	}
	tests := []struct {
		name          string
		fields        fields
		wantBytesRead int
		wantObject    Object
		assertion     assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Reader{
				data: tt.fields.data,
			}
			gotBytesRead, gotObject, err := r.Read()
			tt.assertion(t, err)
			assert.Equal(t, tt.wantBytesRead, gotBytesRead)
			assert.Equal(t, tt.wantObject, gotObject)
		})
	}
}
