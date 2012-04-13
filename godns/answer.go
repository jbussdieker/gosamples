package dns

import "fmt"
import "bytes"
import "encoding/binary"

type Answer struct {
	Name string
	Type uint16
	Class uint16
	TTL uint32
	DataSize uint16
	Data[] byte
}

func readDnsString(buffer *bytes.Buffer) (str string) {
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
	answer.Name = readDnsString(buffer)
	binary.Read(buffer, binary.BigEndian, &answer.Type)
	binary.Read(buffer, binary.BigEndian, &answer.Class)
	binary.Read(buffer, binary.BigEndian, &answer.TTL)
	binary.Read(buffer, binary.BigEndian, &answer.DataSize)
	answer.Data = buffer.Next(int(answer.DataSize))
	return answer
}

func (answer *Answer) Bytes() []byte {
	// TODO: Implement marshal
	return []byte{0}
}

