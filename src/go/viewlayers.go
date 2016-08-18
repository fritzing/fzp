package fzp

import (
	"errors"
)

// ViewLayers represet a fzp ViewLayers data object
type ViewLayers struct {
	Image string      `xml:"image,attr" json:"image"`
	Layer []ViewLayer `xml:"layer"      json:"layer"`
}

// NewViewLayers returns a ViewLayers object
func NewViewLayers() ViewLayers {
	v := ViewLayers{}
	v.Layer = make([]ViewLayer, 0)
	return v
}

// Check validate the ViewLayers data
func (v *ViewLayers) Check() error {
	// TODO: create check rules

	if v.Image == "" {
		return errors.New("missing layer image")
	}

	return nil
}
