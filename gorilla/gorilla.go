package gorilla

import (
	"encoding/json"
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
	if session, err := gorilla.Store.Get(req, gorilla.SessionName); err == nil {
		if value, ok := session.Values[key]; ok {
			delete(session.Values, key)
			return fmt.Sprint(value)
		}
	}
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

func (gorilla Gorilla) Flash(req *http.Request, message session.Message) error {
	var messages []session.Message
	if err := gorilla.Load(req, "_flashes", &messages); err != nil {
		return err
	}
	messages = append(messages, message)
	return gorilla.Add(req, "_flashes", messages)
}

func (gorilla Gorilla) Load(req *http.Request, key string, result interface{}) error {
	value := gorilla.Get(req, key)
	return json.Unmarshal([]byte(value), result)
}

func (gorilla Gorilla) Save(req *http.Request, w http.ResponseWriter) error {
	session, err := gorilla.Store.Get(req, gorilla.SessionName)
	if err != nil {
		return err
	}
	return session.Save(req, w)
}
