package main

import (
	"github.com/xairline/goplane/extra"
	"github.com/xairline/goplane/extra/logging"
)

func main() {
}

func init() {
	plugin := extra.NewPlugin("TestPlugin", "com.github.xairline.goplane.TestPlugin", "TestPlugin")
	plugin.SetPluginStateCallback(onPluginStateChanged)
	logging.MinLevel = logging.Info_Level
}

func onPluginStateChanged(state extra.PluginState, plugin *extra.XPlanePlugin) {
	switch state {
	case extra.PluginStart:
		onPluginStart()
	case extra.PluginEnable:
		onPluginEnable()
	}
}

func onPluginStart() {
	logging.Info("Plugin start")
}

func onPluginEnable() {
	logging.Info("Plugin enable")
}
