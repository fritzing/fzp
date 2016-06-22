package fzp

type Tags []string

func NewTags() *Tags {
	t := Tags{}
	t = make([]string, 0)
	return &t
}
