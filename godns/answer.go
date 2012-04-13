package dns

import "fmt"
import "bytes"
import "encoding/binary"

type Answer struct {
	Name string
	Type uint16
	Class uint16
	TTL uint32
	RDLEN uint16
	RDATA[] byte
}

func readString(buffer *bytes.Buffer) (str string) {
	size := int((buffer.Next(1))[0])
	// Message pointer
	if size == 0xC0 {
		return fmt.Sprintf("[message pointer %v]", buffer.Next(1))
	}
	for size != 0 {
		str += string(buffer.Next(size))
		size = int((buffer.Next(1))[0])
		if size != 0 {
			str += "."
		}
	}
	return
}

func ParseAnswer(buffer *bytes.Buffer) (answer *Answer) {
	answer = &Answer{}
	answer.Name = readString(buffer)
	binary.Read(buffer, binary.BigEndian, &answer.Type)
	binary.Read(buffer, binary.BigEndian, &answer.Class)
	binary.Read(buffer, binary.BigEndian, &answer.TTL)
	binary.Read(buffer, binary.BigEndian, &answer.RDLEN)
	answer.RDATA = buffer.Next(int(answer.RDLEN))
	return answer
}

func (answer *Answer) Bytes() []byte {
	return []byte{0}
}

func (answer *Answer) String() (str string) {
	str += fmt.Sprintf("              Name: %v\n", answer.Name)
	str += fmt.Sprintf("              Type: %v\n", answer.Type)
	str += fmt.Sprintf("             Class: %v\n", answer.Class)
	str += fmt.Sprintf("               TTL: %v\n", answer.TTL)
	str += fmt.Sprintf("             RDLEN: %v\n", answer.RDLEN)
	return
}
