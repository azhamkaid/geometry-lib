package shape

type Regtangle struct {
	Width  float32
	Length float32
}

func (regtangle *Regtangle) Area() float32 {
	return regtangle.Length * regtangle.Width
}

func (regtangle *Regtangle) Perimeter() float32 {
	return 2 * (regtangle.Length + regtangle.Width)

}
