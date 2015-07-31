package fzp

import "errors"

type Bus struct {
	Id         string          `xml:"id,attr"`
	NodeMember []BusNodeMember `xml:"nodeMember"`
}

func (b *Bus) CheckId() error {
	if b.Id == "" {
		return errors.New("bus-id undefined!")
	}
	return nil
}
