package apdu

import "fmt"

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
	// Error will convert the status to an error type, if it represents an error state
	Error() error
	// WarningOrError will convert the status to an error type, if it represents an error or warning state
	WarningOrError() error
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
	// TODO: this function and its sub-functions are a mess of constants that could maybe be pulled out into a const block, but I'm not sure that'd be cleaner/easier to read or fix
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

// Error returns nil, because normal operation is not an error
func (s StatusNormal) Error() error {
	return nil
}

// WarningOrError returns nil, because normal operation is not a warning or error
func (s StatusNormal) WarningOrError() error {
	return nil
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
	RawStatus
	NonVolatileMemoryChanged                 bool
	NoInformationGiven                       bool
	WaitingQueryBytesLength                  uint8
	PossiblePartialReturnedDataCorruption    bool
	EOFBeforeReachedExpectedReturnDataLength bool
	FileDeactivated                          bool
	FileControlInformationMalformed          bool
	FileInTerminationState                   bool
	FileFilledByLastWrite                    bool
	NoInputDataAvailableFromSensor           bool
	ReturnedCounter                          bool
	Counter                                  uint8 // 0 to 15 are the only valid values, from second nibble of SW2
}

// Category is the category of the status, to indicate which implementations should be checked against
func (s StatusWarning) Category() StatusCategory {
	return StatusCategoryWarning
}

// Error returns nil, because a warning is not an error
func (s StatusWarning) Error() error {
	return nil
}

// WarningOrError returns the raw struct, because I can't be bothered writing custom human-readable errors for all possible states
func (s StatusWarning) WarningOrError() error {
	return fmt.Errorf("%#v", s)
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
			s.PossiblePartialReturnedDataCorruption = true
		case r.SW2 == 0x82:
			s.EOFBeforeReachedExpectedReturnDataLength = true
		case r.SW2 == 0x83:
			s.FileDeactivated = true
		case r.SW2 == 0x84:
			s.FileControlInformationMalformed = true
		case r.SW2 == 0x85:
			s.FileInTerminationState = true
		case r.SW2 == 0x86:
			s.NoInputDataAvailableFromSensor = true
		default:
			return StatusInvalid{r}
		}
		out = s
	case 0x63:
		s := StatusWarning{
			RawStatus:                r,
			NonVolatileMemoryChanged: true,
		}
		switch {
		case r.SW2 == 0x00:
			s.NoInformationGiven = true
		case r.SW2 == 0x01:
			s.FileFilledByLastWrite = true
		case (r.SW2 & 0xF0) == 0xC0:
			s.ReturnedCounter = true
			s.Counter = r.SW2 & 0x0F
		default:
			return StatusInvalid{r}
		}
		out = s
	}
	return out
}

// StatusExecError indicates an error executing the command
type StatusExecError struct {
	RawStatus
	NonVolatileMemoryChanged  bool
	WaitingQueryBytesLength   uint8
	NoInformationGiven        bool
	ExecutionError            bool
	ImmediateResponseRequired bool
	MemoryFailure             bool
	SecurityIssues            bool
	SecurityIssuesExtraData   byte // They didn't define 0x66XX at all, beyond "security-related issues"
}

// Category is the category of the status, to indicate which implementations should be checked against
func (s StatusExecError) Category() StatusCategory {
	return StatusCategoryExecError
}

// Error returns the raw struct, because I can't be bothered writing custom human-readable errors for all possible states
func (s StatusExecError) Error() error {
	return fmt.Errorf("%#v", s)
}

// WarningOrError returns the raw struct, because I can't be bothered writing custom human-readable errors for all possible states
func (s StatusExecError) WarningOrError() error {
	return fmt.Errorf("%#v", s)
}

