package fzp

type Views struct {
	Icon       ViewLayers `xml:"iconView>layers"`
	Breadboard ViewLayers `xml:"breadboardView>layers"`
	Pcb        ViewLayers `xml:"pcbView>layers"`
	Schematic  ViewLayers `xml:"schematicView>layers"`
}
