package fzp

import (
	"errors"
	"strconv"
)

// CheckFritzingVersion validate the FritzingVersion data
func (f *Fzp) CheckFritzingVersion() error {
	if f.FritzingVersion == "" {
		return errors.New("fritzingVersion undefined")
	}
	return nil
}

// CheckModuleID validate the ModuleID data
func (f *Fzp) CheckModuleID() error {
	if f.ModuleID == "" {
		return errors.New("moduleId undefined")
	}
	return nil
}

// TODO: is the referenceFile required?

// CheckVersion validate the Version data
func (f *Fzp) CheckVersion() error {
	if f.Version == "" {
		return errors.New("version undefined")
	}
	return nil
}

// CheckTitle validate the Title data
func (f *Fzp) CheckTitle() error {
	if f.Title == "" {
		return errors.New("title undefined")
	}
	return nil
}

// CheckDescription validate the Description data
func (f *Fzp) CheckDescription() error {
	if f.Description == "" {
		return errors.New("description undefined")
	}
	return nil
}

// CheckAuthor validate the Author data
func (f *Fzp) CheckAuthor() error {
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

// CheckTags validate the Tags data
func (f *Fzp) CheckTags() (error, int) {
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
	}
	errMsg := strconv.Itoa(countBrokenTags) + " tag value/s undefined"
	return errors.New(errMsg), countBrokenTags
}

// CheckProperties validate the Properties data
func (f *Fzp) CheckProperties() error {
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

// CheckViews validate the Views data
func (f *Fzp) CheckViews() error {
	// TODO: ...
	return nil
}

// CheckConnectors validate the Connectors data
func (f *Fzp) CheckConnectors() error {
	if len(f.Connectors) != 0 {
		for _, connector := range f.Connectors {
			if err := connector.Check(); err != nil {
				return err
			}
		}
	}
	return nil
}

// CheckBuses validate the Buses data
func (f *Fzp) CheckBuses() error {
	if len(f.Buses) != 0 {
		for _, bus := range f.Buses {
			if err := bus.CheckID(); err != nil {
				return err
			}
		}
	}
	return nil
}

// Check the whole Fzp data
func (f *Fzp) Check() []error {
	var errList []error

	if err := f.CheckFritzingVersion(); err != nil {
		errList = append(errList, err)
	}
	if err := f.CheckModuleID(); err != nil {
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
