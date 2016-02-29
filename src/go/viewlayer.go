package fzp

type ViewLayer struct {
	LayerId string `xml:"layerId,attr"`
}

func NewViewLayer() ViewLayer {
	v := ViewLayer{}
	return v
}
