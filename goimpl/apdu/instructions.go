package apdu

// Instruction indicates what command is being sent to the card
// NOTE: This type is intended to assist, it is NOT a comprehensive enum of possible commands
// Extension protocols and specifications (such as G.P.) are expected to add their own Instruction constants
type Instruction byte

const (
	// InstructionActivateFile is the ActivateFile instruction byte
	InstructionActivateFile Instruction = 0x44
	// InstructionAppendRecord is the AppendRecord instruction byte
	InstructionAppendRecord Instruction = 0xE2
	// InstructionChangeReferenceData is the ChangeReferenceData instruction byte
	InstructionChangeReferenceData Instruction = 0x24
	// InstructionCreateFile is the CreateFile instruction byte
	InstructionCreateFile Instruction = 0xE0
	// InstructionDeactivateFile is the DeactivateFile instruction byte
	InstructionDeactivateFile Instruction = 0x04
	// InstructionDeleteFile is the DeleteFile instruction byte
	InstructionDeleteFile Instruction = 0xE4
	// InstructionDisableVerificationRequirement is the DisableVerificationRequirement instruction byte
	InstructionDisableVerificationRequirement Instruction = 0x26
	// InstructionEnableVerificationRequirement is the EnableVerificationRequirement instruction byte
	InstructionEnableVerificationRequirement Instruction = 0x28
	// InstructionEnvelope is the Envelope instruction byte
	InstructionEnvelope Instruction = 0xC2
	// InstructionEraseBinary is the EraseBinary instruction byte
	InstructionEraseBinary Instruction = 0x0E
	// InstructionEraseRecordS is the EraseRecordS instruction byte
	InstructionEraseRecordS Instruction = 0x0C
	// InstructionExternalMutualAuthenticate is the ExternalMutualAuthenticate instruction byte
	InstructionExternalMutualAuthenticate Instruction = 0x82
	// InstructionGeneralAuthenticate is the GeneralAuthenticate instruction byte
	InstructionGeneralAuthenticate Instruction = 0x86
	// InstructionGenerateAsymmetricKeyPair is the GenerateAsymmetricKeyPair instruction byte
	InstructionGenerateAsymmetricKeyPair Instruction = 0x46
	// InstructionGetChallenge is the GetChallenge instruction byte
	InstructionGetChallenge Instruction = 0x84
	// InstructionGetData is the GetData instruction byte
	InstructionGetData Instruction = 0xCA
	// InstructionGetResponse is the GetResponse instruction byte
	InstructionGetResponse Instruction = 0xC0
	// InstructionInternalAuthenticate is the InternalAuthenticate instruction byte
	InstructionInternalAuthenticate Instruction = 0x88
	// InstructionManageChannel is the ManageChannel instruction byte
	InstructionManageChannel Instruction = 0x70
	// InstructionManageSecurityEnvironment is the ManageSecurityEnvironment instruction byte
	InstructionManageSecurityEnvironment Instruction = 0x22
	// InstructionPerformScqlOperation is the PerformScqlOperation instruction byte
	InstructionPerformScqlOperation Instruction = 0x10
	// InstructionPerformSecurityOperation is the PerformSecurityOperation instruction byte
	InstructionPerformSecurityOperation Instruction = 0x2A
	// InstructionPerformTransactionOperation is the PerformTransactionOperation instruction byte
	InstructionPerformTransactionOperation Instruction = 0x12
	// InstructionPerformUserOperation is the PerformUserOperation instruction byte
	InstructionPerformUserOperation Instruction = 0x14
	// InstructionPutData is the PutData instruction byte
	InstructionPutData Instruction = 0xDA
	// InstructionReadBinary is the ReadBinary instruction byte
	InstructionReadBinary Instruction = 0xB0
	// InstructionReadRecordS is the ReadRecordS instruction byte
	InstructionReadRecordS Instruction = 0xB2
	// InstructionResetRetryCounter is the ResetRetryCounter instruction byte
	InstructionResetRetryCounter Instruction = 0x2C
	// InstructionSearchBinary is the SearchBinary instruction byte
	InstructionSearchBinary Instruction = 0xA0
	// InstructionSearchRecord is the SearchRecord instruction byte
	InstructionSearchRecord Instruction = 0xA2
	// InstructionSelect is the Select instruction byte
	InstructionSelect Instruction = 0xA4
	// InstructionTerminateCardUsage is the TerminateCardUsage instruction byte
	InstructionTerminateCardUsage Instruction = 0xFE
	// InstructionTerminateDf is the TerminateDf instruction byte
	InstructionTerminateDf Instruction = 0xE6
	// InstructionTerminateEf is the TerminateEf instruction byte
	InstructionTerminateEf Instruction = 0xE8
	// InstructionUpdateBinary is the UpdateBinary instruction byte
	InstructionUpdateBinary Instruction = 0xD6
	// InstructionUpdateRecord is the UpdateRecord instruction byte
	InstructionUpdateRecord Instruction = 0xDC
	// InstructionVerify is the Verify instruction byte
	InstructionVerify Instruction = 0x20
	// InstructionWriteBinary is the WriteBinary instruction byte
	InstructionWriteBinary Instruction = 0xD0
	// InstructionWriteRecord is the WriteRecord instruction byte
	InstructionWriteRecord Instruction = 0xD2
)

// InstructionSetBERTLV sets b1 to indicate BER-TLV encoding for the instruction, where allowed
func InstructionSetBERTLV(ins byte) byte {
	return ins | b1
}
