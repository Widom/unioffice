// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"log"
)

type EG_LineDashProperties struct {
	PrstDash *CT_PresetLineDashProperties
	CustDash *CT_DashStopList
}

func NewEG_LineDashProperties() *EG_LineDashProperties {
	ret := &EG_LineDashProperties{}
	return ret
}
func (m *EG_LineDashProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.PrstDash != nil {
		seprstDash := xml.StartElement{Name: xml.Name{Local: "a:prstDash"}}
		e.EncodeElement(m.PrstDash, seprstDash)
	}
	if m.CustDash != nil {
		secustDash := xml.StartElement{Name: xml.Name{Local: "a:custDash"}}
		e.EncodeElement(m.CustDash, secustDash)
	}
	return nil
}
func (m *EG_LineDashProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_LineDashProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "prstDash":
				m.PrstDash = NewCT_PresetLineDashProperties()
				if err := d.DecodeElement(m.PrstDash, &el); err != nil {
					return err
				}
			case "custDash":
				m.CustDash = NewCT_DashStopList()
				if err := d.DecodeElement(m.CustDash, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_LineDashProperties
		case xml.CharData:
		}
	}
	return nil
}
func (m *EG_LineDashProperties) Validate() error {
	return m.ValidateWithPath("EG_LineDashProperties")
}
func (m *EG_LineDashProperties) ValidateWithPath(path string) error {
	if m.PrstDash != nil {
		if err := m.PrstDash.ValidateWithPath(path + "/PrstDash"); err != nil {
			return err
		}
	}
	if m.CustDash != nil {
		if err := m.CustDash.ValidateWithPath(path + "/CustDash"); err != nil {
			return err
		}
	}
	return nil
}