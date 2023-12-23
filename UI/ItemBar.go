package UI

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ItemBar struct {
	X              int32
	Y              int32
	ItemCount      int32
	Selected       int32
	Sprite         rl.Texture2D
	SpriteSelected rl.Texture2D
}

func NewItemBar(x, y int32, spritePath, spriteSelectedPath string) *ItemBar {
	sprite := rl.LoadTexture(spritePath)
	spriteSelected := rl.LoadTexture(spriteSelectedPath)
	return &ItemBar{
		X:              x,
		Y:              y,
		Sprite:         sprite,
		SpriteSelected: spriteSelected,
		Selected:       0,
		ItemCount:      6,
	}
}

func (i *ItemBar) DrawItemBar() {

	rl.DrawTexture(i.Sprite, i.X, i.Y, rl.White)
	rl.DrawTexture(i.SpriteSelected, i.X+16*i.Selected, i.Y, rl.White)
}
func (i *ItemBar) SetPos(x, y float32) {
	i.X = int32(x)
	i.Y = int32(y)
}
