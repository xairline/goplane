//Copyright (c) 2015. The goplane AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package navigation

/*
#cgo CFLAGS: -I../../SDK/CHeaders -fPIC -DSIMDATA_EXPORTS -DXPLM200=1 -DXPLM210=1 -DXPLM300=1 -DXPLM301=1 -DXPLM302=1 -DXPLM303=1
#cgo LDFLAGS: -shared
#include <XPLM/XPLMNavigation.h>
#include <stdlib.h>
#include <string.h>
*/
import "C"

func GetGPSDestinationType() NavType {
	return NavType(C.XPLMGetGPSDestinationType())
}

func GetGPSDestination() NavRef {
	return NavRef(C.XPLMGetGPSDestination())
}
