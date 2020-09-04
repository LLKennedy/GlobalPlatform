package gpapdu

const (
	privSecurityDomain            = 0x80
	privDAPVerification           = 0x40
	privDelegatedManagement       = 0x20
	privCardLock                  = 0x10
	privCardTerminate             = 0x08
	privCardReset                 = 0x04
	privCVMManagement             = 0x02
	privMandatedDAPVerification   = 0x01
	privTrustedPath               = 0x80
	privAuthorizedManagement      = 0x40
	privTokenManagement           = 0x20
	privGlobalDelete              = 0x10
	privGlobalLock                = 0x08
	privGlobalRegistry            = 0x04
	privFinalApplication          = 0x02
	privGlobalService             = 0x01
	privReceiptGeneration         = 0x80
	privCipheredLoadFileDataBlock = 0x40
	privContactlessActivation     = 0x20
	privContactlessSelfActivation = 0x10
	privReserved1                 = 0x08
	privReserved2                 = 0x04
	privReserved3                 = 0x02
	privReserved4                 = 0x01
)

// Privileges are the privileges granted to an application or domain
type Privileges struct {
	SecurityDomain                             bool
	DAPVerification                            bool // Implies SecurityDomain
	DelegatedManagement                        bool // Implies SecurityDomain
	CardLock                                   bool
	CardTerminate                              bool
	CardReset                                  bool
	CVMManagement                              bool
	MandatedDAPVerification                    bool // Implies DAPVerification and SecurityDomain
	TrustedPath                                bool
	AuthorizedManagement                       bool // Implies SecurityDomain
	TokenManagement                            bool
	GlobalDelete                               bool
	GlobalLock                                 bool
	GlobalRegistry                             bool
	FinalApplication                           bool
	GlobalService                              bool
	ReceiptGeneration                          bool
	CipheredLoadFileDataBlock                  bool
	ContactlessActivation                      bool
	ContactlessSelfActivation                  bool
	reserved1, reserved2, reserved3, reserved4 bool
}

// ToBytes converts a Privileges struct to the APDU bytes representation
func (p Privileges) ToBytes() [3]byte {
	out := [3]byte{0, 0, 0}
	// First byte's flags
	if p.SecurityDomain || p.DAPVerification || p.DelegatedManagement || p.MandatedDAPVerification || p.AuthorizedManagement {
		out[0] = out[0] | (privSecurityDomain)
	}
	if p.DAPVerification || p.MandatedDAPVerification {
		out[0] = out[0] | privDAPVerification
	}
	if p.DelegatedManagement {
		out[0] = out[0] | privDelegatedManagement
	}
	if p.CardLock {
		out[0] = out[0] | privCardLock
	}
	if p.CardTerminate {
		out[0] = out[0] | privCardTerminate
	}
	if p.CardReset {
		out[0] = out[0] | privCardReset
	}
	if p.CVMManagement {
		out[0] = out[0] | privCVMManagement
	}
	if p.MandatedDAPVerification {
		out[0] = out[0] | privMandatedDAPVerification
	}
	// Second byte's flags
	if p.TrustedPath {
		out[1] = out[1] | privTrustedPath
	}
	if p.AuthorizedManagement {
		out[1] = out[1] | privAuthorizedManagement
	}
	if p.TokenManagement {
		out[1] = out[1] | privTokenManagement
	}
	if p.GlobalDelete {
		out[1] = out[1] | privGlobalDelete
	}
	if p.GlobalLock {
		out[1] = out[1] | privGlobalLock
	}
	if p.GlobalRegistry {
		out[1] = out[1] | privGlobalRegistry
	}
	if p.FinalApplication {
		out[1] = out[1] | privFinalApplication
	}
	if p.GlobalService {
		out[1] = out[1] | privGlobalService
	}
	// Third byte's flags
	if p.ReceiptGeneration {
		out[2] = out[2] | privReceiptGeneration
	}
	if p.CipheredLoadFileDataBlock {
		out[2] = out[2] | privCipheredLoadFileDataBlock
	}
	if p.ContactlessActivation {
		out[2] = out[2] | privContactlessActivation
	}
	if p.ContactlessSelfActivation {
		out[2] = out[2] | privContactlessSelfActivation
	}
	if p.reserved1 {
		out[2] = out[2] | privReserved1
	}
	if p.reserved2 {
		out[2] = out[2] | privReserved2
	}
	if p.reserved3 {
		out[2] = out[2] | privReserved3
	}
	if p.reserved4 {
		out[2] = out[2] | privReserved4
	}
	return out
}

// PrivilegesFromBytes converts 3 bytes to a set of Privilegs flags
func PrivilegesFromBytes(in [3]byte) Privileges {
	out := Privileges{}
	// First byte's flags
	if in[0]&privSecurityDomain == privSecurityDomain {
		out.SecurityDomain = true
	}
	if in[0]&privDAPVerification == privDAPVerification {
		out.DAPVerification = true
		out.SecurityDomain = true
	}
	if in[0]&privDelegatedManagement == privDelegatedManagement {
		out.DelegatedManagement = true
		out.SecurityDomain = true
	}
	if in[0]&privCardLock == privCardLock {
		out.CardLock = true
	}
	if in[0]&privCardTerminate == privCardTerminate {
		out.CardTerminate = true
	}
	if in[0]&privCardReset == privCardReset {
		out.CardReset = true
	}
	if in[0]&privCVMManagement == privCVMManagement {
		out.CVMManagement = true
	}
	if in[0]&privMandatedDAPVerification == privMandatedDAPVerification {
		out.MandatedDAPVerification = true
		out.DAPVerification = true
		out.SecurityDomain = true
	}
	if in[1]&privTrustedPath == privTrustedPath {
		out.TrustedPath = true
	}
	if in[1]&privAuthorizedManagement == privAuthorizedManagement {
		out.AuthorizedManagement = true
		out.SecurityDomain = true
	}
	if in[1]&privTokenManagement == privTokenManagement {
		out.TokenManagement = true
	}
	if in[1]&privGlobalDelete == privGlobalDelete {
		out.GlobalDelete = true
	}
	if in[1]&privGlobalLock == privGlobalLock {
		out.GlobalLock = true
	}
	if in[1]&privGlobalRegistry == privGlobalRegistry {
		out.GlobalRegistry = true
	}
	if in[1]&privFinalApplication == privFinalApplication {
		out.FinalApplication = true
	}
	if in[1]&privGlobalService == privGlobalService {
		out.GlobalService = true
	}
	if in[2]&privReceiptGeneration == privReceiptGeneration {
		out.ReceiptGeneration = true
	}
	if in[2]&privCipheredLoadFileDataBlock == privCipheredLoadFileDataBlock {
		out.CipheredLoadFileDataBlock = true
	}
	if in[2]&privContactlessActivation == privContactlessActivation {
		out.ContactlessActivation = true
	}
	if in[2]&privContactlessSelfActivation == privContactlessSelfActivation {
		out.ContactlessSelfActivation = true
	}
	if in[2]&privReserved1 == privReserved1 {
		out.reserved1 = true
	}
	if in[2]&privReserved2 == privReserved2 {
		out.reserved2 = true
	}
	if in[2]&privReserved3 == privReserved3 {
		out.reserved3 = true
	}
	if in[2]&privReserved4 == privReserved4 {
		out.reserved4 = true
	}
	return out
}
