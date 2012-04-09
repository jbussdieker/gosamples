package libxml

/*
#cgo pkg-config: libxml-2.0
#include <stdlib.h>
#include <libxml/HTMLtree.h>
*/
import "C"
import "unsafe"

func HtmlNewDoc(uri string, external_id string) unsafe.Pointer {
	curi := cToXmlChar(C.CString(uri))
	cexternal_id := cToXmlChar(C.CString(external_id))
	cdoc := C.htmlNewDoc(curi, cexternal_id)
	C.free(unsafe.Pointer(curi))
	C.free(unsafe.Pointer(cexternal_id))
	return unsafe.Pointer(cdoc)
}

func HtmlNodeDump(n unsafe.Pointer) string {
	buf := C.xmlBufferCreate()
	C.htmlNodeDump(buf, nil, C.xmlNodePtr(n))
	cstr := C.xmlBufferContent(buf)
	str := C.GoString(xmlCharToC(cstr))
	C.xmlBufferFree(buf)
	return str
}
