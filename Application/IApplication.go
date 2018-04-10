package Application

type IApplication interface {
	Mount() (bool, error)
	Mounted() (bool, error)
	Dismount() (bool, error)
	Dismounted() (bool, error)
	Initialize() (bool, error)
}

type Application struct {
}

func (a Application) Mount() (bool, error) {

	go a.Initialize()
	return true, nil
}

func (a Application) Mounted() (bool, error) {
	return true, nil
}

func (a Application) Dismount() (bool, error) {
	return true, nil
}
func (a Application) Dismounted() (bool, error) {
	return true, nil
}
func (a Application) Initialize() (bool, error) {
	return true, nil
}
