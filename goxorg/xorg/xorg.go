package xorg

/*
#cgo pkg-config: x11
#include <X11/Xlib.h>

int DefaultScreenMacro(Display *display) { return DefaultScreen(display); }

Window RootWindowMacro(Display *display, int screen_number) {
	return RootWindow(display, screen_number);
}

unsigned long BlackPixelMacro(Display *display, int screen_number) {
	return BlackPixel(display, screen_number);
}

unsigned long WhitePixelMacro(Display *display, int screen_number) {
	return WhitePixel(display, screen_number);
}

GC DefaultGCMacro(Display *display, int screen_number) {
	return DefaultGC(display, screen_number);
}

int GetTypeMacro(XEvent *e) {
	return e->type;
}
*/
import "C"

type Display *C.Display
type Window C.Window
type Drawable C.Drawable
type Event C.XEvent
type GC C.GC

const ExposureMask = C.ExposureMask
const KeyPressMask = C.KeyPressMask
const KeyPress = C.KeyPress
const Expose = C.Expose

func OpenDisplay() (d Display) {
	d = C.XOpenDisplay(nil)
	if d == nil {
		println("Error opening display")
	}
	return
}

func CloseDisplay(d Display) {
	C.XCloseDisplay(d)
}

func DefaultScreen(d Display) int {
	s := C.DefaultScreenMacro(d)
	return int(s)
}

func RootWindow(d Display, screen_number int) Window {
	w := C.RootWindowMacro(d, C.int(screen_number))
	return Window(w)
}

func BlackPixel(d Display, screen_number int) int {
	color := C.BlackPixelMacro(d, C.int(screen_number))
	return int(color)
}

func WhitePixel(d Display, screen_number int) int {
	color := C.WhitePixelMacro(d, C.int(screen_number))
	return int(color)
}

func CreateSimpleWindow(d Display, root Window, x int, y int, width int, height int, something int, color1 int, color2 int) Window {
	w := C.XCreateSimpleWindow(d, C.Window(root), C.int(x), C.int(y), C.uint(width), C.uint(height), C.uint(something), C.ulong(color1), C.ulong(color2))
	return Window(w)
}

func SelectInput(d Display, w Window, mask int) {
	C.XSelectInput(d, C.Window(w), C.long(mask));
}

func MapWindow(d Display, w Window) {
	C.XMapWindow(d, C.Window(w))
}

func NextEvent(d Display) Event {
	var e Event
	C.XNextEvent(d, (*C.XEvent)(&e))
	return e
}

func FillRectangle(d Display, obj Drawable, gc GC, x int, y int, width int, height int) {
	C.XFillRectangle(d, C.Drawable(obj), gc, C.int(x), C.int(y), C.uint(width), C.uint(height))
}

func DrawString(d Display, obj Drawable, gc GC, x int, y int, str string) {
	C.XDrawString(d, C.Drawable(obj), gc, C.int(x), C.int(y), C.CString(str), C.int(len(str)))
}

func DefaultGC(d Display, screen_number int) GC {
	gc := C.DefaultGCMacro(d, C.int(screen_number))
	return GC(gc) 
}

func (e Event) Type() int {
	t := C.GetTypeMacro((*C.XEvent)(&e))
	return int(t) 
}

