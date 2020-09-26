package gpapdu

import "github.com/llkennedy/globalplatform/goimpl/apdu"

// Client is a GP client wrapping APDU commands
type Client struct {
	transport apdu.Transport
}

// NewClient creates a new Client on the provided Transport
func NewClient(transport apdu.Transport) *Client {
	return &Client{transport}
}

func (c *Client) Delete(deleteRelatedObjects bool, cmd DeleteCommand) (confirmation []byte, err error) {
	switch gotCmd := cmd.(type) {
	case *DeleteCardContent:
		gotCmd.
	case *DeleteKey:
	default:
		return nil, fmt.Errorf("invalid command, must be DeleteKey or DeleteCommand")
	}
}