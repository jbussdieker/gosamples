package xml

import "testing"

func TestNodeString(t *testing.T) {
	if New("1.0").Child().String() != "" {
		t.Fail()
	}
}

func TestNodeAddChild(t *testing.T) {
	doc := New("1.0")
	if doc.Child().String() != "" {
		t.Fail()
	}
	child := doc.NewChild("test", "ing")
	if child == nil {
		t.Fail()
	}
	if child.String() != "<test>ing</test>" {
		t.Fail()
	}
}

func TestChild(t *testing.T) {
	doc := New("1.0")
	if doc.Child().String() != "" {
		t.Fail()
	}
	doc.NewChild("test", "ing")
	if doc.Child().String() != "<test>ing</test>" {
		t.Fail()
	}
}

/*
func TestNodeChildren(t *testing.T) {
	doc := NewFromString("<html><body><a></a><b></b><p></p></body></html>", "", "UTF-8", 0)
	p := doc.Child().Child()
	collect := ""
	for n := range p.Children() {
		collect += n.String()
	}
	if collect != "<a/><b/><p/>" {
		t.Fatal(collect)
	}
}
*/
