package Entities

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
	"time"
)

type Butterfly struct {
	Pos         rl.Rectangle
	Src         rl.Rectangle
	Sprite      rl.Texture2D
	State       int32
	spriteCount int32
	Color       int32
	Rand        *rand.Rand
	speed       float32
	dir         rl.Vector2
}

func NewButterfly(x, y int32, spritePath string) *Butterfly {
	sprite := rl.LoadTexture(spritePath)
	pos := rl.NewRectangle(float32(x), float32(y), 16, 16)
	src := rl.NewRectangle(0, 16, pos.Width, pos.Height)
	Dir := rl.NewVector2(rand.Float32()-0.5, rand.Float32()-0.5)
	rand.NewSource(time.Now().UnixNano())
	return &Butterfly{
		Pos:         pos,
		Src:         src,
		Sprite:      sprite,
		State:       0,
		spriteCount: 0,
		Color:       rand.Int31() % 4,
		dir:         Dir,
		speed:       0.25,
		Rand:        rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}
func (b *Butterfly) Move() {

	b.Pos.X += b.dir.X * b.speed
	b.Pos.Y += b.dir.Y * b.speed

	b.dir.X += rand.Float32()/4 - 0.125
	b.dir.Y += rand.Float32()/8 - 0.0625
	b.dir = rl.Vector2Normalize(b.dir)
}

func (b *Butterfly) Draw(framecount int) {
	if framecount%10 == 0 {
		b.spriteCount = (b.spriteCount + 1) % 4
	}

	b.Src.Y = float32(16 * b.spriteCount)
	b.Src.X = float32(16 * b.Color)

	rl.DrawTexturePro(b.Sprite, b.Src, b.Pos, rl.NewVector2(0, 0), 0, rl.White)
}

func (b *Butterfly) GetX() float32 {
	return b.Pos.X
}
func (b *Butterfly) GetY() float32 {
	return b.Pos.Y
}

func (b *Butterfly) GetWidth() float32 {
	return b.Pos.Width
}
func (b *Butterfly) GetHeight() float32 {
	return b.Pos.Height
}
