// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package diagram

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_DiagramDefinitionHeaderLst struct {
	LayoutDefHdr []*CT_DiagramDefinitionHeader
}

func NewCT_DiagramDefinitionHeaderLst() *CT_DiagramDefinitionHeaderLst {
	ret := &CT_DiagramDefinitionHeaderLst{}
	return ret
}
func (m *CT_DiagramDefinitionHeaderLst) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.LayoutDefHdr != nil {
		selayoutDefHdr := xml.StartElement{Name: xml.Name{Local: "layoutDefHdr"}}
		e.EncodeElement(m.LayoutDefHdr, selayoutDefHdr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_DiagramDefinitionHeaderLst) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_DiagramDefinitionHeaderLst:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "layoutDefHdr":
				tmp := NewCT_DiagramDefinitionHeader()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.LayoutDefHdr = append(m.LayoutDefHdr, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DiagramDefinitionHeaderLst
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_DiagramDefinitionHeaderLst) Validate() error {
	return m.ValidateWithPath("CT_DiagramDefinitionHeaderLst")
}
func (m *CT_DiagramDefinitionHeaderLst) ValidateWithPath(path string) error {
	for i, v := range m.LayoutDefHdr {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/LayoutDefHdr[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}