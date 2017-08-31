// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"fmt"
)

// ST_DepthPercent is a union type
type ST_DepthPercent struct {
	ST_DepthPercentWithSymbol *string
	ST_DepthPercentUShort     *uint16
}

func (m *ST_DepthPercent) Validate() error {
	return m.ValidateWithPath("")
}
func (m *ST_DepthPercent) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_DepthPercentWithSymbol != nil {
		mems = append(mems, "ST_DepthPercentWithSymbol")
	}
	if m.ST_DepthPercentUShort != nil {
		mems = append(mems, "ST_DepthPercentUShort")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}
func (m ST_DepthPercent) String() string {
	if m.ST_DepthPercentWithSymbol != nil {
		return fmt.Sprintf("%v", *m.ST_DepthPercentWithSymbol)
	}
	if m.ST_DepthPercentUShort != nil {
		return fmt.Sprintf("%v", *m.ST_DepthPercentUShort)
	}
	return ""
}