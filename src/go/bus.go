package fzp

import (
	"errors"
)

// Bus represet a fzp bus data object
type Bus struct {
	ID         string    `xml:"id,attr"    json:"id"`
	NodeMember []BusNode `xml:"nodeMember" json:"nodeMember"`
}

// NewBus returns a Bus object
func NewBus(id string) Bus {
	b := Bus{}
	b.ID = id
	return b
}

// Check validate the Bus data
func (b *Bus) Check() error {
	errMsg := ""
	err := b.CheckID()
	if err != nil {
		errMsg = err.Error()
	}
	err = b.CheckNodeMembers()
	if err != nil {
		errMsg += ", " + err.Error()
		return errors.New(errMsg)
	}
	return nil
}

// CheckID validate the Bus ID data
func (b *Bus) CheckID() error {
	if b.ID == "" {
		return errors.New("Bus id undefined")
	}
	return nil
}

// CheckNodeMembers validate the Bus NodeMember data
func (b *Bus) CheckNodeMembers() error {
	if len(b.NodeMember) == 0 {
		return errors.New("bus nodemembers undefined")
	}
	for _, n := range b.NodeMember {
		err := n.Check()
		if err != nil {
			return err
		}
	}
	return nil
}

// AddNodeMember adds a BusNode to the NodeMember array
func (b *Bus) AddNodeMember(id string) {
	nMember := BusNode{}
	nMember.ConnectorID = id
	b.NodeMember = append(b.NodeMember, nMember)
}
