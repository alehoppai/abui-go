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
}

type SizeFields struct {
	width   *string
	height  *string
	padding *string
}

type flex struct {
	children []Node
	LayoutFields
	SizeFields
}

func Flex(children ...Node) *flex {
	dir := v.Column
	align := v.AlignStart
	justify := v.JustifyStart

	layout := LayoutFields{direction: &dir, align: &align, justify: &justify}
	size := SizeFields{}

	return &flex{children: children, LayoutFields: layout, SizeFields: size}
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
	if s.LayoutFields.direction != nil {
		classList += fmt.Sprintf(" %s", string(*s.direction))
	}
	if s.LayoutFields.align != nil {
		classList += fmt.Sprintf(" %s", string(*s.align))
	}
	if s.LayoutFields.justify != nil {
		classList += fmt.Sprintf(" %s", string(*s.justify))
	}
	return strings.TrimSpace(classList)
}
