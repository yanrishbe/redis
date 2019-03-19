package flagsClient

import "flag"

func InitFlags(port, defaultPort int, host, defaultHost, usagePort, usageHost string){

	flag.IntVar(&port, "port", defaultPort, usagePort)
	flag.IntVar(&port, "p", defaultPort, "shorthand for --port")
	flag.StringVar(&host, "host", defaultHost, usageHost)
	flag.StringVar(&host, "h", defaultHost, "shorthand for --host")
	flag.Parse()
}
