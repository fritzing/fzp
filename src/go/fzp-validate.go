package fzp

import (
	"errors"
	"fmt"
	"strconv"
)

// Check the whole Fzp data
func (f *Fzp) Check() []error {
	var errList []error

	if err := f.CheckFritzingVersion(); err != nil {
		errList = append(errList, err)
	}
	if err := f.CheckModuleID(); err != nil {
		errList = append(errList, err)
	}
	tmp := f.CheckVersion()
	if tmp != "" {
		fmt.Println(tmp)
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

// CheckFritzingVersion validate the FritzingVersion data
// type error
func (f *Fzp) CheckFritzingVersion() error {
	if f.FritzingVersion == "" {
		return errors.New("fritzingVersion undefined")
	}
	return nil
}

// CheckModuleID validate the ModuleID data
// type error
func (f *Fzp) CheckModuleID() error {
	if f.ModuleID == "" {
		return errors.New("moduleId undefined")
	}
	return nil
}

// TODO: is the referenceFile required?

// CheckVersion validate the Version data
// Version is not a part critical property.
// type warning
func (f *Fzp) CheckVersion() string {
	if f.Version == "" {
		return "==> WARN version undefined"
	}
	return ""
}

// CheckTitle validate the Title data
// type error
func (f *Fzp) CheckTitle() error {
	if f.Title == "" {
		return errors.New("title undefined")
	}
	return nil
}

// CheckDescription validate the Description data.
// Description is not a part critical property.
// type warning
func (f *Fzp) CheckDescription() string {
	if f.Description == "" {
		return "==> WARN description undefined" + f.Title

	}
	return ""
}

// CheckAuthor validate the Author data.
// Author is not a part critical property.
// type warning
func (f *Fzp) CheckAuthor() string {
	if f.Author == "" {
		return "==> WARN author undefined"

	}
	return ""
}

// Check Date ?
// Check URL ?
// Check Label ?
// Check Taxonomy ?
// Check Family ?
// Check Variant ?

// CheckTags validate the Tags data
// if no tag exist, the part cannot be found by the fritzing app.
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
	return f.Properties.Check()

	// if len(f.Properties) == 0 {
	// 	return errors.New("Missing property family!")
	// }
	//
	// for _, property := range f.Properties {
	// 	if err := property.Check(); err != nil {
	// 		return err
	// 	}
	// }
	//
	// return err
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
