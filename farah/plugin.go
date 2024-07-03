package farah

import "github.com/daarlabs/hirokit/hiro"

func Plugin() hiro.Plugin {
	return hiro.Plugin{
		Name:    "Farah",
		Locales: locales,
	}
}
