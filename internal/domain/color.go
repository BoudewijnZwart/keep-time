package domain

type Color string

const (
	ColorRed   Color = "red"
	ColorBlue  Color = "blue"
	ColorGreen Color = "green"
)

func (c Color) Hex() string {
	switch c {
	case ColorRed:
		return "#FF0000"
	case ColorBlue:
		return "#0000FF"
	case ColorGreen:
		return "#00FF00"
	default:
		return "#000000"
	}
}
