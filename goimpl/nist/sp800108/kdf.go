package sp800108

import (
	"crypto/subtle"
	"encoding/binary"
	"fmt"
	"math"
)

// Counter defines a valid counter length
type Counter uint8

const (
	// CounterLength8 is an 8-bit counter
	CounterLength8 Counter = 8
	// CounterLength16 is a 16-bit counter
	CounterLength16 Counter = 16
	// CounterLength24 is a 24-bit counter
	CounterLength24 Counter = 24
	// CounterLength32 is a 32-bit counter
	CounterLength32 Counter = 32
)

// InputStringOrdering defines an element in the KDF Input String
type InputStringOrdering uint

const (
	// InputOrderLabel is the label
	InputOrderLabel InputStringOrdering = iota
	// InputOrderContext is the context
	InputOrderContext
	// InputOrderL is the output data length as big-endian binary integer L
	InputOrderL
	// InputOrderEmptySeparator is 0x00
	InputOrderEmptySeparator
	// InputOrderCounter is the counter variable
	InputOrderCounter
)

func validateOrdering(order []InputStringOrdering) error {
	if order == nil {
		return fmt.Errorf("completely empty input data is invalid")
	}
	if len(order) > 5 {
		return fmt.Errorf("there are only 5 valid input string element types, ordering is invalid with %d elements", len(order))
	}
	foundCounter := false
	for i, orderElem := range order {
		switch orderElem {
		case InputOrderContext, InputOrderCounter, InputOrderEmptySeparator, InputOrderL, InputOrderLabel:
			// Valid value, now check it's unique in this order
			for j := i + 1; j < len(order); j++ {
				if orderElem == order[j] {
					return fmt.Errorf("duplicate input string ordering element at indices %d and %d", i, j)
				}
			}
			if orderElem == InputOrderCounter {
				foundCounter = true
			}
		default:
			return fmt.Errorf("invalid input string ordering element: %d", orderElem)
		}
	}
	if !foundCounter {
		return fmt.Errorf("counter is mandatory in input string ordering")
	}
	return nil
}

// KDF is a Key Derivation Function
type KDF interface {
	// Derive uses the key and keying material to generate a derived key(set) of the desired size. Output size is big-endian
	Derive(prf PRF, counterLengthR Counter, inputKey, label, context, outputSizeBits []byte, inputOrder []InputStringOrdering) ([]byte, error)
}

// CounterKBKDF is the NIST Key Based Key Derivation Function in counter mode
type CounterKBKDF struct{}

// Derive uses the key and keying material to generate a derived key(set) of the desired size
func (c *CounterKBKDF) Derive(prf PRF, counterLengthR Counter, inputKey, label, context, outputSizeBits []byte, inputOrder []InputStringOrdering) ([]byte, error) {
	if prf == nil {
		return nil, fmt.Errorf("must provide PRF")
	}
	var nLimit uint32
	switch {
	case counterLengthR == CounterLength8:
		nLimit = 0xFF
	case counterLengthR == CounterLength16:
		nLimit = 0xFFFF
	case counterLengthR == CounterLength24:
		nLimit = 0xFFFFFF
	case counterLengthR == CounterLength32:
		nLimit = 0xFFFFFFFF
	case counterLengthR > 32:
		// r <= 32 by definition
		return nil, fmt.Errorf("counter length too large, must be exactly 8, 16, 24 or 32")
	default:
		// Allowed by the spec but not supported since we're working with whole bytes oonly
		return nil, fmt.Errorf("invalid counter length, must be exactly 8, 16, 24 or 32")
	}
	if err := validateOrdering(inputOrder); err != nil {
		return nil, fmt.Errorf("invalid input string ordering: %w", err)
	}
	hBytes := prf.OutputSizeBytes()
	h := hBytes * 8
	l2 := make([]byte, len(outputSizeBits))
	subtle.ConstantTimeCopy(1, l2, outputSizeBits)
	for len(l2) < 8 {
		l2 = append([]byte{0}, l2...)
	}
	L := binary.BigEndian.Uint64(l2)
	outputSizeBytes := uint64(math.Ceil(float64(L) / 8))
	if h == 0 {
		return nil, fmt.Errorf("PRF must return non-zero data")
	}
	// Step 1: n := ceil(L/h)
	n := uint32(math.Ceil(float64(L) / float64(h)))
	// Step 2: If n > (2^r) -1, abort
	if n > nLimit {
		return nil, fmt.Errorf("invalid output length for given PRF and counter size")
	}
	// Step 3: initialise result vector to L bits
	result := make([]byte, uint(n)*hBytes)
	// Step 4: For i = 1 to n, do K(i) := PRF(K_I, [i]_2 || Label || 0x00 || Context || [L]_2)
	// Allow optional omission of some fixed input fields though
	prefix := []byte{}
	suffix := []byte{}
	foundCounterInOrdering := false
	appendToCurrent := func(next []byte) {
		if foundCounterInOrdering {
			suffix = append(suffix, next...)
		} else {
			prefix = append(prefix, next...)
		}
	}
	for _, orderElem := range inputOrder {
		switch orderElem {
		case InputOrderCounter:
			foundCounterInOrdering = true
		case InputOrderContext:
			appendToCurrent(context)
		case InputOrderEmptySeparator:
			appendToCurrent([]byte{0x00})
		case InputOrderL:
			appendToCurrent(outputSizeBits)
		case InputOrderLabel:
			appendToCurrent(label)
		}
	}
	for i := uint32(1); i <= n; i++ {
		iBytes := make([]byte, 4)
		binary.BigEndian.PutUint32(iBytes, i)
		// Trim counter binary representation accordingly
		switch counterLengthR {
		case CounterLength8:
			iBytes = iBytes[3:]
		case CounterLength16:
			iBytes = iBytes[2:]
		case CounterLength24:
			iBytes = iBytes[1:]
		}
		nextInput := append([]byte{}, prefix...)
		nextInput = append(nextInput, iBytes...)
		nextInput = append(nextInput, suffix...)
		nextOutput, err := prf.Compute(inputKey, nextInput)
		if err != nil {
			return nil, fmt.Errorf("PRF error: %w", err)
		}
		subtle.ConstantTimeCopy(1, result[uint(i-1)*hBytes:uint(i)*hBytes], nextOutput)
	}
	return result[:outputSizeBytes], nil
}
