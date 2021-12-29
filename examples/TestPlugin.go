package main

import (
	"fmt"

	"github.com/xairline/goplane/extra"
	"github.com/xairline/goplane/extra/logging"
	"github.com/xairline/goplane/xplm/menus"
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
	case extra.PluginStop:
		onPluginStop()
	case extra.PluginEnable:
		onPluginEnable()
	case extra.PluginDisable:
		onPluginDisable()
	}
}

func onPluginStart() {
	logging.Info("Plugin started")
	menuId := menus.FindPluginsMenu()
	logging.Info(fmt.Sprintf("menuId: %x", menuId))

	menus.ClearAllMenuItems(menuId)

	menuContainerId := menus.AppendMenuItem(menuId, "TestPlugin Menu", 0, false)
	logging.Info(fmt.Sprintf("menuContainerId: %x", menuContainerId))
}

func onPluginStop() {
	logging.Info("Plugin stopped")
}

func onPluginEnable() {
	logging.Info("Plugin enabled")
}

func onPluginDisable() {
	logging.Info("Plugin disabled")
}
