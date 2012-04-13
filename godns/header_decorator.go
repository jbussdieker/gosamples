package dns

import "fmt"

func (h *Header) String() (str string) {
	str += fmt.Sprintf("                ID: %d\n", h.ID)
	str += fmt.Sprintf("             Query: %v\n", h.Query)
	str += fmt.Sprintf("            OpCode: %d\n", h.OpCode)
	str += fmt.Sprintf("     Authoritative: %v\n", h.Authoritative)
	str += fmt.Sprintf("         Truncated: %v\n", h.Truncated)
	str += fmt.Sprintf("         Recursion: %v\n", h.Recursion)
	str += fmt.Sprintf("RecursionSupported: %v\n", h.RecursionSupported)
	str += fmt.Sprintf("          Reserved: %v\n", h.Reserved)
	str += fmt.Sprintf("      ResponseCode: %v\n", h.ResponseCode)
	str += fmt.Sprintf("           QDCOUNT: %v\n", h.QuestionCount)
	str += fmt.Sprintf("           ANCOUNT: %v\n", h.AnswerCount)
	str += fmt.Sprintf("           NSCOUNT: %v\n", h.NSCOUNT)
	str += fmt.Sprintf("           ARCOUNT: %v\n", h.ARCOUNT)
	return
}