package session

import "net/http"

// ManagerInterface session manager interface
type ManagerInterface interface {
	Add(req *http.Request, key string, value interface{})
	Pop(req *http.Request, key string) string
	Get(req *http.Request, key string) string

	Flash(req *http.Request, message Message)
	Load(req *http.Request, key string, result interface{})

	Save(req *http.Request, w http.ResponseWriter)
}

// Message message struct
type Message struct {
	Message string
	Type    string
}
