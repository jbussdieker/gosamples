package libxml
/*
#cgo CFLAGS: -I/usr/include/libxml2
#cgo LDFLAGS: -lxml2
void DoIt();
*/
import "C"

func DoStuff() {
	println("Start C")
	C.DoIt()
	println("End C")
}

