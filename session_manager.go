package session

import "net/http"

// ManagerInterface session manager interface
type ManagerInterface interface {
	Add(req *http.Request, key string, value interface{}) error
	Pop(req *http.Request, key string) string
	Get(req *http.Request, key string) string

	Flash(req *http.Request, message Message) error
	Load(req *http.Request, key string, result interface{}) error

	Save(req *http.Request, w http.ResponseWriter) error
}

// Message message struct
type Message struct {
	Message string
	Type    string
}
