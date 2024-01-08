package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Paddle struct {
	x, y, width, height float32
	colour              color.RGBA
}

func (p *Paddle) Init() {
	p.x = (screenWidth / 2) - 50
	p.y = screenHeight - 40
	p.width = 100
	p.height = 10
	p.colour = color.RGBA{
		R: uint8(0x00),
		G: uint8(0x00),
		B: uint8(0xff),
		A: uint8(0xff)}

}

func (p *Paddle) Update() {

}

func (p *Paddle) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, p.x, p.y, p.width, p.height, p.colour, true)
}
