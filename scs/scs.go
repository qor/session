package scs

import (
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
	return scssession.PutObject(req, key, value)
}

func (scs SCS) Pop(req *http.Request, key string) string {
	return ""
}

func (scs SCS) Get(req *http.Request, key string) string {
	result, _ := scssession.GetString(req, key)
	return result
}

func (scs SCS) Flash(req *http.Request, message session.Message) error {
	return nil
}

func (scs SCS) Load(req *http.Request, key string, result interface{}) error {
	return nil
}

func (scs SCS) Save(req *http.Request, w http.ResponseWriter) error {
	return nil
}
