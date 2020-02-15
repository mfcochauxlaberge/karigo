package plugin

import "net"

// Plugin ...
type Plugin struct {
	Actions map[string]map[string]Action
}

// Run ...
func (p *Plugin) Serve() error {
	ln, err := net.Listen("unix", "127.0.0.1:0")
	if err != nil {
		return err
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}

		go handler(conn)
	}

	return nil
}

func handler(conn net.Conn) {
	return nil
}
