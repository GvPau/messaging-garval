package bootstrap

import "github.com/GvPau/messaging-garval/internal/platform/server"

const (
	port = 8000
	host = "localhost"
)

func Run() error {
	server := server.New(host, port)
	return server.Run()
}
