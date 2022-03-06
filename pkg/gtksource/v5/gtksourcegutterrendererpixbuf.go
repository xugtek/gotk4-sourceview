// Code generated by girgen. DO NOT EDIT.

package gtksource

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtksourceview/gtksource.h>
import "C"

// glib.Type values for gtksourcegutterrendererpixbuf.go.
var GTypeGutterRendererPixbuf = externglib.Type(C.gtk_source_gutter_renderer_pixbuf_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeGutterRendererPixbuf, F: marshalGutterRendererPixbuf},
	})
}

// GutterRendererPixbufOverrider contains methods that are overridable.
type GutterRendererPixbufOverrider interface {
}

type GutterRendererPixbuf struct {
	_ [0]func() // equal guard
	GutterRenderer
}

var (
	_ GutterRendererer = (*GutterRendererPixbuf)(nil)
)

func classInitGutterRendererPixbuffer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapGutterRendererPixbuf(obj *externglib.Object) *GutterRendererPixbuf {
	return &GutterRendererPixbuf{
		GutterRenderer: GutterRenderer{
			Widget: gtk.Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				Object: obj,
				Accessible: gtk.Accessible{
					Object: obj,
				},
				Buildable: gtk.Buildable{
					Object: obj,
				},
				ConstraintTarget: gtk.ConstraintTarget{
					Object: obj,
				},
			},
		},
	}
}

func marshalGutterRendererPixbuf(p uintptr) (interface{}, error) {
	return wrapGutterRendererPixbuf(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewGutterRendererPixbuf: create a new SourceGutterRendererPixbuf.
//
// The function returns the following values:
//
//    - gutterRendererPixbuf: SourceGutterRenderer.
//
func NewGutterRendererPixbuf() *GutterRendererPixbuf {
	var _cret *C.GtkSourceGutterRenderer // in

	_cret = C.gtk_source_gutter_renderer_pixbuf_new()

	var _gutterRendererPixbuf *GutterRendererPixbuf // out

	_gutterRendererPixbuf = wrapGutterRendererPixbuf(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _gutterRendererPixbuf
}

// GIcon: get the gicon of the renderer.
//
// The function returns the following values:
//
//    - icon: #GIcon.
//
func (renderer *GutterRendererPixbuf) GIcon() gio.Iconner {
	var _arg0 *C.GtkSourceGutterRendererPixbuf // out
	var _cret *C.GIcon                         // in

	_arg0 = (*C.GtkSourceGutterRendererPixbuf)(unsafe.Pointer(externglib.InternObject(renderer).Native()))

	_cret = C.gtk_source_gutter_renderer_pixbuf_get_gicon(_arg0)
	runtime.KeepAlive(renderer)

	var _icon gio.Iconner // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.Iconner is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(gio.Iconner)
			return ok
		})
		rv, ok := casted.(gio.Iconner)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.Iconner")
		}
		_icon = rv
	}

	return _icon
}

// The function returns the following values:
//
func (renderer *GutterRendererPixbuf) IconName() string {
	var _arg0 *C.GtkSourceGutterRendererPixbuf // out
	var _cret *C.gchar                         // in

	_arg0 = (*C.GtkSourceGutterRendererPixbuf)(unsafe.Pointer(externglib.InternObject(renderer).Native()))

	_cret = C.gtk_source_gutter_renderer_pixbuf_get_icon_name(_arg0)
	runtime.KeepAlive(renderer)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Paintable gets a Paintable that was set with
// gtk_source_gutter_renderer_pixbuf_set_paintable().
//
// The function returns the following values:
//
//    - paintable (optional) or NULL.
//
func (renderer *GutterRendererPixbuf) Paintable() gdk.Paintabler {
	var _arg0 *C.GtkSourceGutterRendererPixbuf // out
	var _cret *C.GdkPaintable                  // in

	_arg0 = (*C.GtkSourceGutterRendererPixbuf)(unsafe.Pointer(externglib.InternObject(renderer).Native()))

	_cret = C.gtk_source_gutter_renderer_pixbuf_get_paintable(_arg0)
	runtime.KeepAlive(renderer)

	var _paintable gdk.Paintabler // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(gdk.Paintabler)
				return ok
			})
			rv, ok := casted.(gdk.Paintabler)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Paintabler")
			}
			_paintable = rv
		}
	}

	return _paintable
}

