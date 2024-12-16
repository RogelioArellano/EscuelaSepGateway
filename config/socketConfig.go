package config

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

const (
	SEPHost        = "localhost"     // Dirección del microservicio EscuelaGateway
	SEPPort        = "7776"          // Puerto de conexión (ajusta según tu configuración)
	ConnectTimeout = 5 * time.Second // Timeout de conexión
)

type SocketClient struct {
	conn   net.Conn
	reader *bufio.Reader
	writer *bufio.Writer
}

// NewSocketClient crea una nueva instancia de SocketClient y establece la conexión con EscuelaGateway
func NewSocketClient() (*SocketClient, error) {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", SEPHost, SEPPort), ConnectTimeout)
	if err != nil {
		return nil, err
	}

	return &SocketClient{
		conn:   conn,
		reader: bufio.NewReader(conn),
		writer: bufio.NewWriter(conn),
	}, nil
}

// Send envía un mensaje a través del socket
func (s *SocketClient) Send(message []byte) error {
	_, err := s.writer.Write(message)
	if err != nil {
		return err
	}
	return s.writer.Flush()
}

// Receive recibe un mensaje desde el socket
func (s *SocketClient) Receive() ([]byte, error) {
	// Lee hasta un delimitador específico (ajusta según tu protocolo)
	message, err := s.reader.ReadBytes('\n')
	if err != nil {
		return nil, err
	}
	return message, nil
}

// Close cierra la conexión del socket
func (s *SocketClient) Close() error {
	return s.conn.Close()
}
