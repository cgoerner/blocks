package main

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const ballRadius = 10
const initialBallSpeed = 5
const maxBounceRandomness = 20

type Ball struct {
	x, y, radius float32
	colour       color.RGBA
	moving       bool
	heading      float32
	speed        float32
}

func (b *Ball) Init() {
	b.x = screenWidth / 2
	b.y = screenHeight - 50
	b.radius = ballRadius
	b.colour = color.RGBA{
		R: uint8(0xff),
		G: uint8(0xff),
		B: uint8(0xff),
		A: uint8(0xff)}
	b.moving = false
	b.heading = 0
	b.speed = initialBallSpeed
}

func (b *Ball) Update() {
	radians := b.heading * (math.Pi / 180)
	if b.moving {
		b.x = b.x + float32(math.Sin(float64(radians)))*b.speed
		b.y = b.y - float32(math.Cos(float64(radians)))*b.speed
	}
}

func (b *Ball) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, b.x, b.y, b.radius, b.colour, true)
}

func (b *Ball) switchDirection() {
	b.heading = 180 - b.heading + randomness()
	b.speed += 0.1
}

func (b *Ball) bounceOffWall() {
	b.heading = 360 - b.heading + randomness()
	b.speed += 0.1
}

func randomness() float32 {
	return rand.Float32()*maxBounceRandomness - (maxBounceRandomness / 2)
}
