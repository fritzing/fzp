package fzp

// ViewLayer represet a fzp ViewLayer data object
type ViewLayer struct {
	LayerID string `xml:"layerId,attr" json:"layerId"`
}

// NewViewLayer returns a ViewLayer object
func NewViewLayer() ViewLayer {
	v := ViewLayer{}
	return v
}

// Check validate the ViewLayer data
func (v *ViewLayer) Check() error {
	// TODO: create check rules
	return nil
}
