package gpids

// RegisteredApplicationProviderIdentifier is the RID assigned to GlobalPlatform
const RegisteredApplicationProviderIdentifier uint64 = 0xA000000151

// DefaultIssuerSecurityDomainAID is the default AID of the security issuer domain, based on the GP RID
const DefaultIssuerSecurityDomainAID uint64 = RegisteredApplicationProviderIdentifier << (3 * 8)
