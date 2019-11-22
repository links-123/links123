package cmd

import "flag"

type Flags struct {
	Kind    string
	EnvFile string
	Address string

	ShowVersionOnly bool
}

func LoadFlags() *Flags {
	flags := &Flags{}

	flag.StringVar(&flags.Kind, "kind", "", "Kind of micro service")
	flag.StringVar(&flags.EnvFile, "env", "", "Environment file")
	flag.StringVar(&flags.Address, "address", ":8080", "Web API address")
	flag.BoolVar(&flags.ShowVersionOnly, "version", false, "App version")

	flag.Parse()

	return flags
}
