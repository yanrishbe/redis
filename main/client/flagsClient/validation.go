package flagsClient

import (
	"errors"
	"log"
	"regexp"
)

func ValidFlags (port int, host string) (bool, error){
	if port < 1 || port > 65535 {
		log.Fatalln(func() error {
			return errors.New("incorrect port info")
		}())
	}

	patternHost := "^((\\d|[1-9]\\d|1(\\d){2}|2[0-4]\\d|25[0-5])\\.){3}(\\d|[1-9]\\d|1[(\\d){2}|2[0-4]\\d|25[0-5])$"
	matchHost, errHost := regexp.MatchString(patternHost, host)
	return matchHost, errHost
	//if !matchHost || errHost != nil {
	//	log.Fatalln(func() error {
	//		return errors.New("incorrect hort info")
	//	}())
	//}
}
