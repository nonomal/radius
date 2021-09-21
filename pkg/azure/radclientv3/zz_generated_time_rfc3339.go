// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package radclientv3



import (
	"regexp"
	"strings"
	"time"
)

const (
	utcLayoutJSON = `"2006-01-02T15:04:05.999999999"`
	utcLayout     = "2006-01-02T15:04:05.999999999"
	rfc3339JSON   = `"` + time.RFC3339Nano + `"`
)

// Azure reports time in UTC but it doesn't include the 'Z' time zone suffix in some cases.
var tzOffsetRegex = regexp.MustCompile(`(Z|z|\+|-)(\d+:\d+)*"*$`)

type timeRFC3339 time.Time

func (t timeRFC3339) MarshalJSON() (json []byte, err error) {
	tt := time.Time(t)
	return tt.MarshalJSON()
}

func (t timeRFC3339) MarshalText() (text []byte, err error) {
	tt := time.Time(t)
	return tt.MarshalText()
}

func (t *timeRFC3339) UnmarshalJSON(data []byte) error {
	layout := utcLayoutJSON
	if tzOffsetRegex.Match(data) {
		layout = rfc3339JSON
	}
	return t.Parse(layout, string(data))
}

func (t *timeRFC3339) UnmarshalText(data []byte) (err error) {
	layout := utcLayout
	if tzOffsetRegex.Match(data) {
		layout = time.RFC3339Nano
	}
	return t.Parse(layout, string(data))
}

func (t *timeRFC3339) Parse(layout, value string) error {
	p, err := time.Parse(layout, strings.ToUpper(value))
	*t = timeRFC3339(p)
	return err
}
