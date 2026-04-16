package processor

func Compiler(text string) string {

	text = Hex(text)
	text = Bin(text)

	text = Article(text)
	text = Fixpunctuations(text)
	text = Quote(text)
	text = Case(text)
	return text
}
