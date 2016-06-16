package fzp

import (
	"errors"
)

type BusNode struct {
	ConnectorId string `xml:"connectorId,attr"`
}

func NewBusNode(cId string) BusNode {
	b := BusNode{}
	b.ConnectorId = cId
	return b
}

func (b *BusNode) Check() error {
	if b.ConnectorId == "" {
		return errors.New("BusNode ConnectorId undefined")
	}
	return nil
}
