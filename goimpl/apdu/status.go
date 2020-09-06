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
	firstSW1Nibble := r.SW1 & 0xF0
	if (firstSW1Nibble != 0x60 && firstSW1Nibble != 0x90) || r.SW1 == 0x60 {
		// Defined explicitly invalid
		return StatusInvalid{r}
	}
	return StatusInvalid{r} // Not implemented yet
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

// StatusWarning indicates a warning processing the command
type StatusWarning struct {
	RawStatus
}

// Category is the category of the status, to indicate which implementations should be checked against
func (s StatusWarning) Category() StatusCategory {
	return StatusCategoryWarning
}

// StatusExecError indicates an error executing the command
type StatusExecError struct {
	RawStatus
}

// Category is the category of the status, to indicate which implementations should be checked against
func (s StatusExecError) Category() StatusCategory {
	return StatusCategoryExecError
}

// StatusCheckError indicates an error checking the command
type StatusCheckError struct {
	RawStatus
}

// Category is the category of the status, to indicate which implementations should be checked against
func (s StatusCheckError) Category() StatusCategory {
	return StatusCategoryCheckError
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
