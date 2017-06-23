# Session

Session Management for QOR

It wrap other session management into a common interface, allow you to use  https://github.com/alexedwards/scs, http://www.gorillatoolkit.org/pkg/sessions as the backend


Add(request, name, value string/interface{})
Pop(request, name) string
Get(request, name) string
Load(request, name, struct{})

Save(r, w) // middleware
