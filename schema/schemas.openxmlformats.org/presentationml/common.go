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
	"time"

	"baliance.com/gooxml"
	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

func ParseUnionST_TLTime(s string) (ST_TLTime, error) {
	return ST_TLTime{}, nil
}

func ParseUnionST_FixedPercentage(s string) (drawingml.ST_FixedPercentage, error) {
	return drawingml.ParseUnionST_FixedPercentage(s)
}
func ParseStdlibTime(s string) (time.Time, error) {
	return drawingml.ParseStdlibTime(s)
}
func ParseUnionST_Percentage(s string) (drawingml.ST_Percentage, error) {
	return drawingml.ParseUnionST_Percentage(s)
}
func ParseUnionST_Coordinate32(s string) (drawingml.ST_Coordinate32, error) {
	return drawingml.ParseUnionST_Coordinate32(s)
}
func ParseUnionST_PositiveFixedPercentage(s string) (drawingml.ST_PositiveFixedPercentage, error) {
	return drawingml.ParseUnionST_PositiveFixedPercentage(s)
}
func ParseUnionST_TLTimeAnimateValueTime(s string) (ST_TLTimeAnimateValueTime, error) {
	return ST_TLTimeAnimateValueTime{}, nil
}
func ParseUnionST_PositivePercentage(s string) (drawingml.ST_PositivePercentage, error) {
	return drawingml.ParseUnionST_PositivePercentage(s)
}
func ParseUnionST_TransitionEightDirectionType(s string) (ST_TransitionEightDirectionType, error) {
	return ST_TransitionEightDirectionType{}, nil
}

type ST_TransitionSideDirectionType byte

const (
	ST_TransitionSideDirectionTypeUnset ST_TransitionSideDirectionType = 0
	ST_TransitionSideDirectionTypeL     ST_TransitionSideDirectionType = 1
	ST_TransitionSideDirectionTypeU     ST_TransitionSideDirectionType = 2
	ST_TransitionSideDirectionTypeR     ST_TransitionSideDirectionType = 3
	ST_TransitionSideDirectionTypeD     ST_TransitionSideDirectionType = 4
)

func (e ST_TransitionSideDirectionType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TransitionSideDirectionTypeUnset:
		attr.Value = ""
	case ST_TransitionSideDirectionTypeL:
		attr.Value = "l"
	case ST_TransitionSideDirectionTypeU:
		attr.Value = "u"
	case ST_TransitionSideDirectionTypeR:
		attr.Value = "r"
	case ST_TransitionSideDirectionTypeD:
		attr.Value = "d"
	}
	return attr, nil
}
func (e *ST_TransitionSideDirectionType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "l":
		*e = 1
	case "u":
		*e = 2
	case "r":
		*e = 3
	case "d":
		*e = 4
	}
	return nil
}
func (m ST_TransitionSideDirectionType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TransitionSideDirectionType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "l":
			*m = 1
		case "u":
			*m = 2
		case "r":
			*m = 3
		case "d":
			*m = 4
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_TransitionSideDirectionType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "l"
	case 2:
		return "u"
	case 3:
		return "r"
	case 4:
		return "d"
	}
	return ""
}
func (m ST_TransitionSideDirectionType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TransitionSideDirectionType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TransitionCornerDirectionType byte

const (
	ST_TransitionCornerDirectionTypeUnset ST_TransitionCornerDirectionType = 0
	ST_TransitionCornerDirectionTypeLu    ST_TransitionCornerDirectionType = 1
	ST_TransitionCornerDirectionTypeRu    ST_TransitionCornerDirectionType = 2
	ST_TransitionCornerDirectionTypeLd    ST_TransitionCornerDirectionType = 3
	ST_TransitionCornerDirectionTypeRd    ST_TransitionCornerDirectionType = 4
)

func (e ST_TransitionCornerDirectionType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TransitionCornerDirectionTypeUnset:
		attr.Value = ""
	case ST_TransitionCornerDirectionTypeLu:
		attr.Value = "lu"
	case ST_TransitionCornerDirectionTypeRu:
		attr.Value = "ru"
	case ST_TransitionCornerDirectionTypeLd:
		attr.Value = "ld"
	case ST_TransitionCornerDirectionTypeRd:
		attr.Value = "rd"
	}
	return attr, nil
}
func (e *ST_TransitionCornerDirectionType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "lu":
		*e = 1
	case "ru":
		*e = 2
	case "ld":
		*e = 3
	case "rd":
		*e = 4
	}
	return nil
}
func (m ST_TransitionCornerDirectionType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TransitionCornerDirectionType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "lu":
			*m = 1
		case "ru":
			*m = 2
		case "ld":
			*m = 3
		case "rd":
			*m = 4
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_TransitionCornerDirectionType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "lu"
	case 2:
		return "ru"
	case 3:
		return "ld"
	case 4:
		return "rd"
	}
	return ""
}
func (m ST_TransitionCornerDirectionType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TransitionCornerDirectionType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TransitionInOutDirectionType byte

const (
	ST_TransitionInOutDirectionTypeUnset ST_TransitionInOutDirectionType = 0
	ST_TransitionInOutDirectionTypeOut   ST_TransitionInOutDirectionType = 1
	ST_TransitionInOutDirectionTypeIn    ST_TransitionInOutDirectionType = 2
)

func (e ST_TransitionInOutDirectionType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TransitionInOutDirectionTypeUnset:
		attr.Value = ""
	case ST_TransitionInOutDirectionTypeOut:
		attr.Value = "out"
	case ST_TransitionInOutDirectionTypeIn:
		attr.Value = "in"
	}
	return attr, nil
}
func (e *ST_TransitionInOutDirectionType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "out":
		*e = 1
	case "in":
		*e = 2
	}
	return nil
}
func (m ST_TransitionInOutDirectionType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TransitionInOutDirectionType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "out":
			*m = 1
		case "in":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_TransitionInOutDirectionType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "out"
	case 2:
		return "in"
	}
	return ""
}
func (m ST_TransitionInOutDirectionType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TransitionInOutDirectionType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TransitionSpeed byte

const (
	ST_TransitionSpeedUnset ST_TransitionSpeed = 0
	ST_TransitionSpeedSlow  ST_TransitionSpeed = 1
	ST_TransitionSpeedMed   ST_TransitionSpeed = 2
	ST_TransitionSpeedFast  ST_TransitionSpeed = 3
)

func (e ST_TransitionSpeed) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TransitionSpeedUnset:
		attr.Value = ""
	case ST_TransitionSpeedSlow:
		attr.Value = "slow"
	case ST_TransitionSpeedMed:
		attr.Value = "med"
	case ST_TransitionSpeedFast:
		attr.Value = "fast"
	}
	return attr, nil
}
func (e *ST_TransitionSpeed) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "slow":
		*e = 1
	case "med":
		*e = 2
	case "fast":
		*e = 3
	}
	return nil
}
func (m ST_TransitionSpeed) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TransitionSpeed) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "slow":
			*m = 1
		case "med":
			*m = 2
		case "fast":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_TransitionSpeed) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "slow"
	case 2:
		return "med"
	case 3:
		return "fast"
	}
	return ""
}
func (m ST_TransitionSpeed) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TransitionSpeed) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TLTimeIndefinite byte

const (
	ST_TLTimeIndefiniteUnset      ST_TLTimeIndefinite = 0
	ST_TLTimeIndefiniteIndefinite ST_TLTimeIndefinite = 1
)

func (e ST_TLTimeIndefinite) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TLTimeIndefiniteUnset:
		attr.Value = ""
	case ST_TLTimeIndefiniteIndefinite:
		attr.Value = "indefinite"
	}
	return attr, nil
}
func (e *ST_TLTimeIndefinite) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "indefinite":
		*e = 1
	}
	return nil
}
func (m ST_TLTimeIndefinite) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TLTimeIndefinite) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "indefinite":
			*m = 1
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_TLTimeIndefinite) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "indefinite"
	}
	return ""
}
func (m ST_TLTimeIndefinite) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TLTimeIndefinite) ValidateWithPath(path string) error {
	switch m {
	case 0, 1:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_IterateType byte

const (
	ST_IterateTypeUnset ST_IterateType = 0
	ST_IterateTypeEl    ST_IterateType = 1
	ST_IterateTypeWd    ST_IterateType = 2
	ST_IterateTypeLt    ST_IterateType = 3
)

func (e ST_IterateType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_IterateTypeUnset:
		attr.Value = ""
	case ST_IterateTypeEl:
		attr.Value = "el"
	case ST_IterateTypeWd:
		attr.Value = "wd"
	case ST_IterateTypeLt:
		attr.Value = "lt"
	}
	return attr, nil
}
func (e *ST_IterateType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "el":
		*e = 1
	case "wd":
		*e = 2
	case "lt":
		*e = 3
	}
	return nil
}
func (m ST_IterateType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_IterateType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "el":
			*m = 1
		case "wd":
			*m = 2
		case "lt":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_IterateType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "el"
	case 2:
		return "wd"
	case 3:
		return "lt"
	}
	return ""
}
func (m ST_IterateType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_IterateType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TLChartSubelementType byte

const (
	ST_TLChartSubelementTypeUnset        ST_TLChartSubelementType = 0
	ST_TLChartSubelementTypeGridLegend   ST_TLChartSubelementType = 1
	ST_TLChartSubelementTypeSeries       ST_TLChartSubelementType = 2
	ST_TLChartSubelementTypeCategory     ST_TLChartSubelementType = 3
	ST_TLChartSubelementTypePtInSeries   ST_TLChartSubelementType = 4
	ST_TLChartSubelementTypePtInCategory ST_TLChartSubelementType = 5
)

func (e ST_TLChartSubelementType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TLChartSubelementTypeUnset:
		attr.Value = ""
	case ST_TLChartSubelementTypeGridLegend:
		attr.Value = "gridLegend"
	case ST_TLChartSubelementTypeSeries:
		attr.Value = "series"
	case ST_TLChartSubelementTypeCategory:
		attr.Value = "category"
	case ST_TLChartSubelementTypePtInSeries:
		attr.Value = "ptInSeries"
	case ST_TLChartSubelementTypePtInCategory:
		attr.Value = "ptInCategory"
	}
	return attr, nil
}
func (e *ST_TLChartSubelementType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "gridLegend":
		*e = 1
	case "series":
		*e = 2
	case "category":
		*e = 3
	case "ptInSeries":
		*e = 4
	case "ptInCategory":
		*e = 5
	}
	return nil
}
func (m ST_TLChartSubelementType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TLChartSubelementType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "gridLegend":
			*m = 1
		case "series":
			*m = 2
		case "category":
			*m = 3
		case "ptInSeries":
			*m = 4
		case "ptInCategory":
			*m = 5
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_TLChartSubelementType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "gridLegend"
	case 2:
		return "series"
	case 3:
		return "category"
	case 4:
		return "ptInSeries"
	case 5:
		return "ptInCategory"
	}
	return ""
}
func (m ST_TLChartSubelementType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TLChartSubelementType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TLTriggerRuntimeNode byte

const (
	ST_TLTriggerRuntimeNodeUnset ST_TLTriggerRuntimeNode = 0
	ST_TLTriggerRuntimeNodeFirst ST_TLTriggerRuntimeNode = 1
	ST_TLTriggerRuntimeNodeLast  ST_TLTriggerRuntimeNode = 2
	ST_TLTriggerRuntimeNodeAll   ST_TLTriggerRuntimeNode = 3
)

func (e ST_TLTriggerRuntimeNode) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TLTriggerRuntimeNodeUnset:
		attr.Value = ""
	case ST_TLTriggerRuntimeNodeFirst:
		attr.Value = "first"
	case ST_TLTriggerRuntimeNodeLast:
		attr.Value = "last"
	case ST_TLTriggerRuntimeNodeAll:
		attr.Value = "all"
	}
	return attr, nil
}
func (e *ST_TLTriggerRuntimeNode) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "first":
		*e = 1
	case "last":
		*e = 2
	case "all":
		*e = 3
	}
	return nil
}
func (m ST_TLTriggerRuntimeNode) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TLTriggerRuntimeNode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "first":
			*m = 1
		case "last":
			*m = 2
		case "all":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_TLTriggerRuntimeNode) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "first"
	case 2:
		return "last"
	case 3:
		return "all"
	}
	return ""
}
func (m ST_TLTriggerRuntimeNode) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TLTriggerRuntimeNode) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TLTriggerEvent byte

const (
	ST_TLTriggerEventUnset       ST_TLTriggerEvent = 0
	ST_TLTriggerEventOnBegin     ST_TLTriggerEvent = 1
	ST_TLTriggerEventOnEnd       ST_TLTriggerEvent = 2
	ST_TLTriggerEventBegin       ST_TLTriggerEvent = 3
	ST_TLTriggerEventEnd         ST_TLTriggerEvent = 4
	ST_TLTriggerEventOnClick     ST_TLTriggerEvent = 5
	ST_TLTriggerEventOnDblClick  ST_TLTriggerEvent = 6
	ST_TLTriggerEventOnMouseOver ST_TLTriggerEvent = 7
	ST_TLTriggerEventOnMouseOut  ST_TLTriggerEvent = 8
	ST_TLTriggerEventOnNext      ST_TLTriggerEvent = 9
	ST_TLTriggerEventOnPrev      ST_TLTriggerEvent = 10
	ST_TLTriggerEventOnStopAudio ST_TLTriggerEvent = 11
)

