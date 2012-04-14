package log4c

/*
#cgo LDFLAGS: -L/usr/local/lib -llog4c
#cgo CFLAGS:  -I/usr/local/include
#include <stdlib.h>
#include "log4c.h"
static inline void log_string(log4c_category_t *cat, int priority, char *str) {
	log4c_category_log(cat, priority, str);
}
*/
import "C"
import "unsafe"

type Layout struct {
	Ptr *C.log4c_layout_t
}

func makeLayout(ptr *C.log4c_layout_t) *Layout {
	if ptr == nil {
		return nil
	}
	return &Layout{ptr}
}

func GetLayout(name string) *Layout {
	ptr := C.CString(name)
	defer C.free(unsafe.Pointer(ptr))
	return makeLayout(C.log4c_layout_get(ptr))
}


