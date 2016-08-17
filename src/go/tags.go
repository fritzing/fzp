package fzp

import (
	"errors"
	// "fmt"
	"sort"
)

type Tags []string

func NewTags() Tags {
	t := Tags{}
	t = make([]string, 0)
	return t
}

func (t *Tags) Total() int {
	return len(*t)
}

func (t *Tags) Exist(tag string) int {
	count := 0
	for _, v := range *t {
		if v == tag {
			count++
		}
	}
	return count
}

func (t *Tags) Add(tag string) error {
	if t.Exist(tag) != 0 {
		return errors.New("tag name '" + tag + "' already exist")
	}
	*t = append(*t, tag)
	return nil
}

func (t *Tags) Check() error {
	for _, v := range *t {
		// TODO: check if upper / lowercase version exist
		total := sort.SearchStrings(*t, v)
		// fmt.Println("total:", total)
		if total != 0 {
			// fmt.Println("DUPLICATE", v)

			// TODO: collect all errors and return after for loop was finished
			return errors.New("duplicated tag: " + v)
		}
	}
	return nil
}
