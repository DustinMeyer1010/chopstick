package chopstick

// Item to draw on the screen and where it will be drawn
type Ingredients struct {
	Position Position
	Value    string
}

// Draws the ingredient
//
// Note: Leaves chopstick at end of drawing
func (i *Ingredients) Draw(c *chopstick) {
	c.MoveTo(i.Position)
	c.DrawText(i.Value)
}

// Draws the ingredient
//
// Note: Returns chopstick to orginal position
func (i *Ingredients) DrawWithReturn(c *chopstick) {
	prevPosition := Position{X: c.position.X, Y: c.position.Y}
	c.MoveTo(i.Position)
	c.DrawText(i.Value)
	c.MoveTo(prevPosition)
}
