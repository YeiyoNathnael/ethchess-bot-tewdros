package gemini

func markdownToMarkdownV2(message string) string {

	isDoubleAsterix := false

	newMessage := message
	for i := 1; i < len(newMessage); i++ {

		if newMessage[i] == '*' && newMessage[i-1] == '*' {
			isDoubleAsterix = true
		}

		if isDoubleAsterix {
			newMessage = newMessage[:i] + newMessage[i+1:]
			isDoubleAsterix = false
		}
	}

	return newMessage
}
