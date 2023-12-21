package Entities

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Y      int32
	X      int32
	Health int32
	Speed  int32
	Width  float32
	Height float32
	Sprite rl.Texture2D
}

func NewPlayer(x, y, health, speed int32, spritePath string) *Player {
	sprite := rl.LoadTexture(spritePath)
	return &Player{
		X:      x,
		Y:      y,
		Health: health,
		Speed:  speed,
		Sprite: sprite,
		Width:  float32(sprite.Width),
		Height: float32(sprite.Height),
	}
}

func (p *Player) Move(x, y int32) {
	p.X += x
	p.Y += y
}

func (p *Player) ToString() string {
	s := fmt.Sprint("X: ", p.X, "Y: ", p.Y)
	return s
}
func (p *Player) DrawPlayer(framecount, state int) {
	rl.DrawTexture(p.Sprite, p.X, p.Y, rl.White)
}
