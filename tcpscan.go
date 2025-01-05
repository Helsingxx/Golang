package main

import (
	"fmt"
	"net"
	"os"
	"sort"
)

func worker(ports, result chan int) {
	for i := range ports {
		conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%v", i))
		if err == nil {
			result <- i
			conn.Close()
		} else {
			result <- 0
		}
	}
}

func main() {
	portarr, err := parse("1-65535")
	ports := make(chan int, 100)
	result := make(chan int)
	if err != nil {
		os.Exit(1)
	}
	var portlist []int
	for i := 0; i < cap(ports); i++ {
		go worker(ports, result)
	}
	go func() {
		for _, i := range portarr {
			ports <- i
		}
	}()
	for i := 0; i < len(portarr); i++ {
		port := <- result
		if port != 0 {
			portlist = append(portlist, port)
		}
	}
	close(ports)
	close(result)
	sort.Ints(portlist)
	for _, port := range portlist {
		fmt.Printf("Port %v is open.\n", port)
	}
}
