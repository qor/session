package scs_test

import (
	"net/http"
	"testing"

	"github.com/alexedwards/scs/engine/memstore"
	"github.com/qor/session"
	"github.com/qor/session/scs"
)

var SessionManager session.ManagerInterface

func init() {
	engine := memstore.New(0)
	SessionManager = scs.New(engine)
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