func (e ST_TLTriggerEvent) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TLTriggerEventUnset:
		attr.Value = ""
	case ST_TLTriggerEventOnBegin:
		attr.Value = "onBegin"
	case ST_TLTriggerEventOnEnd:
		attr.Value = "onEnd"
	case ST_TLTriggerEventBegin:
		attr.Value = "begin"
	case ST_TLTriggerEventEnd:
		attr.Value = "end"
	case ST_TLTriggerEventOnClick:
		attr.Value = "onClick"
	case ST_TLTriggerEventOnDblClick:
		attr.Value = "onDblClick"
	case ST_TLTriggerEventOnMouseOver:
		attr.Value = "onMouseOver"
	case ST_TLTriggerEventOnMouseOut:
		attr.Value = "onMouseOut"
	case ST_TLTriggerEventOnNext:
		attr.Value = "onNext"
	case ST_TLTriggerEventOnPrev:
		attr.Value = "onPrev"
	case ST_TLTriggerEventOnStopAudio:
		attr.Value = "onStopAudio"
	}
	return attr, nil
}
func (e *ST_TLTriggerEvent) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "onBegin":
		*e = 1
	case "onEnd":
		*e = 2
	case "begin":
		*e = 3
	case "end":
		*e = 4
	case "onClick":
		*e = 5
	case "onDblClick":
		*e = 6
	case "onMouseOver":
		*e = 7
	case "onMouseOut":
		*e = 8
	case "onNext":
		*e = 9
	case "onPrev":
		*e = 10
	case "onStopAudio":
		*e = 11
	}
	return nil
}
func (m ST_TLTriggerEvent) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TLTriggerEvent) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "onBegin":
			*m = 1
		case "onEnd":
			*m = 2
		case "begin":
			*m = 3
		case "end":
			*m = 4
		case "onClick":
			*m = 5
		case "onDblClick":
			*m = 6
		case "onMouseOver":
			*m = 7
		case "onMouseOut":
			*m = 8
		case "onNext":
			*m = 9
		case "onPrev":
			*m = 10
		case "onStopAudio":
			*m = 11
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_TLTriggerEvent) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "onBegin"
	case 2:
		return "onEnd"
	case 3:
		return "begin"
	case 4:
		return "end"
	case 5:
		return "onClick"
	case 6:
		return "onDblClick"
	case 7:
		return "onMouseOver"
	case 8:
		return "onMouseOut"
	case 9:
		return "onNext"
	case 10:
		return "onPrev"
	case 11:
		return "onStopAudio"
	}
	return ""
}
func (m ST_TLTriggerEvent) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TLTriggerEvent) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TLTimeNodePresetClassType byte

const (
	ST_TLTimeNodePresetClassTypeUnset     ST_TLTimeNodePresetClassType = 0
	ST_TLTimeNodePresetClassTypeEntr      ST_TLTimeNodePresetClassType = 1
	ST_TLTimeNodePresetClassTypeExit      ST_TLTimeNodePresetClassType = 2
	ST_TLTimeNodePresetClassTypeEmph      ST_TLTimeNodePresetClassType = 3
	ST_TLTimeNodePresetClassTypePath      ST_TLTimeNodePresetClassType = 4
	ST_TLTimeNodePresetClassTypeVerb      ST_TLTimeNodePresetClassType = 5
	ST_TLTimeNodePresetClassTypeMediacall ST_TLTimeNodePresetClassType = 6
)

func (e ST_TLTimeNodePresetClassType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TLTimeNodePresetClassTypeUnset:
		attr.Value = ""
	case ST_TLTimeNodePresetClassTypeEntr:
		attr.Value = "entr"
	case ST_TLTimeNodePresetClassTypeExit:
		attr.Value = "exit"
	case ST_TLTimeNodePresetClassTypeEmph:
		attr.Value = "emph"
	case ST_TLTimeNodePresetClassTypePath:
		attr.Value = "path"
	case ST_TLTimeNodePresetClassTypeVerb:
		attr.Value = "verb"
	case ST_TLTimeNodePresetClassTypeMediacall:
		attr.Value = "mediacall"
	}
	return attr, nil
}
func (e *ST_TLTimeNodePresetClassType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "entr":
		*e = 1
	case "exit":
		*e = 2
	case "emph":
		*e = 3
	case "path":
		*e = 4
	case "verb":
		*e = 5
	case "mediacall":
		*e = 6
	}
	return nil
}
func (m ST_TLTimeNodePresetClassType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TLTimeNodePresetClassType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "entr":
			*m = 1
		case "exit":
			*m = 2
		case "emph":
			*m = 3
		case "path":
			*m = 4
		case "verb":
			*m = 5
		case "mediacall":
			*m = 6
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_TLTimeNodePresetClassType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "entr"
	case 2:
		return "exit"
	case 3:
		return "emph"
	case 4:
		return "path"
	case 5:
		return "verb"
	case 6:
		return "mediacall"
	}
	return ""
}
func (m ST_TLTimeNodePresetClassType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TLTimeNodePresetClassType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TLTimeNodeRestartType byte

const (
	ST_TLTimeNodeRestartTypeUnset         ST_TLTimeNodeRestartType = 0
	ST_TLTimeNodeRestartTypeAlways        ST_TLTimeNodeRestartType = 1
	ST_TLTimeNodeRestartTypeWhenNotActive ST_TLTimeNodeRestartType = 2
	ST_TLTimeNodeRestartTypeNever         ST_TLTimeNodeRestartType = 3
)

func (e ST_TLTimeNodeRestartType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TLTimeNodeRestartTypeUnset:
		attr.Value = ""
	case ST_TLTimeNodeRestartTypeAlways:
		attr.Value = "always"
	case ST_TLTimeNodeRestartTypeWhenNotActive:
		attr.Value = "whenNotActive"
	case ST_TLTimeNodeRestartTypeNever:
		attr.Value = "never"
	}
	return attr, nil
}
func (e *ST_TLTimeNodeRestartType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "always":
		*e = 1
	case "whenNotActive":
		*e = 2
	case "never":
		*e = 3
	}
	return nil
}
func (m ST_TLTimeNodeRestartType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TLTimeNodeRestartType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "always":
			*m = 1
		case "whenNotActive":
			*m = 2
		case "never":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_TLTimeNodeRestartType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "always"
	case 2:
		return "whenNotActive"
	case 3:
		return "never"
	}
	return ""
}
func (m ST_TLTimeNodeRestartType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TLTimeNodeRestartType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TLTimeNodeFillType byte

const (
	ST_TLTimeNodeFillTypeUnset      ST_TLTimeNodeFillType = 0
	ST_TLTimeNodeFillTypeRemove     ST_TLTimeNodeFillType = 1
	ST_TLTimeNodeFillTypeFreeze     ST_TLTimeNodeFillType = 2
	ST_TLTimeNodeFillTypeHold       ST_TLTimeNodeFillType = 3
	ST_TLTimeNodeFillTypeTransition ST_TLTimeNodeFillType = 4
)

func (e ST_TLTimeNodeFillType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TLTimeNodeFillTypeUnset:
		attr.Value = ""
	case ST_TLTimeNodeFillTypeRemove:
		attr.Value = "remove"
	case ST_TLTimeNodeFillTypeFreeze:
		attr.Value = "freeze"
	case ST_TLTimeNodeFillTypeHold:
		attr.Value = "hold"
	case ST_TLTimeNodeFillTypeTransition:
		attr.Value = "transition"
	}
	return attr, nil
}
func (e *ST_TLTimeNodeFillType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "remove":
		*e = 1
	case "freeze":
		*e = 2
	case "hold":
		*e = 3
	case "transition":
		*e = 4
	}
	return nil
}
func (m ST_TLTimeNodeFillType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TLTimeNodeFillType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "remove":
			*m = 1
		case "freeze":
			*m = 2
		case "hold":
			*m = 3
		case "transition":
			*m = 4
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_TLTimeNodeFillType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "remove"
	case 2:
		return "freeze"
	case 3:
		return "hold"
	case 4:
		return "transition"
	}
	return ""
}
func (m ST_TLTimeNodeFillType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TLTimeNodeFillType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TLTimeNodeSyncType byte

const (
	ST_TLTimeNodeSyncTypeUnset   ST_TLTimeNodeSyncType = 0
	ST_TLTimeNodeSyncTypeCanSlip ST_TLTimeNodeSyncType = 1
	ST_TLTimeNodeSyncTypeLocked  ST_TLTimeNodeSyncType = 2
)

func (e ST_TLTimeNodeSyncType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TLTimeNodeSyncTypeUnset:
		attr.Value = ""
	case ST_TLTimeNodeSyncTypeCanSlip:
		attr.Value = "canSlip"
	case ST_TLTimeNodeSyncTypeLocked:
		attr.Value = "locked"
	}
	return attr, nil
}
func (e *ST_TLTimeNodeSyncType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "canSlip":
		*e = 1
	case "locked":
		*e = 2
	}
	return nil
}
func (m ST_TLTimeNodeSyncType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TLTimeNodeSyncType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "canSlip":
			*m = 1
		case "locked":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_TLTimeNodeSyncType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "canSlip"
	case 2:
		return "locked"
	}
	return ""
}
func (m ST_TLTimeNodeSyncType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TLTimeNodeSyncType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TLTimeNodeMasterRelation byte

const (
	ST_TLTimeNodeMasterRelationUnset     ST_TLTimeNodeMasterRelation = 0
	ST_TLTimeNodeMasterRelationSameClick ST_TLTimeNodeMasterRelation = 1
	ST_TLTimeNodeMasterRelationLastClick ST_TLTimeNodeMasterRelation = 2
	ST_TLTimeNodeMasterRelationNextClick ST_TLTimeNodeMasterRelation = 3
)

func (e ST_TLTimeNodeMasterRelation) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TLTimeNodeMasterRelationUnset:
		attr.Value = ""
	case ST_TLTimeNodeMasterRelationSameClick:
		attr.Value = "sameClick"
	case ST_TLTimeNodeMasterRelationLastClick:
		attr.Value = "lastClick"
	case ST_TLTimeNodeMasterRelationNextClick:
		attr.Value = "nextClick"
	}
	return attr, nil
}
func (e *ST_TLTimeNodeMasterRelation) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "sameClick":
		*e = 1
	case "lastClick":
		*e = 2
	case "nextClick":
		*e = 3
	}
	return nil
}
func (m ST_TLTimeNodeMasterRelation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TLTimeNodeMasterRelation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "sameClick":
			*m = 1
		case "lastClick":
			*m = 2
		case "nextClick":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_TLTimeNodeMasterRelation) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "sameClick"
	case 2:
		return "lastClick"
	case 3:
		return "nextClick"
	}
	return ""
}
func (m ST_TLTimeNodeMasterRelation) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TLTimeNodeMasterRelation) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TLTimeNodeType byte

const (
	ST_TLTimeNodeTypeUnset          ST_TLTimeNodeType = 0
	ST_TLTimeNodeTypeClickEffect    ST_TLTimeNodeType = 1
	ST_TLTimeNodeTypeWithEffect     ST_TLTimeNodeType = 2
	ST_TLTimeNodeTypeAfterEffect    ST_TLTimeNodeType = 3
	ST_TLTimeNodeTypeMainSeq        ST_TLTimeNodeType = 4
	ST_TLTimeNodeTypeInteractiveSeq ST_TLTimeNodeType = 5
	ST_TLTimeNodeTypeClickPar       ST_TLTimeNodeType = 6
	ST_TLTimeNodeTypeWithGroup      ST_TLTimeNodeType = 7
	ST_TLTimeNodeTypeAfterGroup     ST_TLTimeNodeType = 8
	ST_TLTimeNodeTypeTmRoot         ST_TLTimeNodeType = 9
)

