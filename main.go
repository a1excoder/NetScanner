package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("[warning] after bin name write params, <--help>")
		return
	}

	switch os.Args[1] {
	case "--help":
		fmt.Println("check only one port <./NetScanner -s -ip 192.168.0.1 -port 443>")
		fmt.Println("check many(from to) ports <./NetScanner -sm -ip 192.168.0.1 -pr 22 444>")
	case "-s":
		host, port, err := GetHostAndPort(os.Args)
		if err != nil {
			fmt.Printf("[warning] %s\n", err)
			return
		}

		fmt.Printf("%s host is scanning\n", host)
		status, err := CheckPort(host, port)
		if err != nil {
			fmt.Println(err)
		}
		if status {
			fmt.Printf("port %s is opened\n", port)
		} else {
			fmt.Printf("port %s is closed\n", port)
		}
	case "-sm":
		host, port_s, port_e, err := GetHostAndPorts(os.Args)
		if err != nil {
			fmt.Printf("[warning] %s\n", err)
			return
		}

		fmt.Printf("%s host is scanning\n", host)
		var status bool
		for port := port_s; port <= port_e; port++ {
			status, err = CheckPort(host, strconv.Itoa(port))
			if err != nil {
				fmt.Println(err)
			}
			if status {
				fmt.Printf("port %d is opened\n", port)
			} else {
				fmt.Printf("port %d is closed\n", port)
			}
		}
	default:
		fmt.Println("[warning] unknown command")
	}
}

func GetHostAndPort(Args []string) (string, string, error) {
	var host, port string

	for i, box := range Args {
		switch box {
		case "-ip":
			host = Args[i+1]
		case "-port":
			port = Args[i+1]
		}
	}

	if host == "" || port == "" {
		return "", "", fmt.Errorf("host or port not found")
	}

	return host, port, nil
}

func GetHostAndPorts(Args []string) (string, int, int, error) {
	var host string
	var port_s, port_e int
	var err error

	for i, box := range Args {
		switch box {
		case "-ip":
			host = Args[i+1]
		case "-pr":
			port_s, err = strconv.Atoi(Args[i+1])
			if err != nil {
				return "", 0, 0, err
			}
			port_e, err = strconv.Atoi(Args[i+2])
			if err != nil {
				return "", 0, 0, err
			}
		}
	}

	return host, port_s, port_e, nil
}

func CheckPort(host, port string) (bool, error) {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), 3*time.Second)

	if err != nil {
		return false, nil
	}

	if conn != nil {
		defer conn.Close()
		return true, nil
	}

	return false, nil
}