func identifyExecError(r RawStatus) Status {
	var out Status = StatusInvalid{r}
	switch r.SW1 {
	case 0x64:
		s := StatusExecError{
			RawStatus:                r,
			NonVolatileMemoryChanged: false,
		}
		switch {
		case r.SW2 == 0x00:
			s.ExecutionError = true
		case r.SW2 >= 0x02 && r.SW2 <= 0x80:
			s.WaitingQueryBytesLength = r.SW2
		case r.SW2 == 0x01:
			s.ImmediateResponseRequired = true
		default:
			return StatusInvalid{r}
		}
		out = s
	case 0x65:
		s := StatusExecError{
			RawStatus:                r,
			NonVolatileMemoryChanged: true,
		}
		switch {
		case r.SW2 == 0x00:
			s.NoInformationGiven = true
		case r.SW2 == 0x81:
			s.MemoryFailure = true
		default:
			return StatusInvalid{r}
		}
		out = s
	case 0x66:
		out = StatusExecError{
			RawStatus:                r,
			NonVolatileMemoryChanged: false,
			SecurityIssues:           true,
			SecurityIssuesExtraData:  r.SW2,
		}
	}
	return out
}

// StatusCheckError indicates an error checking the command
type StatusCheckError struct {
	RawStatus
	WrongLengthNoFurtherIndication   bool
	ClassFunctionsNotSupported       bool
	ClassFunctionsNotSupportedDetail struct {
		NoInformationGiven          bool
		LogicalChannelNotSupported  bool
		SecureMessagingNotSupported bool
		LastCommandOfChainExpected  bool
		CommandChainingNotSupported bool
	}
	CommandNotAllowed       bool
	CommandNotAllowedDetail struct {
		NoInformationGiven                        bool
		CommandIncompatibleWithFileStructure      bool
		SecurityStatusNotSatisfied                bool
		AuthenticationModeBlocked                 bool
		ReferenceDataNotUsable                    bool
		ConditionsOfUseNotSatisfied               bool
		CommandNotAllowedNoCurrentEF              bool
		ExpectedSecureMessagingDataObjectsMissing bool
		IncorrectSecureMessagingObjects           bool
	}
	WrongParametersNoFurtherIndication bool
	WrongParametersWithDetail          bool
	WrongParametersDetail              struct {
		NoInformationGiven                    bool
		IncorrectParametersInCommandData      bool
		FunctionNotSupported                  bool
		FileOrApplicationNotFound             bool
		RecordNotFound                        bool
		NotEnoughMemorySpaceInFile            bool
		NcInconsistentWithTLV                 bool
		IncorrectParametersP1P2               bool
		NcInconsistentWithParametersP1P2      bool
		ReferencedDataOrReferenceDataNotFound bool
		FileAlreadyExists                     bool
		DFNameAlreadyExists                   bool
	}
	WrongLeField                         bool
	WrongLeFieldAvailableBytes           uint8
	InstructionCodeNotSupportedOrInvalid bool
	ClassNotSupported                    bool
	NoPreciseDiagnosis                   bool
}

// Category is the category of the status, to indicate which implementations should be checked against
func (s StatusCheckError) Category() StatusCategory {
	return StatusCategoryCheckError
}

// Error returns the raw struct, because I can't be bothered writing custom human-readable errors for all possible states
func (s StatusCheckError) Error() error {
	return fmt.Errorf("%#v", s)
}

// WarningOrError returns the raw struct, because I can't be bothered writing custom human-readable errors for all possible states
func (s StatusCheckError) WarningOrError() error {
	return fmt.Errorf("%#v", s)
}