func (e ST_TLTimeNodeType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TLTimeNodeTypeUnset:
		attr.Value = ""
	case ST_TLTimeNodeTypeClickEffect:
		attr.Value = "clickEffect"
	case ST_TLTimeNodeTypeWithEffect:
		attr.Value = "withEffect"
	case ST_TLTimeNodeTypeAfterEffect:
		attr.Value = "afterEffect"
	case ST_TLTimeNodeTypeMainSeq:
		attr.Value = "mainSeq"
	case ST_TLTimeNodeTypeInteractiveSeq:
		attr.Value = "interactiveSeq"
	case ST_TLTimeNodeTypeClickPar:
		attr.Value = "clickPar"
	case ST_TLTimeNodeTypeWithGroup:
		attr.Value = "withGroup"
	case ST_TLTimeNodeTypeAfterGroup:
		attr.Value = "afterGroup"
	case ST_TLTimeNodeTypeTmRoot:
		attr.Value = "tmRoot"
	}
	return attr, nil
}
func (e *ST_TLTimeNodeType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "clickEffect":
		*e = 1
	case "withEffect":
		*e = 2
	case "afterEffect":
		*e = 3
	case "mainSeq":
		*e = 4
	case "interactiveSeq":
		*e = 5
	case "clickPar":
		*e = 6
	case "withGroup":
		*e = 7
	case "afterGroup":
		*e = 8
	case "tmRoot":
		*e = 9
	}
	return nil
}
func (m ST_TLTimeNodeType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TLTimeNodeType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "clickEffect":
			*m = 1
		case "withEffect":
			*m = 2
		case "afterEffect":
			*m = 3
		case "mainSeq":
			*m = 4
		case "interactiveSeq":
			*m = 5
		case "clickPar":
			*m = 6
		case "withGroup":
			*m = 7
		case "afterGroup":
			*m = 8
		case "tmRoot":
			*m = 9
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_TLTimeNodeType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "clickEffect"
	case 2:
		return "withEffect"
	case 3:
		return "afterEffect"
	case 4:
		return "mainSeq"
	case 5:
		return "interactiveSeq"
	case 6:
		return "clickPar"
	case 7:
		return "withGroup"
	case 8:
		return "afterGroup"
	case 9:
		return "tmRoot"
	}
	return ""
}
func (m ST_TLTimeNodeType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TLTimeNodeType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TLNextActionType byte

const (
	ST_TLNextActionTypeUnset ST_TLNextActionType = 0
	ST_TLNextActionTypeNone  ST_TLNextActionType = 1
	ST_TLNextActionTypeSeek  ST_TLNextActionType = 2
)

func (e ST_TLNextActionType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TLNextActionTypeUnset:
		attr.Value = ""
	case ST_TLNextActionTypeNone:
		attr.Value = "none"
	case ST_TLNextActionTypeSeek:
		attr.Value = "seek"
	}
	return attr, nil
}
func (e *ST_TLNextActionType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "seek":
		*e = 2
	}
	return nil
}
func (m ST_TLNextActionType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TLNextActionType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "none":
			*m = 1
		case "seek":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_TLNextActionType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "seek"
	}
	return ""
}
func (m ST_TLNextActionType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TLNextActionType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TLPreviousActionType byte

const (
	ST_TLPreviousActionTypeUnset     ST_TLPreviousActionType = 0
	ST_TLPreviousActionTypeNone      ST_TLPreviousActionType = 1
	ST_TLPreviousActionTypeSkipTimed ST_TLPreviousActionType = 2
)

func (e ST_TLPreviousActionType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TLPreviousActionTypeUnset:
		attr.Value = ""
	case ST_TLPreviousActionTypeNone:
		attr.Value = "none"
	case ST_TLPreviousActionTypeSkipTimed:
		attr.Value = "skipTimed"
	}
	return attr, nil
}
func (e *ST_TLPreviousActionType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "skipTimed":
		*e = 2
	}
	return nil
}
func (m ST_TLPreviousActionType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TLPreviousActionType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "none":
			*m = 1
		case "skipTimed":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_TLPreviousActionType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "skipTimed"
	}
	return ""
}
func (m ST_TLPreviousActionType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TLPreviousActionType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TLBehaviorAdditiveType byte

const (
	ST_TLBehaviorAdditiveTypeUnset ST_TLBehaviorAdditiveType = 0
	ST_TLBehaviorAdditiveTypeBase  ST_TLBehaviorAdditiveType = 1
	ST_TLBehaviorAdditiveTypeSum   ST_TLBehaviorAdditiveType = 2
	ST_TLBehaviorAdditiveTypeRepl  ST_TLBehaviorAdditiveType = 3
	ST_TLBehaviorAdditiveTypeMult  ST_TLBehaviorAdditiveType = 4
	ST_TLBehaviorAdditiveTypeNone  ST_TLBehaviorAdditiveType = 5
)

func (e ST_TLBehaviorAdditiveType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TLBehaviorAdditiveTypeUnset:
		attr.Value = ""
	case ST_TLBehaviorAdditiveTypeBase:
		attr.Value = "base"
	case ST_TLBehaviorAdditiveTypeSum:
		attr.Value = "sum"
	case ST_TLBehaviorAdditiveTypeRepl:
		attr.Value = "repl"
	case ST_TLBehaviorAdditiveTypeMult:
		attr.Value = "mult"
	case ST_TLBehaviorAdditiveTypeNone:
		attr.Value = "none"
	}
	return attr, nil
}
func (e *ST_TLBehaviorAdditiveType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "base":
		*e = 1
	case "sum":
		*e = 2
	case "repl":
		*e = 3
	case "mult":
		*e = 4
	case "none":
		*e = 5
	}
	return nil
}
func (m ST_TLBehaviorAdditiveType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TLBehaviorAdditiveType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "base":
			*m = 1
		case "sum":
			*m = 2
		case "repl":
			*m = 3
		case "mult":
			*m = 4
		case "none":
			*m = 5
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_TLBehaviorAdditiveType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "base"
	case 2:
		return "sum"
	case 3:
		return "repl"
	case 4:
		return "mult"
	case 5:
		return "none"
	}
	return ""
}
func (m ST_TLBehaviorAdditiveType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TLBehaviorAdditiveType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TLBehaviorAccumulateType byte

const (
	ST_TLBehaviorAccumulateTypeUnset  ST_TLBehaviorAccumulateType = 0
	ST_TLBehaviorAccumulateTypeNone   ST_TLBehaviorAccumulateType = 1
	ST_TLBehaviorAccumulateTypeAlways ST_TLBehaviorAccumulateType = 2
)

func (e ST_TLBehaviorAccumulateType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TLBehaviorAccumulateTypeUnset:
		attr.Value = ""
	case ST_TLBehaviorAccumulateTypeNone:
		attr.Value = "none"
	case ST_TLBehaviorAccumulateTypeAlways:
		attr.Value = "always"
	}
	return attr, nil
}
func (e *ST_TLBehaviorAccumulateType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "always":
		*e = 2
	}
	return nil
}
func (m ST_TLBehaviorAccumulateType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TLBehaviorAccumulateType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "none":
			*m = 1
		case "always":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_TLBehaviorAccumulateType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "always"
	}
	return ""
}
func (m ST_TLBehaviorAccumulateType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TLBehaviorAccumulateType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TLBehaviorTransformType byte

const (
	ST_TLBehaviorTransformTypeUnset ST_TLBehaviorTransformType = 0
	ST_TLBehaviorTransformTypePt    ST_TLBehaviorTransformType = 1
	ST_TLBehaviorTransformTypeImg   ST_TLBehaviorTransformType = 2
)

func (e ST_TLBehaviorTransformType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TLBehaviorTransformTypeUnset:
		attr.Value = ""
	case ST_TLBehaviorTransformTypePt:
		attr.Value = "pt"
	case ST_TLBehaviorTransformTypeImg:
		attr.Value = "img"
	}
	return attr, nil
}
func (e *ST_TLBehaviorTransformType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "pt":
		*e = 1
	case "img":
		*e = 2
	}
	return nil
}
func (m ST_TLBehaviorTransformType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TLBehaviorTransformType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "pt":
			*m = 1
		case "img":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_TLBehaviorTransformType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "pt"
	case 2:
		return "img"
	}
	return ""
}
func (m ST_TLBehaviorTransformType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TLBehaviorTransformType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TLBehaviorOverrideType byte

const (
	ST_TLBehaviorOverrideTypeUnset      ST_TLBehaviorOverrideType = 0
	ST_TLBehaviorOverrideTypeNormal     ST_TLBehaviorOverrideType = 1
	ST_TLBehaviorOverrideTypeChildStyle ST_TLBehaviorOverrideType = 2
)

func (e ST_TLBehaviorOverrideType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TLBehaviorOverrideTypeUnset:
		attr.Value = ""
	case ST_TLBehaviorOverrideTypeNormal:
		attr.Value = "normal"
	case ST_TLBehaviorOverrideTypeChildStyle:
		attr.Value = "childStyle"
	}
	return attr, nil
}
func (e *ST_TLBehaviorOverrideType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "normal":
		*e = 1
	case "childStyle":
		*e = 2
	}
	return nil
}
func (m ST_TLBehaviorOverrideType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TLBehaviorOverrideType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "normal":
			*m = 1
		case "childStyle":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_TLBehaviorOverrideType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "normal"
	case 2:
		return "childStyle"
	}
	return ""
}
func (m ST_TLBehaviorOverrideType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TLBehaviorOverrideType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TLAnimateBehaviorCalcMode byte

const (
	ST_TLAnimateBehaviorCalcModeUnset    ST_TLAnimateBehaviorCalcMode = 0
	ST_TLAnimateBehaviorCalcModeDiscrete ST_TLAnimateBehaviorCalcMode = 1
	ST_TLAnimateBehaviorCalcModeLin      ST_TLAnimateBehaviorCalcMode = 2
	ST_TLAnimateBehaviorCalcModeFmla     ST_TLAnimateBehaviorCalcMode = 3
)

func (e ST_TLAnimateBehaviorCalcMode) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TLAnimateBehaviorCalcModeUnset:
		attr.Value = ""
	case ST_TLAnimateBehaviorCalcModeDiscrete:
		attr.Value = "discrete"
	case ST_TLAnimateBehaviorCalcModeLin:
		attr.Value = "lin"
	case ST_TLAnimateBehaviorCalcModeFmla:
		attr.Value = "fmla"
	}
	return attr, nil
}
func (e *ST_TLAnimateBehaviorCalcMode) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "discrete":
		*e = 1
	case "lin":
		*e = 2
	case "fmla":
		*e = 3
	}
	return nil
}
func (m ST_TLAnimateBehaviorCalcMode) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TLAnimateBehaviorCalcMode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "discrete":
			*m = 1
		case "lin":
			*m = 2
		case "fmla":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_TLAnimateBehaviorCalcMode) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "discrete"
	case 2:
		return "lin"
	case 3:
		return "fmla"
	}
	return ""
}
func (m ST_TLAnimateBehaviorCalcMode) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TLAnimateBehaviorCalcMode) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TLAnimateBehaviorValueType byte

const (
	ST_TLAnimateBehaviorValueTypeUnset ST_TLAnimateBehaviorValueType = 0
	ST_TLAnimateBehaviorValueTypeStr   ST_TLAnimateBehaviorValueType = 1
	ST_TLAnimateBehaviorValueTypeNum   ST_TLAnimateBehaviorValueType = 2
	ST_TLAnimateBehaviorValueTypeClr   ST_TLAnimateBehaviorValueType = 3
)

func (e ST_TLAnimateBehaviorValueType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TLAnimateBehaviorValueTypeUnset:
		attr.Value = ""
	case ST_TLAnimateBehaviorValueTypeStr:
		attr.Value = "str"
	case ST_TLAnimateBehaviorValueTypeNum:
		attr.Value = "num"
	case ST_TLAnimateBehaviorValueTypeClr:
		attr.Value = "clr"
	}
	return attr, nil
}
func (e *ST_TLAnimateBehaviorValueType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "str":
		*e = 1
	case "num":
		*e = 2
	case "clr":
		*e = 3
	}
	return nil
}
func (m ST_TLAnimateBehaviorValueType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TLAnimateBehaviorValueType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "str":
			*m = 1
		case "num":
			*m = 2
		case "clr":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_TLAnimateBehaviorValueType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "str"
	case 2:
		return "num"
	case 3:
		return "clr"
	}
	return ""
}
func (m ST_TLAnimateBehaviorValueType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TLAnimateBehaviorValueType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TLAnimateColorSpace byte

const (
	ST_TLAnimateColorSpaceUnset ST_TLAnimateColorSpace = 0
	ST_TLAnimateColorSpaceRgb   ST_TLAnimateColorSpace = 1
	ST_TLAnimateColorSpaceHsl   ST_TLAnimateColorSpace = 2
)

func (e ST_TLAnimateColorSpace) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TLAnimateColorSpaceUnset:
		attr.Value = ""
	case ST_TLAnimateColorSpaceRgb:
		attr.Value = "rgb"
	case ST_TLAnimateColorSpaceHsl:
		attr.Value = "hsl"
	}
	return attr, nil
}
func (e *ST_TLAnimateColorSpace) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "rgb":
		*e = 1
	case "hsl":
		*e = 2
	}
	return nil
}
func (m ST_TLAnimateColorSpace) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TLAnimateColorSpace) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "rgb":
			*m = 1
		case "hsl":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_TLAnimateColorSpace) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "rgb"
	case 2:
		return "hsl"
	}
	return ""
}
func (m ST_TLAnimateColorSpace) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TLAnimateColorSpace) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TLAnimateColorDirection byte

const (
	ST_TLAnimateColorDirectionUnset ST_TLAnimateColorDirection = 0
	ST_TLAnimateColorDirectionCw    ST_TLAnimateColorDirection = 1
	ST_TLAnimateColorDirectionCcw   ST_TLAnimateColorDirection = 2
)

func (e ST_TLAnimateColorDirection) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TLAnimateColorDirectionUnset:
		attr.Value = ""
	case ST_TLAnimateColorDirectionCw:
		attr.Value = "cw"
	case ST_TLAnimateColorDirectionCcw:
		attr.Value = "ccw"
	}
	return attr, nil
}
func (e *ST_TLAnimateColorDirection) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "cw":
		*e = 1
	case "ccw":
		*e = 2
	}
	return nil
}
func (m ST_TLAnimateColorDirection) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TLAnimateColorDirection) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "cw":
			*m = 1
		case "ccw":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_TLAnimateColorDirection) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "cw"
	case 2:
		return "ccw"
	}
	return ""
}
func (m ST_TLAnimateColorDirection) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TLAnimateColorDirection) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TLAnimateEffectTransition byte

const (
	ST_TLAnimateEffectTransitionUnset ST_TLAnimateEffectTransition = 0
	ST_TLAnimateEffectTransitionIn    ST_TLAnimateEffectTransition = 1
	ST_TLAnimateEffectTransitionOut   ST_TLAnimateEffectTransition = 2
	ST_TLAnimateEffectTransitionNone  ST_TLAnimateEffectTransition = 3
)

