package elements

import (
	"fmt"

	v "github.com/alehoppai/abui-go/pkg/abui/values"
)
type flex struct {
	children []Node
	classList []string
	attributes map[string]string
}

func Flex(children ...Node) *flex {
	return &flex{children: children, classList: []string{}, attributes: make(map[string]string)}
}

func (s *flex) Render() string {
	html := "<div"
	if len(s.classList) > 0 {
		html += fmt.Sprintf(" class=\"%s\"", s.classList)
	}
	if len(s.attributes) > 0 {
		// TODO: handle attributes
	}
	html += ">"

	for _, child := range s.children {
		html += child.Render()
	}

	html += "</div>"
	return html
}

// Layout
func (s *flex) Direction(v v.Direction) *flex {
	return s
}

func (s *flex) Align(v v.Alignment) *flex {
	return s
}

func (s *flex) Justify(v v.Justification) *flex {
	return s
}

func (s *flex) Gap(v v.ScaleValue) *flex {
	return s
}

// End Layout
