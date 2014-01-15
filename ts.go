package ts

type Size struct {
	row           uint16
	col           uint16
	xpixel        uint16
	ypixel        uint16
}

func (w Size) Col() int {
	return int(w.row)
}

func (w Size) Row() int {
	return int(w.row)
}

func (w Size) Xpixel() int {
	return int(w.xpixel)
}

func (w Size) Ypixel() int {
	return int(w.ypixel)
}



