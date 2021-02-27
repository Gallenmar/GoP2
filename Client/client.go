package main

import (
	"fmt"
	"net"
	"encoding/binary"
	"bytes"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:10234")
	if err != nil {
		fmt.Println(err)
		return
	}

	var data struct {
		X int
		Y int
	}
	data.X = 6
	data.Y = 12

	var buf bytes.Buffer
	err = binary.Write(&buf, binary.LittleEndian, data)

	_, err = conn.Write(buf.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
	conn.Close()
}
