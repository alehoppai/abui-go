package elements

import "fmt"

type Direction string
const (
	Column Direction = "flex-col"
	Row Direction = "flex-row"
	ColumnReverse Direction = "flex-col-reverse"
	RowReverse Direction = "flex-row-reverse"
)

type Alignment string
const (
	AlignStart Alignment = "items-start"
	AlignCenter Alignment = "items-center"
	AlignEnd Alignment = "items-end"
)

type Justification string
const (
	JustifyStart Justification = "justify-start"
	JustifyCenter Justification = "justify-center"
	JustifyEnd Justification = "justify-end"
)

type Flex struct {
	Class *string
	Children []Node
	Direction *Direction
	Align *Alignment
	Justify *Justification
}

func (s Flex) Render() string {
	class := s.buildClass()
	html := "<div"

	if len(class) > 0 {
		html += fmt.Sprintf(" class=\"%s\"", class)
	}

	html += ">"

	for _, child := range s.Children {
		html += child.Render()
	}

	html += "</div>"
	return html
}

func (s *Flex) buildClass() string {
	class := ""
	if s.Class != nil {
		class += *s.Class
	}
	if s.Direction != nil {
		class += fmt.Sprintf(", %s", *s.Direction)
	}
	if s.Align != nil {
		class += fmt.Sprintf(", %s", *s.Align)
	}
	if s.Justify != nil {
		class += fmt.Sprintf(", %s", *s.Justify)
	}

	return class
}
