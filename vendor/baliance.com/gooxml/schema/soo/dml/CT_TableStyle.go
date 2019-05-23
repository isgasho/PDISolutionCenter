// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package dml

import (
	"encoding/xml"
	"fmt"

	"baliance.com/gooxml"
	"baliance.com/gooxml/schema/soo/ofc/sharedTypes"
)

type CT_TableStyle struct {
	StyleIdAttr   string
	StyleNameAttr string
	TblBg         *CT_TableBackgroundStyle
	WholeTbl      *CT_TablePartStyle
	Band1H        *CT_TablePartStyle
	Band2H        *CT_TablePartStyle
	Band1V        *CT_TablePartStyle
	Band2V        *CT_TablePartStyle
	LastCol       *CT_TablePartStyle
	FirstCol      *CT_TablePartStyle
	LastRow       *CT_TablePartStyle
	SeCell        *CT_TablePartStyle
	SwCell        *CT_TablePartStyle
	FirstRow      *CT_TablePartStyle
	NeCell        *CT_TablePartStyle
	NwCell        *CT_TablePartStyle
	ExtLst        *CT_OfficeArtExtensionList
}

func NewCT_TableStyle() *CT_TableStyle {
	ret := &CT_TableStyle{}
	ret.StyleIdAttr = "{00000000-0000-0000-0000-000000000000}"
	return ret
}

func (m *CT_TableStyle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "styleId"},
		Value: fmt.Sprintf("%v", m.StyleIdAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "styleName"},
		Value: fmt.Sprintf("%v", m.StyleNameAttr)})
	e.EncodeToken(start)
	if m.TblBg != nil {
		setblBg := xml.StartElement{Name: xml.Name{Local: "a:tblBg"}}
		e.EncodeElement(m.TblBg, setblBg)
	}
	if m.WholeTbl != nil {
		sewholeTbl := xml.StartElement{Name: xml.Name{Local: "a:wholeTbl"}}
		e.EncodeElement(m.WholeTbl, sewholeTbl)
	}
	if m.Band1H != nil {
		seband1H := xml.StartElement{Name: xml.Name{Local: "a:band1H"}}
		e.EncodeElement(m.Band1H, seband1H)
	}
	if m.Band2H != nil {
		seband2H := xml.StartElement{Name: xml.Name{Local: "a:band2H"}}
		e.EncodeElement(m.Band2H, seband2H)
	}
	if m.Band1V != nil {
		seband1V := xml.StartElement{Name: xml.Name{Local: "a:band1V"}}
		e.EncodeElement(m.Band1V, seband1V)
	}
	if m.Band2V != nil {
		seband2V := xml.StartElement{Name: xml.Name{Local: "a:band2V"}}
		e.EncodeElement(m.Band2V, seband2V)
	}
	if m.LastCol != nil {
		selastCol := xml.StartElement{Name: xml.Name{Local: "a:lastCol"}}
		e.EncodeElement(m.LastCol, selastCol)
	}
	if m.FirstCol != nil {
		sefirstCol := xml.StartElement{Name: xml.Name{Local: "a:firstCol"}}
		e.EncodeElement(m.FirstCol, sefirstCol)
	}
	if m.LastRow != nil {
		selastRow := xml.StartElement{Name: xml.Name{Local: "a:lastRow"}}
		e.EncodeElement(m.LastRow, selastRow)
	}
	if m.SeCell != nil {
		seseCell := xml.StartElement{Name: xml.Name{Local: "a:seCell"}}
		e.EncodeElement(m.SeCell, seseCell)
	}
	if m.SwCell != nil {
		seswCell := xml.StartElement{Name: xml.Name{Local: "a:swCell"}}
		e.EncodeElement(m.SwCell, seswCell)
	}
	if m.FirstRow != nil {
		sefirstRow := xml.StartElement{Name: xml.Name{Local: "a:firstRow"}}
		e.EncodeElement(m.FirstRow, sefirstRow)
	}
	if m.NeCell != nil {
		seneCell := xml.StartElement{Name: xml.Name{Local: "a:neCell"}}
		e.EncodeElement(m.NeCell, seneCell)
	}
	if m.NwCell != nil {
		senwCell := xml.StartElement{Name: xml.Name{Local: "a:nwCell"}}
		e.EncodeElement(m.NwCell, senwCell)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TableStyle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.StyleIdAttr = "{00000000-0000-0000-0000-000000000000}"
	for _, attr := range start.Attr {
		if attr.Name.Local == "styleName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.StyleNameAttr = parsed
			continue
		}
		if attr.Name.Local == "styleId" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.StyleIdAttr = parsed
			continue
		}
	}
