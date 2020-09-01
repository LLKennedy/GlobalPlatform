package sp800108

import (
	"crypto/hmac"
	"fmt"
	"hash"
)

// PRF is a Pseudorandom Function acceptable to the NIST KBKDF
type PRF interface {
	// Compute generates new data from a key and keying material data.
	Compute(key, data []byte) ([]byte, error)
	// OutputSizeBytes indicates the output size of the PRF in bytes, non-byte multiples of output length are not supported
	OutputSizeBytes() uint
	unimplementableNISTSP800108PRF()
}

// PRFHMAC is the HMAC implementation of PRF
type PRFHMAC struct {
	hashGen func() hash.Hash
}

// NewPRFHMAC creates a new HMAC PRF
func NewPRFHMAC(hashGen func() hash.Hash) *PRFHMAC {
	return &PRFHMAC{
		hashGen: hashGen,
	}
}

// OutputSizeBytes indicates the output size of the PRF in bytes, non-byte multiples of output length are not supported
func (p *PRFHMAC) OutputSizeBytes() uint {
	if p == nil || p.hashGen == nil {
		return 0
	}
	h := p.hashGen()
	if h == nil {
		return 0
	}
	return uint(h.Size())
}

// Compute generatesd new data from a key and keying material data
func (p *PRFHMAC) Compute(key, data []byte) ([]byte, error) {
	if p == nil || p.hashGen == nil {
		return nil, fmt.Errorf("invalid PRF")
	}
	h := hmac.New(p.hashGen, key)
	_, err := h.Write(data)
	if err != nil {
		return nil, err
	}
	out := h.Sum(nil)
	return out, nil
}

func (p *PRFHMAC) unimplementableNISTSP800108PRF() {}
