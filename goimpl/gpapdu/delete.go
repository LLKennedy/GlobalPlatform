package gpapdu

// DeleteCommand is a Delete command, either DeleteKey or DeleteCardContent
type DeleteCommand interface {
	unimplementableDeleteCommandToBytes() []byte
}

// DeleteCardContent is Delete [card content] command
type DeleteCardContent struct {
	ELFileOrAppID []byte
	CRTFDS        *ControlReferenceTemplateForDigitalSignature
}

// ControlReferenceTemplateForDigitalSignature is a control reference template for digital signature
type ControlReferenceTemplateForDigitalSignature struct {
	SecurityDomainID          []byte
	SecurityDomainImageNumber []byte
	ApplicationProviderID     []byte
	TokenID                   []byte
}

func (d DeleteCardContent) unimplementableDeleteCommand() {}

// DeleteKey is a Delete [key] command
type DeleteKey struct {
	IncludeKeyIdentifer     bool
	KeyIdentifier           byte
	IncludeKeyVersionNumber bool
	KeyVersionNumber        byte
}

func (d DeleteKey) unimplementableDeleteCommand() {}
