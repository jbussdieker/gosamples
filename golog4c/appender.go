package log4c

/*
#cgo LDFLAGS: -L/usr/local/lib -llog4c
#cgo CFLAGS:  -I/usr/local/include
#include <stdlib.h>
#include "log4c.h"
static inline void log_string(log4c_category_t *cat, int priority, char *str) {
	log4c_category_error(cat, str);
}
*/
import "C"
import "unsafe"

type Appender struct {
	Ptr *C.log4c_appender_t
}

func makeAppender(ptr *C.log4c_appender_t) *Appender {
	if ptr == nil {
		return nil
	}
	return &Appender{ptr}
}

func NewAppender(name string) *Appender {
	ptr := C.CString(name)
	defer C.free(unsafe.Pointer(ptr))
	return makeAppender(C.log4c_appender_new(ptr))
}

func GetAppender(name string) *Appender {
	ptr := C.CString(name)
	defer C.free(unsafe.Pointer(ptr))
	return makeAppender(C.log4c_appender_get(ptr))
}

