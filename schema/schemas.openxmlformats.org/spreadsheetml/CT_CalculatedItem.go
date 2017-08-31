// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
)

type CT_CalculatedItem struct {
	// Field Index
	FieldAttr *uint32
	// Calculated Item Formula
	FormulaAttr *string
	// Calculated Item Location
	PivotArea *CT_PivotArea
	// Future Feature Data Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_CalculatedItem() *CT_CalculatedItem {
	ret := &CT_CalculatedItem{}
	ret.PivotArea = NewCT_PivotArea()
	return ret
}
func (m *CT_CalculatedItem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.FieldAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "field"},
			Value: fmt.Sprintf("%v", *m.FieldAttr)})
	}
	if m.FormulaAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "formula"},
			Value: fmt.Sprintf("%v", *m.FormulaAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	sepivotArea := xml.StartElement{Name: xml.Name{Local: "x:pivotArea"}}
	e.EncodeElement(m.PivotArea, sepivotArea)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "x:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_CalculatedItem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.PivotArea = NewCT_PivotArea()
	for _, attr := range start.Attr {
		if attr.Name.Local == "field" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.FieldAttr = &pt
		}
		if attr.Name.Local == "formula" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FormulaAttr = &parsed
		}
	}
lCT_CalculatedItem:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "pivotArea":
				if err := d.DecodeElement(m.PivotArea, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CalculatedItem
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_CalculatedItem) Validate() error {
	return m.ValidateWithPath("CT_CalculatedItem")
}
func (m *CT_CalculatedItem) ValidateWithPath(path string) error {
	if err := m.PivotArea.ValidateWithPath(path + "/PivotArea"); err != nil {
		return err
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}