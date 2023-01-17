# LanguageManager

Simple LanguageManager library for [dragonfly](https://github.com/df-mc/dragonfly) servers.

## Usage
First of all you need to load your language files
```go
import (
	languagemanager "github.com/Aboshxm2/df-languagemanager"
)

folder := "/path/to/your/languages/"
languagemanager.LoadLanguages(folder)
```

Then translate your messages like this
```go
import (
	languagemanager "github.com/Aboshxm2/df-languagemanager"
)

message := languagemanager.Translate(messageId, []string{...params}, languageCode)
```

## Examples
see https://github.com/Aboshxm2/df-languagemanager/examples/main.go
