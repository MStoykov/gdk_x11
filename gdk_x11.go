package gdk_x11

// #cgo pkg-config: gdk-3.0 gdk-x11-3.0
// #include <gdk/gdk.h>
// #include <gdk/gdkx.h>
import "C"
import (
	"github.com/conformal/gotk3/gdk"
	"unsafe"
)

func GetGDKX11WindowId(window *gdk.Window) uint {
	id := C.gdk_x11_window_get_xid((*C.GdkWindow)(unsafe.Pointer(window.Native())))
	return uint(id)
}
