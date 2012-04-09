package xml

/*
#cgo pkg-config: libxml-2.0
#include <stdlib.h>
#include <libxml/tree.h>
*/
import "C"
import "unsafe"
import . "libxml"
import . "libxml/element"
import . "libxml/node"

type XmlDocument struct {
	ptr unsafe.Pointer
}

const (
	XML_PARSE_RECOVER    = C.XML_PARSE_RECOVER    //: recover on errors
	XML_PARSE_NOENT      = C.XML_PARSE_NOENT      //: substitute entities
	XML_PARSE_DTDLOAD    = C.XML_PARSE_DTDLOAD    //: load the external subset
	XML_PARSE_DTDATTR    = C.XML_PARSE_DTDATTR    //: default DTD attributes
	XML_PARSE_DTDVALID   = C.XML_PARSE_DTDVALID   //: validate with the DTD
	XML_PARSE_NOERROR    = C.XML_PARSE_NOERROR    //: suppress error reports
	XML_PARSE_NOWARNING  = C.XML_PARSE_NOWARNING  //: suppress warning reports
	XML_PARSE_PEDANTIC   = C.XML_PARSE_PEDANTIC   //: pedantic error reporting
	XML_PARSE_NOBLANKS   = C.XML_PARSE_NOBLANKS   //: remove blank nodes
	XML_PARSE_SAX1       = C.XML_PARSE_SAX1       //: use the SAX1 interface internally
	XML_PARSE_XINCLUDE   = C.XML_PARSE_XINCLUDE   //: Implement XInclude substitition
	XML_PARSE_NONET      = C.XML_PARSE_NONET      //: Forbid network access
	XML_PARSE_NODICT     = C.XML_PARSE_NODICT     //: Do not reuse the context dictionnary
	XML_PARSE_NSCLEAN    = C.XML_PARSE_NSCLEAN    //: remove redundant namespaces declarations
	XML_PARSE_NOCDATA    = C.XML_PARSE_NOCDATA    //: merge CDATA as text nodes
	XML_PARSE_NOXINCNODE = C.XML_PARSE_NOXINCNODE //: do not generate XINCLUDE START/END nodes
	XML_PARSE_COMPACT    = C.XML_PARSE_COMPACT    //: compact small text nodes; no modification of the tree allowed afterwards (will possibly crash if you try to modify the tree)
	XML_PARSE_OLD10      = C.XML_PARSE_OLD10      //: parse using XML-1.0 before update 5
	XML_PARSE_NOBASEFIX  = C.XML_PARSE_NOBASEFIX  //: do not fixup XINCLUDE xml:base uris
	XML_PARSE_HUGE       = C.XML_PARSE_HUGE       //: relax any hardcoded limit from the parser
	XML_PARSE_OLDSAX     = C.XML_PARSE_OLDSAX     //: parse using SAX2 interface from before 2.7.0
)

func New(version string) Document {
	return &XmlDocument{ptr: XmlNewDoc(version)}
}

func NewFromString(buffer string, url string, encoding string, options XmlParseOptions) Document {
	return &XmlDocument{ptr: XmlReadMemory(buffer, url, encoding, options)}
}

func (doc *XmlDocument) String() string {
	return XmlNodeDump(doc.ptr)
}

func (doc *XmlDocument) GetRootElement() Element {
	cnode := XmlDocGetRootElement(doc.ptr)
	return ElementFromDoc(cnode)
}

func (doc *XmlDocument) GetAttribute(name string) string {
	return NodeFromDoc(doc.ptr).GetAttribute(name)
}

func (doc *XmlDocument) NewChild(name string, content string) Node {
	return NodeFromDoc(doc.ptr).NewChild(name, content)
}

func (doc *XmlDocument) Child() Element {
	return ElementFromDoc(doc.ptr).Child()
}

func (doc *XmlDocument) Children() chan Element {
	return ElementFromDoc(doc.ptr).Children()
}

func (doc *XmlDocument) Name() string {
	return ElementFromDoc(doc.ptr).Name()
}


