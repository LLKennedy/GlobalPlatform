package scp03

import (
	"fmt"

	"github.com/llkennedy/globalplatform/goimpl/nist/sp800108"
)

// KDF is a wrapper around SP800-108's KBKDF with SCP03-specified parameters
type KDF struct{}

// DataDerivationConstant is a data derivation constant for KDF functions
type DataDerivationConstant byte

const (
	// DDCCardCryptogram is a Card Cryptogram Data Derivation Constant
	DDCCardCryptogram DataDerivationConstant = 0b00000000
	// DDCHostCryptogram is a Host Cryptogram Data Derivation Constant
	DDCHostCryptogram DataDerivationConstant = 0b00000001
	// DDCCardChallengeGeneration is a Card Challenge Generation Data Derivation Constant
	DDCCardChallengeGeneration DataDerivationConstant = 0b00000010
	// DDCSENC is a S-ENC derivation Data Derivation Constant
	DDCSENC DataDerivationConstant = 0b00000100
	// DDCSMAC is a S-MAC derivation Data Derivation Constant
	DDCSMAC DataDerivationConstant = 0b00000110
	// DDCSRMAC is a S-RMAC derivation Data Derivation Constant
	DDCSRMAC DataDerivationConstant = 0b00000111
)

// Derive derives data from a base key and input data
func (k *KDF) Derive(key []byte, rawKDF sp800108.KDF, label [11]byte, ddc DataDerivationConstant) ([]byte, error) {
	if key == nil || rawKDF == nil {
		return nil, fmt.Errorf("KDF: nil parameters")
	}
	switch ddc {
	case DDCCardChallengeGeneration, DDCCardCryptogram, DDCHostCryptogram, DDCSENC, DDCSMAC, DDCSRMAC:
		// Supported value
	default:
		return nil, fmt.Errorf("KDF: Invalid data derivation constant: %x", ddc)
	}

	return nil, fmt.Errorf("not implemented")
}
