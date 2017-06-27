package scs_test

import (
	"github.com/alexedwards/scs/engine/memstore"
	"github.com/qor/session"
	"github.com/qor/session/scs"
)

var SessionManager session.ManagerInterface

func init() {
	engine := memstore.New(0)
	SessionManager = scs.New(engine)
}
