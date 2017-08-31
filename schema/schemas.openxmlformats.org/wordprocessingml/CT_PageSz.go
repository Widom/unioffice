// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/sharedTypes"
)

type CT_PageSz struct {
	// Page Width
	WAttr *sharedTypes.ST_TwipsMeasure
	// Page Height
	HAttr *sharedTypes.ST_TwipsMeasure
	// Page Orientation
	OrientAttr ST_PageOrientation
	// Printer Paper Code
	CodeAttr *int64
}

func NewCT_PageSz() *CT_PageSz {
	ret := &CT_PageSz{}
	return ret
}
func (m *CT_PageSz) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.WAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:w"},
			Value: fmt.Sprintf("%v", *m.WAttr)})
	}
	if m.HAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:h"},
			Value: fmt.Sprintf("%v", *m.HAttr)})
	}
	if m.OrientAttr != ST_PageOrientationUnset {
		attr, err := m.OrientAttr.MarshalXMLAttr(xml.Name{Local: "w:orient"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.CodeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:code"},
			Value: fmt.Sprintf("%v", *m.CodeAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_PageSz) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "w" {
			parsed, err := ParseUnionST_TwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.WAttr = &parsed
		}
		if attr.Name.Local == "h" {
			parsed, err := ParseUnionST_TwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.HAttr = &parsed
		}
		if attr.Name.Local == "orient" {
			m.OrientAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "code" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.CodeAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_PageSz: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_PageSz) Validate() error {
	return m.ValidateWithPath("CT_PageSz")
}
func (m *CT_PageSz) ValidateWithPath(path string) error {
	if m.WAttr != nil {
		if err := m.WAttr.ValidateWithPath(path + "/WAttr"); err != nil {
			return err
		}
	}
	if m.HAttr != nil {
		if err := m.HAttr.ValidateWithPath(path + "/HAttr"); err != nil {
			return err
		}
	}
	if err := m.OrientAttr.ValidateWithPath(path + "/OrientAttr"); err != nil {
		return err
	}
	return nil
}