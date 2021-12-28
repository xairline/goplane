package extra

/*
#include <stdlib.h>
#include <string.h>
*/
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/xairline/goplane/extra/logging"
	"github.com/xairline/goplane/xplm/plugins"
	"github.com/xairline/goplane/xplm/processing"
	"github.com/xairline/goplane/xplm/utilities"
)

//basicStructureForAnXPlanePlugin
type XPlanePlugin struct {
	id                 plugins.PluginId
	name               string
	signature          string
	description        string
	messageHandler     plugins.MessageHandler
	flightLoop         processing.FlightLoopFunc
	flightLoopInterval float32
	errorCallback      utilities.ErrorCallback
	stateCallback      PluginStateCallback
}

//statusOfAPlugInForTheCallbackFunktion
type PluginState int

const (
	PluginStart   PluginState = 0 // Plugin is started
	PluginEnable  PluginState = 1 // Plugin is activated
	PluginDisable PluginState = 2 // Plugin is deactivated
	PluginStop    PluginState = 3 // Plugin is stopped
)

//Callback-Function for the status change of a plugin.
type PluginStateCallback func(state PluginState, plugin *XPlanePlugin)

//Message handler that writes the received message as a debug message.
func DebugMessageHandler(msg plugins.Message) {
	name, _, _, _ := plugins.GetPluginInfo(msg.PluginId)
	logging.Debug(fmt.Sprintf("receive message from %v (ID: %v): %v", name, msg.PluginId, msg.MessageId))
}

var (
	plugin *XPlanePlugin // Global plugin instance
)

// creates a new plugin.
func NewPlugin(name, signature, description string) *XPlanePlugin {
	logging.PluginName = name
	plugin = &XPlanePlugin{plugins.NO_PLUGIN_ID, name, signature, description, nil, nil, 1.0, nil, nil}
	logging.Info("================================================================")
	logging.Info(fmt.Sprintf("Plugin %v initialized", name))
	logging.Info(fmt.Sprintf("  signature: %v", signature))
	logging.Info(fmt.Sprintf("  description: %v", description))
	logging.Info("================================================================")
	return plugin
}

// Returns the ID of the plugin.
func (self *XPlanePlugin) GetId() plugins.PluginId {
	if self.id == -plugins.NO_PLUGIN_ID {
		self.id = plugins.GetMyId()
	}
	return self.id
}

// supplies the name of the plugin.
func (self *XPlanePlugin) GetName() string {
	return self.name
}

// Returns the description of the plugin.
func (self *XPlanePlugin) GetDescription() string {
	return self.description
}

// delivers the signature of the plugin.
func (self *XPlanePlugin) GetSignature() string {
	return self.signature
}

// Returns the messager of the plugin.
func (self *XPlanePlugin) GetMessageHandler() plugins.MessageHandler {
	return self.messageHandler
}

// Sets the gauge handerer of the plugin.
func (self *XPlanePlugin) SetMessageHandler(handler plugins.MessageHandler) {
	self.messageHandler = handler
}

// Sets the Flightloop feature that is to be automatically registered when the plug-in starts.
func (self *XPlanePlugin) SetFlightLoopFunc(flightLoopFunc processing.FlightLoopFunc, interval float32) {
	self.flightLoop = flightLoopFunc
	self.flightLoopInterval = interval
}

// Sets the ErrorCallback function that is to be automatically registered when plug-ins starts.
func (self *XPlanePlugin) SetErrorCallback(callback utilities.ErrorCallback) {
	self.errorCallback = callback
}

// Sets the callback function for plugin status change
func (self *XPlanePlugin) SetPluginStateCallback(callback PluginStateCallback) {
	self.stateCallback = callback
}

func (self *XPlanePlugin) onStart(name, sig, desc *C.char) {
	copyStringToCPointer(self.name, name)
	copyStringToCPointer(self.signature, sig)
	copyStringToCPointer(self.description, desc)
	if self.errorCallback != nil {
		utilities.SetErrorCallback(self.errorCallback)
	}
	if self.flightLoop != nil {
		processing.RegisterFlightLoopCallback(self.flightLoop, self.flightLoopInterval, self)
	}
}

func copyStringToCPointer(text string, target *C.char) {
	cMsg := C.CString(text)
	defer C.free(unsafe.Pointer(cMsg))
	C.strcpy(target, cMsg)
}

func (self *XPlanePlugin) String() string {
	return fmt.Sprintf("%v (singature: %v, id: %v)", self.GetName(), self.GetSignature(), self.GetId())
}

// external interface method addressed by X-Plane when receiving a message
//export XPluginReceiveMessage
func XPluginReceiveMessage(pluginId C.int, messageId C.int, messageData unsafe.Pointer) {
	if plugin.messageHandler != nil {
		plugin.messageHandler(plugins.Message{plugins.PluginId(pluginId), plugins.MessageId(messageId), messageData})
	}
}

// external interface method addressed by X-Plane when starting the plugin
//export XPluginStart
func XPluginStart(outName *C.char, outSig *C.char, outDesc *C.char) int {
	plugin.onStart(outName, outSig, outDesc)
	if plugin.stateCallback != nil {
		plugin.stateCallback(PluginStart, plugin)
	}
	return 1
}

// external interface method addressed by X-Plane when activating the plugin
//export XPluginEnable
func XPluginEnable() int {
	if plugin.stateCallback != nil {
		plugin.stateCallback(PluginEnable, plugin)
	}
	return 1
}

// external interface method addressed by X-Plane when deactivating the plugin
//export XPluginDisable
func XPluginDisable() {
	if plugin.stateCallback != nil {
		plugin.stateCallback(PluginDisable, plugin)
	}

}

// external interface method addressed by X-Plane when stopping the plugin
//export XPluginStop
func XPluginStop() {
	if plugin.stateCallback != nil {
		plugin.stateCallback(PluginStop, plugin)
	}
}
