package test

import (
	"net/http"
	"testing"

	"github.com/qor/session"
)

func TestAll(manager session.ManagerInterface, t *testing.T) {
	TestSesionManagerAddAndGet(manager, t)
	TestSesionManagerAddAndPop(manager, t)
}

func TestSesionManagerAddAndGet(manager session.ManagerInterface, t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	if err := manager.Add(req, "key", "value"); err != nil {
		t.Errorf("Should add session correctly, but got %v", err)
	}

	if value := manager.Get(req, "key"); value != "value" {
		t.Errorf("failed to fetch saved session value, got %#v", value)
	}

	if value := manager.Get(req, "key"); value != "value" {
		t.Errorf("possible to re-fetch saved session value, got %#v", value)
	}
}

func TestSesionManagerAddAndPop(manager session.ManagerInterface, t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	if err := manager.Add(req, "key", "value"); err != nil {
		t.Errorf("Should add session correctly, but got %v", err)
	}

	if value := manager.Pop(req, "key"); value != "value" {
		t.Errorf("failed to fetch saved session value, got %#v", value)
	}

	if value := manager.Pop(req, "key"); value == "value" {
		t.Errorf("can't re-fetch saved session value after get with Pop, got %#v", value)
	}
}
