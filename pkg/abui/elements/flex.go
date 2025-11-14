package elements

import (
	"fmt"
	"strings"

	v "github.com/alehoppai/abui-go/pkg/abui/values"
)

type LayoutFields struct {
	direction *v.Direction
	align     *v.Alignment
	justify   *v.Justification
	gap       *v.ScaleValue
}

// width can be set with not just ScaleValue but
// also with fractionalScale values
type SizeFields struct {
	width     string
	minWidth  string
	maxWidth  string
	height    string
	minHeight string
	maxHeight string
}

type OffsetFields struct {
	padding string
	margin  string
}

type flex struct {
	children []Node
	layout   LayoutFields
	size     SizeFields
	offset   OffsetFields
}

func Flex(children ...Node) *flex {
	dir := v.Column
	align := v.AlignStart
	justify := v.JustifyStart

	layout := LayoutFields{direction: &dir, align: &align, justify: &justify}
	size := SizeFields{}
	offset := OffsetFields{}

	return &flex{children: children, layout: layout, size: size, offset: offset}
}

func (s flex) Render() string {
	classList := s.buildClassList()
	html := "<div"
	if len(classList) > 0 {
		html += fmt.Sprintf(" class=\"%s\"", classList)
	}
	html += ">"

	for _, child := range s.children {
		html += child.Render()
	}

	html += "</div>"
	return html
}

func (s *flex) buildClassList() string {
	classList := ""
	if s.layout.direction != nil {
		classList += fmt.Sprintf(" %s", string(*s.layout.direction))
	}
	if s.layout.align != nil {
		classList += fmt.Sprintf(" %s", string(*s.layout.align))
	}
	if s.layout.justify != nil {
		classList += fmt.Sprintf(" %s", string(*s.layout.justify))
	}
	if s.layout.gap != nil {
		classList += fmt.Sprintf(" gap-%f", *s.layout.gap)
	}

	if s.size.width != nil {
		classList += fmt.Sprintf(" %s", *s.size.width)
	}
	if s.size.maxWidth != nil {
		classList += fmt.Sprintf(" %s", *s.size.maxWidth)
	}
	if s.size.minWidth != nil {
		classList += fmt.Sprintf(" %s", *s.size.minWidth)
	}
	if s.size.height != nil {
		classList += fmt.Sprintf(" %s", *s.size.height)
	}
	if s.size.maxHeight != nil {
		classList += fmt.Sprintf(" %s", *s.size.maxHeight)
	}
	if s.size.minHeight != nil {
		classList += fmt.Sprintf(" %s", *s.size.minHeight)
	}

	if s.offset.margin != nil {
		classList += fmt.Sprintf(" %s", *s.offset.margin)
	}
	if s.offset.padding != nil {
		classList += fmt.Sprintf(" %s", *s.offset.padding)
	}

	return strings.TrimSpace(classList)
}

// Layout
func (s *flex) Direction(v v.Direction) *flex {
	s.layout.direction = &v
	return s
}

func (s *flex) Align(v v.Alignment) *flex {
	s.layout.align = &v
	return s
}

func (s *flex) Justify(v v.Justification) *flex {
	s.layout.justify = &v
	return s
}

func (s *flex) Gap(v v.ScaleValue) *flex {
	s.layout.gap = &v
	return s
}

// End Layout

// Size
func (s *flex) Width(v v.ScaleValue) *flex {
	str := fmt.Sprintf("w-%v", v)
	s.size.width = &str
	return s
}

func (s *flex) WidthFrac(v v.FractionalScaleValue) *flex {
	str := fmt.Sprintf("w-%v", v)
	s.size.width = &str
	return s
}

func (s *flex) WidthContent(v v.MinMaxContentValue) *flex {
	str := fmt.Sprintf("w-%v", v)
	s.size.width = &str
	return s
}

func (s *flex) MinWidth(v v.ScaleValue) *flex {
	str := fmt.Sprintf("min-w-%v", v)
	s.size.minWidth = &str
	return s
}

func (s *flex) MinWidthFrac(v v.FractionalScaleValue) *flex {
	str := fmt.Sprintf("min-w-%v", v)
	s.size.minWidth = &str
	return s
}

func (s *flex) MinWidthContent(v v.MinMaxContentValue) *flex {
	str := fmt.Sprintf("min-w-%v", v)
	s.size.minWidth = &str
	return s
}

func (s *flex) MaxWidth(v v.ScaleValue) *flex {
	str := fmt.Sprintf("max-w-%v", v)
	s.size.maxWidth = &str
	return s
}

func (s *flex) MaxWidthFrac(v v.FractionalScaleValue) *flex {
	str := fmt.Sprintf("max-w-%v", v)
	s.size.maxWidth = &str
	return s
}

func (s *flex) MaxWidthContent(v v.MinMaxContentValue) *flex {
	str := fmt.Sprintf("max-w-%v", v)
	s.size.maxWidth = &str
	return s
}

func (s *flex) Height(v v.ScaleValue) *flex {
	str := fmt.Sprintf("h-%v", v)
	s.size.height = &str
	return s
}

func (s *flex) HeightFrac(v v.FractionalScaleValue) *flex {
	str := fmt.Sprintf("h-%v", v)
	s.size.height = &str
	return s
}

func (s *flex) HeightContent(v v.MinMaxContentValue) *flex {
	str := fmt.Sprintf("h-%v", v)
	s.size.height = &str
	return s
}

func (s *flex) MinHeight(v v.ScaleValue) *flex {
	str := fmt.Sprintf("min-h-%v", v)
	s.size.minHeight = &str
	return s
}

func (s *flex) MinHeightFrac(v v.FractionalScaleValue) *flex {
	str := fmt.Sprintf("min-h-%v", v)
	s.size.minHeight = &str
	return s
}

func (s *flex) MinHeightContent(v v.MinMaxContentValue) *flex {
	str := fmt.Sprintf("min-h-%v", v)
	s.size.minHeight = &str
	return s
}

func (s *flex) MaxHeight(v v.ScaleValue) *flex {
	str := fmt.Sprintf("maxh-%v", v)
	s.size.maxHeight = &str
	return s
}

func (s *flex) MaxHeightFrac(v v.FractionalScaleValue) *flex {
	str := fmt.Sprintf("maxh-%v", v)
	s.size.maxHeight = &str
	return s
}

func (s *flex) MaxHeightContent(v v.MinMaxContentValue) *flex {
	str := fmt.Sprintf("maxh-%v", v)
	s.size.maxHeight = &str
	return s
}

func (s *flex) Size(v v.ScaleValue) *flex {
	wstr := fmt.Sprintf("w-%v", v)
	hstr := fmt.Sprintf("h-%v", v)
	s.size.width = &wstr
	s.size.height = &hstr
	return s
}

// End Size

// Offset

func (s *flex) Padding(v v.ScaleValue, o *v.Offset) *flex {
	if o != nil {
		str := fmt.Sprintf("p-%v", v)
		s.offset.padding = &str
	} else {
		str := fmt.Sprintf("p%v-%v", &o, v)
		s.offset.padding = &str
	}

	return s
}

func (s *flex) Margin(v v.ScaleValue, o *v.Offset) *flex {
	if o != nil {
		str := fmt.Sprintf("m-%v", v)
		s.offset.margin = &str
	} else {
		str := fmt.Sprintf("m%v-%v", &o, v)
		s.offset.margin = &str
	}

	return s
}

// End Offset
