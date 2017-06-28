package gorilla

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	gorillaContext "github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"github.com/qor/qor/utils"
	"github.com/qor/session"
)

func New(sessionName string, store sessions.Store) *Gorilla {
	return &Gorilla{SessionName: sessionName, Store: store}
}

type Gorilla struct {
	SessionName string
	Store       sessions.Store
}

var writer utils.ContextKey = "writer"

func (gorilla Gorilla) saveSession(req *http.Request) {
	if session, err := gorilla.Store.Get(req, gorilla.SessionName); err == nil {
		if w, ok := req.Context().Value(writer).(http.ResponseWriter); ok {
			session.Save(req, w)
		}
	}
}

func (gorilla Gorilla) Add(req *http.Request, key string, value interface{}) error {
	defer gorilla.saveSession(req)

	session, err := gorilla.Store.Get(req, gorilla.SessionName)

	if err != nil {
		return err
	}

	if str, ok := value.(string); ok {
		session.Values[key] = str
	} else {
		result, _ := json.Marshal(value)
		session.Values[key] = string(result)
	}

	return nil
}

func (gorilla Gorilla) Pop(req *http.Request, key string) string {
	defer gorilla.saveSession(req)

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

func (gorilla Gorilla) Flashes(req *http.Request) []session.Message {
	var messages []session.Message
	gorilla.PopLoad(req, "_flashes", &messages)
	return messages
}

func (gorilla Gorilla) Load(req *http.Request, key string, result interface{}) error {
	value := gorilla.Get(req, key)
	if value != "" {
		return json.Unmarshal([]byte(value), result)
	}
	return nil
}

func (gorilla Gorilla) PopLoad(req *http.Request, key string, result interface{}) error {
	value := gorilla.Pop(req, key)
	if value != "" {
		return json.Unmarshal([]byte(value), result)
	}
	return nil
}

// Middleware session middleware
func (gorilla Gorilla) Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer gorillaContext.Clear(req)

		handler.ServeHTTP(w, req.WithContext(context.WithValue(req.Context(), writer, w)))
	})
}
