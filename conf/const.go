package conf

import (
	"net/http"
)

var (
	Debug      bool   // is debug command
	Help       bool   // is help command
	Version    bool   // is print version command
	ConfigFile string // config file
	SkipUpdate bool   // skip update

	Client *http.Client // request client

	Origins []string // allow origins
)

var Conf = new(Config)

const (
	VERSION = "v0.1"
)
