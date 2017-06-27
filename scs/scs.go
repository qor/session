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

func (scs SCS) Add(req *http.Request, key string, value interface{}) {
}

func (scs SCS) Pop(req *http.Request, key string) string {
	return ""
}

func (scs SCS) Get(req *http.Request, key string) string {
	return ""
}

func (scs SCS) Flash(req *http.Request, message session.Message) {
}

func (scs SCS) Load(req *http.Request, key string, result interface{}) {
}

func (scs SCS) Save(req *http.Request, w http.ResponseWriter) {
}
