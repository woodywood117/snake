package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type bodypart struct {
	xoff, yoff float64
	dir        Direction
	lastdir    Direction
}

type snake struct {
	tile      *pixel.Sprite
	bodyparts []bodypart
}

func newSnake(sprite *pixel.Sprite) *snake {
	return &snake{sprite, []bodypart{
		{410, 350, Up, Up},
		{410, 330, Up, Up},
		{410, 310, Up, Up},
	}}
}

func (s *snake) updateBodyDir(direction Direction) {
	switch {
	case s.bodyparts[0].dir == Up && direction == Down:
		return
	case s.bodyparts[0].dir == Down && direction == Up:
		return
	case s.bodyparts[0].dir == Left && direction == Right:
		return
	case s.bodyparts[0].dir == Right && direction == Left:
		return
	}
	s.bodyparts[0].dir = direction
}

func (s *snake) move() {
	NextDirection := s.bodyparts[0].dir
	for part := range s.bodyparts {
		s.bodyparts[part].lastdir = s.bodyparts[part].dir
		switch s.bodyparts[part].dir {
		case Up:
			s.bodyparts[part].yoff += s.tile.Frame().H()
			s.bodyparts[part].dir = NextDirection
			NextDirection = Up
		case Down:
			s.bodyparts[part].yoff -= s.tile.Frame().H()
			s.bodyparts[part].dir = NextDirection
			NextDirection = Down
		case Left:
			s.bodyparts[part].xoff -= s.tile.Frame().W()
			s.bodyparts[part].dir = NextDirection
			NextDirection = Left
		case Right:
			s.bodyparts[part].xoff += s.tile.Frame().W()
			s.bodyparts[part].dir = NextDirection
			NextDirection = Right
		}

		if s.bodyparts[part].xoff > winx {
			s.bodyparts[part].xoff = 10
		}
		if s.bodyparts[part].xoff < 0 {
			s.bodyparts[part].xoff = winx - 10
		}
		if s.bodyparts[part].yoff > winy {
			s.bodyparts[part].yoff = 10
		}
		if s.bodyparts[part].yoff < 0 {
			s.bodyparts[part].yoff = winy - 10
		}
	}
}

func (s *snake) draw(win *pixelgl.Window) {
	for part := range s.bodyparts {
		s.tile.Draw(win, pixel.IM.Moved(pixel.V(s.bodyparts[part].xoff, s.bodyparts[part].yoff)))
	}
}

func (s *snake) collide(a *apple) bool {
	return s.bodyparts[0].xoff == a.xoff && s.bodyparts[0].yoff == a.yoff
}

func (s *snake) selfCollide() bool {
	for i := 1; i < len(s.bodyparts); i++ {
		if s.bodyparts[0].xoff == s.bodyparts[i].xoff && s.bodyparts[0].yoff == s.bodyparts[i].yoff {
			return true
		}
	}
	return false
}

func (s *snake) addPart() {
	newpart := s.bodyparts[len(s.bodyparts)-1]
	newpart.dir = newpart.lastdir
	switch newpart.dir {
	case Up:
		newpart.yoff -= s.tile.Frame().H()
	case Down:
		newpart.yoff += s.tile.Frame().H()
	case Left:
		newpart.xoff += s.tile.Frame().W()
	case Right:
		newpart.xoff -= s.tile.Frame().W()
	}

	s.bodyparts = append(s.bodyparts, newpart)
}
