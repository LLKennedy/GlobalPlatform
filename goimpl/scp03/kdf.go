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

// KDFOutputLength is a valid output length for the KDF
type KDFOutputLength uint16

const (
	// KDFOutput64 is a 64 bit output from the KDF
	KDFOutput64 = 0x0040
	// KDFOutput128 is a 128 bit output from the KDF
	KDFOutput128 = 0x0080
	// KDFOutput192 is a 192 bit output from the KDF
	KDFOutput192 = 0x00C0
	// KDFOutput256 is a 256 bit output from the KDF
	KDFOutput256 = 0x0100
)

// Derive derives data from a base key and input data
func (k *KDF) Derive(key []byte, rawKDF sp800108.KDF, label [11]byte, ddc DataDerivationConstant, length KDFOutputLength, context []byte /*FIXME: "further specified"*/) ([]byte, error) {
	if key == nil || rawKDF == nil {
		return nil, fmt.Errorf("KDF: nil parameters")
	}
	switch ddc {
	case DDCCardChallengeGeneration, DDCCardCryptogram, DDCHostCryptogram, DDCSENC, DDCSMAC, DDCSRMAC:
		// Supported value
	default:
		return nil, fmt.Errorf("KDF: Invalid data derivation constant: %x", ddc)
	}
	lengthData := []byte{0, 0}
	switch length {
	case KDFOutput64:
		lengthData[1] = 0x40
	case KDFOutput128:
		lengthData[1] = 0x80
	case KDFOutput192:
		lengthData[1] = 0xC0
	case KDFOutput256:
		lengthData[0] = 0x01
	default:
		return nil, fmt.Errorf("KDF: Invalid output length: %d", length)
	}
	prf := &sp800108.PRFCMAC{}
	ordering := []sp800108.InputStringOrdering{sp800108.InputOrderLabel, sp800108.InputOrderEmptySeparator, sp800108.InputOrderL, sp800108.InputOrderCounter, sp800108.InputOrderCounter}
	return rawKDF.Derive(prf, sp800108.CounterLength8, key, label[:], context, lengthData, ordering)
}
