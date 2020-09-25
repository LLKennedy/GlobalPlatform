package bertlv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	A int `bertlv:"abc"`
	b bool
}

type myint uint64

func TestMarshal(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name      string
		args      args
		want      []byte
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "string",
			args: args{
				v: myint(12),
			},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Marshal(tt.args.v)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
