package util

import "flag"

var (
	ConfigPath string
)

func init() {
	flag.StringVar(&ConfigPath, "config file path", "./conf/server.json", "record the path of config file")
}

func Parse() {
	flag.Parse()

}
