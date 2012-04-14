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

const (
	FATAL = C.LOG4C_PRIORITY_FATAL
	ERROR = C.LOG4C_PRIORITY_ERROR
	INFO = C.LOG4C_PRIORITY_INFO
)

func Init() int {
	return int(C.log4c_init())
}

func Free() {
	C.log4c_fini()
}


