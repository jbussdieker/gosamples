package hello

/*
#cgo pkg-config: libhello

#include <hello.h>

int _Hello() {
	return Hello();
}
*/
import "C"

func Hello() {
	a := C._Hello()
	println(a)
}

