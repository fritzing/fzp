package fzp

// Views represet a fzp Views data object
type Views struct {
	Icon       ViewLayers `xml:"iconView>layers"       json:"icon"`
	Breadboard ViewLayers `xml:"breadboardView>layers" json:"breadboard"`
	Pcb        ViewLayers `xml:"pcbView>layers"        json:"pcb"`
	Schematic  ViewLayers `xml:"schematicView>layers"  json:"schematic"`
}

// NewViews returns a Views object
func NewViews() Views {
	v := Views{}
	return v
}

// Check validate the Views data
func (v *Views) Check() error {
	// TODO: create check rules
	return nil
}
