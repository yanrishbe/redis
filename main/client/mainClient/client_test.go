package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/yanrishbe/redis/main/client/flagsClient"
	"testing"
)

type testpair struct {
	port int
	host string
}

var testvalues = []testpair {
//{65555, "127.0.0.1:"},
{0,"127.0.0.1:"},
{9090,"127.0.0.1:"},git 
}

func TestFlags(t *testing.T) {
	for _,value := range testvalues{
		assert.Equal(t, flagsClient.ValidFlags(value.port, value.host), {true, !=nil})
		//if boolVal, errVal:= flagsClient.ValidFlags(value.port, value.host); !boolVal || errVal != nil{
		//	t.Error(
		//		"For", value,
		//		"expected","correct port and host",
		//		"got", value.port, value.host,
		//	)
		//}
	}
}
