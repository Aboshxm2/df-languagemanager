package languagemanager

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

var Languages []*Language

const DefaultLanguage = "en"

// return the translated message
// if language is not found it will try to translate from the default language
// if messageId is not found it will return the messageId
func Translate(messageId string, params []string, languageCode string) string {
	if languageCode == "" {
		languageCode = DefaultLanguage
	}

	for _, lang := range Languages {
		if lang.languageCode == languageCode {
			message, ok := lang.Translate(messageId, params)
			if ok {
				return message
			}
		}
	}

	// try translate with the default language
	if languageCode != DefaultLanguage {
		Translate(messageId, params, DefaultLanguage)
	}

	return messageId
}

func LoadLanguages(folder string) {
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
			languageCode := strings.TrimSuffix(file.Name(), ".json")

			bytes, err := ioutil.ReadFile(folder + file.Name())
			if err != nil {
				panic(err)
			}

			var m map[string]string
			errr := json.Unmarshal(bytes, &m)

			if errr != nil {
				panic(err)
			}

			lang := Language{
				languageCode: languageCode,
				messages:     m,
			}

			Languages = append(Languages, &lang)
		}
	}
}
