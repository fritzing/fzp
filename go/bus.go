package fzp

type Bus struct {
	Id         string          `xml:"id,attr"`
	NodeMember []BusNodeMember `xml:"nodeMember"`
}