func identifyCheckError(r RawStatus) Status {
	var out Status
	combined := (uint16(r.SW1) << 8) | uint16(r.SW2)
	switch combined {
	case 0x6700:
		out = StatusCheckError{
			RawStatus:                      r,
			WrongLengthNoFurtherIndication: true,
		}
	case 0x6B00:
		out = StatusCheckError{
			RawStatus:                          r,
			WrongParametersNoFurtherIndication: true,
		}
	case 0x6D00:
		out = StatusCheckError{
			RawStatus:                            r,
			InstructionCodeNotSupportedOrInvalid: true,
		}
	case 0x6E00:
		out = StatusCheckError{
			RawStatus:         r,
			ClassNotSupported: true,
		}
	case 0x6F00:
		out = StatusCheckError{
			RawStatus:          r,
			NoPreciseDiagnosis: true,
		}
	}
	if out != nil {
		return out
	}
	out = StatusInvalid{r}
	switch r.SW1 {
	case 0x68:
		s := StatusCheckError{
			RawStatus:                  r,
			ClassFunctionsNotSupported: true,
		}
		detail := s.ClassFunctionsNotSupportedDetail
		switch r.SW2 {
		case 0x00:
			detail.NoInformationGiven = true
		case 0x81:
			detail.LogicalChannelNotSupported = true
		case 0x82:
			detail.SecureMessagingNotSupported = true
		case 0x83:
			detail.LastCommandOfChainExpected = true
		case 0x84:
			detail.CommandChainingNotSupported = true
		default:
			return StatusInvalid{r}
		}
		s.ClassFunctionsNotSupportedDetail = detail
		out = s
	case 0x69:
		s := StatusCheckError{
			RawStatus:         r,
			CommandNotAllowed: true,
		}
		detail := s.CommandNotAllowedDetail
		switch r.SW2 {
		case 0x00:
			detail.NoInformationGiven = true
		case 0x81:
			detail.CommandIncompatibleWithFileStructure = true
		case 0x82:
			detail.SecurityStatusNotSatisfied = true
		case 0x83:
			detail.AuthenticationModeBlocked = true
		case 0x84:
			detail.ReferenceDataNotUsable = true
		case 0x85:
			detail.ConditionsOfUseNotSatisfied = true
		case 0x86:
			detail.CommandNotAllowedNoCurrentEF = true
		case 0x87:
			detail.ExpectedSecureMessagingDataObjectsMissing = true
		case 0x88:
			detail.IncorrectSecureMessagingObjects = true
		default:
			return StatusInvalid{r}
		}
		s.CommandNotAllowedDetail = detail
		out = s
	case 0x6A:
		s := StatusCheckError{
			RawStatus:                 r,
			WrongParametersWithDetail: true,
		}
		detail := s.WrongParametersDetail
		switch r.SW2 {
		case 0x00:
			detail.NoInformationGiven = true
		case 0x80:
			detail.IncorrectParametersInCommandData = true
		case 0x81:
			detail.FunctionNotSupported = true
		case 0x82:
			detail.FileOrApplicationNotFound = true
		case 0x83:
			detail.RecordNotFound = true
		case 0x84:
			detail.NotEnoughMemorySpaceInFile = true
		case 0x85:
			detail.NcInconsistentWithTLV = true
		case 0x86:
			detail.IncorrectParametersP1P2 = true
		case 0x87:
			detail.NcInconsistentWithParametersP1P2 = true
		case 0x88:
			detail.ReferencedDataOrReferenceDataNotFound = true
		case 0x89:
			detail.FileAlreadyExists = true
		case 0x8A:
			detail.DFNameAlreadyExists = true
		default:
			return StatusInvalid{r}
		}
		s.WrongParametersDetail = detail
		out = s
	case 0x6C:
		out = StatusCheckError{
			RawStatus:                  r,
			WrongLeField:               true,
			WrongLeFieldAvailableBytes: r.SW2,
		}
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

// Error always returns a simple error, because proprietary status returns should override this via composition
func (s StatusProprietary) Error() error {
	return fmt.Errorf("Unknown proprietary status with bytes %X%X", s.SW1, s.SW2)
}

// WarningOrError always returns a simple error, because proprietary status returns should override this via composition
func (s StatusProprietary) WarningOrError() error {
	return fmt.Errorf("Unknown proprietary status with bytes %X%X", s.SW1, s.SW2)
}

// StatusInvalid is an invalid status
type StatusInvalid struct {
	RawStatus
}

// Category is the category of the status, to indicate which implementations should be checked against
func (s StatusInvalid) Category() StatusCategory {
	return StatusCategoryInvalid
}

// Error always returns a simple error, because this status is invalid
func (s StatusInvalid) Error() error {
	return fmt.Errorf("Invalid status with bytes %X%X", s.SW1, s.SW2)
}

// WarningOrError always returns a simple error, because this status is invalid
func (s StatusInvalid) WarningOrError() error {
	return fmt.Errorf("Invalid status with bytes %X%X", s.SW1, s.SW2)
}
