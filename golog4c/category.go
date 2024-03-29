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

type Category struct {
	Ptr *C.log4c_category_t
}

func NewCategory(name string) *Category {
	if C.log4c_init() != 0 {
		return nil
	}
	ptr := C.CString(name)
	defer C.free(unsafe.Pointer(ptr))
	return &Category{C.log4c_category_new(ptr)}
}

func GetCategory(name string) *Category {
	ptr := C.CString(name)
	defer C.free(unsafe.Pointer(ptr))
	return &Category{C.log4c_category_get(ptr)}
}

func (logger *Category) Logf(format string, data ...interface{}) {
	ptr := C.CString(format)
	defer C.free(unsafe.Pointer(ptr))
	C.log_string(logger.Ptr, INFO, ptr);
}

func (logger *Category) Errorf(format string, data ...interface{}) {
	ptr := C.CString(format)
	defer C.free(unsafe.Pointer(ptr))
	C.log_string(logger.Ptr, ERROR, ptr);
}

func (category *Category) SetAppender(appender *Appender) {
	C.log4c_category_set_appender(category.Ptr, appender.Ptr)
}

func (category *Category) SetPriority(priority int) int {
	return int(C.log4c_category_set_priority(category.Ptr, C.int(priority)))
}
