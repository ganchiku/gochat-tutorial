package chat

import (
	"log"

	"golang.org/x/net/context"
)

// Server is sturct
type Server struct {
}

// SayHello is the function to return response as server
func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message body from client :%s", message.Body)
	return &Message{Body: "Hello from the server!"}, nil
}
