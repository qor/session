package gorilla_test

import (
	"net/http"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/qor/session"
	"github.com/qor/session/gorilla"
)

var SessionManager session.ManagerInterface

func init() {
	engine := sessions.NewCookieStore([]byte("something-very-secret"))
	SessionManager = gorilla.New("_session", engine)
}

func TestSesionManager(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	if err := SessionManager.Add(req, "key", "value"); err != nil {
		t.Errorf("Should add session correctly, but got %v", err)
	}

	if value := SessionManager.Get(req, "key"); value != "value" {
		t.Errorf("failed to fetch saved session value, got %#v", value)
	}
}
