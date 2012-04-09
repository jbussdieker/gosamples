package node

import "unsafe"

import . "libxml"
import . "libxml/element"

type XmlNode struct {
	doc unsafe.Pointer
	ptr unsafe.Pointer
}

func NodeFromDoc(doc unsafe.Pointer) Node {
	return &XmlNode{doc: doc, ptr: doc}
}

func (n *XmlNode) NewChild(name string, content string) Node {
	newnode := XmlNewDocRawNode(n.doc, nil, name, content)
	XmlAddChild(n.ptr, newnode)
	return &XmlNode{doc: n.doc, ptr: newnode}
}

func (n *XmlNode) GetAttribute(name string) string {
	return XmlGetProp(n.ptr, name)
}

func (n *XmlNode) Child() Element {
	return ElementFromNode(n.doc, n.ptr).Child()
}

func (n *XmlNode) Children() chan Element {
	return ElementFromNode(n.doc, n.ptr).Children()
}

func (n *XmlNode) Name() string {
	return ElementFromNode(n.doc, n.ptr).Name()
}

func (n *XmlNode) String() string {
	return XmlNodeDump(n.ptr)
}
