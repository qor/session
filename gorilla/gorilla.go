package gorilla

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/qor/session"
)

func New(sessionName string, store sessions.Store) *Gorilla {
	return &Gorilla{SessionName: sessionName, Store: store}
}

type Gorilla struct {
	SessionName string
	Store       sessions.Store
}

func (gorilla Gorilla) Add(req *http.Request, key string, value interface{}) error {
	if session, err := gorilla.Store.Get(req, gorilla.SessionName); err == nil {
		session.Values[key] = value
	} else {
		return err
	}
	return nil
}

func (gorilla Gorilla) Pop(req *http.Request, key string) string {
	return ""
}

func (gorilla Gorilla) Get(req *http.Request, key string) string {
	if session, err := gorilla.Store.Get(req, gorilla.SessionName); err == nil {
		if value, ok := session.Values[key]; ok {
			return fmt.Sprint(value)
		}
	}
	return ""
}

func (gorilla Gorilla) Flash(req *http.Request, message session.Message) {
}

func (gorilla Gorilla) Load(req *http.Request, key string, result interface{}) {
}

func (gorilla Gorilla) Save(req *http.Request, w http.ResponseWriter) error {
	return nil
}