func (e ST_TLAnimateEffectTransition) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TLAnimateEffectTransitionUnset:
		attr.Value = ""
	case ST_TLAnimateEffectTransitionIn:
		attr.Value = "in"
	case ST_TLAnimateEffectTransitionOut:
		attr.Value = "out"
	case ST_TLAnimateEffectTransitionNone:
		attr.Value = "none"
	}
	return attr, nil
}
func (e *ST_TLAnimateEffectTransition) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "in":
		*e = 1
	case "out":
		*e = 2
	case "none":
		*e = 3
	}
	return nil
}
func (m ST_TLAnimateEffectTransition) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TLAnimateEffectTransition) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "in":
			*m = 1
		case "out":
			*m = 2
		case "none":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_TLAnimateEffectTransition) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "in"
	case 2:
		return "out"
	case 3:
		return "none"
	}
	return ""
}
func (m ST_TLAnimateEffectTransition) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TLAnimateEffectTransition) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TLAnimateMotionBehaviorOrigin byte

const (
	ST_TLAnimateMotionBehaviorOriginUnset  ST_TLAnimateMotionBehaviorOrigin = 0
	ST_TLAnimateMotionBehaviorOriginParent ST_TLAnimateMotionBehaviorOrigin = 1
	ST_TLAnimateMotionBehaviorOriginLayout ST_TLAnimateMotionBehaviorOrigin = 2
)

func (e ST_TLAnimateMotionBehaviorOrigin) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TLAnimateMotionBehaviorOriginUnset:
		attr.Value = ""
	case ST_TLAnimateMotionBehaviorOriginParent:
		attr.Value = "parent"
	case ST_TLAnimateMotionBehaviorOriginLayout:
		attr.Value = "layout"
	}
	return attr, nil
}
func (e *ST_TLAnimateMotionBehaviorOrigin) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "parent":
		*e = 1
	case "layout":
		*e = 2
	}
	return nil
}
func (m ST_TLAnimateMotionBehaviorOrigin) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TLAnimateMotionBehaviorOrigin) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "parent":
			*m = 1
		case "layout":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_TLAnimateMotionBehaviorOrigin) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "parent"
	case 2:
		return "layout"
	}
	return ""
}
func (m ST_TLAnimateMotionBehaviorOrigin) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TLAnimateMotionBehaviorOrigin) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TLAnimateMotionPathEditMode byte

const (
	ST_TLAnimateMotionPathEditModeUnset    ST_TLAnimateMotionPathEditMode = 0
	ST_TLAnimateMotionPathEditModeRelative ST_TLAnimateMotionPathEditMode = 1
	ST_TLAnimateMotionPathEditModeFixed    ST_TLAnimateMotionPathEditMode = 2
)

func (e ST_TLAnimateMotionPathEditMode) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TLAnimateMotionPathEditModeUnset:
		attr.Value = ""
	case ST_TLAnimateMotionPathEditModeRelative:
		attr.Value = "relative"
	case ST_TLAnimateMotionPathEditModeFixed:
		attr.Value = "fixed"
	}
	return attr, nil
}
func (e *ST_TLAnimateMotionPathEditMode) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "relative":
		*e = 1
	case "fixed":
		*e = 2
	}
	return nil
}
func (m ST_TLAnimateMotionPathEditMode) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TLAnimateMotionPathEditMode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "relative":
			*m = 1
		case "fixed":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_TLAnimateMotionPathEditMode) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "relative"
	case 2:
		return "fixed"
	}
	return ""
}
func (m ST_TLAnimateMotionPathEditMode) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TLAnimateMotionPathEditMode) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TLCommandType byte

const (
	ST_TLCommandTypeUnset ST_TLCommandType = 0
	ST_TLCommandTypeEvt   ST_TLCommandType = 1
	ST_TLCommandTypeCall  ST_TLCommandType = 2
	ST_TLCommandTypeVerb  ST_TLCommandType = 3
)

func (e ST_TLCommandType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TLCommandTypeUnset:
		attr.Value = ""
	case ST_TLCommandTypeEvt:
		attr.Value = "evt"
	case ST_TLCommandTypeCall:
		attr.Value = "call"
	case ST_TLCommandTypeVerb:
		attr.Value = "verb"
	}
	return attr, nil
}
func (e *ST_TLCommandType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "evt":
		*e = 1
	case "call":
		*e = 2
	case "verb":
		*e = 3
	}
	return nil
}
func (m ST_TLCommandType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TLCommandType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "evt":
			*m = 1
		case "call":
			*m = 2
		case "verb":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_TLCommandType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "evt"
	case 2:
		return "call"
	case 3:
		return "verb"
	}
	return ""
}
func (m ST_TLCommandType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TLCommandType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TLParaBuildType byte

const (
	ST_TLParaBuildTypeUnset     ST_TLParaBuildType = 0
	ST_TLParaBuildTypeAllAtOnce ST_TLParaBuildType = 1
	ST_TLParaBuildTypeP         ST_TLParaBuildType = 2
	ST_TLParaBuildTypeCust      ST_TLParaBuildType = 3
	ST_TLParaBuildTypeWhole     ST_TLParaBuildType = 4
)

func (e ST_TLParaBuildType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TLParaBuildTypeUnset:
		attr.Value = ""
	case ST_TLParaBuildTypeAllAtOnce:
		attr.Value = "allAtOnce"
	case ST_TLParaBuildTypeP:
		attr.Value = "p"
	case ST_TLParaBuildTypeCust:
		attr.Value = "cust"
	case ST_TLParaBuildTypeWhole:
		attr.Value = "whole"
	}
	return attr, nil
}
func (e *ST_TLParaBuildType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "allAtOnce":
		*e = 1
	case "p":
		*e = 2
	case "cust":
		*e = 3
	case "whole":
		*e = 4
	}
	return nil
}
func (m ST_TLParaBuildType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TLParaBuildType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "allAtOnce":
			*m = 1
		case "p":
			*m = 2
		case "cust":
			*m = 3
		case "whole":
			*m = 4
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_TLParaBuildType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "allAtOnce"
	case 2:
		return "p"
	case 3:
		return "cust"
	case 4:
		return "whole"
	}
	return ""
}
func (m ST_TLParaBuildType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TLParaBuildType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TLDiagramBuildType byte

const (
	ST_TLDiagramBuildTypeUnset         ST_TLDiagramBuildType = 0
	ST_TLDiagramBuildTypeWhole         ST_TLDiagramBuildType = 1
	ST_TLDiagramBuildTypeDepthByNode   ST_TLDiagramBuildType = 2
	ST_TLDiagramBuildTypeDepthByBranch ST_TLDiagramBuildType = 3
	ST_TLDiagramBuildTypeBreadthByNode ST_TLDiagramBuildType = 4
	ST_TLDiagramBuildTypeBreadthByLvl  ST_TLDiagramBuildType = 5
	ST_TLDiagramBuildTypeCw            ST_TLDiagramBuildType = 6
	ST_TLDiagramBuildTypeCwIn          ST_TLDiagramBuildType = 7
	ST_TLDiagramBuildTypeCwOut         ST_TLDiagramBuildType = 8
	ST_TLDiagramBuildTypeCcw           ST_TLDiagramBuildType = 9
	ST_TLDiagramBuildTypeCcwIn         ST_TLDiagramBuildType = 10
	ST_TLDiagramBuildTypeCcwOut        ST_TLDiagramBuildType = 11
	ST_TLDiagramBuildTypeInByRing      ST_TLDiagramBuildType = 12
	ST_TLDiagramBuildTypeOutByRing     ST_TLDiagramBuildType = 13
	ST_TLDiagramBuildTypeUp            ST_TLDiagramBuildType = 14
	ST_TLDiagramBuildTypeDown          ST_TLDiagramBuildType = 15
	ST_TLDiagramBuildTypeAllAtOnce     ST_TLDiagramBuildType = 16
	ST_TLDiagramBuildTypeCust          ST_TLDiagramBuildType = 17
)

func (e ST_TLDiagramBuildType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TLDiagramBuildTypeUnset:
		attr.Value = ""
	case ST_TLDiagramBuildTypeWhole:
		attr.Value = "whole"
	case ST_TLDiagramBuildTypeDepthByNode:
		attr.Value = "depthByNode"
	case ST_TLDiagramBuildTypeDepthByBranch:
		attr.Value = "depthByBranch"
	case ST_TLDiagramBuildTypeBreadthByNode:
		attr.Value = "breadthByNode"
	case ST_TLDiagramBuildTypeBreadthByLvl:
		attr.Value = "breadthByLvl"
	case ST_TLDiagramBuildTypeCw:
		attr.Value = "cw"
	case ST_TLDiagramBuildTypeCwIn:
		attr.Value = "cwIn"
	case ST_TLDiagramBuildTypeCwOut:
		attr.Value = "cwOut"
	case ST_TLDiagramBuildTypeCcw:
		attr.Value = "ccw"
	case ST_TLDiagramBuildTypeCcwIn:
		attr.Value = "ccwIn"
	case ST_TLDiagramBuildTypeCcwOut:
		attr.Value = "ccwOut"
	case ST_TLDiagramBuildTypeInByRing:
		attr.Value = "inByRing"
	case ST_TLDiagramBuildTypeOutByRing:
		attr.Value = "outByRing"
	case ST_TLDiagramBuildTypeUp:
		attr.Value = "up"
	case ST_TLDiagramBuildTypeDown:
		attr.Value = "down"
	case ST_TLDiagramBuildTypeAllAtOnce:
		attr.Value = "allAtOnce"
	case ST_TLDiagramBuildTypeCust:
		attr.Value = "cust"
	}
	return attr, nil
}
func (e *ST_TLDiagramBuildType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "whole":
		*e = 1
	case "depthByNode":
		*e = 2
	case "depthByBranch":
		*e = 3
	case "breadthByNode":
		*e = 4
	case "breadthByLvl":
		*e = 5
	case "cw":
		*e = 6
	case "cwIn":
		*e = 7
	case "cwOut":
		*e = 8
	case "ccw":
		*e = 9
	case "ccwIn":
		*e = 10
	case "ccwOut":
		*e = 11
	case "inByRing":
		*e = 12
	case "outByRing":
		*e = 13
	case "up":
		*e = 14
	case "down":
		*e = 15
	case "allAtOnce":
		*e = 16
	case "cust":
		*e = 17
	}
	return nil
}
func (m ST_TLDiagramBuildType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TLDiagramBuildType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "whole":
			*m = 1
		case "depthByNode":
			*m = 2
		case "depthByBranch":
			*m = 3
		case "breadthByNode":
			*m = 4
		case "breadthByLvl":
			*m = 5
		case "cw":
			*m = 6
		case "cwIn":
			*m = 7
		case "cwOut":
			*m = 8
		case "ccw":
			*m = 9
		case "ccwIn":
			*m = 10
		case "ccwOut":
			*m = 11
		case "inByRing":
			*m = 12
		case "outByRing":
			*m = 13
		case "up":
			*m = 14
		case "down":
			*m = 15
		case "allAtOnce":
			*m = 16
		case "cust":
			*m = 17
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_TLDiagramBuildType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "whole"
	case 2:
		return "depthByNode"
	case 3:
		return "depthByBranch"
	case 4:
		return "breadthByNode"
	case 5:
		return "breadthByLvl"
	case 6:
		return "cw"
	case 7:
		return "cwIn"
	case 8:
		return "cwOut"
	case 9:
		return "ccw"
	case 10:
		return "ccwIn"
	case 11:
		return "ccwOut"
	case 12:
		return "inByRing"
	case 13:
		return "outByRing"
	case 14:
		return "up"
	case 15:
		return "down"
	case 16:
		return "allAtOnce"
	case 17:
		return "cust"
	}
	return ""
}
func (m ST_TLDiagramBuildType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TLDiagramBuildType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TLOleChartBuildType byte

const (
	ST_TLOleChartBuildTypeUnset      ST_TLOleChartBuildType = 0
	ST_TLOleChartBuildTypeAllAtOnce  ST_TLOleChartBuildType = 1
	ST_TLOleChartBuildTypeSeries     ST_TLOleChartBuildType = 2
	ST_TLOleChartBuildTypeCategory   ST_TLOleChartBuildType = 3
	ST_TLOleChartBuildTypeSeriesEl   ST_TLOleChartBuildType = 4
	ST_TLOleChartBuildTypeCategoryEl ST_TLOleChartBuildType = 5
)

func (e ST_TLOleChartBuildType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TLOleChartBuildTypeUnset:
		attr.Value = ""
	case ST_TLOleChartBuildTypeAllAtOnce:
		attr.Value = "allAtOnce"
	case ST_TLOleChartBuildTypeSeries:
		attr.Value = "series"
	case ST_TLOleChartBuildTypeCategory:
		attr.Value = "category"
	case ST_TLOleChartBuildTypeSeriesEl:
		attr.Value = "seriesEl"
	case ST_TLOleChartBuildTypeCategoryEl:
		attr.Value = "categoryEl"
	}
	return attr, nil
}
func (e *ST_TLOleChartBuildType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "allAtOnce":
		*e = 1
	case "series":
		*e = 2
	case "category":
		*e = 3
	case "seriesEl":
		*e = 4
	case "categoryEl":
		*e = 5
	}
	return nil
}
func (m ST_TLOleChartBuildType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_TLOleChartBuildType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "allAtOnce":
			*m = 1
		case "series":
			*m = 2
		case "category":
			*m = 3
		case "seriesEl":
			*m = 4
		case "categoryEl":
			*m = 5
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_TLOleChartBuildType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "allAtOnce"
	case 2:
		return "series"
	case 3:
		return "category"
	case 4:
		return "seriesEl"
	case 5:
		return "categoryEl"
	}
	return ""
}
func (m ST_TLOleChartBuildType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_TLOleChartBuildType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_Direction byte

const (
	ST_DirectionUnset ST_Direction = 0
	ST_DirectionHorz  ST_Direction = 1
	ST_DirectionVert  ST_Direction = 2
)

func (e ST_Direction) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_DirectionUnset:
		attr.Value = ""
	case ST_DirectionHorz:
		attr.Value = "horz"
	case ST_DirectionVert:
		attr.Value = "vert"
	}
	return attr, nil
}
func (e *ST_Direction) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "horz":
		*e = 1
	case "vert":
		*e = 2
	}
	return nil
}
func (m ST_Direction) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_Direction) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "horz":
			*m = 1
		case "vert":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_Direction) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "horz"
	case 2:
		return "vert"
	}
	return ""
}
func (m ST_Direction) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_Direction) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_OleObjectFollowColorScheme byte

