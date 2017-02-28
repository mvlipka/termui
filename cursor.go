package termui

import tm "github.com/nsf/termbox-go"

// Cursor is a cursor object used to navigate the terminal
type Cursor struct {
	x int
	y int
}

// NewCursor creates a new cursor at X,Y
func NewCursor(x, y int) *Cursor {
	return &Cursor{x: x, y: y}
}

// GetPosition gets the cursor's current position
func (c *Cursor) GetPosition() (int, int) {
	return c.x, c.y
}

// MoveLeft moves the cursor left, within terminal bounds
func (c *Cursor) MoveLeft() {
	if c.x <= 0 {
		return
	}
	c.x--
	tm.SetCursor(c.x, c.y)
}

// MoveRight moves the cursor right, within terminal bounds
func (c *Cursor) MoveRight() {
	x, _ := tm.Size()
	if c.x >= x {
		return
	}
	c.x++
	tm.SetCursor(c.x, c.y)
}

// MoveDown moves the cursor down, within terminal bounds
func (c *Cursor) MoveDown() {
	_, y := tm.Size()
	if c.y >= y {
		return
	}
	c.y++
	tm.SetCursor(c.x, c.y)
}

// MoveUp moves the cursor up, within terminal bounds
func (c *Cursor) MoveUp() {
	if c.y <= 0 {
		return
	}
	c.y--
	tm.SetCursor(c.x, c.y)
}
