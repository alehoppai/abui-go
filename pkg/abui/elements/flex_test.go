package elements

import (
	"testing"

	v "github.com/alehoppai/abui-go/pkg/abui/values"
)

func TestFlexRender_NoChildren(t *testing.T) {
	tree := Flex()

	expected := "<div class=\"flex-col items-start justify-start\"></div>"
	got := tree.Render()

	if got != expected {
		t.Errorf("Render() = %q; want %q", got, expected)
	}
}

func TestFlexRender_WithChildren(t *testing.T) {
	tree := Flex(Flex(), Flex(Flex()))

	expected := "<div class=\"flex-col items-start justify-start\"><div class=\"flex-col items-start justify-start\"></div><div class=\"flex-col items-start justify-start\"><div class=\"flex-col items-start justify-start\"></div></div></div>"
	got := tree.Render()

	if got != expected {
		t.Errorf("Render() = %q; want %q", got, expected)
	}
}

func TestFlexRender_LayoutSetup(t *testing.T) {
	tree := Flex().Direction(v.Row).Align(v.AlignCenter).Justify(v.JustifyCenter)

	expected := "<div class=\"flex-row items-center justify-center\"></div>"
	got := tree.Render()

	if got != expected {
		t.Errorf("Render() = %q; want %q", got, expected)
	}
}

func TestFlexRender_SizeExactSetup(t *testing.T) {
	tree := Flex().Width(v.S_20).Height(v.S_16)

	expected := "<div class=\"flex-col items-start justify-start w-5 h-4\"></div>"
	got := tree.Render()

	if got != expected {
		t.Errorf("Render() = %q; want %q", got, expected)
	}
}

func TestFlexRender_SizeFracSetup(t *testing.T) {
	tree := Flex().WidthFrac(v.FS_1_2).HeightFrac(v.FS_1_3)

	expected := "<div class=\"flex-col items-start justify-start w-1/2 h-1/3\"></div>"
	got := tree.Render()

	if got != expected {
		t.Errorf("Render() = %q; want %q", got, expected)
	}
}

func TestFlexRender_SizeMinMaxSetup(t *testing.T) {
	tree := Flex().WidthContent(v.S_Auto).HeightContent(v.S_Fit)

	expected := "<div class=\"flex-col items-start justify-start w-auto h-fit-content\"></div>"
	got := tree.Render()

	if got != expected {
		t.Errorf("Render() = %q; want %q", got, expected)
	}
}
