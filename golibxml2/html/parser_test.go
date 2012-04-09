package html

import "testing"
import "fmt"
import . "libxml"

type parserTest struct {
	input   string
	output  string
	options HtmlParseOptions
}

var tests []parserTest = []parserTest{
	{
		input:   "",
		output:  "",
		options: 0,
	},
	{
		input:   "<p>a</p>",
		output:  "<!DOCTYPE html PUBLIC \"-//W3C//DTD HTML 4.0 Transitional//EN\" \"http://www.w3.org/TR/REC-html40/loose.dtd\">\n<html><body><p>a</p></body></html>\n",
		options: 0,
	},
	{
		input:   "<p></p>",
		output:  "<html><body><p></p></body></html>\n",
		options: HTML_PARSE_NODEFDTD,
	},
	// TODO: Find out if this is a bug
	{
		input:   "<p></p>",
		output:  "<html><body><p></p></body></html>\n",
		options: HTML_PARSE_NODEFDTD | HTML_PARSE_NOIMPLIED,
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
