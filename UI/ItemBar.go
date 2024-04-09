package UI

import (
	"GoGoGo/Inventory"
	rl "github.com/gen2brain/raylib-go/raylib"
	"strconv"
)

type ItemBar struct {
	X              int32
	Y              int32
	ItemCount      int
	Selected       int32
	ItemSprites    rl.Texture2D
	Sprite         rl.Texture2D
	SpriteSelected rl.Texture2D
	Inventory      []Inventory.Item
}

func NewItemBar(x, y int32, spritePath, spriteSelectedPath, itemSpritesPath string, inventory []Inventory.Item) *ItemBar {
	sprite := rl.LoadTexture(spritePath)
	spriteSelected := rl.LoadTexture(spriteSelectedPath)
	items := rl.LoadTexture(itemSpritesPath)
	return &ItemBar{
		X:              x,
		Y:              y,
		Sprite:         sprite,
		SpriteSelected: spriteSelected,
		ItemSprites:    items,
		Selected:       0,
		ItemCount:      6,
		Inventory:      inventory,
	}
}

func (i *ItemBar) DrawItemBar() {

	rl.DrawTexture(i.Sprite, i.X, i.Y, rl.White)
	rl.DrawTexture(i.SpriteSelected, i.X+64*i.Selected, i.Y, rl.White)
	for j := 0; j < i.ItemCount; j++ {
		if i.Inventory[j].Id != int32(0) {
			rl.DrawTexture(i.Inventory[j].Sprite, i.X+64*int32(j), i.Y, rl.White)
			rl.DrawText(strconv.Itoa(int(i.Inventory[j].Count)), i.X+64*int32(j)+25, i.Y+45, 20, rl.White)
		}
	}
}
func (i *ItemBar) SetPos(x, y float32) {
	i.X = int32(x)
	i.Y = int32(y)
}