// Pixbuf: get the pixbuf of the renderer.
//
// The function returns the following values:
//
//    - pixbuf: Pixbuf.
//
func (renderer *GutterRendererPixbuf) Pixbuf() *gdkpixbuf.Pixbuf {
	var _arg0 *C.GtkSourceGutterRendererPixbuf // out
	var _cret *C.GdkPixbuf                     // in

	_arg0 = (*C.GtkSourceGutterRendererPixbuf)(unsafe.Pointer(externglib.InternObject(renderer).Native()))

	_cret = C.gtk_source_gutter_renderer_pixbuf_get_pixbuf(_arg0)
	runtime.KeepAlive(renderer)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	{
		obj := externglib.Take(unsafe.Pointer(_cret))
		_pixbuf = &gdkpixbuf.Pixbuf{
			Object: obj,
			LoadableIcon: gio.LoadableIcon{
				Icon: gio.Icon{
					Object: obj,
				},
			},
		}
	}

	return _pixbuf
}

// OverlayPaintable allows overlaying a paintable on top of any other image that
// has been set for the pixbuf. This will be applied when the widget is next
// snapshot.
//
// The function takes the following parameters:
//
//    - paintable: Paintable.
//
func (renderer *GutterRendererPixbuf) OverlayPaintable(paintable gdk.Paintabler) {
	var _arg0 *C.GtkSourceGutterRendererPixbuf // out
	var _arg1 *C.GdkPaintable                  // out

	_arg0 = (*C.GtkSourceGutterRendererPixbuf)(unsafe.Pointer(externglib.InternObject(renderer).Native()))
	_arg1 = (*C.GdkPaintable)(unsafe.Pointer(externglib.InternObject(paintable).Native()))

	C.gtk_source_gutter_renderer_pixbuf_overlay_paintable(_arg0, _arg1)
	runtime.KeepAlive(renderer)
	runtime.KeepAlive(paintable)
}

// The function takes the following parameters:
//
//    - icon (optional): icon, or NULL.
//
func (renderer *GutterRendererPixbuf) SetGIcon(icon gio.Iconner) {
	var _arg0 *C.GtkSourceGutterRendererPixbuf // out
	var _arg1 *C.GIcon                         // out

	_arg0 = (*C.GtkSourceGutterRendererPixbuf)(unsafe.Pointer(externglib.InternObject(renderer).Native()))
	if icon != nil {
		_arg1 = (*C.GIcon)(unsafe.Pointer(externglib.InternObject(icon).Native()))
	}

	C.gtk_source_gutter_renderer_pixbuf_set_gicon(_arg0, _arg1)
	runtime.KeepAlive(renderer)
	runtime.KeepAlive(icon)
}

// The function takes the following parameters:
//
//    - iconName (optional): icon name, or NULL.
//
func (renderer *GutterRendererPixbuf) SetIconName(iconName string) {
	var _arg0 *C.GtkSourceGutterRendererPixbuf // out
	var _arg1 *C.gchar                         // out

	_arg0 = (*C.GtkSourceGutterRendererPixbuf)(unsafe.Pointer(externglib.InternObject(renderer).Native()))
	if iconName != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(iconName)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_source_gutter_renderer_pixbuf_set_icon_name(_arg0, _arg1)
	runtime.KeepAlive(renderer)
	runtime.KeepAlive(iconName)
}

// The function takes the following parameters:
//
//    - paintable (optional): paintable, or NULL.
//
func (renderer *GutterRendererPixbuf) SetPaintable(paintable gdk.Paintabler) {
	var _arg0 *C.GtkSourceGutterRendererPixbuf // out
	var _arg1 *C.GdkPaintable                  // out

	_arg0 = (*C.GtkSourceGutterRendererPixbuf)(unsafe.Pointer(externglib.InternObject(renderer).Native()))
	if paintable != nil {
		_arg1 = (*C.GdkPaintable)(unsafe.Pointer(externglib.InternObject(paintable).Native()))
	}

	C.gtk_source_gutter_renderer_pixbuf_set_paintable(_arg0, _arg1)
	runtime.KeepAlive(renderer)
	runtime.KeepAlive(paintable)
}

// The function takes the following parameters:
//
//    - pixbuf (optional): pixbuf, or NULL.
//
func (renderer *GutterRendererPixbuf) SetPixbuf(pixbuf *gdkpixbuf.Pixbuf) {
	var _arg0 *C.GtkSourceGutterRendererPixbuf // out
	var _arg1 *C.GdkPixbuf                     // out

	_arg0 = (*C.GtkSourceGutterRendererPixbuf)(unsafe.Pointer(externglib.InternObject(renderer).Native()))
	if pixbuf != nil {
		_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(externglib.InternObject(pixbuf).Native()))
	}

	C.gtk_source_gutter_renderer_pixbuf_set_pixbuf(_arg0, _arg1)
	runtime.KeepAlive(renderer)
	runtime.KeepAlive(pixbuf)
}
