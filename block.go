package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const blockWidth = (1024 - (2 * (blockColumns + 1))) / blockColumns
const blockHeight = 20

type Block struct {
	x, y, width, height float32
	colour              color.RGBA
	hit                 bool
}

func (b *Block) Init(x, y float32, colour color.RGBA) {
	b.x = x
	b.y = y
	b.width = blockWidth
	b.height = blockHeight
	b.colour = colour
}

func (b *Block) Update() {
	if b.hit {
		b.colour = color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x00}
	}
}

func (b *Block) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, b.x, b.y, b.width, b.height, b.colour, true)
}
