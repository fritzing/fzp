package fzp

// ViewLayer represet a fzp ViewLayer data object
type ViewLayer struct {
	LayerID string `xml:"layerId,attr" json:"layerId"`

	// BreadboardLayer    string `xml:"layerId,attr" json:"breadboard"`
	// SchematicLayer     string `xml:"layerId,attr" json:"schematic"`
	// PcbLayerCopper0    string `xml:"layerId,attr" json:"copper0"`
	// PcbLayerCopper1    string `xml:"layerId,attr" json:"copper1"`
	// PcbLayerSilkscreen string `xml:"layerId,attr" json:"silkscreen"`
	// IconLayer          string `xml:"layerId,attr" json:"icon"`
}

// NewViewLayer returns a ViewLayer object
func NewViewLayer(id string) ViewLayer {
	v := ViewLayer{}
	v.LayerID = id
	return v
}

// Check validate the ViewLayer data
func (v *ViewLayer) Check() error {
	// TODO: create check rules
	return nil
}
