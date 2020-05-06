package icapclient

import (
	"bufio"
	"net"
)

// transport represents the transport layer data
type transport struct {
	network string
	addr    string
	sckt    net.Conn
}

// Dial fires up a tcp socket
func (t *transport) dial() error {
	sckt, err := net.Dial(t.network, t.addr)

	if err != nil {
		return err
	}

	t.sckt = sckt

	return nil
}

// Write writes data to the server
func (t *transport) write(data string) (int, error) {
	return t.sckt.Write([]byte(data))
}

// Read reads data from server
func (t *transport) read() (string, error) {

	respMsg, err := bufio.NewReader(t.sckt).ReadString('\n')

	if err != nil {
		return "", err
	}

	return respMsg, nil
}

func (t *transport) close() error {
	return t.sckt.Close()
}
