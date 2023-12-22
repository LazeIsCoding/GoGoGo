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
		Width:  32,
		Height: 32,
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
	//rl.DrawTexture(p.Sprite, p.X, p.Y, rl.White)
	fmt.Println(float32(p.X), " ", p.Y)
	rl.DrawTexturePro(p.Sprite,
		rl.NewRectangle(0, 64, p.Width, p.Height),
		rl.NewRectangle(float32(p.X), float32(p.Y), p.Width, p.Height),
		rl.NewVector2(float32(p.X), float32(p.Y)),
		0,
		rl.White)

}
