package sp800108

import (
	"crypto/hmac"
	"fmt"
	"hash"
)

// PRF is a Pseudorandom Function acceptable to the NIST KBKDF
type PRF interface {
	// Compute generates new data from a key and keying material data.
	Compute(data []byte) ([]byte, error)
	unimplementableNISTSP800108PRF()
}

// PRFHMAC is the HMAC implementation of PRF
type PRFHMAC struct {
	hash hash.Hash
}

// NewPRFHMAC creates a new HMAC PRF
func NewPRFHMAC(hashGen func() hash.Hash, key []byte) *PRFHMAC {
	return &PRFHMAC{
		hash: hmac.New(hashGen, key),
	}
}

// Compute generatesd new data from a key and keying material data
func (p *PRFHMAC) Compute(data []byte) ([]byte, error) {
	if p == nil || p.hash == nil {
		return nil, fmt.Errorf("invalid PRF")
	}
	p.hash.Reset()
	_, err := p.hash.Write(data)
	if err != nil {
		return nil, err
	}
	out := p.hash.Sum(nil)
	return out, nil
}

func (p *PRFHMAC) unimplementableNISTSP800108PRF() {}
