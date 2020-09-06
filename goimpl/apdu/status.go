package apdu

// StatusCategory is a category of APDU status
type StatusCategory int

const (
	// StatusCategoryInvalid indicates an invalid status non-compliant with the spec
	StatusCategoryInvalid StatusCategory = iota
	// StatusCategoryNormal indicates normal processing without warning or error
	StatusCategoryNormal
	// StatusCategoryWarning indicates a warning processing the command
	StatusCategoryWarning
	// StatusCategoryExecError indicates an error executing the command
	StatusCategoryExecError
	// StatusCategoryCheckError indicates an error checking the command
	StatusCategoryCheckError
	// StatusCategoryProprietary indicates a proprietary status defined outside the ISO-IEC 7816 specification
	StatusCategoryProprietary
)

// Status is an APDU response status, it must be identified as one of the concrete types defined here to get more information
type Status interface {
	// Category is the category of the status, to indicate which implementations should be checked against
	Category() StatusCategory
	// Raw is the raw status bytes for further information or manual re-parsing
	Raw() RawStatus
}

// RawStatus is raw status bytes on their own, used to simplify more specific type definitions and conversions.
type RawStatus struct {
	SW1 byte // First status byte
	SW2 byte // Second status byte
}

// Raw is the raw status bytes for further information or manual re-parsing
func (r RawStatus) Raw() RawStatus {
	return r
}

// Identify parses status bytes and as a specific category of status response, and processes specific data for valid non-proprietary status bytes.
// Proprietary status definitions should wrap this functionality with additional processing for proprietary status information.
func (r RawStatus) Identify() Status {
	// TODO: this function is a mess of constants that could maybe be pulled out into a const block, but I'm not sure that'd be cleaner/easier to read or fix
	firstSW1Nibble := r.SW1 & 0xF0
	var combined uint16 = (uint16(r.SW1) << 8) | uint16(r.SW2)
	var out Status = StatusInvalid{r}
	switch {
	case firstSW1Nibble != 0x60 && firstSW1Nibble != 0x90, r.SW1 == 0x60:
		// Defined explicitly invalid
		out = StatusInvalid{r}
	case combined == 0x9000:
		// Normal exception to the proprietary rules below
		out = identifyNormal(r)
	case combined == 0x6700, combined == 0x6B00, combined == 0x6D00, combined == 0x6E00, combined == 0x6F00:
		// Checking error exceptions to the proprietary rules below
		out = identifyCheckError(r)
	case r.SW1 == 0x67, r.SW1 == 0x6B, r.SW1 == 0x6D, r.SW1 == 0x6E, r.SW1 == 0x6F, firstSW1Nibble == 0x90:
		// All values ruled proprietary, cannot further parse within this package
		out = StatusProprietary{r}
	case r.SW1 == 0x61:
		// Normal processing with data
		out = identifyNormal(r)
	case r.SW1 == 0x62, r.SW1 == 0x63:
		// All warning values
		out = identifyWarning(r)
	case r.SW1 == 0x64, r.SW1 == 0x65, r.SW1 == 0x66:
		// All exec error values
		out = identifyExecError(r)
	case r.SW1 == 0x68, r.SW1 == 0x69, r.SW1 == 0x6A, r.SW1 == 0x6C:
		// All check error values except the proprietary exceptions we caught earlier
		out = identifyCheckError(r)
	}
	return out
}

// StatusNormal indicates normal processing without warning or error
type StatusNormal struct {
	RawStatus
	NoFurtherQualification bool  // 9000
	RemainingDataLength    uint8 // 61XX
}

// Category is the category of the status, to indicate which implementations should be checked against
func (s StatusNormal) Category() StatusCategory {
	return StatusCategoryNormal
}

func identifyNormal(r RawStatus) Status {
	var out Status = StatusInvalid{r}
	switch r.SW1 {
	case 0x90:
		out = StatusNormal{
			RawStatus:              r,
			NoFurtherQualification: true,
		}
	case 0x61:
		out = StatusNormal{
			RawStatus:           r,
			RemainingDataLength: r.SW2,
		}
	}
	return out
}

// StatusWarning indicates a warning processing the command
type StatusWarning struct {
	NonVolatileMemoryChanged bool
	NoInformationGiven       bool
	WaitingQueryBytesLength  uint8
	RawStatus
}

// Category is the category of the status, to indicate which implementations should be checked against
func (s StatusWarning) Category() StatusCategory {
	return StatusCategoryWarning
}

func identifyWarning(r RawStatus) Status {
	var out Status = StatusInvalid{r}
	switch r.SW1 {
	case 0x62:
		s := StatusWarning{
			RawStatus:                r,
			NonVolatileMemoryChanged: false,
		}
		switch {
		case r.SW2 == 0x00:
			s.NoInformationGiven = true
		case r.SW2 >= 0x02 && r.SW2 <= 0x80:
			s.WaitingQueryBytesLength = r.SW2
		case r.SW2 == 0x81:
		default:
			return StatusInvalid{r}
		}
		out = s
	case 0x63:
		out = StatusWarning{
			RawStatus:                r,
			NonVolatileMemoryChanged: true,
		}
	}
	return out
}

// StatusExecError indicates an error executing the command
type StatusExecError struct {
	NonVolatileMemoryChanged bool
	WaitingQueryBytesLength  uint8
	RawStatus
}

// Category is the category of the status, to indicate which implementations should be checked against
func (s StatusExecError) Category() StatusCategory {
	return StatusCategoryExecError
}

func identifyExecError(r RawStatus) Status {
	var out Status = StatusInvalid{r}
	switch r.SW1 {
	}
	return out
}

// StatusCheckError indicates an error checking the command
type StatusCheckError struct {
	RawStatus
}

// Category is the category of the status, to indicate which implementations should be checked against
func (s StatusCheckError) Category() StatusCategory {
	return StatusCategoryCheckError
}

func identifyCheckError(r RawStatus) Status {
	var out Status = StatusInvalid{r}
	switch r.SW1 {
	}
	return out
}

// StatusProprietary indicates a proprietary status defined outside the ISO-IEC 7816 specification
type StatusProprietary struct {
	RawStatus
}

// Category is the category of the status, to indicate which implementations should be checked against
func (s StatusProprietary) Category() StatusCategory {
	return StatusCategoryProprietary
}

// StatusInvalid is an invalid status
type StatusInvalid struct {
	RawStatus
}

// Category is the category of the status, to indicate which implementations should be checked against
func (s StatusInvalid) Category() StatusCategory {
	return StatusCategoryInvalid
}
