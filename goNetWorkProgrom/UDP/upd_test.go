package udp

import "testing"

func TestUDPServer(t *testing.T) {
	UDPServer()
}

func TestUDPClient(t *testing.T) {
	UDPClient()
}

func TestUDPUploadClient(t *testing.T) {
	UDPFileClient()
}

func TestUDPUploadServer(t *testing.T) {
	UDPFileServer()
}
