package AppShell

import (
	"github.com/sckelemen/grpc/Application"
	"github.com/sckelemen/grpc/UpdateManager"
)

type ApplicationShell struct {
	Application   Application.Application
	UpdateManager UpdateManager.UpdateManager
}

func (as ApplicationShell) Reload() (bool, error) {

}

func (as ApplicationShell) LoadApplication(id string) (bool, error) {

}
