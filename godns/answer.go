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

func readString(buf []byte) (str string) {
	size := buf[0]
	// Message pointer
	if size == 0xC0 {
		println("MESSAGE POINTER:", buf[1])
		return ""
	}
	var index byte = 0
	for size != 0 {
		index++
		str += string(buf[index:index+size])
		index += size
		size = buf[index]
		if size != 0 {
			str += "."
		}
	}
	return
}

func ParseAnswer(buf []byte) (*Answer, []byte) {
	answer := &Answer{}
	answer.Name = readString(buf)

	buffer := bytes.NewBuffer(buf[len(answer.Name)+2:])
	binary.Read(buffer, binary.BigEndian, &answer.Type)
	binary.Read(buffer, binary.BigEndian, &answer.Class)
	binary.Read(buffer, binary.BigEndian, &answer.TTL)
	binary.Read(buffer, binary.BigEndian, &answer.RDLEN)
	answer.RDATA = buffer.Next(int(answer.RDLEN))
	return answer, buffer.Bytes()
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
