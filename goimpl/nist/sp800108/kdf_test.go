package sp800108

import (
	"crypto/sha512"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounterKBKDF_Derive(t *testing.T) {
	type args struct {
		prf                   PRF
		counterLengthR        Counter
		inputKey              []byte
		label                 []byte
		context               []byte
		outputSizeBits        []byte
		includeEmptySeparator bool
	}
	tests := []struct {
		name      string
		c         *CounterKBKDF
		args      args
		want      []byte
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "test vector 1",
			args: args{
				prf:                   NewPRFHMAC(sha512.New384),
				label:                 blindHexDecode("bc2c728f9dc6db426dd4e85fdb493826a31fec0607644209f9bf2264b6401b5db3"),
				context:               blindHexDecode("4c1a76aa08d93f08d3d9e2ba434b682e480004fb0d9271a8e8cd"),
				includeEmptySeparator: true,
				counterLengthR:        CounterLength16,
				inputKey:              blindHexDecode("26ef897e4b617b597f766ec8d8ccf44c543e790a7d218f029dcb4a3695ae2caccce9d3e935f6741581f2f53e49cd46f8"),
				outputSizeBits:        []byte{128},
			},
			// Edited from the provided vectors because things are weird
			// https://crypto.stackexchange.com/questions/83675/inconsistency-between-nist-sp800-108-and-kbkdfvs
			want:      []byte{0x93, 0x14, 0xad, 0x6, 0xb2, 0x49, 0x69, 0x13, 0xb4, 0x25, 0x84, 0xf6, 0x7b, 0x8f, 0x7e, 0x34},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CounterKBKDF{}
			got, err := c.Derive(tt.args.prf, tt.args.counterLengthR, tt.args.inputKey, tt.args.label, tt.args.context, tt.args.outputSizeBits, tt.args.includeEmptySeparator)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func blindHexDecode(hexData string) []byte {
	data, err := hex.DecodeString(hexData)
	if err != nil {
		panic(err)
	}
	return data
}
