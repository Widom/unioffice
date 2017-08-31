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
	"strconv"
)

type CT_Hyperlink struct {
	IdAttr             *string
	InvalidUrlAttr     *string
	ActionAttr         *string
	TgtFrameAttr       *string
	TooltipAttr        *string
	HistoryAttr        *bool
	HighlightClickAttr *bool
	EndSndAttr         *bool
	Snd                *CT_EmbeddedWAVAudioFile
	ExtLst             *CT_OfficeArtExtensionList
}

func NewCT_Hyperlink() *CT_Hyperlink {
	ret := &CT_Hyperlink{}
	return ret
}
func (m *CT_Hyperlink) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}
	if m.InvalidUrlAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "invalidUrl"},
			Value: fmt.Sprintf("%v", *m.InvalidUrlAttr)})
	}
	if m.ActionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "action"},
			Value: fmt.Sprintf("%v", *m.ActionAttr)})
	}
	if m.TgtFrameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "tgtFrame"},
			Value: fmt.Sprintf("%v", *m.TgtFrameAttr)})
	}
	if m.TooltipAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "tooltip"},
			Value: fmt.Sprintf("%v", *m.TooltipAttr)})
	}
	if m.HistoryAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "history"},
			Value: fmt.Sprintf("%v", *m.HistoryAttr)})
	}
	if m.HighlightClickAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "highlightClick"},
			Value: fmt.Sprintf("%v", *m.HighlightClickAttr)})
	}
	if m.EndSndAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "endSnd"},
			Value: fmt.Sprintf("%v", *m.EndSndAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.Snd != nil {
		sesnd := xml.StartElement{Name: xml.Name{Local: "a:snd"}}
		e.EncodeElement(m.Snd, sesnd)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Hyperlink) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
		}
		if attr.Name.Local == "invalidUrl" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.InvalidUrlAttr = &parsed
		}
		if attr.Name.Local == "action" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ActionAttr = &parsed
		}
		if attr.Name.Local == "tgtFrame" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TgtFrameAttr = &parsed
		}
		if attr.Name.Local == "tooltip" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TooltipAttr = &parsed
		}
		if attr.Name.Local == "history" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.HistoryAttr = &parsed
		}
		if attr.Name.Local == "highlightClick" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.HighlightClickAttr = &parsed
		}
		if attr.Name.Local == "endSnd" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.EndSndAttr = &parsed
		}
	}
lCT_Hyperlink:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "snd":
				m.Snd = NewCT_EmbeddedWAVAudioFile()
				if err := d.DecodeElement(m.Snd, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_OfficeArtExtensionList()
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
			break lCT_Hyperlink
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Hyperlink) Validate() error {
	return m.ValidateWithPath("CT_Hyperlink")
}
func (m *CT_Hyperlink) ValidateWithPath(path string) error {
	if m.Snd != nil {
		if err := m.Snd.ValidateWithPath(path + "/Snd"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}