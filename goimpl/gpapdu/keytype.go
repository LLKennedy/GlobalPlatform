package gpapdu

// KeyType is a key type indicator
type KeyType byte

const (
	// DESWithImplicitMode is the DESWithImplicitMode key type
	DESWithImplicitMode KeyType = 0x80
	// PreSharedKeyForTransportLayerSecurity is the PreSharedKeyForTransportLayerSecuritykey type
	PreSharedKeyForTransportLayerSecurity KeyType = 0x85
	// AES is the AES key type
	AES KeyType = 0x88
	// HMACSHA1WithImplictLength is the HMACSHA1WithImplictLength key type
	HMACSHA1WithImplictLength KeyType = 0x90
	// HMACSHA1Length160Bits is the HMACSHA1Length160Bits key type
	HMACSHA1Length160Bits KeyType = 0x91
	// RSAPublicKeyPubExponentEClearText is the RSAPublicKeyPubExponentEClearText key type
	RSAPublicKeyPubExponentEClearText KeyType = 0xA0
	// RSAPublicKeyModulesNClearText is the RSAPublicKeyModulesNClearText key type
	RSAPublicKeyModulesNClearText KeyType = 0xA1
	// RSAPrivateKeyModulusN is the RSAPrivateKeyModulusN key type
	RSAPrivateKeyModulusN KeyType = 0xA2
	// RSAPrivateKeyPrivateExponentD is the RSAPrivateKeyPrivateExponentD key type
	RSAPrivateKeyPrivateExponentD KeyType = 0xA3
	// RSAPrivateKeyChineseRemainderP is the RSAPrivateKeyChineseRemainderP key type
	RSAPrivateKeyChineseRemainderP KeyType = 0xA4
	// RSAPrivateKeyChineseRemainderQ is the RSAPrivateKeyChineseRemainderQ key type
	RSAPrivateKeyChineseRemainderQ KeyType = 0xA5
	// RSAPrivateKeyChineseRemainderPQ is the RSAPrivateKeyChineseRemainderPQ key type
	RSAPrivateKeyChineseRemainderPQ KeyType = 0xA6
	// RSAPrivateKeyChineseRemainderDP1 is the RSAPrivateKeyChineseRemainderDP1 key type
	RSAPrivateKeyChineseRemainderDP1 KeyType = 0xA7
	// RSAPrivateKeyChineseRemainderDQ1 is the RSAPrivateKeyChineseRemainderDQ1 key type
	RSAPrivateKeyChineseRemainderDQ1 KeyType = 0xA8
	// ECCPublicKey is the ECCPublicKey key type
	ECCPublicKey KeyType = 0xB0
	// ECCPrivateKey is the ECCPrivateKey key type
	ECCPrivateKey KeyType = 0xB1
	// ECCFieldParameterP is the ECCFieldParameterP key type
	ECCFieldParameterP KeyType = 0xB2
	// ECCFieldParameterA is the ECCFieldParameterA key type
	ECCFieldParameterA KeyType = 0xB3
	// ECCFieldParameterB is the ECCFieldParameterB key type
	ECCFieldParameterB KeyType = 0xB4
	// ECCFieldParameterG is the ECCFieldParameterG key type
	ECCFieldParameterG KeyType = 0xB5
	// ECCFieldParameterN is the ECCFieldParameterN key type
	ECCFieldParameterN KeyType = 0xB6
	// ECCFieldParameterK is the ECCFieldParameterK key type
	ECCFieldParameterK KeyType = 0xB7
	// ECCKeyParametersReference is the ECCKeyParametersReference key type
	ECCKeyParametersReference KeyType = 0xF0
	// ExtendedFormat is the ExtendedFormat key type
	ExtendedFormat KeyType = 0xFF
)
