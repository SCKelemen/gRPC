package PluginManager

import (
	"fmt"
)

type IPluginManager interface {
}

type PluginManager struct {
	Plugins map[string]IPlugin
}

func (pm PluginManager) Mount(name string) (bool, error) {
	plugin := NewPlugin(name)
	ok, err := plugin.Mount()
	if !ok {
		fmt.Println(err)
		return false, err
	}
	pm.Plugins[name] = plugin
	return plugin.Mounted()
}

func (pm PluginManager) Dismount(name string) (bool, error) {
	plugin, ok := pm.Plugins[name]
	if !ok {
		return false, nil
	}
	ok, err := plugin.Dismount()
	if !ok {
		fmt.Println(err)
		return false, nil
	}
	delete(pm.Plugins, name)
	return plugin.Dismounted()
}

func (pm PluginManager) Dispatch(plugin string, action string) (bool, error) {
	plugIn, ok := pm.Plugins[plugin]
	if !ok {
		return false, nil
	}
	// might want to run this in a different thread
	return plugIn.Dispatch(action)
}

func (pm PluginManager) DispatchAll(action string) (bool, map[string]error) {
	var result map[string]error
	good := true
	for plugin := range pm.Plugins {
		ok, err := pm.Plugins[plugin].Dispatch(action)
		if !ok {
			good = false
			result[plugin] = err
		}
	}
	return good, result

}
