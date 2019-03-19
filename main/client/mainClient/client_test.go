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
{65555, "127.0.0.1"},
{0,"172.17.0.1"},
{-2,"17.0.0.1"},
}

func TestFlags(t *testing.T) {
	//assert := assert.New(t)
	for _,value := range testvalues{
		boolHost, errHost:= flagsClient.ValidHost(value.host)
		errPort := flagsClient.ValidPort(value.port)
		assert.Error(t, errPort, "Port error")
		assert.NoError(t, errHost)
		assert.Equal(t,boolHost, true )
	}
}