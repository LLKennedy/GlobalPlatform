package apdu

import "fmt"

// CLASecureMessaging indicates the status and format of Secure Messaging for the command
type CLASecureMessaging byte

const (
	// CLASMNone is no SM or no indication of SM (may still be implied from context, e.g. secure session)
	CLASMNone = 0x00
	// CLASMProprietary is proprietary secure messaging
	CLASMProprietary = 0x01
	// CLASMISONoHeaderProcessing is secure messaging as defined in ISO-IEC 7816-4 section 6, with the command header not processed
	CLASMISONoHeaderProcessing = 0x02
	// CLASMISOHeaderAuth is secure messaging as defined in ISO-IEC 7816-4 section 6, with the command header authenticated
	CLASMISOHeaderAuth = 0x03
)

// Class is any data type which can be converted to a Class byte
// TODO: do we need setters, or will all processing be done before it gets passed along as the interface type?
type Class interface {
	// ToClassByte converts the class to a class byte
	ToClassByte() byte
}

const (
	firstInterindustryClassBase    byte = 0x00
	furtherInterindustroyClassBase byte = 0x40
	longChannelsStart                   = 4
	longChannelsEnd                     = 19
)

// InterindustryClass represents the APDU class byte and its many flags and functions
type InterindustryClass struct {
	NotLastCommandOfChain bool // True = more messages will be sent as part of this command, e.g. data overflow requiring multiple messages
	SecureMessaging       CLASecureMessaging
	LogicalChannelNumber  uint8 // Must be 0-19, values 4-19 restrict options for SecureMessaging to None or ISONoHeaderProcessing
}

// InterindustryClassFromByte parses a single byte to an InterindustryClass
func InterindustryClassFromByte(in byte) (class InterindustryClass, err error) {
	switch {
	case (in & (b6 | b7 | b8)) == 0:
		class.LogicalChannelNumber = in & 0x03
		class.SecureMessaging = CLASecureMessaging((in & (b4 | b3)) >> 2)
		class.NotLastCommandOfChain = (in & b5) == b5
	case (in & (b7 | b8)) == b7:
		class.LogicalChannelNumber = (in & 0x0F) + 4
		if (in & b6) == b6 {
			class.SecureMessaging = CLASMISONoHeaderProcessing
		} else {
			class.SecureMessaging = CLASMNone
		}
		class.NotLastCommandOfChain = (in & b5) == b5
	default:
		err = fmt.Errorf("lead bits did not match any Interindustry class byte format")
	}
	return
}

// ToClassByte converts the class to a class byte
func (c InterindustryClass) ToClassByte() byte {
	var out byte
	shortChannels := true
	// >19 is invalid and will be treated as zero
	if c.LogicalChannelNumber >= longChannelsStart && c.LogicalChannelNumber <= longChannelsEnd {
		out = furtherInterindustroyClassBase
		shortChannels = false
	} else {
		out = firstInterindustryClassBase
	}
	if c.NotLastCommandOfChain {
		out = out | b5
	}
	// We'll interpret an invalid SM setting as none/no indication, and none means all zeroes, so no action here in that case
	if shortChannels {
		if c.SecureMessaging < 4 {
			out = out | byte(c.SecureMessaging<<2)
		}
	} else if c.SecureMessaging == CLASMISONoHeaderProcessing {
		out = out | b6
	}
	// Same here, >19 == 0 == no action
	if shortChannels {
		if c.LogicalChannelNumber < longChannelsStart {
			out = out | c.LogicalChannelNumber
		}
	} else if c.LogicalChannelNumber >= longChannelsStart && c.LogicalChannelNumber <= longChannelsEnd {
		out = out | (c.LogicalChannelNumber - longChannelsStart)
	}
	return out
}

// IsLastCommand returns whether this is the last command in the chian
func (c InterindustryClass) IsLastCommand() bool {
	return !c.NotLastCommandOfChain
}

// GetSMIndication returns the indication of Secure Messaging
func (c InterindustryClass) GetSMIndication() CLASecureMessaging {
	if c.LogicalChannelNumber < longChannelsStart && c.SecureMessaging < 4 {
		return c.SecureMessaging
	}
	if c.LogicalChannelNumber >= longChannelsStart && c.LogicalChannelNumber <= longChannelsEnd && c.SecureMessaging == CLASMISONoHeaderProcessing {
		return CLASMISONoHeaderProcessing
	}
	return CLASMNone
}

// GetLogicalChannel returns the logical channel number
func (c InterindustryClass) GetLogicalChannel() uint8 {
	if c.LogicalChannelNumber > longChannelsEnd {
		return 0
	}
	return c.LogicalChannelNumber
}
