// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package content_types

import (
	"encoding/xml"
	"fmt"
)

type CT_Default struct {
	ExtensionAttr   string
	ContentTypeAttr string
}

func NewCT_Default() *CT_Default {
	ret := &CT_Default{}
	ret.ExtensionAttr = "xml"
	ret.ContentTypeAttr = "application/xml"
	return ret
}
func (m *CT_Default) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "Extension"},
		Value: fmt.Sprintf("%v", m.ExtensionAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ContentType"},
		Value: fmt.Sprintf("%v", m.ContentTypeAttr)})
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Default) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ExtensionAttr = "xml"
	m.ContentTypeAttr = "application/xml"
	for _, attr := range start.Attr {
		if attr.Name.Local == "Extension" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ExtensionAttr = parsed
		}
		if attr.Name.Local == "ContentType" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ContentTypeAttr = parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Default: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_Default) Validate() error {
	return m.ValidateWithPath("CT_Default")
}
func (m *CT_Default) ValidateWithPath(path string) error {
	if !ST_ExtensionPatternRe.MatchString(m.ExtensionAttr) {
		return fmt.Errorf(`%s/m.ExtensionAttr must match '%s' (have %v)`, path, ST_ExtensionPatternRe, m.ExtensionAttr)
	}
	if !ST_ContentTypePatternRe.MatchString(m.ContentTypeAttr) {
		return fmt.Errorf(`%s/m.ContentTypeAttr must match '%s' (have %v)`, path, ST_ContentTypePatternRe, m.ContentTypeAttr)
	}
	return nil
}