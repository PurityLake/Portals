package math

type Rect struct {
	X, Y, W, H int
}

func (rect Rect) Contains(x, y int) bool {
	return x >= rect.X && x < rect.X+rect.W && y >= rect.Y && y < rect.Y+rect.H
}

func (rect Rect) Intersect(other Rect) bool {
	return rect.X < other.X+other.W && rect.X+rect.W > other.X && rect.Y < other.Y+other.H && rect.Y+rect.H > other.Y
}

func (rect Rect) InterestNear(other Rect, padding int) bool {
	padded_rect := Rect{other.X - padding, other.Y - padding, other.W + padding*2, other.H + padding*2}
	return padded_rect.Intersect(other)
}
