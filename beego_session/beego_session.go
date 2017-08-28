package beego_session

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	beego_session "github.com/astaxie/beego/session"
	"github.com/qor/qor/utils"
	"github.com/qor/session"
)

var writer utils.ContextKey = "gorilla_writer"

// New initialize session manager for BeegoSession
func New(engine *beego_session.Manager) *BeegoSession {
	return &BeegoSession{Manager: engine}
}

// BeegoSession session manager struct for BeegoSession
type BeegoSession struct {
	*beego_session.Manager
}

func (beegosession BeegoSession) getSession(req *http.Request) (beego_session.Store, error) {
	if w, ok := req.Context().Value(writer).(http.ResponseWriter); ok {
		return beegosession.Manager.SessionStart(w, req)
	}
	panic("middleware not applied")
}

// Add value to session data, if value is not string, will marshal it into JSON encoding and save it into session data.
func (beegosession BeegoSession) Add(req *http.Request, key string, value interface{}) error {
	sess, _ := beegosession.getSession(req)
	defer sess.SessionRelease(nil)

	if str, ok := value.(string); ok {
		return sess.Set(key, str)
	}
	result, _ := json.Marshal(value)
	return sess.Set(key, string(result))
}

// Pop value from session data
func (beegosession BeegoSession) Pop(req *http.Request, key string) string {
	sess, _ := beegosession.getSession(req)
	defer sess.SessionRelease(nil)

	result := fmt.Sprint(sess.Get(key))

	sess.Delete(key)
	return result
}

// Get value from session data
func (beegosession BeegoSession) Get(req *http.Request, key string) string {
	sess, _ := beegosession.getSession(req)
	return fmt.Sprint(sess.Get(key))
}

// Flash add flash message to session data
func (beegosession BeegoSession) Flash(req *http.Request, message session.Message) error {
	var messages []session.Message
	if err := beegosession.Load(req, "_flashes", &messages); err != nil {
		return err
	}
	messages = append(messages, message)
	return beegosession.Add(req, "_flashes", messages)
}

// Flashes returns a slice of flash messages from session data
func (beegosession BeegoSession) Flashes(req *http.Request) []session.Message {
	var messages []session.Message
	beegosession.PopLoad(req, "_flashes", &messages)
	return messages
}

// Load get value from session data and unmarshal it into result
func (beegosession BeegoSession) Load(req *http.Request, key string, result interface{}) error {
	value := beegosession.Get(req, key)
	if value != "" {
		return json.Unmarshal([]byte(value), result)
	}
	return nil
}

// PopLoad pop value from session data and unmarshal it into result
func (beegosession BeegoSession) PopLoad(req *http.Request, key string, result interface{}) error {
	value := beegosession.Pop(req, key)
	if value != "" {
		return json.Unmarshal([]byte(value), result)
	}
	return nil
}

// Middleware returns a new session manager middleware instance
func (beegosession BeegoSession) Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx := context.WithValue(req.Context(), writer, w)
		handler.ServeHTTP(w, req.WithContext(ctx))
	})
}
