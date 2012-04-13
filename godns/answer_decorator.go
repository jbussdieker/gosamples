package dns

import "fmt"

////////////////////////////////////////////////////////////////////////////////
// Public functions
////////////////////////////////////////////////////////////////////////////////

func NewAnswer() *Answer {
	return &Answer{}
}

////////////////////////////////////////////////////////////////////////////////
// Method functions
////////////////////////////////////////////////////////////////////////////////

func (answer *Answer) String() (str string) {
	// Print a string (usefull for debugging)
	str += fmt.Sprintf("              Name: %v\n", answer.Name)
	str += fmt.Sprintf("              Type: %v\n", answer.Type)
	str += fmt.Sprintf("             Class: %v\n", answer.Class)
	str += fmt.Sprintf("               TTL: %v\n", answer.TTL)
	str += fmt.Sprintf("          DataSize: %v\n", answer.DataSize)
	if answer.Type == RECORD_TYPE_TXT {
		str += fmt.Sprintf("              Data: %v\n", string(answer.Data))
	}

	return
}
