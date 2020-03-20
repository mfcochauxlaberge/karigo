package freeport

import "net"

// Find finds and returns a free network port.
func Find() uint {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0
	}
	defer l.Close()

	return uint(l.Addr().(*net.TCPAddr).Port)
}
