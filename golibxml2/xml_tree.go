package libxml

/*
#cgo pkg-config: libxml-2.0
#include <stdlib.h>
#include <libxml/tree.h>
*/
import "C"
import "unsafe"

type ElementType uint32

func XmlDocGetRootElement(doc unsafe.Pointer) unsafe.Pointer {
	cnode := C.xmlDocGetRootElement(C.xmlDocPtr(doc))
	return unsafe.Pointer(cnode)
}

func XmlNewDocRawNode(doc unsafe.Pointer, ns unsafe.Pointer, name string, content string) unsafe.Pointer {
	cname := cToXmlChar(C.CString(name))
	ccontent := cToXmlChar(C.CString(content))
	cnode := C.xmlNewDocRawNode(C.xmlDocPtr(doc), C.xmlNsPtr(ns), cname, ccontent)
	C.free(unsafe.Pointer(cname))
	C.free(unsafe.Pointer(ccontent))
	return unsafe.Pointer(cnode)
}

func XmlAddChild(node unsafe.Pointer, child unsafe.Pointer) {
	C.xmlAddChild((*C.xmlNode)(node), (*C.xmlNode)(child))
}

func XmlGetProp(n unsafe.Pointer, name string) string {
	cname := C.CString(name)
	cstr := C.xmlGetProp(C.xmlNodePtr(n), cToXmlChar(cname))
	str := C.GoString(xmlCharToC(cstr))
	C.free(unsafe.Pointer(cname))
	return str
}

func XmlNodeDump(n unsafe.Pointer) string {
	buf := C.xmlBufferCreate()
	C.xmlNodeDump(buf, nil, C.xmlNodePtr(n), 0, 0)
	cstr := C.xmlBufferContent(buf)
	str := C.GoString(xmlCharToC(cstr))
	C.xmlBufferFree(buf)
	return str
}

func XmlNewDoc(version string) unsafe.Pointer {
	cversion := cToXmlChar(C.CString(version))
	cdoc := C.xmlNewDoc(cversion)
	C.free(unsafe.Pointer(cversion))
	return unsafe.Pointer(cdoc)
}

func XmlFirstElementChild(n unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.xmlFirstElementChild(C.xmlNodePtr(n)))
}

func XmlNextElementSibling(n unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.xmlNextElementSibling(C.xmlNodePtr(n)))
}
