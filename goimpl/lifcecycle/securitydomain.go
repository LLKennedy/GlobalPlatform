package lifecycle

// SecurityDomain is a Security Domain life cycle state
type SecurityDomain int

const (
	// SecurityDomainInstalled is the INSTALLED security domain life cycle state
	SecurityDomainInstalled SecurityDomain = 0x03
	// SecurityDomainSelectable is the SELECTABLE security domain life cycle state
	SecurityDomainSelectable SecurityDomain = 0x07
	// SecurityDomainPersonalized is the PERSONALIZED security domain life cycle state
	SecurityDomainPersonalized SecurityDomain = 0x0F
	// SecurityDomainLocked is the LOCKED security domain life cycle state
	SecurityDomainLocked SecurityDomain = 0x83
)
