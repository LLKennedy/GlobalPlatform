package apdu

import (
	"fmt"
	"sync"
)

// Transport is an APDU transport system, the TPDU layer is completely abstracted and not within the scope of this project
type Transport interface {
	// Send attempts to send a Command over the transport, returning the Response.
	// APDU response level errors (error bytes in SW1 and SW2) should result in error = nil, error should only be used for failure to complete the transmission (e.g. card removed from reader).
	// NOTE: It is the responsibility of the Transport to prohibit command interleaving, this includes thread safety measures to prevent simultaneous use of the Send command before completion of each pair.
	// If thread-safety is not and cannot be built into the underlying Transport implementation T, use &TransportWrapper{Transport: T} to implement these safety features for you.
	Send(Command) (Response, error)
}

// TransportWrapper wraps a Transport implementation in thread-safe features to prevent interleaving of command response pairs according to ISO-IEC 7816-4
type TransportWrapper struct {
	sync.Mutex
	Transport Transport
}

// Send attempts to send a Command over the transport, returning the Response.
func (T *TransportWrapper) Send(cmd Command) (Response, error) {
	if T == nil {
		return Response{}, fmt.Errorf("invalid transport wrapper, must be non-nil")
	}
	T.Lock()
	defer T.Unlock()
	if T.Transport == nil {
		return Response{}, fmt.Errorf("invalid transport wrapper, wrapped transport must be non-nil")
	}
	return T.Transport.Send(cmd)
}
