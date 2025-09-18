package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/alecthomas/kong"
	"go.bug.st/serial"
	"go.bug.st/serial/enumerator"
)

type CLI struct {
	BaudRate int    `name:"baudRate" help:"Baud rate for the connection" default:"115200"`
	PortName string `arg:"" optional:""`
}

func listPorts() {
	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		fmt.Println("No serial ports found!")
		return
	}
	fmt.Println("Found ports:")
	for _, port := range ports {
		fmt.Printf("- %s\n", port.Name)
		if port.IsUSB {
			fmt.Printf("\tProduct    %s\n", port.Product)
			fmt.Printf("\tUSB ID     %s:%s\n", port.VID, port.PID)
			fmt.Printf("\tUSB serial %s\n", port.SerialNumber)
		}
	}
}

func main() {
	var cli CLI
	_ = kong.Parse(&cli)

	if len(cli.PortName) == 0 {
		listPorts()
		return
	}

	var wg sync.WaitGroup

	mode := &serial.Mode{
		BaudRate: cli.BaudRate,
	}

	port, err := serial.Open(cli.PortName, mode)
	if err != nil {
		log.Fatal(err)
	}

	wg.Add(1)
	go (func() {
		w := io.Writer(os.Stdout)
		buff := make([]byte, 1024)
		for {
			n, err := port.Read(buff)
			if err != nil {
				log.Fatal(err)
				break
			}
			if n == 0 {
				break
			}
			w.Write(buff[0:n])
		}
		wg.Done()
	})()

	wg.Add(1)
	go (func() {
		r := io.Reader(os.Stdin)
		buff := make([]byte, 1024)
		for {
			n, err := r.Read(buff)
			if err != nil {
				log.Fatal(err)
				break
			}
			if n == 0 {
				break
			}
			port.Write(buff[0:n])
		}
		wg.Done()
	})()

	wg.Wait()
}
