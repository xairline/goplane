//Copyright (c) 2015. The goplane AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package scenery

/*
#cgo CFLAGS: -I../../SDK/CHeaders -fPIC -DSIMDATA_EXPORTS -DXPLM200=1 -DXPLM210=1 -DXPLM300=1 -DXPLM301=1 -DXPLM302=1 -DXPLM303=1
#include <XPLM/XPLMScenery.h>
#include <stdlib.h>
#include <string.h>

extern void objectLoadedCallback(XPLMObjectRef inObject, void* inRefcon);
*/
import "C"
import (
	"unsafe"

	"github.com/xairline/goplane/xplm/graphics"
)

type ProbeType int
type ProbeResult int
type ProbeRef unsafe.Pointer

type ProbeInfo struct {
	size      int32
	LocationX float32
	LocationY float32
	LocationZ float32
	NormalX   float32
	NormalY   float32
	NormalZ   float32
	VelocityX float32
	VelocityY float32
	velocityZ float32
	Is_wet    int32
}

const (
	ProbeY          ProbeType   = 0
	ProbeHitTerrain ProbeResult = 0
	ProbeError      ProbeResult = 1
	ProbeMissed     ProbeResult = 2
)

func CreateProbe(probeType ProbeType) ProbeRef {
	return ProbeRef(C.XPLMCreateProbe(C.XPLMProbeType(probeType)))
}

func DestroyProbe(probeRef ProbeRef) {
	C.XPLMDestroyProbe(C.XPLMProbeRef(probeRef))
}

func ProbeTerrainXYZ(probeRef ProbeRef, x, y, z float32) (ProbeResult, ProbeInfo) {
	info := ProbeInfo{}
	info.size = int32(unsafe.Sizeof(info))
	result := ProbeResult(C.XPLMProbeTerrainXYZ(C.XPLMProbeRef(probeRef), C.float(x), C.float(y), C.float(z), (*C.XPLMProbeInfo_t)(unsafe.Pointer(&info))))
	return result, info
}

func ProbeTerrainLatLonAlt(probeRef ProbeRef, lat, lon, alt float32) (ProbeResult, ProbeInfo) {
	x64, y64, z64 := graphics.WorldToLocal(float64(lat), float64(lon), float64(alt))
	return ProbeTerrainXYZ(probeRef, float32(x64), float32(y64), float32(z64))
}
