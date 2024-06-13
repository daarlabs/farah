package farah

import "github.com/daarlabs/arcanum/mirage"

func Plugin() mirage.Plugin {
	return mirage.Plugin{
		Name:    "Farah",
		Locales: locales,
	}
}
