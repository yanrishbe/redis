package flagsClient

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yanrishbe/redis/main/client/getData"
	"net"
	"os"
	"testing"
	"time"
)

type testpair struct {
	port int
	host string
}

var testvalues = []testpair{
	{65555, "127.0.0.1"},
	{0, "172.17.0.1"},
	{-2, "17.0.0.1"},
}

func TestFlags(t *testing.T) {
	assert := assert.New(t)
	for _, value := range testvalues {
		boolHost, errHost := ValidHost(value.host)
		errPort := ValidPort(value.port)
		assert.Error(errPort, "Port error")
		assert.NoError(errHost)
		assert.Equal(boolHost, true)
	}

}

func MockServer(t *testing.T) {
	li, err := net.Listen("tcp", ":9092")
	require.NoError(t, err)
	conn, err := li.Accept()
	require.NoError(t, err, conn.Write([]byte("test")))
	defer func() {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}()

}

func TestOutputToClient(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)
	go MockServer(t)
	time.Sleep(time.Millisecond * 100)
	conn, err := net.Dial("tcp", ":9092")
	require.NoError(err)
	defer func() {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}()

	file, _ := os.Create("test.txt")
	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()

	assert.NotEmpty(t, getData.OutputToClient(conn, file))
}
