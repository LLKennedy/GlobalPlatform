package gpapdu

import (
	"fmt"

	"github.com/llkennedy/globalplatform/goimpl/apdu"
)

// SendOnTransport sends the GP command over the APDU transport, handling conversion errors
func SendOnTransport(t apdu.Transport, cmd Command) (apdu.Response, error) {
	if t == nil {
		return apdu.Response{}, fmt.Errorf("cannot send command on nil transport")
	}
	toSend, err := cmd.ToAPDU()
	if err != nil {
		return apdu.Response{}, fmt.Errorf("failed to convert command to APDU: %w", err)
	}
	return t.Send(toSend)
}
