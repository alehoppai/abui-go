package elements

import "testing"

func TestDivRender_NoChildren(t *testing.T) {
	tree := Flex{
		Children: []Node{},
	}

	expected := "<div></div>"
	got := tree.Render()

	if got != expected {
        t.Errorf("Render() = %q; want %q", got, expected)
    }
}
