package dns

import "bytes"
import "strings"
import "fmt"
import "encoding/binary"

type Question struct {
	Name string
	Type RecordType
	Class ClassType
}

func NewQuestion(name string, rtype RecordType, class ClassType) *Question {
	return &Question{
		Name: name,
		Type: rtype,
		Class: class,
	}
}

func ParseQuestion(buffer *bytes.Buffer) (q *Question) {
	q = &Question{}
	q.Name = readString(buffer)
	binary.Read(buffer, binary.BigEndian, &q.Type)
	binary.Read(buffer, binary.BigEndian, &q.Class)
	return q
}

func (q *Question) Bytes() []byte {
	buf := new(bytes.Buffer)
	parts := strings.Split(q.Name, ".")
	for _, part := range parts {
		write8(buf, uint8(len(part)))
		buf.Write([]byte(part))
	}
	write8(buf, 0)
	write16(buf, uint16(q.Type))
	write16(buf, uint16(q.Class))
	return buf.Bytes()
}

func (q *Question) String() (str string) {
	str += fmt.Sprintf("              Name: %s\n", q.Name)
	str += fmt.Sprintf("              Type: %d\n", q.Type)
	str += fmt.Sprintf("             Class: %d\n", q.Class)
	return
}
