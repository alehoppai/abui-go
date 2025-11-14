package elements

import "testing"

func TestDivRender_NoChildren(t *testing.T) {
	tree := Flex()

	expected := "<div class=\"flex-col items-start justify-start\"></div>"
	got := tree.Render()

	if got != expected {
		t.Errorf("Render() = %q; want %q", got, expected)
	}
}
