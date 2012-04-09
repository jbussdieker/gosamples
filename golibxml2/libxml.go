package libxml

/*
#cgo pkg-config: libxml-2.0
#include <stdlib.h>
#include <libxml/tree.h>
*/
import "C"
import "unsafe"

type Element interface {
	String() string
	Name() string
	Child() Element
	Children() chan Element
}

type Node interface {
	Element
	GetAttribute(name string) string
	NewChild(name string, content string) Node
}

type Document interface {
	Node
	GetRootElement() Element
}

func xmlCharToC(c *C.xmlChar) *C.char {
	return (*C.char)(unsafe.Pointer(c))
}

func cToXmlChar(c *C.char) *C.xmlChar {
	return (*C.xmlChar)(unsafe.Pointer(c))
}
