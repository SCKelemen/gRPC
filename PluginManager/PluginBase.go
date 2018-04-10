package PluginManager

type PluginBase struct {
	ID string
}

func (pb PluginBase) Name() string {
	return pb.ID
}

func (pb PluginBase) Mount() (bool, error) {
	return true, nil
}

func (pb PluginBase) Mounted() (bool, error) {
	return true, nil
}

func (pb PluginBase) Dismount() (bool, error) {
	return true, nil
}

func (pb PluginBase) Dismounted() (bool, error) {
	return true, nil
}

func (pb PluginBase) Dispatch(message string) (bool, error) {
	return true, nil
}

func NewPlugin(name string) PluginBase {
	return PluginBase{ID: name}
}
