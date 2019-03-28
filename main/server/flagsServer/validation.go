package flagsServer

import (
	"errors"
	"flag"
)

func InitFlags(port *int, defaultPort int, mode *string, defaultMode, usagePort, usageMode string){

	flag.IntVar(port, "port", defaultPort, usagePort)
	flag.IntVar(port, "p", defaultPort, "shorthand for --port")
	flag.StringVar(mode, "mode", defaultMode, usageMode)
	flag.StringVar(mode, "m", defaultMode, "shorthand for --mode")
	flag.Parse()
}

func ValidPort(port int) error {
	var errPort error
	if port < 1 || port > 65535 {
		errPort = errors.New("incorrect port info")
	}
	return errPort
}