const (
	ST_OleObjectFollowColorSchemeUnset             ST_OleObjectFollowColorScheme = 0
	ST_OleObjectFollowColorSchemeNone              ST_OleObjectFollowColorScheme = 1
	ST_OleObjectFollowColorSchemeFull              ST_OleObjectFollowColorScheme = 2
	ST_OleObjectFollowColorSchemeTextAndBackground ST_OleObjectFollowColorScheme = 3
)

func (e ST_OleObjectFollowColorScheme) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_OleObjectFollowColorSchemeUnset:
		attr.Value = ""
	case ST_OleObjectFollowColorSchemeNone:
		attr.Value = "none"
	case ST_OleObjectFollowColorSchemeFull:
		attr.Value = "full"
	case ST_OleObjectFollowColorSchemeTextAndBackground:
		attr.Value = "textAndBackground"
	}
	return attr, nil
}
func (e *ST_OleObjectFollowColorScheme) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "full":
		*e = 2
	case "textAndBackground":
		*e = 3
	}
	return nil
}
func (m ST_OleObjectFollowColorScheme) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_OleObjectFollowColorScheme) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "none":
			*m = 1
		case "full":
			*m = 2
		case "textAndBackground":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_OleObjectFollowColorScheme) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "full"
	case 3:
		return "textAndBackground"
	}
	return ""
}
func (m ST_OleObjectFollowColorScheme) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_OleObjectFollowColorScheme) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_PhotoAlbumLayout byte

const (
	ST_PhotoAlbumLayoutUnset      ST_PhotoAlbumLayout = 0
	ST_PhotoAlbumLayoutFitToSlide ST_PhotoAlbumLayout = 1
	ST_PhotoAlbumLayout1pic       ST_PhotoAlbumLayout = 2
	ST_PhotoAlbumLayout2pic       ST_PhotoAlbumLayout = 3
	ST_PhotoAlbumLayout4pic       ST_PhotoAlbumLayout = 4
	ST_PhotoAlbumLayout1picTitle  ST_PhotoAlbumLayout = 5
	ST_PhotoAlbumLayout2picTitle  ST_PhotoAlbumLayout = 6
	ST_PhotoAlbumLayout4picTitle  ST_PhotoAlbumLayout = 7
)

func (e ST_PhotoAlbumLayout) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_PhotoAlbumLayoutUnset:
		attr.Value = ""
	case ST_PhotoAlbumLayoutFitToSlide:
		attr.Value = "fitToSlide"
	case ST_PhotoAlbumLayout1pic:
		attr.Value = "1pic"
	case ST_PhotoAlbumLayout2pic:
		attr.Value = "2pic"
	case ST_PhotoAlbumLayout4pic:
		attr.Value = "4pic"
	case ST_PhotoAlbumLayout1picTitle:
		attr.Value = "1picTitle"
	case ST_PhotoAlbumLayout2picTitle:
		attr.Value = "2picTitle"
	case ST_PhotoAlbumLayout4picTitle:
		attr.Value = "4picTitle"
	}
	return attr, nil
}
func (e *ST_PhotoAlbumLayout) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "fitToSlide":
		*e = 1
	case "1pic":
		*e = 2
	case "2pic":
		*e = 3
	case "4pic":
		*e = 4
	case "1picTitle":
		*e = 5
	case "2picTitle":
		*e = 6
	case "4picTitle":
		*e = 7
	}
	return nil
}
func (m ST_PhotoAlbumLayout) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_PhotoAlbumLayout) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "fitToSlide":
			*m = 1
		case "1pic":
			*m = 2
		case "2pic":
			*m = 3
		case "4pic":
			*m = 4
		case "1picTitle":
			*m = 5
		case "2picTitle":
			*m = 6
		case "4picTitle":
			*m = 7
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_PhotoAlbumLayout) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "fitToSlide"
	case 2:
		return "1pic"
	case 3:
		return "2pic"
	case 4:
		return "4pic"
	case 5:
		return "1picTitle"
	case 6:
		return "2picTitle"
	case 7:
		return "4picTitle"
	}
	return ""
}
func (m ST_PhotoAlbumLayout) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_PhotoAlbumLayout) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_PhotoAlbumFrameShape byte

const (
	ST_PhotoAlbumFrameShapeUnset       ST_PhotoAlbumFrameShape = 0
	ST_PhotoAlbumFrameShapeFrameStyle1 ST_PhotoAlbumFrameShape = 1
	ST_PhotoAlbumFrameShapeFrameStyle2 ST_PhotoAlbumFrameShape = 2
	ST_PhotoAlbumFrameShapeFrameStyle3 ST_PhotoAlbumFrameShape = 3
	ST_PhotoAlbumFrameShapeFrameStyle4 ST_PhotoAlbumFrameShape = 4
	ST_PhotoAlbumFrameShapeFrameStyle5 ST_PhotoAlbumFrameShape = 5
	ST_PhotoAlbumFrameShapeFrameStyle6 ST_PhotoAlbumFrameShape = 6
	ST_PhotoAlbumFrameShapeFrameStyle7 ST_PhotoAlbumFrameShape = 7
)

func (e ST_PhotoAlbumFrameShape) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_PhotoAlbumFrameShapeUnset:
		attr.Value = ""
	case ST_PhotoAlbumFrameShapeFrameStyle1:
		attr.Value = "frameStyle1"
	case ST_PhotoAlbumFrameShapeFrameStyle2:
		attr.Value = "frameStyle2"
	case ST_PhotoAlbumFrameShapeFrameStyle3:
		attr.Value = "frameStyle3"
	case ST_PhotoAlbumFrameShapeFrameStyle4:
		attr.Value = "frameStyle4"
	case ST_PhotoAlbumFrameShapeFrameStyle5:
		attr.Value = "frameStyle5"
	case ST_PhotoAlbumFrameShapeFrameStyle6:
		attr.Value = "frameStyle6"
	case ST_PhotoAlbumFrameShapeFrameStyle7:
		attr.Value = "frameStyle7"
	}
	return attr, nil
}
func (e *ST_PhotoAlbumFrameShape) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "frameStyle1":
		*e = 1
	case "frameStyle2":
		*e = 2
	case "frameStyle3":
		*e = 3
	case "frameStyle4":
		*e = 4
	case "frameStyle5":
		*e = 5
	case "frameStyle6":
		*e = 6
	case "frameStyle7":
		*e = 7
	}
	return nil
}
func (m ST_PhotoAlbumFrameShape) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_PhotoAlbumFrameShape) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "frameStyle1":
			*m = 1
		case "frameStyle2":
			*m = 2
		case "frameStyle3":
			*m = 3
		case "frameStyle4":
			*m = 4
		case "frameStyle5":
			*m = 5
		case "frameStyle6":
			*m = 6
		case "frameStyle7":
			*m = 7
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_PhotoAlbumFrameShape) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "frameStyle1"
	case 2:
		return "frameStyle2"
	case 3:
		return "frameStyle3"
	case 4:
		return "frameStyle4"
	case 5:
		return "frameStyle5"
	case 6:
		return "frameStyle6"
	case 7:
		return "frameStyle7"
	}
	return ""
}
func (m ST_PhotoAlbumFrameShape) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_PhotoAlbumFrameShape) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_SlideSizeType byte

const (
	ST_SlideSizeTypeUnset       ST_SlideSizeType = 0
	ST_SlideSizeTypeScreen4x3   ST_SlideSizeType = 1
	ST_SlideSizeTypeLetter      ST_SlideSizeType = 2
	ST_SlideSizeTypeA4          ST_SlideSizeType = 3
	ST_SlideSizeType35mm        ST_SlideSizeType = 4
	ST_SlideSizeTypeOverhead    ST_SlideSizeType = 5
	ST_SlideSizeTypeBanner      ST_SlideSizeType = 6
	ST_SlideSizeTypeCustom      ST_SlideSizeType = 7
	ST_SlideSizeTypeLedger      ST_SlideSizeType = 8
	ST_SlideSizeTypeA3          ST_SlideSizeType = 9
	ST_SlideSizeTypeB4ISO       ST_SlideSizeType = 10
	ST_SlideSizeTypeB5ISO       ST_SlideSizeType = 11
	ST_SlideSizeTypeB4JIS       ST_SlideSizeType = 12
	ST_SlideSizeTypeB5JIS       ST_SlideSizeType = 13
	ST_SlideSizeTypeHagakiCard  ST_SlideSizeType = 14
	ST_SlideSizeTypeScreen16x9  ST_SlideSizeType = 15
	ST_SlideSizeTypeScreen16x10 ST_SlideSizeType = 16
)

