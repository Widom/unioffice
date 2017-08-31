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
	"log"
	"strconv"
)

type CT_Num struct {
	// Numbering Definition Instance ID
	NumIdAttr int64
	// Abstract Numbering Definition Reference
	AbstractNumId *CT_DecimalNumber
	// Numbering Level Definition Override
	LvlOverride []*CT_NumLvl
}

func NewCT_Num() *CT_Num {
	ret := &CT_Num{}
	ret.AbstractNumId = NewCT_DecimalNumber()
	return ret
}
func (m *CT_Num) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:numId"},
		Value: fmt.Sprintf("%v", m.NumIdAttr)})
	e.EncodeToken(start)
	start.Attr = nil
	seabstractNumId := xml.StartElement{Name: xml.Name{Local: "w:abstractNumId"}}
	e.EncodeElement(m.AbstractNumId, seabstractNumId)
	if m.LvlOverride != nil {
		selvlOverride := xml.StartElement{Name: xml.Name{Local: "w:lvlOverride"}}
		e.EncodeElement(m.LvlOverride, selvlOverride)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Num) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.AbstractNumId = NewCT_DecimalNumber()
	for _, attr := range start.Attr {
		if attr.Name.Local == "numId" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.NumIdAttr = parsed
		}
	}
lCT_Num:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "abstractNumId":
				if err := d.DecodeElement(m.AbstractNumId, &el); err != nil {
					return err
				}
			case "lvlOverride":
				tmp := NewCT_NumLvl()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.LvlOverride = append(m.LvlOverride, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Num
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Num) Validate() error {
	return m.ValidateWithPath("CT_Num")
}
func (m *CT_Num) ValidateWithPath(path string) error {
	if err := m.AbstractNumId.ValidateWithPath(path + "/AbstractNumId"); err != nil {
		return err
	}
	for i, v := range m.LvlOverride {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/LvlOverride[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}