package flagsClient

import (
	"errors"
	"regexp"
)

func ValidPort(port int) error {
	var errPort error
	if port < 1 || port > 65535 {
		errPort = errors.New("incorrect port info")
	}
	return errPort
}

func ValidHost(host string) (bool, error) {
	patternHost := "^((\\d|[1-9]\\d|1(\\d){2}|2[0-4]\\d|25[0-5])\\.){3}(\\d|[1-9]\\d|1[(\\d){2}|2[0-4]\\d|25[0-5])$"
		matchHost, errHost := regexp.MatchString(patternHost, host)
		return matchHost, errHost
}
