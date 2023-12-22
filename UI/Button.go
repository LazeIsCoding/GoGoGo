package UI

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	Pause    int32 = 0
	Continue       = 1
	Option         = 2
	Exit           = 3
)

type Button struct {
	X             int32
	Y             int32
	Active        bool
	Hovered       bool
	Pressed       bool
	Type          int32
	Bounds        rl.Rectangle
	Sprite        rl.Texture2D
	SpritePressed rl.Texture2D
}

func NewButton(x, y, Type int32, spritePath, spritePressedPath string) *Button {
	sprite := rl.LoadTexture(spritePath)
	spritePressed := rl.LoadTexture(spritePressedPath)

	return &Button{
		X:             x,
		Y:             y,
		Sprite:        sprite,
		SpritePressed: spritePressed,
		Active:        true,
		Hovered:       false,
		Pressed:       false,
		Type:          Type,
	}
}

func (b *Button) DrawButton(framecount int) {
	if !b.Pressed {
		rl.DrawTexture(b.Sprite, b.X, b.Y, rl.White)
	} else {
		rl.DrawTexture(b.SpritePressed, b.X, b.Y, rl.White)
	}
}
func (b *Button) SetPos(x, y float32) {
	b.X = int32(x)
	b.Y = int32(y)
}

func (b *Button) OnClick() {
	b.Pressed = true
	action(b.Type)
}

func action(Type int32) {
	switch Type {
	case Pause:
		fmt.Print("Pause")
	case Continue:
		fmt.Print("Continue")
	case Option:
		fmt.Print("Option")
	case Exit:
		fmt.Print("Exit")
	}
}
