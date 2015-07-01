package main

import "fmt"

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color string

type Box struct {
	width, height, depath float64
	color Color
}

type BoxList []Box //a slice of boxes

func (b Box) Volume() float64 {
	return b.width * b.height * b.depath
}

func (b *Box) SetColor(c Color) {
	b.color = c
}

func (b1 BoxList) BiggestsColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range b1 {
		if b.Volume() > v {
			v = b.Volume()
			k = b.color
		}
	}
	return k
}

func (b1 BoxList) PaintItBlack() {
	for i, _ := range b1 {
		b1[i].SetColor("BLACK")
	}
}

func (c Color) String() string {
	strings := []string {"_WHITE_", "_BLACK_", "_BLUE_", "_RED_", "_YELLOW_"}
	return strings[c]
}



func main() {
	boxes := BoxList {
		Box{14, 24, 34, "RED"},
		Box{15, 25, 35, "YELLOW"},
		Box{16, 26, 36, "BLACK"},
		Box{17, 27, 37, "BLUE"},
		Box{18, 28, 38, "WHITE"},
		Box{19, 29, 39, "YELLOW"},
	}

	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("The volume of the first one is ", boxes[0].Volume(), " cm")
	fmt.Println("The color of the last one is ", boxes[len(boxes)-1].color.String())
	fmt.Println("The biggest one is ", boxes.BiggestsColor().String())
	fmt.Println("Let's paint them all black")
	boxes.PaintItBlack()
	fmt.Println("The color of the second one is ", boxes[1].color.String())
	fmt.Println("Obviously, now, the biggest one is ", boxes.BiggestsColor().String())

}

