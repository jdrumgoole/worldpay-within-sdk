package rpc

type Configuration struct {

	Protocol string
	Framed bool
	Buffered bool
	Host string
	Port int
	Secure bool
	BufferSize int
}