package chopstick

// Collection of ingredients to build a element or UI
type Bento []Ingredients

// Create a new Bento with all the ingredients
func NewBento(ingredients ...Ingredients) Bento {
	return Bento(ingredients)
}

// Draws the bento but returns chopstick to orginal position
func (b *Bento) DrawWithReturn(c *chopstick) {
	prevPosition := Position{c.position.X, c.position.Y}
	for _, i := range *b {
		i.Draw(c)
	}
	c.MoveTo(prevPosition)
}

// Draws the bento but leaves the chopstick at end of drawing
func (b *Bento) Draw(c *chopstick) {
	for _, i := range *b {
		i.Draw(c)
	}
}
