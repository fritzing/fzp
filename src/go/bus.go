package fzp

import (
	"errors"
	"strings"
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
	b.NodeMember = make([]BusNode, 0)
	return b
}

// Check validate the Bus data
func (b *Bus) Check() error {
	errMsg := ""
	if err := b.CheckID(); err != nil {
		errMsg = err.Error()
	}
	if err := b.CheckNodeMembers(); err != nil {
		errMsg += ", " + err.Error()
		return errors.New(errMsg)
	}
	return nil
}

// CheckID validate the Bus ID data
func (b *Bus) CheckID() error {
	if b.ID == "" {
		return errors.New("bus id undefined")
	}
	return nil
}

// CheckNodeMembers validate the Bus NodeMember data
func (b *Bus) CheckNodeMembers() error {
	if len(b.NodeMember) == 0 {
		return errors.New("bus nodemembers undefined")
	}
	var errConcat []string
	for _, n := range b.NodeMember {
		err := n.Check()
		if err != nil {
			errConcat = append(errConcat, err.Error())
		}
	}
	if len(errConcat) != 0 {
		return errors.New(strings.Join(errConcat, "\n"))
	}
	return nil
}

// AddNodeMember adds a BusNode to the NodeMember array
func (b *Bus) AddNodeMember(id string) {
	nMember := BusNode{}
	nMember.ConnectorID = id
	b.NodeMember = append(b.NodeMember, nMember)
}
