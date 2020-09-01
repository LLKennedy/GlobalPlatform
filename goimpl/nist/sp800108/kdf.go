package sp800108

// Counter is valid storage for a counter-mode counter variable
type Counter = uint32

// KDF is a Key Derivation Function
type KDF interface {
	// Derive uses the key and keying material to generate a derived key(set) of the desired size
	Derive(data []byte, size int) []byte
}
