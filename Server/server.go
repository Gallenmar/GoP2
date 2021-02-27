package main

import (
	"fmt"
	"net"
	"encoding/binary"
	"bytes"
	"github.com/nsf/termbox-go"
)

func main() {
	termbox.Init()
	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(3, 3, '*', termbox.ColorRed, termbox.ColorBlue)
	termbox.Flush()
	
	adr, err := net.ResolveUDPAddr("udp", "127.0.0.1:10234")
	if err != nil {
		fmt.Println(err)
		return
	}

	listener, err := net.ListenUDP("udp", adr)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		handleConnection(listener)
	}



	termbox.Close()
}

func handleConnection(con *net.UDPConn) {


	buf := make([]byte, 2000)
	n, err := con.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	buff := bytes.NewReader(buf[0:n])

	var data struct {
		X int
		Y int
	}
	err = binary.Read(buff, binary.LittleEndian, &data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(data.X)





}