func (e ST_SlideSizeType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_SlideSizeTypeUnset:
		attr.Value = ""
	case ST_SlideSizeTypeScreen4x3:
		attr.Value = "screen4x3"
	case ST_SlideSizeTypeLetter:
		attr.Value = "letter"
	case ST_SlideSizeTypeA4:
		attr.Value = "A4"
	case ST_SlideSizeType35mm:
		attr.Value = "35mm"
	case ST_SlideSizeTypeOverhead:
		attr.Value = "overhead"
	case ST_SlideSizeTypeBanner:
		attr.Value = "banner"
	case ST_SlideSizeTypeCustom:
		attr.Value = "custom"
	case ST_SlideSizeTypeLedger:
		attr.Value = "ledger"
	case ST_SlideSizeTypeA3:
		attr.Value = "A3"
	case ST_SlideSizeTypeB4ISO:
		attr.Value = "B4ISO"
	case ST_SlideSizeTypeB5ISO:
		attr.Value = "B5ISO"
	case ST_SlideSizeTypeB4JIS:
		attr.Value = "B4JIS"
	case ST_SlideSizeTypeB5JIS:
		attr.Value = "B5JIS"
	case ST_SlideSizeTypeHagakiCard:
		attr.Value = "hagakiCard"
	case ST_SlideSizeTypeScreen16x9:
		attr.Value = "screen16x9"
	case ST_SlideSizeTypeScreen16x10:
		attr.Value = "screen16x10"
	}
	return attr, nil
}
func (e *ST_SlideSizeType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "screen4x3":
		*e = 1
	case "letter":
		*e = 2
	case "A4":
		*e = 3
	case "35mm":
		*e = 4
	case "overhead":
		*e = 5
	case "banner":
		*e = 6
	case "custom":
		*e = 7
	case "ledger":
		*e = 8
	case "A3":
		*e = 9
	case "B4ISO":
		*e = 10
	case "B5ISO":
		*e = 11
	case "B4JIS":
		*e = 12
	case "B5JIS":
		*e = 13
	case "hagakiCard":
		*e = 14
	case "screen16x9":
		*e = 15
	case "screen16x10":
		*e = 16
	}
	return nil
}
func (m ST_SlideSizeType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_SlideSizeType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "screen4x3":
			*m = 1
		case "letter":
			*m = 2
		case "A4":
			*m = 3
		case "35mm":
			*m = 4
		case "overhead":
			*m = 5
		case "banner":
			*m = 6
		case "custom":
			*m = 7
		case "ledger":
			*m = 8
		case "A3":
			*m = 9
		case "B4ISO":
			*m = 10
		case "B5ISO":
			*m = 11
		case "B4JIS":
			*m = 12
		case "B5JIS":
			*m = 13
		case "hagakiCard":
			*m = 14
		case "screen16x9":
			*m = 15
		case "screen16x10":
			*m = 16
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_SlideSizeType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "screen4x3"
	case 2:
		return "letter"
	case 3:
		return "A4"
	case 4:
		return "35mm"
	case 5:
		return "overhead"
	case 6:
		return "banner"
	case 7:
		return "custom"
	case 8:
		return "ledger"
	case 9:
		return "A3"
	case 10:
		return "B4ISO"
	case 11:
		return "B5ISO"
	case 12:
		return "B4JIS"
	case 13:
		return "B5JIS"
	case 14:
		return "hagakiCard"
	case 15:
		return "screen16x9"
	case 16:
		return "screen16x10"
	}
	return ""
}
func (m ST_SlideSizeType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_SlideSizeType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_WebColorType byte

const (
	ST_WebColorTypeUnset              ST_WebColorType = 0
	ST_WebColorTypeNone               ST_WebColorType = 1
	ST_WebColorTypeBrowser            ST_WebColorType = 2
	ST_WebColorTypePresentationText   ST_WebColorType = 3
	ST_WebColorTypePresentationAccent ST_WebColorType = 4
	ST_WebColorTypeWhiteTextOnBlack   ST_WebColorType = 5
	ST_WebColorTypeBlackTextOnWhite   ST_WebColorType = 6
)

func (e ST_WebColorType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_WebColorTypeUnset:
		attr.Value = ""
	case ST_WebColorTypeNone:
		attr.Value = "none"
	case ST_WebColorTypeBrowser:
		attr.Value = "browser"
	case ST_WebColorTypePresentationText:
		attr.Value = "presentationText"
	case ST_WebColorTypePresentationAccent:
		attr.Value = "presentationAccent"
	case ST_WebColorTypeWhiteTextOnBlack:
		attr.Value = "whiteTextOnBlack"
	case ST_WebColorTypeBlackTextOnWhite:
		attr.Value = "blackTextOnWhite"
	}
	return attr, nil
}
func (e *ST_WebColorType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "browser":
		*e = 2
	case "presentationText":
		*e = 3
	case "presentationAccent":
		*e = 4
	case "whiteTextOnBlack":
		*e = 5
	case "blackTextOnWhite":
		*e = 6
	}
	return nil
}
func (m ST_WebColorType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_WebColorType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "none":
			*m = 1
		case "browser":
			*m = 2
		case "presentationText":
			*m = 3
		case "presentationAccent":
			*m = 4
		case "whiteTextOnBlack":
			*m = 5
		case "blackTextOnWhite":
			*m = 6
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_WebColorType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "browser"
	case 3:
		return "presentationText"
	case 4:
		return "presentationAccent"
	case 5:
		return "whiteTextOnBlack"
	case 6:
		return "blackTextOnWhite"
	}
	return ""
}
func (m ST_WebColorType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_WebColorType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_WebScreenSize byte

const (
	ST_WebScreenSizeUnset     ST_WebScreenSize = 0
	ST_WebScreenSize544x376   ST_WebScreenSize = 1
	ST_WebScreenSize640x480   ST_WebScreenSize = 2
	ST_WebScreenSize720x512   ST_WebScreenSize = 3
	ST_WebScreenSize800x600   ST_WebScreenSize = 4
	ST_WebScreenSize1024x768  ST_WebScreenSize = 5
	ST_WebScreenSize1152x882  ST_WebScreenSize = 6
	ST_WebScreenSize1152x900  ST_WebScreenSize = 7
	ST_WebScreenSize1280x1024 ST_WebScreenSize = 8
	ST_WebScreenSize1600x1200 ST_WebScreenSize = 9
	ST_WebScreenSize1800x1400 ST_WebScreenSize = 10
	ST_WebScreenSize1920x1200 ST_WebScreenSize = 11
)

func (e ST_WebScreenSize) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_WebScreenSizeUnset:
		attr.Value = ""
	case ST_WebScreenSize544x376:
		attr.Value = "544x376"
	case ST_WebScreenSize640x480:
		attr.Value = "640x480"
	case ST_WebScreenSize720x512:
		attr.Value = "720x512"
	case ST_WebScreenSize800x600:
		attr.Value = "800x600"
	case ST_WebScreenSize1024x768:
		attr.Value = "1024x768"
	case ST_WebScreenSize1152x882:
		attr.Value = "1152x882"
	case ST_WebScreenSize1152x900:
		attr.Value = "1152x900"
	case ST_WebScreenSize1280x1024:
		attr.Value = "1280x1024"
	case ST_WebScreenSize1600x1200:
		attr.Value = "1600x1200"
	case ST_WebScreenSize1800x1400:
		attr.Value = "1800x1400"
	case ST_WebScreenSize1920x1200:
		attr.Value = "1920x1200"
	}
	return attr, nil
}
func (e *ST_WebScreenSize) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "544x376":
		*e = 1
	case "640x480":
		*e = 2
	case "720x512":
		*e = 3
	case "800x600":
		*e = 4
	case "1024x768":
		*e = 5
	case "1152x882":
		*e = 6
	case "1152x900":
		*e = 7
	case "1280x1024":
		*e = 8
	case "1600x1200":
		*e = 9
	case "1800x1400":
		*e = 10
	case "1920x1200":
		*e = 11
	}
	return nil
}
func (m ST_WebScreenSize) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_WebScreenSize) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "544x376":
			*m = 1
		case "640x480":
			*m = 2
		case "720x512":
			*m = 3
		case "800x600":
			*m = 4
		case "1024x768":
			*m = 5
		case "1152x882":
			*m = 6
		case "1152x900":
			*m = 7
		case "1280x1024":
			*m = 8
		case "1600x1200":
			*m = 9
		case "1800x1400":
			*m = 10
		case "1920x1200":
			*m = 11
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_WebScreenSize) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "544x376"
	case 2:
		return "640x480"
	case 3:
		return "720x512"
	case 4:
		return "800x600"
	case 5:
		return "1024x768"
	case 6:
		return "1152x882"
	case 7:
		return "1152x900"
	case 8:
		return "1280x1024"
	case 9:
		return "1600x1200"
	case 10:
		return "1800x1400"
	case 11:
		return "1920x1200"
	}
	return ""
}
func (m ST_WebScreenSize) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_WebScreenSize) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_PrintWhat byte

const (
	ST_PrintWhatUnset     ST_PrintWhat = 0
	ST_PrintWhatSlides    ST_PrintWhat = 1
	ST_PrintWhatHandouts1 ST_PrintWhat = 2
	ST_PrintWhatHandouts2 ST_PrintWhat = 3
	ST_PrintWhatHandouts3 ST_PrintWhat = 4
	ST_PrintWhatHandouts4 ST_PrintWhat = 5
	ST_PrintWhatHandouts6 ST_PrintWhat = 6
	ST_PrintWhatHandouts9 ST_PrintWhat = 7
	ST_PrintWhatNotes     ST_PrintWhat = 8
	ST_PrintWhatOutline   ST_PrintWhat = 9
)

func (e ST_PrintWhat) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_PrintWhatUnset:
		attr.Value = ""
	case ST_PrintWhatSlides:
		attr.Value = "slides"
	case ST_PrintWhatHandouts1:
		attr.Value = "handouts1"
	case ST_PrintWhatHandouts2:
		attr.Value = "handouts2"
	case ST_PrintWhatHandouts3:
		attr.Value = "handouts3"
	case ST_PrintWhatHandouts4:
		attr.Value = "handouts4"
	case ST_PrintWhatHandouts6:
		attr.Value = "handouts6"
	case ST_PrintWhatHandouts9:
		attr.Value = "handouts9"
	case ST_PrintWhatNotes:
		attr.Value = "notes"
	case ST_PrintWhatOutline:
		attr.Value = "outline"
	}
	return attr, nil
}
func (e *ST_PrintWhat) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "slides":
		*e = 1
	case "handouts1":
		*e = 2
	case "handouts2":
		*e = 3
	case "handouts3":
		*e = 4
	case "handouts4":
		*e = 5
	case "handouts6":
		*e = 6
	case "handouts9":
		*e = 7
	case "notes":
		*e = 8
	case "outline":
		*e = 9
	}
	return nil
}
func (m ST_PrintWhat) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_PrintWhat) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "slides":
			*m = 1
		case "handouts1":
			*m = 2
		case "handouts2":
			*m = 3
		case "handouts3":
			*m = 4
		case "handouts4":
			*m = 5
		case "handouts6":
			*m = 6
		case "handouts9":
			*m = 7
		case "notes":
			*m = 8
		case "outline":
			*m = 9
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_PrintWhat) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "slides"
	case 2:
		return "handouts1"
	case 3:
		return "handouts2"
	case 4:
		return "handouts3"
	case 5:
		return "handouts4"
	case 6:
		return "handouts6"
	case 7:
		return "handouts9"
	case 8:
		return "notes"
	case 9:
		return "outline"
	}
	return ""
}
func (m ST_PrintWhat) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_PrintWhat) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_PrintColorMode byte

const (
	ST_PrintColorModeUnset ST_PrintColorMode = 0
	ST_PrintColorModeBw    ST_PrintColorMode = 1
	ST_PrintColorModeGray  ST_PrintColorMode = 2
	ST_PrintColorModeClr   ST_PrintColorMode = 3
)

func (e ST_PrintColorMode) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_PrintColorModeUnset:
		attr.Value = ""
	case ST_PrintColorModeBw:
		attr.Value = "bw"
	case ST_PrintColorModeGray:
		attr.Value = "gray"
	case ST_PrintColorModeClr:
		attr.Value = "clr"
	}
	return attr, nil
}
func (e *ST_PrintColorMode) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "bw":
		*e = 1
	case "gray":
		*e = 2
	case "clr":
		*e = 3
	}
	return nil
}
func (m ST_PrintColorMode) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_PrintColorMode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "bw":
			*m = 1
		case "gray":
			*m = 2
		case "clr":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_PrintColorMode) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "bw"
	case 2:
		return "gray"
	case 3:
		return "clr"
	}
	return ""
}
func (m ST_PrintColorMode) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_PrintColorMode) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_PlaceholderType byte

const (
	ST_PlaceholderTypeUnset    ST_PlaceholderType = 0
	ST_PlaceholderTypeTitle    ST_PlaceholderType = 1
	ST_PlaceholderTypeBody     ST_PlaceholderType = 2
	ST_PlaceholderTypeCtrTitle ST_PlaceholderType = 3
	ST_PlaceholderTypeSubTitle ST_PlaceholderType = 4
	ST_PlaceholderTypeDt       ST_PlaceholderType = 5
	ST_PlaceholderTypeSldNum   ST_PlaceholderType = 6
	ST_PlaceholderTypeFtr      ST_PlaceholderType = 7
	ST_PlaceholderTypeHdr      ST_PlaceholderType = 8
	ST_PlaceholderTypeObj      ST_PlaceholderType = 9
	ST_PlaceholderTypeChart    ST_PlaceholderType = 10
	ST_PlaceholderTypeTbl      ST_PlaceholderType = 11
	ST_PlaceholderTypeClipArt  ST_PlaceholderType = 12
	ST_PlaceholderTypeDgm      ST_PlaceholderType = 13
	ST_PlaceholderTypeMedia    ST_PlaceholderType = 14
	ST_PlaceholderTypeSldImg   ST_PlaceholderType = 15
	ST_PlaceholderTypePic      ST_PlaceholderType = 16
)

func (e ST_PlaceholderType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_PlaceholderTypeUnset:
		attr.Value = ""
	case ST_PlaceholderTypeTitle:
		attr.Value = "title"
	case ST_PlaceholderTypeBody:
		attr.Value = "body"
	case ST_PlaceholderTypeCtrTitle:
		attr.Value = "ctrTitle"
	case ST_PlaceholderTypeSubTitle:
		attr.Value = "subTitle"
	case ST_PlaceholderTypeDt:
		attr.Value = "dt"
	case ST_PlaceholderTypeSldNum:
		attr.Value = "sldNum"
	case ST_PlaceholderTypeFtr:
		attr.Value = "ftr"
	case ST_PlaceholderTypeHdr:
		attr.Value = "hdr"
	case ST_PlaceholderTypeObj:
		attr.Value = "obj"
	case ST_PlaceholderTypeChart:
		attr.Value = "chart"
	case ST_PlaceholderTypeTbl:
		attr.Value = "tbl"
	case ST_PlaceholderTypeClipArt:
		attr.Value = "clipArt"
	case ST_PlaceholderTypeDgm:
		attr.Value = "dgm"
	case ST_PlaceholderTypeMedia:
		attr.Value = "media"
	case ST_PlaceholderTypeSldImg:
		attr.Value = "sldImg"
	case ST_PlaceholderTypePic:
		attr.Value = "pic"
	}
	return attr, nil
}
func (e *ST_PlaceholderType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "title":
		*e = 1
	case "body":
		*e = 2
	case "ctrTitle":
		*e = 3
	case "subTitle":
		*e = 4
	case "dt":
		*e = 5
	case "sldNum":
		*e = 6
	case "ftr":
		*e = 7
	case "hdr":
		*e = 8
	case "obj":
		*e = 9
	case "chart":
		*e = 10
	case "tbl":
		*e = 11
	case "clipArt":
		*e = 12
	case "dgm":
		*e = 13
	case "media":
		*e = 14
	case "sldImg":
		*e = 15
	case "pic":
		*e = 16
	}
	return nil
}
func (m ST_PlaceholderType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_PlaceholderType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "title":
			*m = 1
		case "body":
			*m = 2
		case "ctrTitle":
			*m = 3
		case "subTitle":
			*m = 4
		case "dt":
			*m = 5
		case "sldNum":
			*m = 6
		case "ftr":
			*m = 7
		case "hdr":
			*m = 8
		case "obj":
			*m = 9
		case "chart":
			*m = 10
		case "tbl":
			*m = 11
		case "clipArt":
			*m = 12
		case "dgm":
			*m = 13
		case "media":
			*m = 14
		case "sldImg":
			*m = 15
		case "pic":
			*m = 16
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_PlaceholderType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "title"
	case 2:
		return "body"
	case 3:
		return "ctrTitle"
	case 4:
		return "subTitle"
	case 5:
		return "dt"
	case 6:
		return "sldNum"
	case 7:
		return "ftr"
	case 8:
		return "hdr"
	case 9:
		return "obj"
	case 10:
		return "chart"
	case 11:
		return "tbl"
	case 12:
		return "clipArt"
	case 13:
		return "dgm"
	case 14:
		return "media"
	case 15:
		return "sldImg"
	case 16:
		return "pic"
	}
	return ""
}
func (m ST_PlaceholderType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_PlaceholderType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_PlaceholderSize byte

const (
	ST_PlaceholderSizeUnset   ST_PlaceholderSize = 0
	ST_PlaceholderSizeFull    ST_PlaceholderSize = 1
	ST_PlaceholderSizeHalf    ST_PlaceholderSize = 2
	ST_PlaceholderSizeQuarter ST_PlaceholderSize = 3
)

func (e ST_PlaceholderSize) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_PlaceholderSizeUnset:
		attr.Value = ""
	case ST_PlaceholderSizeFull:
		attr.Value = "full"
	case ST_PlaceholderSizeHalf:
		attr.Value = "half"
	case ST_PlaceholderSizeQuarter:
		attr.Value = "quarter"
	}
	return attr, nil
}
func (e *ST_PlaceholderSize) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "full":
		*e = 1
	case "half":
		*e = 2
	case "quarter":
		*e = 3
	}
	return nil
}
func (m ST_PlaceholderSize) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_PlaceholderSize) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "full":
			*m = 1
		case "half":
			*m = 2
		case "quarter":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_PlaceholderSize) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "full"
	case 2:
		return "half"
	case 3:
		return "quarter"
	}
	return ""
}
func (m ST_PlaceholderSize) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_PlaceholderSize) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_SlideLayoutType byte

