package flex

/*
#cgo CFLAGS: -I${SRCDIR}/../../third_party/tensorflow/headers
#ifndef GO_FLEX_TFLITE_H
#include "flex.go.h"
#endif
#cgo android,linux LDFLAGS: -lflex_delegate
#cgo linux LDFLAGS: -ldl -lflex_delegate
#cgo android LDFLAGS: -llog -lm
#cgo !android,!darwin,linux LDFLAGS: -L${SRCDIR}/../../third_party/tensorflow/libs/linux_x86/
#cgo android,arm64 LDFLAGS: -L${SRCDIR}/../../third_party/tensorflow/libs/android/arm64-v8a/
#cgo android,arm LDFLAGS: -L${SRCDIR}/../../third_party/tensorflow/libs/android/armeabi-v7a/
*/
import "C"
import (
	"unsafe"

	"github.com/iAmPlus/go-tflite/delegates"
)

// FlexDelegate implement Delegater
type FlexDelegate struct {
	d *C.TfLiteDelegate
}

func NewFlex() delegates.Delegater {
	d := C.TfLiteFlexDelegateCreate()
	if d == nil {
		return nil
	}

	return &FlexDelegate{d: d}
}

func (d *FlexDelegate) Delete() {
	C.TfLiteFlexDelegateDelete(d.d)
}

func (d *FlexDelegate) Ptr() unsafe.Pointer {
	return unsafe.Pointer(d.d)
}
