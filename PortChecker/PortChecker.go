package portchecker

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

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
