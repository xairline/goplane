//Copyright (c) 2015. The goplane AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package plugins

/*
#cgo CFLAGS: -I../../SDK/CHeaders -fPIC -DSIMDATA_EXPORTS -DXPLM200=1 -DXPLM210=1 -DXPLM300=1 -DXPLM301=1 -DXPLM302=1 -DXPLM303=1
#include <XPLM/XPLMPlugin.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

type MessageId C.int

type Message struct {
	PluginId  PluginId
	MessageId MessageId
	Data      unsafe.Pointer
}

type MessageHandler func(Message)

const (
	MSG_PLANE_CRASHED          MessageId = 101
	MSG_PLANE_LOADED           MessageId = 102
	MSG_AIRPORT_LOADED         MessageId = 103
	MSG_SCENERY_LOADED         MessageId = 104
	MSG_AIRPLANE_COUNT_CHANGED MessageId = 105
	MSG_PLANE_UNLOADED         MessageId = 106
	MSG_WILL_WRITE_PREFS       MessageId = 107
	MSG_LIVERY_LOADED          MessageId = 108
)

func SendMessageToPlugin(msg Message) {
	C.XPLMSendMessageToPlugin(C.XPLMPluginID(msg.PluginId), C.int(msg.MessageId), msg.Data)
}
