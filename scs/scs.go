package scs

import (
	"encoding/json"
	"net/http"

	scssession "github.com/alexedwards/scs/session"
	"github.com/qor/session"
)

func New(engine scssession.Engine) *SCS {
	return &SCS{Engine: engine}
}

type SCS struct {
	scssession.Engine
}

func (scs SCS) Add(req *http.Request, key string, value interface{}) error {
	if str, ok := value.(string); ok {
		return scssession.PutString(req, key, str)
	}
	result, _ := json.Marshal(value)
	return scssession.PutString(req, key, string(result))
}

func (scs SCS) Pop(req *http.Request, key string) string {
	result, _ := scssession.PopString(req, key)
	return result
}

func (scs SCS) Get(req *http.Request, key string) string {
	result, _ := scssession.GetString(req, key)
	return result
}

func (scs SCS) Flash(req *http.Request, message session.Message) error {
	var messages []session.Message
	if err := scs.Load(req, "_flashes", &messages); err != nil {
		return err
	}
	messages = append(messages, message)
	return scs.Add(req, "_flashes", messages)
}

func (scs SCS) Flashes(req *http.Request) []session.Message {
	var messages []session.Message
	scs.PopLoad(req, "_flashes", &messages)
	return messages
}

func (scs SCS) Load(req *http.Request, key string, result interface{}) error {
	value := scs.Get(req, key)
	return json.Unmarshal([]byte(value), result)
}

func (scs SCS) PopLoad(req *http.Request, key string, result interface{}) error {
	value := scs.Pop(req, key)
	return json.Unmarshal([]byte(value), result)
}

func (scs SCS) Middleware(handler http.Handler) http.Handler {
	return scssession.Manage(scs.Engine)(handler)
}
