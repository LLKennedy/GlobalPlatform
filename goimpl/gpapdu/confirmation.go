package gpapdu

import (
	"fmt"

	"github.com/llkennedy/globalplatform/goimpl/apdu"
)

// ResponseConfirmation is a Confirmation from a response message
type ResponseConfirmation struct {
	Receipt             []byte
	ConfirmationCounter uint16
	SDUniqueData        []byte
	TokenIdentifier     []byte
	TokenDataDigest     []byte
}

// ConfirmationFromResponse converts an apdu.Response to a Confirmation message
func ConfirmationFromResponse(in apdu.Response) (res ResponseConfirmation, err error) {
	dataLen := len(in.Data)
	if dataLen < 6 {
		err = fmt.Errorf("cannot extract confirmation from too short response data, must be at least 6 bytes")
		return
	}
	receiptLen := int(in.Data[0])
	receiptStart := 1
	// What's the point of this? We still only encode 0-255, but we arbitrarily encode 128-255 on a different length than 0-127
	if receiptLen&b8 > 0 {
		receiptLen = int(in.Data[1])
		receiptStart = 2
	}
	if receiptLen > 0 {
		if dataLen < (receiptStart + receiptLen + 5) {
			err = fmt.Errorf("cannot extract confirmation receipt of length %d from response data of length %d", receiptLen, dataLen)
			return
		}
		res.Receipt = make([]byte, receiptLen)
		// Copy receipt data
		copy(res.Receipt, in.Data[receiptStart:receiptStart+receiptLen])
	}
	confirmationStart := receiptStart + receiptLen
	confirmationCounterLen := int(in.Data[confirmationStart])
	_ = confirmationCounterLen // FIXME: finish implementing this
	err = fmt.Errorf("not implemented")
	return
}
