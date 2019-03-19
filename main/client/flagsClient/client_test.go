package flagsClient

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/yanrishbe/redis/main/client/getData"
	"net"
	"os"

	"testing"
)

type testpair struct {
	port int
	host string
}

var testvalues = []testpair {
{65555, "127.0.0.1"},
{0,"172.17.0.1"},
{-2,"17.0.0.1"},
}

func TestFlags(t *testing.T) {
	assert := assert.New(t)
	for _,value := range testvalues{
		boolHost, errHost:= ValidHost(value.host)
		errPort := ValidPort(value.port)
		assert.Error(errPort, "Port error")
		assert.NoError(errHost)
		assert.Equal(boolHost, true )
	}

}
 func TestOutputToClient(t *testing.T){
 	file, _ := os.Create("test.txt")
 	defer func(){
 	file.Close()
	}()
	 conn, _ := net.Dial("tcp", "127.0.0.1:9090")
	 fmt.Fprintf(conn, "set a asf\n")
	 assert.NotEmpty(t, getData.OutputToClient(conn, file))
 }