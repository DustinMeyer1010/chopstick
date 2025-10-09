package chopstick

// Collection of ingridents to build a element or UI
type Bento []Ingredients

// Item to draw on the screen and where it will be drawn
type Ingredients struct {
	Position Position
	Value    string
}

func (b *Bento) Draw(c *chopstick) {
	prevPosition := Position{c.position.X, c.position.Y}
	for _, i := range *b {
		i.Draw(c)
	}
	c.MoveTo(prevPosition)
}

func (i *Ingredients) Draw(c *chopstick) {
	c.MoveTo(i.Position)
	c.DrawTextWithReturn(i.Value)
}
