package gpapdu

import (
	"bytes"

	"github.com/llkennedy/globalplatform/goimpl/bertlv"
)

const (
	tagELFileOrAppID                 = 0x4F
	tagCRTFDS                        = 0xB6
	tagSecurityDomainID              = 0x42
	tagSecurityDomainImageNumber     = 0x45
	tagApplicationProviderIdentifier = 0x5F20
	tagTokenID                       = 0x9E
)

// DeleteCommand is a Delete command, either DeleteKey or DeleteCardContent
type DeleteCommand interface {
	unimplementableDeleteCommandToBytes() []byte
}

// ControlReferenceTemplateForDigitalSignature is a control reference template for digital signature
type ControlReferenceTemplateForDigitalSignature struct {
	SecurityDomainID          []byte
	SecurityDomainImageNumber []byte
	ApplicationProviderID     []byte
	TokenID                   []byte
}

// ToBerTlv encodes the data as BER-TLV
func (c ControlReferenceTemplateForDigitalSignature) ToBerTlv() (obj bertlv.Object) {
	obj.Tag = bertlv.TagFromUintForced(tagCRTFDS)
	buf := bytes.NewBuffer(nil)
	writer, _ := bertlv.NewWriter(buf)
	if c.SecurityDomainID != nil {
		writer.Write(bertlv.Object{
			Tag:   bertlv.TagFromUintForced(tagSecurityDomainID),
			Value: c.SecurityDomainID,
		})
	}
	if c.SecurityDomainImageNumber != nil {
		writer.Write(bertlv.Object{
			Tag:   bertlv.TagFromUintForced(tagSecurityDomainImageNumber),
			Value: c.SecurityDomainImageNumber,
		})
	}
	if c.ApplicationProviderID != nil {
		writer.Write(bertlv.Object{
			Tag:   bertlv.TagFromUintForced(tagApplicationProviderIdentifier),
			Value: c.ApplicationProviderID,
		})
	}
	if c.TokenID != nil {
		writer.Write(bertlv.Object{
			Tag:   bertlv.TagFromUintForced(tagTokenID),
			Value: c.TokenID,
		})
	}
	obj.Value = buf.Bytes()
	return
}

// DeleteCardContent is Delete [card content] command
type DeleteCardContent struct {
	ELFileOrAppID []byte
	CRTFDS        *ControlReferenceTemplateForDigitalSignature
}

func (d DeleteCardContent) unimplementableDeleteCommandToBytes() []byte {

	panic("not implemented") // FIXME
}

// DeleteKey is a Delete [key] command
type DeleteKey struct {
	IncludeKeyIdentifer     bool
	KeyIdentifier           byte
	IncludeKeyVersionNumber bool
	KeyVersionNumber        byte
}

func (d DeleteKey) unimplementableDeleteCommandToBytes() []byte {
	panic("not implemented") // FIXME
}
