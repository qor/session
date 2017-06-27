package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/qor/session"
)

var Server *httptest.Server

type Site struct {
	SessionManager session.ManagerInterface
}

func (site Site) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/get":
		value := site.SessionManager.Get(req, req.URL.Query().Get("key"))
		w.Write([]byte(value))
	case "/pop":
		value := site.SessionManager.Pop(req, req.URL.Query().Get("key"))
		w.Write([]byte(value))
	case "/set":
		err := site.SessionManager.Add(req, req.URL.Query().Get("key"), req.URL.Query().Get("value"))
		if err != nil {
			panic(fmt.Sprintf("No error should happe when set session, but got %v", err))
		}
	}
}

func TestAll(manager session.ManagerInterface, t *testing.T) {
	Server = httptest.NewServer(manager.Middleware(Site{SessionManager: manager}))

	req, _ := http.NewRequest("GET", "/", nil)
	TestSesionManagerAddAndGet(req, manager, t)
	TestSesionManagerAddAndPop(req, manager, t)
	TestFlash(req, manager, t)
	TestLoad(req, manager, t)
	TestRequest(manager, t)
}

func TestRequest(manager session.ManagerInterface, t *testing.T) {
	resp, err := http.Get(Server.URL + "/set?key=key1&value=value1")
	if err != nil {
		t.Errorf("no error should happen when request set cookie")
	}

	cookieJar, _ := cookiejar.New(nil)
	url, _ := url.Parse(Server.URL)
	cookieJar.SetCookies(url, resp.Cookies())

	client := &http.Client{
		Jar: cookieJar,
	}

	resp, err = client.Get(Server.URL + "/get?key=key1")
	if err != nil {
		t.Errorf("no error should happend when request get cookie")
	}
	responseData, _ := ioutil.ReadAll(resp.Body)
	if string(responseData) != "value1" {
		t.Errorf("failed to get saved session")
	}
}

func TestSesionManagerAddAndGet(req *http.Request, manager session.ManagerInterface, t *testing.T) {
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

func TestSesionManagerAddAndGet1(req *http.Request, manager session.ManagerInterface, t *testing.T) {
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

func TestSesionManagerAddAndPop(req *http.Request, manager session.ManagerInterface, t *testing.T) {
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

func TestFlash(req *http.Request, manager session.ManagerInterface, t *testing.T) {
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

func TestLoad(req *http.Request, manager session.ManagerInterface, t *testing.T) {
	type result struct {
		Name    string
		Age     int
		Actived bool
	}

	user := result{Name: "jinzhu", Age: 18, Actived: true}
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
