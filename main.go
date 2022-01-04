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
		fmt.Println("check many(from to) ports <./NetScanner -sm -ip 192.168.0.1 -pr 22 444 -t 2>")

		fmt.Println("check many(from to) ports and show only opened <./NetScanner -smo -ip 192.168.0.1 -pr 22 444>")
		fmt.Println("check many(from to) ports and show only opened <./NetScanner -smo -ip 192.168.0.1 -pr 22 444 -t 2>")

	case "-s":
		// scan only one port

		if len(os.Args) < 6 {
			fmt.Println("[warning] for scanning one port program need 5 params")
			return
		}

		host, port, err := GetHostAndPort(os.Args)
		if err != nil {
			fmt.Printf("[warning] %s\n", err)
			return
		}

		fmt.Printf("%s host is scanning\n", host)
		status := CheckPort(host, port, 3*time.Second)

		port_number, _ := strconv.Atoi(port)
		port_name, port_status := GetPortType(port_number)

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

		host, port_s, port_e, timeout, err := GetHostAndPorts(os.Args)
		if err != nil {
			fmt.Printf("[warning] %s\n", err)
			return
		}

		fmt.Printf("%s host is scanning\n", host)
		var status, port_status bool
		var port_name string

		for port := port_s; port <= port_e; port++ {
			if timeout != 0 {
				status = CheckPort(host, strconv.Itoa(port), time.Duration(timeout)*time.Second)
			} else {
				status = CheckPort(host, strconv.Itoa(port), 3*time.Second)
			}

			port_name, port_status = GetPortType(port)
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
			fmt.Println("[warning] for scanning many ports program need min 6 params")
			return
		}

		host, port_s, port_e, timeout, err := GetHostAndPorts(os.Args)
		if err != nil {
			fmt.Printf("[warning] %s\n", err)
			return
		}

		fmt.Printf("%s host is scanning\n", host)
		var status, port_status bool
		var port_name string

		for port := port_s; port <= port_e; port++ {
			if timeout != 0 {
				status = CheckPort(host, strconv.Itoa(port), time.Duration(timeout)*time.Second)
			} else {
				status = CheckPort(host, strconv.Itoa(port), 3*time.Second)
			}

			port_name, port_status = GetPortType(port)
			if status {
				fmt.Printf("port %d is opened", port)
				if port_status {
					fmt.Printf(" (%s)", port_name)
				}
				fmt.Println()
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

func GetHostAndPorts(Args []string) (host string, port_s, port_e, timeout int, err error) {
	for i, box := range Args {
		switch box {
		case "-ip":
			host = Args[i+1]
		case "-pr":
			port_s, err = strconv.Atoi(Args[i+1])
			if err != nil {
				return "", 0, 0, 0, err
			}
			port_e, err = strconv.Atoi(Args[i+2])
			if err != nil {
				return "", 0, 0, 0, err
			}
		case "-t":
			timeout, err = strconv.Atoi(Args[i+1])
			if err != nil {
				return "", 0, 0, 0, err
			}
		}
	}

	return
}

func CheckPort(host, port string, timeout time.Duration) bool {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)

	if err != nil {
		return false
	}

	if conn != nil {
		defer conn.Close()
		return true
	}

	return false
}

func GetPortType(port_number int) (port_name string, status bool) {
	switch port_number {
	case 20:
		return "FTP", true
	case 21:
		return "FTP", true
	case 22:
		return "SSH", true
	case 23:
		return "TELNET", true
	case 25:
		return "SMTP", true
	case 110:
		return "SMTP", true
	case 143:
		return "SMTP", true
	case 80:
		return "HTTP", true
	case 443:
		return "HTTPS", true
	case 53:
		return "DOMAIN", true
	case 81:
		return "HOSTS2-NS", true
	case 194:
		return "IRC", true
	case 4445:
		return "UPNOTIFYP", true
	case 8888:
		return "Althttpd", true
	case 9999:
		return "Crypto or ...", true
	default:
		return "", false
	}
}
