package gpapdu

import "github.com/llkennedy/globalplatform/goimpl/apdu"

// Context is a GP APDU context, which maintains state based on previous calls and handles concurrency safety
type Context struct {
	transport *apdu.TransportWrapper
	s8mode    bool
}
