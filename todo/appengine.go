// +build appengine

// A Google App Engine application providing a web UI for task management.
// DISCLAIMER: This is not intended to be used, since instances can be
// restarted at any time and all tasks would be lost.
package todo

import "github.com/campoy/todo/server"

func init() {
	server.RegisterHandlers()
}
