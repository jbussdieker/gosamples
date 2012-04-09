package gtk
/*
#cgo pkg-config: gtk+-2.0
#include <gtk/gtk.h>
*/
import "C"

type GtkWidget struct {
	Widget *C.GtkWidget
}

type GtkWindowType int
const GTK_WINDOW_TOPLEVEL GtkWindowType = 0
type GtkWindow struct {
	GtkWidget
}

func Init() {
	C.gtk_init(nil, nil)
}

func Window(t GtkWindowType) *GtkWindow {
	return &GtkWindow{GtkWidget{
		C.gtk_window_new(C.GtkWindowType(t))}}
}

func (v *GtkWidget) Show() {
	C.gtk_widget_show(v.Widget)
}

func Main() {
	C.gtk_main()
}

