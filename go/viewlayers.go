package fzp

type ViewLayers struct {
	Image string      `xml:"image,attr"`
	Layer []ViewLayer `xml:"layer"`
}
