package UI

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	Pause    int = 0
	Continue     = 1
	Exit         = 2
)

type Button struct {
	X       int32
	Y       int32
	Active  bool
	Hovered bool
	Type    int32
	Sprite  rl.Texture2D
}

func NewButton(x, y, Type int32, spritePath string) *Button {
	sprite := rl.LoadTexture(spritePath)
	return &Button{
		X:       x,
		Y:       y,
		Sprite:  sprite,
		Active:  true,
		Hovered: false,
		Type:    Type,
	}
}

func (b *Button) DrawButton(framecount int) {
	rl.DrawTexture(b.Sprite, b.X, b.Y, rl.White)
}
func (b *Button) SetPos(x, y float32) {
	b.X = int32(x)
	b.Y = int32(y)
}
