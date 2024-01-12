package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const blockColumns = 14
const blockRows = 8
const startingBlocks = blockColumns * blockRows
const headingHeight = 40

var (
	normalFont font.Face
	score      int
)

type Game struct {
	blocks []Block
	paddle Paddle
	ball   Ball
}

func NewGame() *Game {
	spacer := 2
	var xoffset, yoffset float32
	var colour color.RGBA

	g := &Game{}
	g.blocks = make([]Block, startingBlocks)

	for i := range g.blocks {
		row := (i / blockColumns) + 1
		col := (i + 1) - (blockColumns * (row - 1))

		xoffset = float32((spacer * col) + ((blockWidth * col) - blockWidth))
		yoffset = float32(((blockHeight * row) - blockHeight) + (spacer * row) + headingHeight)

		switch row {
		case 1, 2:
			colour = color.RGBA{R: 0xa3, G: 0x1e, B: 0x0a, A: 0xff}
		case 3, 4:
			colour = color.RGBA{R: 0xc2, G: 0x85, B: 0x0a, A: 0xff}
		case 5, 6:
			colour = color.RGBA{R: 0x0a, G: 0x85, B: 0x33, A: 0xff}
		case 7, 8:
			colour = color.RGBA{R: 0xc2, G: 0xc2, B: 0x29, A: 0xff}

		}

		g.blocks[i].Init(xoffset, yoffset, colour)
		fmt.Println(g.blocks[i].x, g.blocks[i].y)

	}
	g.paddle = Paddle{}
	g.paddle.Init()

	g.ball = Ball{}
	g.ball.Init()

	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	normalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}

	return g
}

func (g *Game) Update() error {

	for i := range g.blocks {
		g.blocks[i].Update()
	}

	g.handleMovement()
	g.checkCollisions()
	g.ball.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	text.Draw(screen, fmt.Sprint(score), normalFont, 10, 30, color.White)
	for _, b := range g.blocks {
		b.Draw(screen)
	}
	g.paddle.Draw(screen)
	g.ball.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) handleMovement() {

	amountToMove := float32(4)

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.paddle.x -= amountToMove
		if !g.ball.moving {
			g.ball.x -= amountToMove
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.paddle.x += amountToMove
		if !g.ball.moving {
			g.ball.x += amountToMove
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.ball.heading = 0
		g.ball.moving = true
	}

}

func (g *Game) checkCollisions() {
	//blocks
	for i, block := range g.blocks {
		if block.hit {
			continue
		}
		if block.x < g.ball.x+g.ball.radius &&
			block.x+block.width > g.ball.x-g.ball.radius &&
			block.y < g.ball.y+g.ball.radius &&
			block.y+block.height > g.ball.y-g.ball.radius {

			if (block.y+block.height)-(g.ball.y-g.ball.radius) < 0 {
				fmt.Println("collision with top surface of", i)
				g.blocks[i].hit = true
				g.ball.switchDirection()
				fmt.Println("heading", g.ball.heading)
				fmt.Println("speed", g.ball.speed)
				incrementScore(block.colour)
				break
			}

			if block.y-(g.ball.y+g.ball.radius) < 0 {
				fmt.Println("collision with bottom surface of", i)
				g.blocks[i].hit = true
				g.ball.switchDirection()
				fmt.Println("heading", g.ball.heading)
				fmt.Println("speed", g.ball.speed)
				incrementScore(block.colour)
				break
			}

			if block.x-(g.ball.x-g.ball.radius) < 0 {
				fmt.Println("collision with left surface of", i)
				g.blocks[i].hit = true
				g.ball.bounceOffWall()
				fmt.Println("heading", g.ball.heading)
				fmt.Println("speed", g.ball.speed)
				incrementScore(block.colour)
				break
			}

			if (block.x+block.width)-(g.ball.x+g.ball.radius) < 0 {
				fmt.Println("collision with right surface of", i)
				g.blocks[i].hit = true
				g.ball.bounceOffWall()
				fmt.Println("heading", g.ball.heading)
				fmt.Println("speed", g.ball.speed)
				incrementScore(block.colour)
				break
			}
		}
	}

	//walls
	if (g.ball.y - g.ball.radius) <= 0 {

		fmt.Println("collision with ceiling")
		g.ball.switchDirection()
		fmt.Println("heading", g.ball.heading)
		fmt.Println("speed", g.ball.speed)
	}

	if (g.ball.x - g.ball.radius) <= 0 {

		fmt.Println("collision with left wall")
		g.ball.bounceOffWall()
		fmt.Println("heading", g.ball.heading)
		fmt.Println("speed", g.ball.speed)
	}

	if (g.ball.x + g.ball.radius) >= screenWidth {

		fmt.Println("collision with right wall")
		g.ball.bounceOffWall()
		fmt.Println("heading", g.ball.heading)
		fmt.Println("speed", g.ball.speed)
	}

	if (g.ball.y + g.ball.radius) >= screenHeight {

		fmt.Println("fell through the floor")
		g.ball.Init()
		g.ball.x = g.paddle.x + (g.paddle.width / 2)
		fmt.Println("heading", g.ball.heading)
		fmt.Println("speed", g.ball.speed)
	}

	//paddle
	if g.paddle.x < g.ball.x+g.ball.radius &&
		g.paddle.x+g.paddle.width > g.ball.x-g.ball.radius &&
		g.paddle.y < g.ball.y+g.ball.radius &&
		g.paddle.y+g.paddle.height > g.ball.y-g.ball.radius {

		fmt.Println("collision with paddle")
		g.ball.switchDirection()
		fmt.Println("heading", g.ball.heading)
		fmt.Println("speed", g.ball.speed)
	}
}
