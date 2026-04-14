package processor

func Compiler(text string) string {

	text = Hex(text)
	text = Bin(text)
	text = Upper(text)

	text = Article(text)
	text = Fixpunctuations(text)
	return Quote(text)
}
