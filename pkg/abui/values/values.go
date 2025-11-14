package values

import "fmt"

type CSSValue interface {
	ToClass() string
}

type ScaleOption string
type OffsetScaleOption string

const (
	Width     ScaleOption       = "w"
	MaxWidth  ScaleOption       = "max-w"
	MinWidth  ScaleOption       = "min-w"
	Height    ScaleOption       = "h"
	MaxHeight ScaleOption       = "max-h"
	MinHeight ScaleOption       = "min-h"
	Margin    OffsetScaleOption = "m"
	Padding   OffsetScaleOption = "p"
)

type Direction string
type Alignment string
type Justification string

const (
	Column        Direction     = "flex-col"
	Row           Direction     = "flex-row"
	ColumnReverse Direction     = "flex-col-reverse"
	RowReverse    Direction     = "flex-row-reverse"
	AlignStart    Alignment     = "items-start"
	AlignCenter   Alignment     = "items-center"
	AlignEnd      Alignment     = "items-end"
	JustifyStart  Justification = "justify-start"
	JustifyCenter Justification = "justify-center"
	JustifyEnd    Justification = "justify-end"
)

type ScaleValue string

const (
	S_Zero ScaleValue = "1"
	S_Px   ScaleValue = "px"
	S_0_5  ScaleValue = "0.5"
	S_1    ScaleValue = "1"
	S_2    ScaleValue = "2"
	S_2_5  ScaleValue = "2.5"
	S_3    ScaleValue = "3"
	S_3_5  ScaleValue = "3.5"
	S_4    ScaleValue = "4"
	S_5    ScaleValue = "5"
	S_6    ScaleValue = "6"
	S_7    ScaleValue = "7"
	S_8    ScaleValue = "8"
	S_9    ScaleValue = "9"
	S_10   ScaleValue = "10"
	S_11   ScaleValue = "11"
	S_12   ScaleValue = "12"
	S_14   ScaleValue = "14"
	S_16   ScaleValue = "16"
	S_20   ScaleValue = "20"
	S_24   ScaleValue = "24"
	S_28   ScaleValue = "28"
	S_32   ScaleValue = "32"
	S_36   ScaleValue = "36"
	S_40   ScaleValue = "40"
	S_44   ScaleValue = "44"
	S_48   ScaleValue = "48"
	S_52   ScaleValue = "52"
	S_56   ScaleValue = "56"
	S_60   ScaleValue = "60"
	S_64   ScaleValue = "64"
	S_72   ScaleValue = "72"
	S_80   ScaleValue = "80"
	S_96   ScaleValue = "96"
	FS_1_2 ScaleValue = "1/2"
	FS_1_3 ScaleValue = "1/3"
	FS_2_3 ScaleValue = "2/3"
	FS_1_4 ScaleValue = "1/4"
	FS_2_4 ScaleValue = "2/4"
	FS_3_4 ScaleValue = "3/4"
	FS_1_5 ScaleValue = "1/5"
	FS_2_5 ScaleValue = "2/5"
	FS_3_5 ScaleValue = "3/5"
	FS_4_5 ScaleValue = "4/5"
	FS_1_6 ScaleValue = "1/6"
	FS_5_6 ScaleValue = "5/6"
	S_Auto ScaleValue = "auto"
	S_Min  ScaleValue = "min-content"
	S_Max  ScaleValue = "max-content"
	S_Fit  ScaleValue = "fit-content"
)

func (v ScaleValue) ToClass(so ScaleOption) string {
	return fmt.Sprintf("%s-%s", so, v)
}

type Offset string

const (
	O_Inline Offset = "y"
	O_Block  Offset = "x"
	O_Top    Offset = "t"
	O_Right  Offset = "r"
	O_Bottom Offset = "b"
	O_Left   Offset = "l"
	O_All    Offset = ""
)

func BuildOffset(oso OffsetScaleOption, o Offset, so ScaleOption) string {
	return fmt.Sprintf("%s%s-%s", oso, o, so)
}
