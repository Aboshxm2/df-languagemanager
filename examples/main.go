package main

import (
	"fmt"

	languagemanager "github.com/Aboshxm2/df-languagemanager"
)

func main() {
	languagemanager.LoadLanguages("./examples/langs/")

	playerName := "Aboshxm2"

	fmt.Println(languagemanager.Translate("something.something", []string{}, ""))        // something.something
	fmt.Println(languagemanager.Translate("core.chat.join", []string{}, ""))             // %0% joined the game
	fmt.Println(languagemanager.Translate("core.chat.join", []string{playerName}, ""))   // Aboshxm2 joined the game
	fmt.Println(languagemanager.Translate("core.chat.join", []string{playerName}, "ar")) // Aboshxm2 دخل اللعبة
}