lCT_TableStyle:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "tblBg"}:
				m.TblBg = NewCT_TableBackgroundStyle()
				if err := d.DecodeElement(m.TblBg, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "wholeTbl"}:
				m.WholeTbl = NewCT_TablePartStyle()
				if err := d.DecodeElement(m.WholeTbl, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "band1H"}:
				m.Band1H = NewCT_TablePartStyle()
				if err := d.DecodeElement(m.Band1H, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "band2H"}:
				m.Band2H = NewCT_TablePartStyle()
				if err := d.DecodeElement(m.Band2H, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "band1V"}:
				m.Band1V = NewCT_TablePartStyle()
				if err := d.DecodeElement(m.Band1V, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "band2V"}:
				m.Band2V = NewCT_TablePartStyle()
				if err := d.DecodeElement(m.Band2V, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "lastCol"}:
				m.LastCol = NewCT_TablePartStyle()
				if err := d.DecodeElement(m.LastCol, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "firstCol"}:
				m.FirstCol = NewCT_TablePartStyle()
				if err := d.DecodeElement(m.FirstCol, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "lastRow"}:
				m.LastRow = NewCT_TablePartStyle()
				if err := d.DecodeElement(m.LastRow, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "seCell"}:
				m.SeCell = NewCT_TablePartStyle()
				if err := d.DecodeElement(m.SeCell, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "swCell"}:
				m.SwCell = NewCT_TablePartStyle()
				if err := d.DecodeElement(m.SwCell, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "firstRow"}:
				m.FirstRow = NewCT_TablePartStyle()
				if err := d.DecodeElement(m.FirstRow, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "neCell"}:
				m.NeCell = NewCT_TablePartStyle()
				if err := d.DecodeElement(m.NeCell, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "nwCell"}:
				m.NwCell = NewCT_TablePartStyle()
				if err := d.DecodeElement(m.NwCell, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "extLst"}:
				m.ExtLst = NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_TableStyle %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TableStyle
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TableStyle and its children
func (m *CT_TableStyle) Validate() error {
	return m.ValidateWithPath("CT_TableStyle")
}

// ValidateWithPath validates the CT_TableStyle and its children, prefixing error messages with path
func (m *CT_TableStyle) ValidateWithPath(path string) error {
	if !sharedTypes.ST_GuidPatternRe.MatchString(m.StyleIdAttr) {
		return fmt.Errorf(`%s/m.StyleIdAttr must match '%s' (have %v)`, path, sharedTypes.ST_GuidPatternRe, m.StyleIdAttr)
	}
	if m.TblBg != nil {
		if err := m.TblBg.ValidateWithPath(path + "/TblBg"); err != nil {
			return err
		}
	}
	if m.WholeTbl != nil {
		if err := m.WholeTbl.ValidateWithPath(path + "/WholeTbl"); err != nil {
			return err
		}
	}
	if m.Band1H != nil {
		if err := m.Band1H.ValidateWithPath(path + "/Band1H"); err != nil {
			return err
		}
	}
	if m.Band2H != nil {
		if err := m.Band2H.ValidateWithPath(path + "/Band2H"); err != nil {
			return err
		}
	}
	if m.Band1V != nil {
		if err := m.Band1V.ValidateWithPath(path + "/Band1V"); err != nil {
			return err
		}
	}
	if m.Band2V != nil {
		if err := m.Band2V.ValidateWithPath(path + "/Band2V"); err != nil {
			return err
		}
	}
	if m.LastCol != nil {
		if err := m.LastCol.ValidateWithPath(path + "/LastCol"); err != nil {
			return err
		}
	}
	if m.FirstCol != nil {
		if err := m.FirstCol.ValidateWithPath(path + "/FirstCol"); err != nil {
			return err
		}
	}
	if m.LastRow != nil {
		if err := m.LastRow.ValidateWithPath(path + "/LastRow"); err != nil {
			return err
		}
	}
	if m.SeCell != nil {
		if err := m.SeCell.ValidateWithPath(path + "/SeCell"); err != nil {
			return err
		}
	}
	if m.SwCell != nil {
		if err := m.SwCell.ValidateWithPath(path + "/SwCell"); err != nil {
			return err
		}
	}
	if m.FirstRow != nil {
		if err := m.FirstRow.ValidateWithPath(path + "/FirstRow"); err != nil {
			return err
		}
	}
	if m.NeCell != nil {
		if err := m.NeCell.ValidateWithPath(path + "/NeCell"); err != nil {
			return err
		}
	}
	if m.NwCell != nil {
		if err := m.NwCell.ValidateWithPath(path + "/NwCell"); err != nil {
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
