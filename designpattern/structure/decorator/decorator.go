package decorator

type Drawer interface {
	Draw() string
}

type Square struct {
}

func (s *Square) Draw() string {
	return "square"
}


// ColorSquare依赖的仍然是Drawer接口
type ColorSquare struct {
	square Drawer 
	color  string
}

func (s *ColorSquare) Draw() string {
	return s.square.Draw() + ":" + s.color
}

func NewColorSquare(square Drawer, color string) *ColorSquare {
	return &ColorSquare{
		square: square,
		color: color,
	}
}
