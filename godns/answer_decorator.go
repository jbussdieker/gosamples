package dns

import "fmt"

func (answer *Answer) String() (str string) {
	str += fmt.Sprintf("              Name: %v\n", answer.Name)
	str += fmt.Sprintf("              Type: %v\n", answer.Type)
	str += fmt.Sprintf("             Class: %v\n", answer.Class)
	str += fmt.Sprintf("               TTL: %v\n", answer.TTL)
	str += fmt.Sprintf("          DataSize: %v\n", answer.DataSize)
	return
}

