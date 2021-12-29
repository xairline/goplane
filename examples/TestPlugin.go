package main

import (
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

var myMenuId menus.MenuID

func onPluginStart() {
	logging.Info("Plugin started")

	menuId := menus.FindPluginsMenu()
	menuContainerId := menus.AppendMenuItem(menuId, "TestPlugin Menu", nil, false)
	myMenuId = menus.CreateMenu("TestPlugin Menu", menuId, menuContainerId, menuHandler, nil)
	menus.AppendMenuItem(myMenuId, "TestPlugin Menu sub 1", "TestPlugin Menu sub 1", false)
	menus.AppendMenuItem(myMenuId, "TestPlugin Menu sub 2", "TestPlugin Menu sub 2", false)
}

func menuHandler(menuRef, itemRef interface{}) {
	logging.Infof("clicked: %+v", itemRef)
}

func onPluginStop() {
	menus.DestroyMenu(myMenuId)
	logging.Info("Plugin stopped")
}

func onPluginEnable() {
	logging.Info("Plugin enabled")
}

func onPluginDisable() {
	logging.Info("Plugin disabled")
}
