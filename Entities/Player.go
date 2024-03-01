package Entities

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Pos         rl.Rectangle
	Src         rl.Rectangle
	Health      int32
	Speed       float32
	Sprite      rl.Texture2D
	State       int32
	spriteCount int32
	Width       int32
	Height      int32
}

func NewPlayer(x, y, health int32, speed float32, spritePath string) *Player {
	sprite := rl.LoadTexture(spritePath)
	pos := rl.NewRectangle(float32(x), float32(y), 52, 76)
	src := rl.NewRectangle(0, 32, pos.Width, pos.Height)
	return &Player{
		Pos:         pos,
		Src:         src,
		Health:      health,
		Speed:       speed,
		Sprite:      sprite,
		spriteCount: 0,
		Width:       128,
		Height:      128,
	}
}

func (p *Player) Move(x, y float32) {
	p.Pos.X += x
	p.Pos.Y += y
}

func (p *Player) ToString() string {
	s := fmt.Sprint("X: ", p.Pos.X, "Y: ", p.Pos.Y)
	return s
}
func (p *Player) DrawPlayer(framecount int) {
	//rl.DrawTexture(p.Sprite, int32(p.GetX()), int32(p.GetY()), rl.White)
	//fmt.Printf()
	//rl.DrawTexturePro(playerSprite, playerSrc, playerDest, rl.NewVector2(playerDest.Width, playerDest.Height), 0, rl.White)

	if framecount%15 == 0 {
		p.spriteCount = (p.spriteCount + 1) % 8
	}
	p.Src.Y = float32(p.Height * p.spriteCount)
	p.Src.X = float32(p.Width * p.State)

	rl.DrawTexturePro(p.Sprite, p.Src, p.Pos, rl.NewVector2(0, 0), 0, rl.White)

}
func (p *Player) GetX() float32 {
	return p.Pos.X
}

func (p *Player) GetY() float32 {
	return p.Pos.Y
}

func (p *Player) GetWidth() float32 {
	return p.Pos.Width
}
func (p *Player) GetHeight() float32 {
	return p.Pos.Height
}
