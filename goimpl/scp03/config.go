package scp03

// Configuration is the Secure Communication Configuration object "i" defined in section 5.1 with some hopefully sensible defaults
type Configuration struct {
	LegacyS8Mode          bool // Whether to drop down to legacy S8 mode
	PseudoRandomChallenge bool // Whether to drop random challenge down to pseudo-random callege
	NoRMACEncryption      bool // Whether all R-MAC and R-Encryption support is disabled
	NoREncryption         bool // Whether only R-Encryption support is disabled
}

// ParseConfiguration parses a byte into the config flags
func ParseConfiguration(in byte) Configuration {
	conf := Configuration{}
	conf.LegacyS8Mode = in&0b00000001 == 0
	conf.PseudoRandomChallenge = in&0b00010000 > 0
	conf.NoRMACEncryption = in&0b00100000 == 0
	if conf.NoRMACEncryption {
		conf.NoREncryption = true
	} else {
		conf.NoREncryption = in&0b01000000 == 0
	}
	return conf
}

// ToByte converts the config flags to a byte
func (c Configuration) ToByte() byte {
	out := byte(0)
	if !c.LegacyS8Mode {
		out = out | 0b00000001
	}
	if c.PseudoRandomChallenge {
		out = out | 0b00010000
	}
	if !c.NoRMACEncryption {
		out = out | 0b00100000
		if !c.NoREncryption {
			out = out | 0b01000000
		}
	}
	return out
}
