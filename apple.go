package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type apple struct {
	tile       *pixel.Sprite
	xoff, yoff float64
}

// TODO: take in snake so that apples don't spawn on snakes
func newApple(sprite *pixel.Sprite) *apple {
	var xoff, yoff float64
	xoff = float64(rand.Intn(int(winx/scale)))*scale + 10
	yoff = float64(rand.Intn(int(winy/scale)))*scale + 10
	return &apple{sprite, xoff, yoff}
}

func (a *apple) draw(win *pixelgl.Window) {
	a.tile.Draw(win, pixel.IM.Moved(pixel.V(a.xoff, a.yoff)))
}
