package libxml

/*
#cgo pkg-config: libxml-2.0
#include <stdlib.h>
#include <string.h>
#include <libxml/parser.h>
*/
import "C"
import "unsafe"

type XmlParseOptions uint32

func XmlReadMemory(buffer string, url string, encoding string, options XmlParseOptions) unsafe.Pointer {
	cbuffer := C.CString(buffer)
	curl := C.CString(url)
	cencoding := C.CString(encoding)
	cdoc := C.xmlReadMemory(cbuffer, C.int(C.strlen(cbuffer)), curl, cencoding, C.int(options))
	C.free(unsafe.Pointer(cbuffer))
	C.free(unsafe.Pointer(curl))
	C.free(unsafe.Pointer(cencoding))
	return unsafe.Pointer(cdoc)
}
