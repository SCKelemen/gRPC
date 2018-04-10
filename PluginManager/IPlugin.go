package PluginManager

type IPlugin interface {
	Name() string
	Mount() (bool, error)
	Mounted() (bool, error)
	Dismount() (bool, error)
	Dismounted() (bool, error)
	Dispatch(message string) (bool, error)
}
