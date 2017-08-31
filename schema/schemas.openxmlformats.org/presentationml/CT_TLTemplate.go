// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
)

type CT_TLTemplate struct {
	// Level
	LvlAttr *uint32
	// Time Node List
	TnLst *CT_TimeNodeList
}

func NewCT_TLTemplate() *CT_TLTemplate {
	ret := &CT_TLTemplate{}
	ret.TnLst = NewCT_TimeNodeList()
	return ret
}
func (m *CT_TLTemplate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.LvlAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "lvl"},
			Value: fmt.Sprintf("%v", *m.LvlAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	setnLst := xml.StartElement{Name: xml.Name{Local: "p:tnLst"}}
	e.EncodeElement(m.TnLst, setnLst)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_TLTemplate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.TnLst = NewCT_TimeNodeList()
	for _, attr := range start.Attr {
		if attr.Name.Local == "lvl" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.LvlAttr = &pt
		}
	}
lCT_TLTemplate:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "tnLst":
				if err := d.DecodeElement(m.TnLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TLTemplate
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_TLTemplate) Validate() error {
	return m.ValidateWithPath("CT_TLTemplate")
}
func (m *CT_TLTemplate) ValidateWithPath(path string) error {
	if err := m.TnLst.ValidateWithPath(path + "/TnLst"); err != nil {
		return err
	}
	return nil
}