const (
	ST_SlideLayoutTypeUnset                   ST_SlideLayoutType = 0
	ST_SlideLayoutTypeTitle                   ST_SlideLayoutType = 1
	ST_SlideLayoutTypeTx                      ST_SlideLayoutType = 2
	ST_SlideLayoutTypeTwoColTx                ST_SlideLayoutType = 3
	ST_SlideLayoutTypeTbl                     ST_SlideLayoutType = 4
	ST_SlideLayoutTypeTxAndChart              ST_SlideLayoutType = 5
	ST_SlideLayoutTypeChartAndTx              ST_SlideLayoutType = 6
	ST_SlideLayoutTypeDgm                     ST_SlideLayoutType = 7
	ST_SlideLayoutTypeChart                   ST_SlideLayoutType = 8
	ST_SlideLayoutTypeTxAndClipArt            ST_SlideLayoutType = 9
	ST_SlideLayoutTypeClipArtAndTx            ST_SlideLayoutType = 10
	ST_SlideLayoutTypeTitleOnly               ST_SlideLayoutType = 11
	ST_SlideLayoutTypeBlank                   ST_SlideLayoutType = 12
	ST_SlideLayoutTypeTxAndObj                ST_SlideLayoutType = 13
	ST_SlideLayoutTypeObjAndTx                ST_SlideLayoutType = 14
	ST_SlideLayoutTypeObjOnly                 ST_SlideLayoutType = 15
	ST_SlideLayoutTypeObj                     ST_SlideLayoutType = 16
	ST_SlideLayoutTypeTxAndMedia              ST_SlideLayoutType = 17
	ST_SlideLayoutTypeMediaAndTx              ST_SlideLayoutType = 18
	ST_SlideLayoutTypeObjOverTx               ST_SlideLayoutType = 19
	ST_SlideLayoutTypeTxOverObj               ST_SlideLayoutType = 20
	ST_SlideLayoutTypeTxAndTwoObj             ST_SlideLayoutType = 21
	ST_SlideLayoutTypeTwoObjAndTx             ST_SlideLayoutType = 22
	ST_SlideLayoutTypeTwoObjOverTx            ST_SlideLayoutType = 23
	ST_SlideLayoutTypeFourObj                 ST_SlideLayoutType = 24
	ST_SlideLayoutTypeVertTx                  ST_SlideLayoutType = 25
	ST_SlideLayoutTypeClipArtAndVertTx        ST_SlideLayoutType = 26
	ST_SlideLayoutTypeVertTitleAndTx          ST_SlideLayoutType = 27
	ST_SlideLayoutTypeVertTitleAndTxOverChart ST_SlideLayoutType = 28
	ST_SlideLayoutTypeTwoObj                  ST_SlideLayoutType = 29
	ST_SlideLayoutTypeObjAndTwoObj            ST_SlideLayoutType = 30
	ST_SlideLayoutTypeTwoObjAndObj            ST_SlideLayoutType = 31
	ST_SlideLayoutTypeCust                    ST_SlideLayoutType = 32
	ST_SlideLayoutTypeSecHead                 ST_SlideLayoutType = 33
	ST_SlideLayoutTypeTwoTxTwoObj             ST_SlideLayoutType = 34
	ST_SlideLayoutTypeObjTx                   ST_SlideLayoutType = 35
	ST_SlideLayoutTypePicTx                   ST_SlideLayoutType = 36
)

func (e ST_SlideLayoutType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_SlideLayoutTypeUnset:
		attr.Value = ""
	case ST_SlideLayoutTypeTitle:
		attr.Value = "title"
	case ST_SlideLayoutTypeTx:
		attr.Value = "tx"
	case ST_SlideLayoutTypeTwoColTx:
		attr.Value = "twoColTx"
	case ST_SlideLayoutTypeTbl:
		attr.Value = "tbl"
	case ST_SlideLayoutTypeTxAndChart:
		attr.Value = "txAndChart"
	case ST_SlideLayoutTypeChartAndTx:
		attr.Value = "chartAndTx"
	case ST_SlideLayoutTypeDgm:
		attr.Value = "dgm"
	case ST_SlideLayoutTypeChart:
		attr.Value = "chart"
	case ST_SlideLayoutTypeTxAndClipArt:
		attr.Value = "txAndClipArt"
	case ST_SlideLayoutTypeClipArtAndTx:
		attr.Value = "clipArtAndTx"
	case ST_SlideLayoutTypeTitleOnly:
		attr.Value = "titleOnly"
	case ST_SlideLayoutTypeBlank:
		attr.Value = "blank"
	case ST_SlideLayoutTypeTxAndObj:
		attr.Value = "txAndObj"
	case ST_SlideLayoutTypeObjAndTx:
		attr.Value = "objAndTx"
	case ST_SlideLayoutTypeObjOnly:
		attr.Value = "objOnly"
	case ST_SlideLayoutTypeObj:
		attr.Value = "obj"
	case ST_SlideLayoutTypeTxAndMedia:
		attr.Value = "txAndMedia"
	case ST_SlideLayoutTypeMediaAndTx:
		attr.Value = "mediaAndTx"
	case ST_SlideLayoutTypeObjOverTx:
		attr.Value = "objOverTx"
	case ST_SlideLayoutTypeTxOverObj:
		attr.Value = "txOverObj"
	case ST_SlideLayoutTypeTxAndTwoObj:
		attr.Value = "txAndTwoObj"
	case ST_SlideLayoutTypeTwoObjAndTx:
		attr.Value = "twoObjAndTx"
	case ST_SlideLayoutTypeTwoObjOverTx:
		attr.Value = "twoObjOverTx"
	case ST_SlideLayoutTypeFourObj:
		attr.Value = "fourObj"
	case ST_SlideLayoutTypeVertTx:
		attr.Value = "vertTx"
	case ST_SlideLayoutTypeClipArtAndVertTx:
		attr.Value = "clipArtAndVertTx"
	case ST_SlideLayoutTypeVertTitleAndTx:
		attr.Value = "vertTitleAndTx"
	case ST_SlideLayoutTypeVertTitleAndTxOverChart:
		attr.Value = "vertTitleAndTxOverChart"
	case ST_SlideLayoutTypeTwoObj:
		attr.Value = "twoObj"
	case ST_SlideLayoutTypeObjAndTwoObj:
		attr.Value = "objAndTwoObj"
	case ST_SlideLayoutTypeTwoObjAndObj:
		attr.Value = "twoObjAndObj"
	case ST_SlideLayoutTypeCust:
		attr.Value = "cust"
	case ST_SlideLayoutTypeSecHead:
		attr.Value = "secHead"
	case ST_SlideLayoutTypeTwoTxTwoObj:
		attr.Value = "twoTxTwoObj"
	case ST_SlideLayoutTypeObjTx:
		attr.Value = "objTx"
	case ST_SlideLayoutTypePicTx:
		attr.Value = "picTx"
	}
	return attr, nil
}
func (e *ST_SlideLayoutType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "title":
		*e = 1
	case "tx":
		*e = 2
	case "twoColTx":
		*e = 3
	case "tbl":
		*e = 4
	case "txAndChart":
		*e = 5
	case "chartAndTx":
		*e = 6
	case "dgm":
		*e = 7
	case "chart":
		*e = 8
	case "txAndClipArt":
		*e = 9
	case "clipArtAndTx":
		*e = 10
	case "titleOnly":
		*e = 11
	case "blank":
		*e = 12
	case "txAndObj":
		*e = 13
	case "objAndTx":
		*e = 14
	case "objOnly":
		*e = 15
	case "obj":
		*e = 16
	case "txAndMedia":
		*e = 17
	case "mediaAndTx":
		*e = 18
	case "objOverTx":
		*e = 19
	case "txOverObj":
		*e = 20
	case "txAndTwoObj":
		*e = 21
	case "twoObjAndTx":
		*e = 22
	case "twoObjOverTx":
		*e = 23
	case "fourObj":
		*e = 24
	case "vertTx":
		*e = 25
	case "clipArtAndVertTx":
		*e = 26
	case "vertTitleAndTx":
		*e = 27
	case "vertTitleAndTxOverChart":
		*e = 28
	case "twoObj":
		*e = 29
	case "objAndTwoObj":
		*e = 30
	case "twoObjAndObj":
		*e = 31
	case "cust":
		*e = 32
	case "secHead":
		*e = 33
	case "twoTxTwoObj":
		*e = 34
	case "objTx":
		*e = 35
	case "picTx":
		*e = 36
	}
	return nil
}
func (m ST_SlideLayoutType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_SlideLayoutType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "title":
			*m = 1
		case "tx":
			*m = 2
		case "twoColTx":
			*m = 3
		case "tbl":
			*m = 4
		case "txAndChart":
			*m = 5
		case "chartAndTx":
			*m = 6
		case "dgm":
			*m = 7
		case "chart":
			*m = 8
		case "txAndClipArt":
			*m = 9
		case "clipArtAndTx":
			*m = 10
		case "titleOnly":
			*m = 11
		case "blank":
			*m = 12
		case "txAndObj":
			*m = 13
		case "objAndTx":
			*m = 14
		case "objOnly":
			*m = 15
		case "obj":
			*m = 16
		case "txAndMedia":
			*m = 17
		case "mediaAndTx":
			*m = 18
		case "objOverTx":
			*m = 19
		case "txOverObj":
			*m = 20
		case "txAndTwoObj":
			*m = 21
		case "twoObjAndTx":
			*m = 22
		case "twoObjOverTx":
			*m = 23
		case "fourObj":
			*m = 24
		case "vertTx":
			*m = 25
		case "clipArtAndVertTx":
			*m = 26
		case "vertTitleAndTx":
			*m = 27
		case "vertTitleAndTxOverChart":
			*m = 28
		case "twoObj":
			*m = 29
		case "objAndTwoObj":
			*m = 30
		case "twoObjAndObj":
			*m = 31
		case "cust":
			*m = 32
		case "secHead":
			*m = 33
		case "twoTxTwoObj":
			*m = 34
		case "objTx":
			*m = 35
		case "picTx":
			*m = 36
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_SlideLayoutType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "title"
	case 2:
		return "tx"
	case 3:
		return "twoColTx"
	case 4:
		return "tbl"
	case 5:
		return "txAndChart"
	case 6:
		return "chartAndTx"
	case 7:
		return "dgm"
	case 8:
		return "chart"
	case 9:
		return "txAndClipArt"
	case 10:
		return "clipArtAndTx"
	case 11:
		return "titleOnly"
	case 12:
		return "blank"
	case 13:
		return "txAndObj"
	case 14:
		return "objAndTx"
	case 15:
		return "objOnly"
	case 16:
		return "obj"
	case 17:
		return "txAndMedia"
	case 18:
		return "mediaAndTx"
	case 19:
		return "objOverTx"
	case 20:
		return "txOverObj"
	case 21:
		return "txAndTwoObj"
	case 22:
		return "twoObjAndTx"
	case 23:
		return "twoObjOverTx"
	case 24:
		return "fourObj"
	case 25:
		return "vertTx"
	case 26:
		return "clipArtAndVertTx"
	case 27:
		return "vertTitleAndTx"
	case 28:
		return "vertTitleAndTxOverChart"
	case 29:
		return "twoObj"
	case 30:
		return "objAndTwoObj"
	case 31:
		return "twoObjAndObj"
	case 32:
		return "cust"
	case 33:
		return "secHead"
	case 34:
		return "twoTxTwoObj"
	case 35:
		return "objTx"
	case 36:
		return "picTx"
	}
	return ""
}
func (m ST_SlideLayoutType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_SlideLayoutType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_SplitterBarState byte

const (
	ST_SplitterBarStateUnset     ST_SplitterBarState = 0
	ST_SplitterBarStateMinimized ST_SplitterBarState = 1
	ST_SplitterBarStateRestored  ST_SplitterBarState = 2
	ST_SplitterBarStateMaximized ST_SplitterBarState = 3
)

func (e ST_SplitterBarState) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_SplitterBarStateUnset:
		attr.Value = ""
	case ST_SplitterBarStateMinimized:
		attr.Value = "minimized"
	case ST_SplitterBarStateRestored:
		attr.Value = "restored"
	case ST_SplitterBarStateMaximized:
		attr.Value = "maximized"
	}
	return attr, nil
}
func (e *ST_SplitterBarState) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "minimized":
		*e = 1
	case "restored":
		*e = 2
	case "maximized":
		*e = 3
	}
	return nil
}
func (m ST_SplitterBarState) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_SplitterBarState) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "minimized":
			*m = 1
		case "restored":
			*m = 2
		case "maximized":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_SplitterBarState) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "minimized"
	case 2:
		return "restored"
	case 3:
		return "maximized"
	}
	return ""
}
func (m ST_SplitterBarState) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_SplitterBarState) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_ViewType byte

