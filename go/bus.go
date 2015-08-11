package fzp

import (
	"errors"
)

type Bus struct {
	Id         string    `xml:"id,attr"`
	NodeMember []BusNode `xml:"nodeMember"`
}

func NewBus(id string) Bus {
	b := Bus{}
	b.Id = id
	return b
}

func (b *Bus) CheckId() error {
	if b.Id == "" {
		return errors.New("Bus id undefined")
	}
	return nil
}

func (b *Bus) CheckNodeMembers() error {
	if len(b.NodeMember) == 0 {
		return errors.New("bus nodemembers undefined")
	} else {
		for _, n := range b.NodeMember {
			err := n.Check()
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func (b *Bus) AddNodeMember(id string) {
	nMember := BusNode{}
	nMember.ConnectorId = id
	b.NodeMember = append(b.NodeMember, nMember)
}
