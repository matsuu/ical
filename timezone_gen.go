package ical

// THIS FILE IS AUTO-GENERATED BY internal/cmd/gentypes/gentypes.go
// DO NOT EDIT. ALL CHANGES WILL BE LOST

import (
	"bytes"
	"strings"

	"github.com/pkg/errors"
)

type Timezone struct {
	entries EntryList
	props   *PropertySet
}

func NewTimezone() *Timezone {
	return &Timezone{
		props: NewPropertySet(),
	}
}

func (v *Timezone) String() string {
	var buf bytes.Buffer
	NewEncoder(&buf).Encode(v)
	return buf.String()
}

func (v Timezone) Type() string {
	return "VTIMEZONE"
}

func (v *Timezone) Entries() <-chan Entry {
	return v.entries.Iterator()
}

func (v *Timezone) GetProperty(name string) (*Property, bool) {
	return v.props.GetFirst(name)
}

func (v *Timezone) Properties() <-chan *Property {
	return v.props.Iterator()
}

func (v *Timezone) AddProperty(key, value string, options ...PropertyOption) error {
	var params Parameters
	var force bool
	for _, option := range options {
		switch option.Name() {
		case "Parameters":
			params = option.Get().(Parameters)
		case "Force":
			force = option.Get().(bool)
		}
	}

	switch key = strings.ToLower(key); key {
	case "tzid", "last-modified", "tzurl":
		v.props.Set(NewProperty(key, value, params))
	default:
		if strings.HasPrefix(key, "x-") || force {
			v.props.Append(NewProperty(key, value, params))
		} else {
			return errors.Errorf(`invalid property %s`, key)
		} /* end if */
	}
	return nil
}
