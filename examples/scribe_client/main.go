package main

import (
	"fmt"
	"net"

	"github.com/FFengIll/go-thrift/examples/scribe"
	"github.com/FFengIll/go-thrift/thrift"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:1463")
	if err != nil {
		panic(err)
	}

	t := thrift.NewTransport(thrift.NewFramedReadWriteCloser(conn, 0), thrift.BinaryProtocol)
	client := thrift.NewClient(t, false)
	scr := scribe.ScribeClient{Client: client}
	res, err := scr.Log([]*scribe.LogEntry{{Category: "category", Message: "message"}})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response: %+v\n", res)
}
