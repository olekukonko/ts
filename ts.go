package ts

// Return System Size
type Size struct {
	row    uint16
	col    uint16
	posX   uint16
	posY   uint16
}

// Get Terminal Width
func (w Size) Col() int {
	return int(w.row)
}

// Get Terminal Height
func (w Size) Row() int {
	return int(w.row)
}

// Get Position X
func (w Size) PosX() int {
	return int(w.posX)
}

// Get Position Y
func (w Size) PosY() int {
	return int(w.posY)
}
