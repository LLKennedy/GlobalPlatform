package apdu

// Context is an APDU context, which maintains state based on previous calls and handles concurrency safety
type Context struct {
	transport *TransportWrapper
}
