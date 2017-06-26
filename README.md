# Session

Session Management for QOR

It wrap other session management into a common interface, allow you to use  https://github.com/alexedwards/scs, http://www.gorillatoolkit.org/pkg/sessions as the backend


Add(request, name, value string/interface{})
Pop(request, name) string
Get(request, name) string

Flash(request, session.Message{Message: message, Type: "warning"}) // alias of Add
Load(request, name, struct{}) // alias of Get

Save(r, w) // in middleware
