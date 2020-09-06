package gpapi

import (
	"fmt"

	"github.com/llkennedy/globalplatform/goimpl/gpapdu"
)

// Card represents a physical card in a reader
type Card struct {
	ctx *gpapdu.Context
}

// StartSCP03 starts an SCP03 session
func (c *Card) StartSCP03() error {
	sess, err := c.ctx.InitializeUpdate(0, 0, nil)
	if err != nil {
		return err
	}
	_ = sess // FIXME: this function is nowhere near ready, it just stands to show the basics of what this package is for before I develop it
	return fmt.Errorf("not implemented")
}
