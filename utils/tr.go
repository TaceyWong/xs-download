package utils

import (
	"os"
	"strings"
)

var LANG = func() string {
	envLang := strings.ToLower(os.Getenv("LANG"))
	var lang string
	switch {
	case strings.Contains(envLang, "en"):
		lang = LangEN
	case strings.Contains(envLang, "zh"):
		lang = LangZH
	default:
		lang = LangZH
	}
	return lang
}

const (
	LangZH = "zh_CN"
	LangEN = "en_US"
)

type Word struct {
	Ori   string
	MapTo map[string]string
}

var words = []*Word{
	{Ori: "中国", MapTo: map[string]string{LangEN: "China"}},
	{Ori: "验证", MapTo: map[string]string{LangEN: "Verify"}},
	{Ori: "获取密钥", MapTo: map[string]string{LangEN: "Get Key"}},
}

var Words = make(map[string]*Word)

func init() {
	for _, w := range words {
		Words[w.Ori] = w
	}
}

func (w Word) String() string {
	return w.Ori
}

func Tr(text string) string {
	if _, ok := Words[text]; ok && LANG() != LangZH {
		return Words[text].MapTo[LANG()]
	} else {
		return text
	}
}
