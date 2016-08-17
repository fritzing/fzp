package fzp

import (
	"errors"
	"strings"
)

type Layers struct {
  BreadboardLayer LayerIDs string `xml:"layerId,breadboard" json:"breadboard"`
  SchematicLayer  LayerIDs string `xml:"layerId,schematic" json:"schematic"`
  PcbLayerCopper0  LayerIDs string `xml:"layerId,copper0" json:"copper0"`
  PcbLayerCopper1  LayerIDs string `xml:"layerId,copper1" json:"copper1"`
  PcbLayerSilkscreen  LayerIDs string `xml:"layerId,silkscreen" json:"silkscreen"`
  IconLayer  LayerIDs string `xml:"layerId,icon" json:"icon"`
  }


  // NewLayers returns a Layers object
  func NewLayers() Layer {
  	v := Layers{}
  	v.BreadboardLayer = NewLayerIDs()
  	v.SchematicLayer = NewLayerIDs()
  	v.PcbLayerCopper0 = NewLayerIDs()
    v.PcbLayerCopper1 = NewLayerIDs()
  	v.PcbLayerSilkscreen = NewLayerIDs()
    v.IconLayer = NewLayerIDs()
  	return v
  }

  //TODO: Check if layerchecks work and ok?
  // Check validate the Layer data. this calls the iconlayer, BreadboardLayer, pcblayer and
  // SchematicLayer check funcs and concatendate the error messages
  func (v *Layers) Check() error {
  	var errConcat []string

  	if err := v.BreadboardLayer.Check(); err != nil {
  		errConcat = append(errConcat, err.Error())
  	}
  	if err := v.SchematicLayer.Check(); err != nil {
  		errConcat = append(errConcat, err.Error())
  	}
  	if err := v.PcbLayerCopper0.Check(); err != nil {
  		errConcat = append(errConcat, err.Error())
  	}
    if err := v.PcbLayerCopper1.Check(); err != nil {
  		errConcat = append(errConcat, err.Error())
  	}
    if err := v.PcbLayerSilkscreen.Check(); err != nil {
  		errConcat = append(errConcat, err.Error())
  	}
  	if err := v.IconLayer.Check(); err != nil {
  		errConcat = append(errConcat, err.Error())
  	}

  	if len(errConcat) != 0 {
  		return errors.New(strings.Join(errConcat, "\n"))
  	}
  	return nil
  }
