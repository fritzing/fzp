package fzp

import (
	"errors"
	"strconv"
)

func (f *FZP) CheckFritzingVersion() error {
	if f.FritzingVersion == "" {
		return errors.New("fritzingVersion undefined")
	}
	return nil
}

func (f *FZP) CheckModuleId() error {
	if f.ModuleId == "" {
		return errors.New("moduleId undefined")
	}
	return nil
}

// TODO: is the referenceFile required?

func (f *FZP) CheckVersion() error {
	if f.Version == "" {
		return errors.New("version undefined")
	}
	return nil
}

func (f *FZP) CheckTitle() error {
	if f.Title == "" {
		return errors.New("title undefined")
	}
	return nil
}

func (f *FZP) CheckDescription() error {
	if f.Description == "" {
		return errors.New("description undefined")
	}
	return nil
}

func (f *FZP) CheckAuthor() error {
	if f.Author == "" {
		return errors.New("author undefined")
	}
	return nil
}

// Check Date ?
// Check URL ?
// Check Label ?
// Check Taxonomy ?
// Check Family ?
// Check Variant ?

func (f *FZP) CheckTags() (error, int) {
	countBrokenTags := 0

	if len(f.Tags) != 0 {
		for _, tag := range f.Tags {
			if tag == "" {
				countBrokenTags++
			}
		}
	}

	if countBrokenTags == 0 {
		return nil, countBrokenTags
	} else {
		errMsg := strconv.Itoa(countBrokenTags) + " tag value/s undefined"
		return errors.New(errMsg), countBrokenTags
	}
}

func (f *FZP) CheckProperties() error {
	if len(f.Properties) != 0 {
		for _, property := range f.Properties {
			if err := property.CheckName(); err != nil {
				return err
			}
			if err := property.CheckValue(); err != nil {
				return err
			}
		}
	}
	return nil
}

func (f *FZP) CheckViews() error {
	// TODO: ...
	return nil
}

func (f *FZP) CheckConnectors() error {
	if len(f.Connectors) != 0 {
		for _, connector := range f.Connectors {
			if err := connector.Check(); err != nil {
				return err
			}
		}
	}
	return nil
}

func (f *FZP) CheckBuses() error {
	if len(f.Buses) != 0 {
		for _, bus := range f.Buses {
			if err := bus.CheckId(); err != nil {
				return err
			}
		}
	}
	return nil
}

// check all the data
func (f *FZP) Check() []error {
	var errList []error

	if err := f.CheckFritzingVersion(); err != nil {
		errList = append(errList, err)
	}
	if err := f.CheckModuleId(); err != nil {
		errList = append(errList, err)
	}
	if err := f.CheckVersion(); err != nil {
		errList = append(errList, err)
	}
	if err := f.CheckTitle(); err != nil {
		errList = append(errList, err)
	}
	errTags, _ := f.CheckTags()
	if errTags != nil {
		errList = append(errList, errTags)
	}
	if err := f.CheckProperties(); err != nil {
		errList = append(errList, err)
	}
	if err := f.CheckViews(); err != nil {
		errList = append(errList, err)
	}
	if err := f.CheckConnectors(); err != nil {
		errList = append(errList, err)
	}
	if err := f.CheckBuses(); err != nil {
		errList = append(errList, err)
	}

	return errList
}
