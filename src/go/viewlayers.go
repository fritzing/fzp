package fzp

type ViewLayers struct {
	Image string      `xml:"image,attr"`
	Layer []ViewLayer `xml:"layer"`
}

func NewViewLayers() ViewLayers {
	v := ViewLayers{}
	return v
}
