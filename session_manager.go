package session

import "net/http"

// ManagerInterface session manager interface
type ManagerInterface interface {
	Add(req *http.Request, key string, value interface{}) error
	Pop(req *http.Request, key string) string
	Get(req *http.Request, key string) string

	Flash(req *http.Request, message Message) error
	Flashes(req *http.Request) []Message
	Load(req *http.Request, key string, result interface{}) error
	PopLoad(req *http.Request, key string, result interface{}) error

	Middleware(http.Handler) http.Handler
}

// Message message struct
type Message struct {
	Message string
	Type    string
}
