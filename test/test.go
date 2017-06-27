package test

import (
	"net/http"
	"testing"

	"github.com/qor/session"
)

func TestAll(manager session.ManagerInterface, t *testing.T) {
	TestSesionManagerAddAndGet(manager, t)
	TestSesionManagerAddAndPop(manager, t)
	TestFlash(manager, t)
	TestLoad(manager, t)
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

func TestFlash(manager session.ManagerInterface, t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	if err := manager.Flash(req, session.Message{
		Message: "hello1",
	}); err != nil {
		t.Errorf("No error should happen when add Flash, but got %v", err)
	}

	if err := manager.Flash(req, session.Message{
		Message: "hello2",
	}); err != nil {
		t.Errorf("No error should happen when add Flash, but got %v", err)
	}

	flashes := manager.Flashes(req)
	if len(flashes) != 2 {
		t.Errorf("should find 2 flash messages")
	}

	flashes2 := manager.Flashes(req)
	if len(flashes2) != 0 {
		t.Errorf("flash should be cleared when fetch it second time, but got %v", len(flashes2))
	}
}

func TestLoad(manager session.ManagerInterface, t *testing.T) {
}
