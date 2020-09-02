package lifecycle

// Card is a life cycle state
type Card int

const (
	// CardOPReady is the OP_READY life cycle state
	CardOPReady Card = 0x01
	// CardInitialized is the INITIALIZED life cycle state
	CardInitialized Card = 0x07
	// CardSecured is the SECURED life cycle state
	CardSecured Card = 0x0F
	// CardCardLocked is the CARD_LOCKED life cycle state
	CardCardLocked Card = 0x7F
	// CardTerminated is the TERMINATED life cycle state
	CardTerminated Card = 0xFF
)
