package apdu

// Transport is an APDU transport system, the TPDU layer is completely abstracted and not within the scope of this project
type Transport interface {
	// Send attempts to send a Command over the transport, returning the Response.
	// APDU response level errors (error bytes in SW1 and SW2) should result in error = nil, error should only be used for failure to complete the transmission (card removed from reader, etc.).
	Send(Command) (Response, error)
}
