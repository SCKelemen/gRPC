package Daemon

import (
	"github.com/takama/daemon"
)

type Service struct {
	daemon.Daemon
}

func (service *Service) Manage() (string, error) {
	usage := "Usage: intern install | remove | start | stop | status"

}

const (
	installPlugin = "searchsvc install <plugin>"  // install a plugin
	removePlugin  = "searchsvc remove <plugin>"   // uninstall a plugin
	update        = "searchsvc update"            // update plugin definitions
	upgrade       = "searchsvc upgrade ?<plugin>" // installs latest versions of plugins
	list          = "searchsvc list"              // list installed plugins
	start         = "searchsvc start"             // start the graph
	stop          = "searchsvc stop"              // stop the graph
	status        = "searchsvc status"            // get the status
)
