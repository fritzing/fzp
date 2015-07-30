package main

import (
	"encoding/xml"
)

type Fzp struct {
	XMLName         xml.Name    `xml:"module"`
	FritzingVersion string      `xml:"fritzingVersion,attr"`
	ModuleId        string      `xml:"moduleId,attr"`
	Version         string      `xml:"version"`
	Title           string      `xml:"title"`
	Description     string      `xml:"description"`
	Author          string      `xml:"author"`
	Date            string      `xml:"date"`
	URL             string      `xml:"url"`
	Label           string      `xml:"label"`
	Taxonomy        string      `xml:"taxonomy"`
	Family          string      `xml:"family"`
	Variant         string      `xml:"variant"`
	Tags            []string    `xml:"tags>tag"`
	Properties      []Property  `xml:"properties>property"`
	Views           Views       `xml:"views"`
	Connectors      []Connector `xml:"connectors>connector"`
	Buses           []Bus       `xml:"buses>bus"`
}

type Property struct {
	Name  string `xml:"name,attr"`
	Value string `xml:",chardata"`
}

type Views struct {
	Icon       ViewLayers `xml:"iconView>layers"`
	Breadboard ViewLayers `xml:"breadboardView>layers"`
	Pcb        ViewLayers `xml:"pcbView>layers"`
	Schematic  ViewLayers `xml:"schematicView>layers"`
}

type ViewLayers struct {
	Image string      `xml:"image,attr"`
	Layer []ViewLayer `xml:"layer"`
}

type ViewLayer struct {
	LayerId string `xml:"layerId,attr"`
}

type Connector struct {
	Id             string           `xml:"id,attr"`
	Name           string           `xml:"name,attr"`
	Type           string           `xml:"type,attr"`
	Description    string           `xml:"description"`
	BreadboardView []ConnectorLayer `xml:"views>breadboardView>p"`
	PcbView        []ConnectorLayer `xml:"views>pcbView>p"`
	SchematicView  []ConnectorLayer `xml:"views>schematicView>p"`
}

type ConnectorLayer struct {
	Layer      string `xml:"layer,attr"`
	SvgId      string `xml:"svgId,attr"`
	TerminalId string `xml:"terminalId,attr"`
}

type Bus struct {
	Id         string          `xml:"id,attr"`
	NodeMember []BusNodeMember `xml:"nodeMember"`
}

type BusNodeMember struct {
	ConnectorId string `xml:"connectorId,attr"`
}
