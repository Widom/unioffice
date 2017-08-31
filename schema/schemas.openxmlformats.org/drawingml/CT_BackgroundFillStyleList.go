// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_BackgroundFillStyleList struct {
	EG_FillProperties []*EG_FillProperties
}

func NewCT_BackgroundFillStyleList() *CT_BackgroundFillStyleList {
	ret := &CT_BackgroundFillStyleList{}
	return ret
}
func (m *CT_BackgroundFillStyleList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	for _, c := range m.EG_FillProperties {
		c.MarshalXML(e, start)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_BackgroundFillStyleList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_BackgroundFillStyleList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "noFill":
				tmpfillproperties := NewEG_FillProperties()
				tmpfillproperties.NoFill = NewCT_NoFillProperties()
				if err := d.DecodeElement(tmpfillproperties.NoFill, &el); err != nil {
					return err
				}
				m.EG_FillProperties = append(m.EG_FillProperties, tmpfillproperties)
			case "solidFill":
				tmpfillproperties := NewEG_FillProperties()
				tmpfillproperties.SolidFill = NewCT_SolidColorFillProperties()
				if err := d.DecodeElement(tmpfillproperties.SolidFill, &el); err != nil {
					return err
				}
				m.EG_FillProperties = append(m.EG_FillProperties, tmpfillproperties)
			case "gradFill":
				tmpfillproperties := NewEG_FillProperties()
				tmpfillproperties.GradFill = NewCT_GradientFillProperties()
				if err := d.DecodeElement(tmpfillproperties.GradFill, &el); err != nil {
					return err
				}
				m.EG_FillProperties = append(m.EG_FillProperties, tmpfillproperties)
			case "blipFill":
				tmpfillproperties := NewEG_FillProperties()
				tmpfillproperties.BlipFill = NewCT_BlipFillProperties()
				if err := d.DecodeElement(tmpfillproperties.BlipFill, &el); err != nil {
					return err
				}
				m.EG_FillProperties = append(m.EG_FillProperties, tmpfillproperties)
			case "pattFill":
				tmpfillproperties := NewEG_FillProperties()
				tmpfillproperties.PattFill = NewCT_PatternFillProperties()
				if err := d.DecodeElement(tmpfillproperties.PattFill, &el); err != nil {
					return err
				}
				m.EG_FillProperties = append(m.EG_FillProperties, tmpfillproperties)
			case "grpFill":
				tmpfillproperties := NewEG_FillProperties()
				tmpfillproperties.GrpFill = NewCT_GroupFillProperties()
				if err := d.DecodeElement(tmpfillproperties.GrpFill, &el); err != nil {
					return err
				}
				m.EG_FillProperties = append(m.EG_FillProperties, tmpfillproperties)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_BackgroundFillStyleList
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_BackgroundFillStyleList) Validate() error {
	return m.ValidateWithPath("CT_BackgroundFillStyleList")
}
func (m *CT_BackgroundFillStyleList) ValidateWithPath(path string) error {
	for i, v := range m.EG_FillProperties {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/EG_FillProperties[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}