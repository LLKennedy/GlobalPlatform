package bertlv

import (
	"encoding/binary"
	"fmt"
	"io"
)

// TagClass is the class of a Tag
type TagClass byte

const (
	// TagClassUniversal is the Universal Tag Class
	TagClassUniversal TagClass = iota
	// TagClassApplication is the Application Tag Class
	TagClassApplication
	// TagClassContextSpecific is a context-specific Tag Class
	TagClassContextSpecific
	// TagClassPrivate is a private Tag Class
	TagClassPrivate
)

// Tag is the Type of Type-Length-Value, always an ASN.1 tag (class and number)
type Tag struct {
	Class               TagClass
	ConstructedEncoding bool
	Number              uint64 // Technically there is no upper limit on the bit size of the tag number.
	// You could have literally infinity octets with all bits set to one and it'd be valid. I'm capping at 64bit for sanity.
}

// ToBytes converts a Tag to bytes
func (t Tag) ToBytes() ([]byte, error) {
	switch t.Class {
	case TagClassUniversal, TagClassApplication, TagClassContextSpecific, TagClassPrivate:
	default:
		return nil, fmt.Errorf("invalid tag class: %d", t.Class)
	}
	data := []byte{byte(t.Class) << 6}
	if t.ConstructedEncoding {
		data[0] = data[0] | b6
	}
	if t.Number <= 30 {
		data[0] = data[0] | byte(t.Number)
		return data, nil
	}
	data[0] = data[0] | 31
	raw := make([]byte, 8)
	binary.BigEndian.PutUint64(raw, t.Number)
	upperLimit := uint64(0x7F)
	requiredBytes := 1
	for upperLimit < t.Number {
		requiredBytes++
		upperLimit = upperLimit << 7
		upperLimit = upperLimit | 0x7F
	}
	formattedNumber := make([]byte, requiredBytes)
	bits := &reverseBitReader{
		data: raw,
	}
	for i := requiredBytes - 1; i >= 0; i-- {
		if i != requiredBytes-1 {
			// Last byte get marked by b8
			formattedNumber[i] = b8
		}
		// Set bits 0-7 according to the real data
		for j := 0; j < 7; j++ {
			nextBit := bits.ReadBit()
			formattedNumber[i] = formattedNumber[i] | (nextBit << j)
		}
	}
	data = append(data, formattedNumber...)
	return data, nil
}

// Write writes the tag to an io.Writer
func (t Tag) Write(w io.Writer) error {
	data, err := t.ToBytes()
	if err != nil {
		return err
	}
	n, err := w.Write(data)
	if err != nil {
		return err
	}
	if n < len(data) {
		return fmt.Errorf("failed to write full tag data, only got %d bytes of %d", n, len(data))
	}
	return nil
}

// TagFromBytes converts a tag from an io.Reader, it will only read more than necessary if the tag is invalid in the number section and will return overread = true if it overread by one byte.
func TagFromBytes(data io.Reader) (tag Tag, overread bool, err error) {
	return
}

type reverseBitReader struct {
	data     []byte
	index    int
	bitindex int
}

func (b *reverseBitReader) ReadBit() byte {
	maxIndex := len(b.data) - 1
	currentByte := b.data[maxIndex-b.index]
	bit := currentByte & (b1 << b.bitindex)
	b.bitindex++
	if b.bitindex > 7 {
		b.bitindex = 0
		b.index++
	}
	if bit > 0 {
		// Always return true as just b1
		return b1
	}
	return 0
}
