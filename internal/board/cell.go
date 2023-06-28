package board

type Cell struct {
	Items []Item
	North *Cell
	East  *Cell
	South *Cell
	West  *Cell
}

func NewCell() *Cell {
	return &Cell{
		Items: nil,
		North: nil,
		East:  nil,
		South: nil,
		West:  nil,
	}
}

func (c *Cell) AddItem(item Item) {
	c.Items = append(c.Items, item)
}

func (c *Cell) RemoveItem(item Item) {
	for i := 0; i < len(c.Items); i++ {
		if c.Items[i] == item {
			c.Items = append(c.Items[:i], c.Items[i+1:]...)
			break
		}
	}
}

func (c *Cell) HasItem(item Item) bool {
	for _, cellItem := range c.Items {
		if cellItem == item {
			return true
		}
	}

	return false
}

// CleamItems removes all the items from the cell
func (c *Cell) CleamItems() {
	c.Items = nil
}
