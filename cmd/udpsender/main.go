package main

import (
	"fmt"
	"log"
	"net"
	"io"
	"bufio"
	"os"
)

func main(){
	UDPAddress, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil{
		log.Fatal(err)
	}
	for{
		conn, err := net.DialUDP("udp", nil, UDPAddress)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		reader := bufio.NewReader(os.Stdin)

		for{
			fmt.Print("> ")

			line, err := reader.ReadString('\n')
			if err != nil{
				if err != io.EOF{
					log.Fatal(err)
				}
				return
			}
			_, err = conn.Write([]byte(line))
			if err != nil{
				log.Println(err)
			}
		}

	}
}