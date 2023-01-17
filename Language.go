package languagemanager

import (
	"strconv"
	"strings"
)

type Language struct {
	languageCode string
	messages     map[string]string
}

// return false if messageId is not found
func (lang *Language) Translate(messageId string, params []string) (string, bool) {
	for id, message := range lang.messages {
		if id == messageId {
			for i, param := range params {
				message = strings.Replace(message, "%"+strconv.Itoa(i)+"%", param, -1)
			}

			return message, true
		}
	}

	return messageId, false
}
