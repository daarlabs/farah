package farah

import "github.com/daarlabs/hirokit/hiro"

func Plugin(props ...Props) hiro.Plugin {
	if len(props) > 0 {
		Config.Spinner = props[0].Spinner
	}
	return hiro.Plugin{
		Name:    "Farah",
		Locales: locales,
	}
}
