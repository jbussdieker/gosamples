package log4c

/*
#cgo LDFLAGS: -L/usr/local/lib -llog4c
#cgo CFLAGS:  -I/usr/local/include
#include <stdlib.h>
#include <unistd.h>
#include "log4c.h"
static inline void log_string(log4c_category_t *cat, int priority, char *str) {
	log4c_category_error(cat, str);
	sleep(10);
}
*/
import "C"

const (
	FATAL = C.LOG4C_PRIORITY_FATAL
	ERROR = C.LOG4C_PRIORITY_ERROR
	INFO = C.LOG4C_PRIORITY_INFO
)

type Logger struct {
}

func New() *Logger {
	ptr := C.CString("log.cfg")
	if C.log4c_init() != 0 {
		return nil
	}
	C.log4c_rc_load(C.log4c_rc, ptr)
	return &Logger{}
}

func Free() {
	C.log4c_fini()
}


