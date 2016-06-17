package fzp

import (
	"errors"
)

// BusNode represet a fzp BusNode data object
type BusNode struct {
	ConnectorID string `xml:"connectorId,attr" json:"connectorId"`
}

// NewBusNode returns a BusNode object
func NewBusNode(id string) BusNode {
	b := BusNode{}
	b.ConnectorID = id
	return b
}

// Check validate the BusNode data
func (b *BusNode) Check() error {
	if b.ConnectorID == "" {
		return errors.New("BusNode ConnectorID undefined")
	}
	return nil
}
