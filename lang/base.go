package lang

import (
	"fmt"
	"go-fiber-react/config"
	"strings"
)

var L l

type l struct{}

func (*l) Convert(msg string, args ...map[string]any) string {
	if len(args) == 0 || args[0] == nil {
		return msg
	}

	newMsg := msg
	for key, value := range args[0] {
		newMsg = strings.ReplaceAll(newMsg, ":"+key, fmt.Sprintf("%v", value))
	}
	return newMsg
}

func (*l) Get() ILang {
	return locale[config.APP_LOCALE]
}

var locale = map[string]ILang{
	"en": EN,
	"id": ID,
}
