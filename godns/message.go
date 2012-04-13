package dns

import "bytes"

////////////////////////////////////////////////////////////////////////////////
// Types
////////////////////////////////////////////////////////////////////////////////

type Message struct {
	*Header
	Questions[] *Question
	Answers[] *Answer
	Nameservers[] *Answer
	Additionals[] *Answer
}

////////////////////////////////////////////////////////////////////////////////
// Public functions
////////////////////////////////////////////////////////////////////////////////

func ParseMessage(buf []byte) (message *Message, err error) {
	buffer := bytes.NewBuffer(buf)

	message = &Message{
		Header: ParseHeader(buffer),
	}
	message.Questions = make([]*Question, message.QuestionCount)
	for i := 0; i < int(message.QuestionCount); i++ {
		message.Questions[i] = ParseQuestion(buffer)
	}
	message.Answers = make([]*Answer, message.AnswerCount)
	for i := 0; i < int(message.AnswerCount); i++ {
		message.Answers[i] = ParseAnswer(buffer)
	}
	message.Nameservers = make([]*Answer, message.NameserverCount)
	for i := 0; i < int(message.NameserverCount); i++ {
		message.Nameservers[i] = ParseAnswer(buffer)
	}
	message.Additionals = make([]*Answer, message.AdditionalCount)
	for i := 0; i < int(message.AdditionalCount); i++ {
		message.Additionals[i] = ParseAnswer(buffer)
	}
	if buffer.Len() > 0 {
		println("ERROR UNPARSED BYTES:", buffer.Len())
	}
	return
}

func (message *Message) Bytes() []byte {
	buf := new(bytes.Buffer)
	buf.Write(message.Header.Bytes())
	for i := 0; i < int(message.QuestionCount); i++ {
		buf.Write(message.Questions[i].Bytes())
	}
	for i := 0; i < int(message.AnswerCount); i++ {
		buf.Write(message.Answers[i].Bytes())
	}
	for i := 0; i < int(message.NameserverCount); i++ {
		buf.Write(message.Nameservers[i].Bytes())
	}
	for i := 0; i < int(message.AdditionalCount); i++ {
		buf.Write(message.Additionals[i].Bytes())
	}
	return buf.Bytes()
}

