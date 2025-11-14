package values

type Direction string

const (
	Column        Direction = "flex-col"
	Row           Direction = "flex-row"
	ColumnReverse Direction = "flex-col-reverse"
	RowReverse    Direction = "flex-row-reverse"
)

type Alignment string

const (
	AlignStart  Alignment = "items-start"
	AlignCenter Alignment = "items-center"
	AlignEnd    Alignment = "items-end"
)

type Justification string

const (
	JustifyStart  Justification = "justify-start"
	JustifyCenter Justification = "justify-center"
	JustifyEnd    Justification = "justify-end"
)

type ScaleValue float32

const (
	S_Zero ScaleValue = 0
	S_Px   ScaleValue = 0.0625
	S_0_5  ScaleValue = 0.125
	S_1    ScaleValue = 0.25
	S_2    ScaleValue = 0.5
	S_2_5  ScaleValue = 0.625
	S_3    ScaleValue = 0.75
	S_3_5  ScaleValue = 0.875
	S_4    ScaleValue = 1
	S_5    ScaleValue = 1.25
	S_6    ScaleValue = 1.5
	S_7    ScaleValue = 1.75
	S_8    ScaleValue = 2
	S_9    ScaleValue = 2.25
	S_10   ScaleValue = 2.5
	S_11   ScaleValue = 2.75
	S_12   ScaleValue = 3
	S_14   ScaleValue = 3.5
	S_16   ScaleValue = 4
	S_20   ScaleValue = 5
	S_24   ScaleValue = 6
	S_28   ScaleValue = 7
	S_32   ScaleValue = 8
	S_36   ScaleValue = 9
	S_40   ScaleValue = 10
	S_44   ScaleValue = 11
	S_48   ScaleValue = 12
	S_52   ScaleValue = 13
	S_56   ScaleValue = 14
	S_60   ScaleValue = 15
	S_64   ScaleValue = 16
	S_72   ScaleValue = 18
	S_80   ScaleValue = 20
	S_96   ScaleValue = 24
)

type FractionalScaleValue float32

const (
	FS_1_2 FractionalScaleValue = 50
	FS_1_3 FractionalScaleValue = 33.333333
	FS_2_3 FractionalScaleValue = 66.666667
	FS_1_4 FractionalScaleValue = 25
	FS_2_4 FractionalScaleValue = 50
	FS_3_4 FractionalScaleValue = 75
	FS_1_5 FractionalScaleValue = 20
	FS_2_5 FractionalScaleValue = 40
	FS_3_5 FractionalScaleValue = 60
	FS_4_5 FractionalScaleValue = 80
	FS_1_6 FractionalScaleValue = 16.666667
	FS_5_6 FractionalScaleValue = 83.333333
)

type SizeVerbalValue string

const (
	S_Auto SizeVerbalValue = "auto"
	S_Min  SizeVerbalValue = "min-content"
	S_Max  SizeVerbalValue = "max-content"
	S_Fit  SizeVerbalValue = "fit-content"
)
