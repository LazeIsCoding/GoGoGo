package Inventory

import rl "github.com/gen2brain/raylib-go/raylib"

type Inventory struct {
	Inv      []Item
	space    int32
	tree     Tree
	Act      int32
	capacity int32
}

func InitInv(size int32) *Inventory {
	return &Inventory{
		space:    size,
		capacity: size,
		Act:      0,
		Inv:      make([]Item, size),
	}
}

func (inv *Inventory) AddItem(count int32, id int32, sprite rl.Texture2D) bool {

	/*index, nd := inv.tree.Search(item.getID())
	if nd != nil {
		inv.Inv[index].count += count
	}
	if !inv.isFull() {
		if &inv.Inv[inv.Act] != nil {
			inv.Inv[inv.Act] = *item
			inv.tree.Insert(item.getID(), inv.Act, nil)
		}a
	}*/
	if inv.Inv[inv.Act].Id == 0 {
		inv.Inv[inv.Act] = *CreateItem(id, sprite)
	} else {
		inv.Inv[inv.Act].Count++
	}
	return true
}

func (inv *Inventory) isFull() bool {
	return inv.space == 0
}
