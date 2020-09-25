package gpids

import "encoding/asn1"

const (
	iso            = 1
	memberbody     = 2
	usa            = 840
	globalplatform = 114283
)

// BaseOID is the base Global Platform OID -
// {iso(1) member-body(2) country-USA(840) GlobalPlatform(114283)}
func BaseOID() asn1.ObjectIdentifier {
	return asn1.ObjectIdentifier{
		iso,
		memberbody,
		usa,
		globalplatform,
	}
}

// CardRecognitionDataOID is the OID for Card Recognition Data, as well as identifying GP as the Tag Allocation Authority for CRD objects -
// {globalPlatform 1}
func CardRecognitionDataOID() asn1.ObjectIdentifier {
	return append(BaseOID(), 1)
}
