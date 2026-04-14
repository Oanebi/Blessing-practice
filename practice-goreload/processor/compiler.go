package processor

func Compiler(text string) string {
	text = Upper(text)

	text = Low(text)

	text = Hex(text)
	text = Bin(text)
	text = Cap(text)

	text = Article(text)
	text = Fixpunctuations(text)
	return Quote(text)
}
