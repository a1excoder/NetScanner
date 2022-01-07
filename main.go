package main

import (
	httpreqchecker "NetScanner/HttpReqChecker"
	portchecker "NetScanner/PortChecker"
	"fmt"
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
		fmt.Println("check many(from to) ports <./NetScanner -sm -ip 192.168.0.1 -pr 22 444 -t 2>")

		fmt.Println("check many(from to) ports and show only opened <./NetScanner -smo -ip 192.168.0.1 -pr 22 444>")
		fmt.Println("check many(from to) ports and show only opened <./NetScanner -smo -ip 192.168.0.1 -pr 22 444 -t 2>")

		fmt.Println("send an http request via Get and get a response <./NetScanner -rg https://jsonplaceholder.typicode.com/posts/1>")
	case "-s":
		// scan only one port

		if len(os.Args) < 6 {
			fmt.Println("[warning] for scanning one port program need 5 params")
			return
		}

		host, port, err := portchecker.GetHostAndPort(os.Args)
		if err != nil {
			fmt.Printf("[warning] %s\n", err)
			return
		}

		fmt.Printf("%s host is scanning\n", host)
		status := portchecker.CheckPort(host, port, 3*time.Second)

		port_number, _ := strconv.Atoi(port)
		port_name, port_status := portchecker.GetPortType(port_number)

		if status {
			fmt.Printf("port %s is opened", port)
			if port_status {
				fmt.Printf(" (%s)", port_name)
			}
			fmt.Println()
		} else {
			fmt.Printf("port %s is closed", port)
			if port_status {
				fmt.Printf(" (%s)", port_name)
			}
			fmt.Println()
		}
	case "-sm":
		// scan many ports

		if len(os.Args) < 7 {
			fmt.Println("[warning] for scanning many ports program need min 6 params")
			return
		}

		host, port_s, port_e, timeout, err := portchecker.GetHostAndPorts(os.Args)
		if err != nil {
			fmt.Printf("[warning] %s\n", err)
			return
		}

		fmt.Printf("%s host is scanning\n", host)
		var status, port_status bool
		var port_name string

		for port := port_s; port <= port_e; port++ {
			if timeout != 0 {
				status = portchecker.CheckPort(host, strconv.Itoa(port), time.Duration(timeout)*time.Second)
			} else {
				status = portchecker.CheckPort(host, strconv.Itoa(port), 3*time.Second)
			}

			port_name, port_status = portchecker.GetPortType(port)
			if status {
				fmt.Printf("port %d is opened", port)
				if port_status {
					fmt.Printf(" (%s)", port_name)
				}
				fmt.Println()
			} else {
				fmt.Printf("port %d is closed", port)
				if port_status {
					fmt.Printf(" (%s)", port_name)
				}
				fmt.Println()
			}
		}
	case "-smo":
		// scan many ports and show only open

		if len(os.Args) < 7 {
			fmt.Println("[warning] to scanning many ports program need min 6 params")
			return
		}

		host, port_s, port_e, timeout, err := portchecker.GetHostAndPorts(os.Args)
		if err != nil {
			fmt.Printf("[warning] %s\n", err)
			return
		}

		fmt.Printf("%s host is scanning\n", host)
		var status, port_status bool
		var port_name string

		for port := port_s; port <= port_e; port++ {
			if timeout != 0 {
				status = portchecker.CheckPort(host, strconv.Itoa(port), time.Duration(timeout)*time.Second)
			} else {
				status = portchecker.CheckPort(host, strconv.Itoa(port), 3*time.Second)
			}

			port_name, port_status = portchecker.GetPortType(port)
			if status {
				fmt.Printf("port %d is opened", port)
				if port_status {
					fmt.Printf(" (%s)", port_name)
				}
				fmt.Println()
			}
		}
	case "-rg":
		// send an http request via Get and get a response

		if len(os.Args) < 3 {
			fmt.Println("[warning] to send Get request need 1 params")
			return
		}

		answer, status_code, err := httpreqchecker.GetReq(os.Args[2])
		if err != nil {
			fmt.Printf("[warning] %s\n", err)
			return
		}

		fmt.Printf("Code status: %d\n", status_code)
		fmt.Println(answer)
	default:
		fmt.Println("[warning] unknown command")
	}
}
