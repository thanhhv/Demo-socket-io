package skio

import (
	"io"
	"net"
	"net/http"
	"net/url"
)

type Conn interface {
	io.Closer

	ID() string
	URL() url.URL
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	RemoteHeader() http.Header

	// Context of this connection. You can save one context for one
	// connection, and share it between all handlers. The handlers
	// is called in one goroutine, so no need to lock context if it
	// only be accessed in one connection.
	Context() interface{}
	SetContext(v interface{})
	Namespace() string
	Emit(msg string, v ...interface{})

	// Broadcast server side apis
	Join(room string)
	Leave(room string)
	LeaveAll()
	Rooms() []string
}

type AppSocket interface {
	Conn
}

type appSocket struct {
	Conn
}

func NewAppSocket(conn Conn) *appSocket {
	return &appSocket{conn}
}
