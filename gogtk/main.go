package main

import . "gtk"

func main() {
	Init()
	w := Window(GTK_WINDOW_TOPLEVEL)
	w.Show()
	Main()
}

