package Application

type IApplication interface {
	Install() (bool, error)
	Uninstall() (bool, error)
	Start() (bool, error)
}
