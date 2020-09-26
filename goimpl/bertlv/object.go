package bertlv

// Object is a full BER-TLV object
type Object struct {
	Tag Tag
	// Length should be purely decorative and merely represent the length of the Value slice, but in the case of a malformed TLV object Length represents the legth which was encoded, regardless of value's real length
	// During encoding, Length is always ignored. You cannot use Length to provide padding, you must do this manually since I can't possibly know what kind of padding you want.
	Length uint64
	Value  []byte
}
