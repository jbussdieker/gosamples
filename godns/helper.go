package dns

import "bytes"
import "fmt"
import "encoding/binary"

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

func write16(buf *bytes.Buffer, value uint16) {
	err := binary.Write(buf, binary.BigEndian, value)
	if err != nil {
		panic("write16 failed")
	}
}

func write8(buf *bytes.Buffer, value uint8) {
	err := binary.Write(buf, binary.BigEndian, value)
	if err != nil {
		panic("write16 failed")
	}
}
