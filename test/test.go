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
	type result struct {
		Name    string
		Age     int
		Actived bool
	}

	user := result{Name: "jinzhu", Age: 18, Actived: true}
	req, _ := http.NewRequest("GET", "/", nil)
	manager.Add(req, "current_user", user)

	var user1 result
	if err := manager.Load(req, "current_user", &user1); err != nil {
		t.Errorf("no error should happen when Load struct")
	}

	if user1.Name != user.Name || user1.Age != user.Age || user1.Actived != user.Actived {
		t.Errorf("Should be able to add, load struct, ")
	}

	var user2 result
	if err := manager.Load(req, "current_user", &user2); err != nil {
		t.Errorf("no error should happen when Load struct")
	}

	if user2.Name != user.Name || user2.Age != user.Age || user2.Actived != user.Actived {
		t.Errorf("Should be able to load struct more than once")
	}

	var user3 result
	if err := manager.PopLoad(req, "current_user", &user3); err != nil {
		t.Errorf("no error should happen when PopLoad struct")
	}

	if user3.Name != user.Name || user3.Age != user.Age || user3.Actived != user.Actived {
		t.Errorf("Should be able to add, pop load struct")
	}

	var user4 result
	if err := manager.Load(req, "current_user", &user4); err != nil {
		t.Errorf("Should return error when fetch data after PopLoad")
	}
}
