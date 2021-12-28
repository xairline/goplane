package extra

/*
#include <stdlib.h>
#include <string.h>
*/
import "C"
import (
	"unsafe"

	"github.com/xairline/goplane/xplm/plugins"
)

// Registers the transferred DataRefs either at the DataRefEditor (http://www.xsquawkbox.net/xpsdk/mediawiki/datareefeditor) or the DataTheTool (https://github.com/leecbaker/datareefttool)
func RegisterDataRefToDataRefEditor(name ...string) bool {
	pluginId := plugins.FindPluginBySignature("xplanesdk.examples.DataRefEditor")
	if pluginId == plugins.NO_PLUGIN_ID {
		// DataRefeditor not available -> evt. DataRefTool?
		pluginId = plugins.FindPluginBySignature("com.leecbaker.datareftool")
		if pluginId == plugins.NO_PLUGIN_ID {
			return false
		}
	}
	for _, current := range name {
		cName := C.CString(current)
		defer C.free(unsafe.Pointer(cName))
		msg := plugins.Message{pluginId, plugins.MessageId(0x01000000), unsafe.Pointer(cName)}
		plugins.SendMessageToPlugin(msg)
	}
	return true
}
