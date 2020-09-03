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
	conf.LegacyS8Mode = in&b1 == 0
	conf.PseudoRandomChallenge = in&b5 > 0
	conf.NoRMACEncryption = in&b6 == 0
	if conf.NoRMACEncryption {
		conf.NoREncryption = true
	} else {
		conf.NoREncryption = in&b7 == 0
	}
	return conf
}

// ToByte converts the config flags to a byte
func (c Configuration) ToByte() byte {
	out := byte(0)
	if !c.LegacyS8Mode {
		out = out | b1
	}
	if c.PseudoRandomChallenge {
		out = out | b5
	}
	if !c.NoRMACEncryption {
		out = out | b6
		if !c.NoREncryption {
			out = out | b7
		}
	}
	return out
}
