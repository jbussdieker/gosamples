package xml

import "testing"
import "fmt"
import . "libxml"

type parserTest struct {
	input   string
	output  string
	options XmlParseOptions
}

var tests []parserTest = []parserTest{
	{
		input:   "",
		output:  "",
		options: 0,
	},
	{
		input:   "<p/>",
		output:  "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<p/>\n",
		options: 0,
	},
	{
		input:   "<p></p>",
		output:  "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<p/>\n",
		options: 0,
	},
}

func TestParser(t *testing.T) {
	for _, test_case := range tests {
		doc := NewFromString(test_case.input, "", "UTF-8", test_case.options)
		if doc.String() != test_case.output {
			fmt.Println("**GOT*******************************")
			fmt.Println(doc.String())
			fmt.Println("**EXPECTED**************************")
			fmt.Println(test_case.output)
			fmt.Println("************************************")
			t.Fail()
		}
	}
}
