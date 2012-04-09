package main

import . "xorg"

func main() {
	d := OpenDisplay()
	s := DefaultScreen(d)
	w := CreateSimpleWindow(d, RootWindow(d, s), 100, 100, 640, 480, 1, BlackPixel(d, s), WhitePixel(d, s))

	SelectInput(d, w, ExposureMask | KeyPressMask)
	MapWindow(d, w)

	for {
		e := NextEvent(d)
		if e.Type() == Expose {
		    FillRectangle(d, Drawable(w), DefaultGC(d, s), 20, 20, 10, 10)
			DrawString(d, Drawable(w), DefaultGC(d, s), 10, 50, "This is a test")
		}
		if e.Type() == KeyPress {
			break
		}
	}

	CloseDisplay(d)
}

