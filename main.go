package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"time"
)

var winx, winy float64 = 800, 600
var scale float64 = 20
var frametime time.Duration = time.Second / 15
var score int

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

func run() {
	// Set up window
	cfg := pixelgl.WindowConfig{
		Title:  "Snake",
		Bounds: pixel.R(0, 0, winx, winy),
		VSync:  false,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// Create snake
	pict := pixel.MakePictureData(pixel.R(0, 0, scale, scale))
	for i := 0; i < len(pict.Pix); i++ {
		pict.Pix[i] = colornames.Greenyellow
	}
	snakeSprite := pixel.NewSprite(pict, pict.Bounds())
	s := newSnake(snakeSprite)

	// Create starting apple
	pict = pixel.MakePictureData(pixel.R(0, 0, scale, scale))
	for i := 0; i < len(pict.Pix); i++ {
		pict.Pix[i] = colornames.Red
	}
	appleSprite := pixel.NewSprite(pict, pict.Bounds())
	a := newApple(appleSprite)

	// Game loop
	t := time.Now()
	for !win.Closed() {
		if win.Pressed(pixelgl.KeyLeft) {
			s.updateBodyDir(Left)
		}
		if win.Pressed(pixelgl.KeyRight) {
			s.updateBodyDir(Right)
		}
		if win.Pressed(pixelgl.KeyDown) {
			s.updateBodyDir(Down)
		}
		if win.Pressed(pixelgl.KeyUp) {
			s.updateBodyDir(Up)
		}

		if time.Since(t) < frametime {
			continue
		}
		win.Clear(colornames.Skyblue)

		s.move()

		if s.selfCollide() {
			fmt.Println("Game over")
			s = newSnake(snakeSprite)
			a = newApple(appleSprite)
			// TODO: clean this up and add gameover screen
			//os.Exit(0)
		}

		if s.collide(a) {
			a = newApple(appleSprite)
			s.addPart()
			score++
		}

		a.draw(win)
		s.draw(win)

		win.Update()
		t = time.Now()
	}
}

func main() {
	pixelgl.Run(run)
}
