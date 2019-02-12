package server

import "flag"

type Option struct {
	Mode string
}

func ParseOption() Option {
	var option Option
	flag.StringVar(&option.Mode, "m", "local", "Specify mode, production or staging or local. default is local.")
	flag.Parse()
	return option
}
