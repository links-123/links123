package cmd

import "flag"

type Flags struct {
	Kind            string
	EnvFile         string
	ShowVersionOnly bool
}

func LoadFlags() *Flags {
	flags := &Flags{}

	flag.StringVar(&flags.Kind, "kind", "", "Kind of micro service")
	flag.StringVar(&flags.EnvFile, "env", "", "Environment file")
	flag.BoolVar(&flags.ShowVersionOnly, "version", false, "App version")

	flag.Parse()

	return flags
}