const (
	ST_ViewTypeUnset            ST_ViewType = 0
	ST_ViewTypeSldView          ST_ViewType = 1
	ST_ViewTypeSldMasterView    ST_ViewType = 2
	ST_ViewTypeNotesView        ST_ViewType = 3
	ST_ViewTypeHandoutView      ST_ViewType = 4
	ST_ViewTypeNotesMasterView  ST_ViewType = 5
	ST_ViewTypeOutlineView      ST_ViewType = 6
	ST_ViewTypeSldSorterView    ST_ViewType = 7
	ST_ViewTypeSldThumbnailView ST_ViewType = 8
)

func (e ST_ViewType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ViewTypeUnset:
		attr.Value = ""
	case ST_ViewTypeSldView:
		attr.Value = "sldView"
	case ST_ViewTypeSldMasterView:
		attr.Value = "sldMasterView"
	case ST_ViewTypeNotesView:
		attr.Value = "notesView"
	case ST_ViewTypeHandoutView:
		attr.Value = "handoutView"
	case ST_ViewTypeNotesMasterView:
		attr.Value = "notesMasterView"
	case ST_ViewTypeOutlineView:
		attr.Value = "outlineView"
	case ST_ViewTypeSldSorterView:
		attr.Value = "sldSorterView"
	case ST_ViewTypeSldThumbnailView:
		attr.Value = "sldThumbnailView"
	}
	return attr, nil
}
func (e *ST_ViewType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "sldView":
		*e = 1
	case "sldMasterView":
		*e = 2
	case "notesView":
		*e = 3
	case "handoutView":
		*e = 4
	case "notesMasterView":
		*e = 5
	case "outlineView":
		*e = 6
	case "sldSorterView":
		*e = 7
	case "sldThumbnailView":
		*e = 8
	}
	return nil
}
func (m ST_ViewType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}
func (m *ST_ViewType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "sldView":
			*m = 1
		case "sldMasterView":
			*m = 2
		case "notesView":
			*m = 3
		case "handoutView":
			*m = 4
		case "notesMasterView":
			*m = 5
		case "outlineView":
			*m = 6
		case "sldSorterView":
			*m = 7
		case "sldThumbnailView":
			*m = 8
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}
func (m ST_ViewType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "sldView"
	case 2:
		return "sldMasterView"
	case 3:
		return "notesView"
	case 4:
		return "handoutView"
	case 5:
		return "notesMasterView"
	case 6:
		return "outlineView"
	case 7:
		return "sldSorterView"
	case 8:
		return "sldThumbnailView"
	}
	return ""
}
func (m ST_ViewType) Validate() error {
	return m.ValidateWithPath("")
}
func (m ST_ViewType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}
func init() {
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_SideDirectionTransition", NewCT_SideDirectionTransition)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_CornerDirectionTransition", NewCT_CornerDirectionTransition)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_EightDirectionTransition", NewCT_EightDirectionTransition)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_OrientationTransition", NewCT_OrientationTransition)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_InOutTransition", NewCT_InOutTransition)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_OptionalBlackTransition", NewCT_OptionalBlackTransition)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_SplitTransition", NewCT_SplitTransition)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_WheelTransition", NewCT_WheelTransition)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TransitionStartSoundAction", NewCT_TransitionStartSoundAction)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TransitionSoundAction", NewCT_TransitionSoundAction)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_SlideTransition", NewCT_SlideTransition)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLIterateIntervalTime", NewCT_TLIterateIntervalTime)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLIterateIntervalPercentage", NewCT_TLIterateIntervalPercentage)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLIterateData", NewCT_TLIterateData)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLSubShapeId", NewCT_TLSubShapeId)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLTextTargetElement", NewCT_TLTextTargetElement)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLOleChartTargetElement", NewCT_TLOleChartTargetElement)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLShapeTargetElement", NewCT_TLShapeTargetElement)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLTimeTargetElement", NewCT_TLTimeTargetElement)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLTriggerTimeNodeID", NewCT_TLTriggerTimeNodeID)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLTriggerRuntimeNode", NewCT_TLTriggerRuntimeNode)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLTimeCondition", NewCT_TLTimeCondition)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLTimeConditionList", NewCT_TLTimeConditionList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TimeNodeList", NewCT_TimeNodeList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLCommonTimeNodeData", NewCT_TLCommonTimeNodeData)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLTimeNodeParallel", NewCT_TLTimeNodeParallel)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLTimeNodeSequence", NewCT_TLTimeNodeSequence)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLTimeNodeExclusive", NewCT_TLTimeNodeExclusive)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLBehaviorAttributeNameList", NewCT_TLBehaviorAttributeNameList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLCommonBehaviorData", NewCT_TLCommonBehaviorData)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLAnimVariantBooleanVal", NewCT_TLAnimVariantBooleanVal)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLAnimVariantIntegerVal", NewCT_TLAnimVariantIntegerVal)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLAnimVariantFloatVal", NewCT_TLAnimVariantFloatVal)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLAnimVariantStringVal", NewCT_TLAnimVariantStringVal)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLAnimVariant", NewCT_TLAnimVariant)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLTimeAnimateValue", NewCT_TLTimeAnimateValue)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLTimeAnimateValueList", NewCT_TLTimeAnimateValueList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLAnimateBehavior", NewCT_TLAnimateBehavior)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLByRgbColorTransform", NewCT_TLByRgbColorTransform)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLByHslColorTransform", NewCT_TLByHslColorTransform)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLByAnimateColorTransform", NewCT_TLByAnimateColorTransform)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLAnimateColorBehavior", NewCT_TLAnimateColorBehavior)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLAnimateEffectBehavior", NewCT_TLAnimateEffectBehavior)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLPoint", NewCT_TLPoint)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLAnimateMotionBehavior", NewCT_TLAnimateMotionBehavior)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLAnimateRotationBehavior", NewCT_TLAnimateRotationBehavior)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLAnimateScaleBehavior", NewCT_TLAnimateScaleBehavior)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLCommandBehavior", NewCT_TLCommandBehavior)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLSetBehavior", NewCT_TLSetBehavior)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLCommonMediaNodeData", NewCT_TLCommonMediaNodeData)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLMediaNodeAudio", NewCT_TLMediaNodeAudio)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLMediaNodeVideo", NewCT_TLMediaNodeVideo)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLTemplate", NewCT_TLTemplate)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLTemplateList", NewCT_TLTemplateList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLBuildParagraph", NewCT_TLBuildParagraph)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLBuildDiagram", NewCT_TLBuildDiagram)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLOleBuildChart", NewCT_TLOleBuildChart)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TLGraphicalObjectBuild", NewCT_TLGraphicalObjectBuild)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_BuildList", NewCT_BuildList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_SlideTiming", NewCT_SlideTiming)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_Empty", NewCT_Empty)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_IndexRange", NewCT_IndexRange)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_SlideRelationshipListEntry", NewCT_SlideRelationshipListEntry)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_SlideRelationshipList", NewCT_SlideRelationshipList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_CustomShowId", NewCT_CustomShowId)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_CustomerData", NewCT_CustomerData)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TagsData", NewCT_TagsData)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_CustomerDataList", NewCT_CustomerDataList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_Extension", NewCT_Extension)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_ExtensionList", NewCT_ExtensionList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_ExtensionListModify", NewCT_ExtensionListModify)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_CommentAuthor", NewCT_CommentAuthor)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_CommentAuthorList", NewCT_CommentAuthorList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_Comment", NewCT_Comment)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_CommentList", NewCT_CommentList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_OleObjectEmbed", NewCT_OleObjectEmbed)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_OleObjectLink", NewCT_OleObjectLink)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_OleObject", NewCT_OleObject)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_Control", NewCT_Control)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_ControlList", NewCT_ControlList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_SlideIdListEntry", NewCT_SlideIdListEntry)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_SlideIdList", NewCT_SlideIdList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_SlideMasterIdListEntry", NewCT_SlideMasterIdListEntry)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_SlideMasterIdList", NewCT_SlideMasterIdList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_NotesMasterIdListEntry", NewCT_NotesMasterIdListEntry)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_NotesMasterIdList", NewCT_NotesMasterIdList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_HandoutMasterIdListEntry", NewCT_HandoutMasterIdListEntry)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_HandoutMasterIdList", NewCT_HandoutMasterIdList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_EmbeddedFontDataId", NewCT_EmbeddedFontDataId)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_EmbeddedFontListEntry", NewCT_EmbeddedFontListEntry)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_EmbeddedFontList", NewCT_EmbeddedFontList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_SmartTags", NewCT_SmartTags)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_CustomShow", NewCT_CustomShow)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_CustomShowList", NewCT_CustomShowList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_PhotoAlbum", NewCT_PhotoAlbum)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_SlideSize", NewCT_SlideSize)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_Kinsoku", NewCT_Kinsoku)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_ModifyVerifier", NewCT_ModifyVerifier)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_Presentation", NewCT_Presentation)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_HtmlPublishProperties", NewCT_HtmlPublishProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_WebProperties", NewCT_WebProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_PrintProperties", NewCT_PrintProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_ShowInfoBrowse", NewCT_ShowInfoBrowse)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_ShowInfoKiosk", NewCT_ShowInfoKiosk)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_ShowProperties", NewCT_ShowProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_PresentationProperties", NewCT_PresentationProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_HeaderFooter", NewCT_HeaderFooter)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_Placeholder", NewCT_Placeholder)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_ApplicationNonVisualDrawingProps", NewCT_ApplicationNonVisualDrawingProps)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_ShapeNonVisual", NewCT_ShapeNonVisual)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_Shape", NewCT_Shape)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_ConnectorNonVisual", NewCT_ConnectorNonVisual)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_Connector", NewCT_Connector)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_PictureNonVisual", NewCT_PictureNonVisual)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_Picture", NewCT_Picture)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_GraphicalObjectFrameNonVisual", NewCT_GraphicalObjectFrameNonVisual)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_GraphicalObjectFrame", NewCT_GraphicalObjectFrame)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_GroupShapeNonVisual", NewCT_GroupShapeNonVisual)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_GroupShape", NewCT_GroupShape)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_Rel", NewCT_Rel)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_BackgroundProperties", NewCT_BackgroundProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_Background", NewCT_Background)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_CommonSlideData", NewCT_CommonSlideData)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_Slide", NewCT_Slide)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_SlideLayout", NewCT_SlideLayout)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_SlideMasterTextStyles", NewCT_SlideMasterTextStyles)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_SlideLayoutIdListEntry", NewCT_SlideLayoutIdListEntry)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_SlideLayoutIdList", NewCT_SlideLayoutIdList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_SlideMaster", NewCT_SlideMaster)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_HandoutMaster", NewCT_HandoutMaster)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_NotesMaster", NewCT_NotesMaster)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_NotesSlide", NewCT_NotesSlide)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_SlideSyncProperties", NewCT_SlideSyncProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_StringTag", NewCT_StringTag)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_TagList", NewCT_TagList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_NormalViewPortion", NewCT_NormalViewPortion)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_NormalViewProperties", NewCT_NormalViewProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_CommonViewProperties", NewCT_CommonViewProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_NotesTextViewProperties", NewCT_NotesTextViewProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_OutlineViewSlideEntry", NewCT_OutlineViewSlideEntry)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_OutlineViewSlideList", NewCT_OutlineViewSlideList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_OutlineViewProperties", NewCT_OutlineViewProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_SlideSorterViewProperties", NewCT_SlideSorterViewProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_Guide", NewCT_Guide)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_GuideList", NewCT_GuideList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_CommonSlideViewProperties", NewCT_CommonSlideViewProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_SlideViewProperties", NewCT_SlideViewProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_NotesViewProperties", NewCT_NotesViewProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "CT_ViewProperties", NewCT_ViewProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "cmAuthorLst", NewCmAuthorLst)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "cmLst", NewCmLst)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "oleObj", NewOleObj)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "presentation", NewPresentation)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "presentationPr", NewPresentationPr)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "sld", NewSld)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "sldLayout", NewSldLayout)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "sldMaster", NewSldMaster)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "handoutMaster", NewHandoutMaster)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "notesMaster", NewNotesMaster)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "notes", NewNotes)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "sldSyncPr", NewSldSyncPr)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "tagLst", NewTagLst)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "viewPr", NewViewPr)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "EG_SlideListChoice", NewEG_SlideListChoice)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "EG_ExtensionList", NewEG_ExtensionList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "EG_ShowType", NewEG_ShowType)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "EG_TopLevelSlide", NewEG_TopLevelSlide)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "EG_ChildSlide", NewEG_ChildSlide)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "EG_Background", NewEG_Background)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "AG_TLBuild", NewAG_TLBuild)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "AG_Ole", NewAG_Ole)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/presentationml/2006/main", "AG_ChildSlide", NewAG_ChildSlide)
}