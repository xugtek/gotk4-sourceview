// Code generated by girgen. DO NOT EDIT.

package gtksource

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtksourceview/gtksource.h>
import "C"

// GType values.
var (
	GTypeGutterRendererPixbuf = coreglib.Type(C.gtk_source_gutter_renderer_pixbuf_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeGutterRendererPixbuf, F: marshalGutterRendererPixbuf},
	})
}

// GutterRendererPixbufOverrides contains methods that are overridable.
type GutterRendererPixbufOverrides struct {
}

func defaultGutterRendererPixbufOverrides(v *GutterRendererPixbuf) GutterRendererPixbufOverrides {
	return GutterRendererPixbufOverrides{}
}

// GutterRendererPixbuf renders a pixbuf in the gutter.
//
// A GtkSourceGutterRendererPixbuf can be used to render an image in a cell of
// gutter.
type GutterRendererPixbuf struct {
	_ [0]func() // equal guard
	GutterRenderer
}

var (
	_ GutterRendererer = (*GutterRendererPixbuf)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*GutterRendererPixbuf, *GutterRendererPixbufClass, GutterRendererPixbufOverrides](
		GTypeGutterRendererPixbuf,
		initGutterRendererPixbufClass,
		wrapGutterRendererPixbuf,
		defaultGutterRendererPixbufOverrides,
	)
}

func initGutterRendererPixbufClass(gclass unsafe.Pointer, overrides GutterRendererPixbufOverrides, classInitFunc func(*GutterRendererPixbufClass)) {
	if classInitFunc != nil {
		class := (*GutterRendererPixbufClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapGutterRendererPixbuf(obj *coreglib.Object) *GutterRendererPixbuf {
	return &GutterRendererPixbuf{
		GutterRenderer: GutterRenderer{
			Widget: gtk.Widget{
				InitiallyUnowned: coreglib.InitiallyUnowned{
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
	return wrapGutterRendererPixbuf(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewGutterRendererPixbuf: create a new SourceGutterRendererPixbuf.
//
// The function returns the following values:
//
//   - gutterRendererPixbuf: SourceGutterRenderer.
func NewGutterRendererPixbuf() *GutterRendererPixbuf {
	var _cret *C.GtkSourceGutterRenderer // in

	_cret = C.gtk_source_gutter_renderer_pixbuf_new()

	var _gutterRendererPixbuf *GutterRendererPixbuf // out

	_gutterRendererPixbuf = wrapGutterRendererPixbuf(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _gutterRendererPixbuf
}

// GIcon: get the gicon of the renderer.
//
// The function returns the following values:
//
//   - icon: #GIcon.
func (renderer *GutterRendererPixbuf) GIcon() *gio.Icon {
	var _arg0 *C.GtkSourceGutterRendererPixbuf // out
	var _cret *C.GIcon                         // in

	_arg0 = (*C.GtkSourceGutterRendererPixbuf)(unsafe.Pointer(coreglib.InternObject(renderer).Native()))

	_cret = C.gtk_source_gutter_renderer_pixbuf_get_gicon(_arg0)
	runtime.KeepAlive(renderer)

	var _icon *gio.Icon // out

	{
		obj := coreglib.Take(unsafe.Pointer(_cret))
		_icon = &gio.Icon{
			Object: obj,
		}
	}

	return _icon
}

func (renderer *GutterRendererPixbuf) IconName() string {
	var _arg0 *C.GtkSourceGutterRendererPixbuf // out
	var _cret *C.gchar                         // in

	_arg0 = (*C.GtkSourceGutterRendererPixbuf)(unsafe.Pointer(coreglib.InternObject(renderer).Native()))

	_cret = C.gtk_source_gutter_renderer_pixbuf_get_icon_name(_arg0)
	runtime.KeepAlive(renderer)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Paintable gets a gdk.Paintable that was set with
// gutterrendererpixbuf.SetPaintable.
//
// The function returns the following values:
//
//   - paintable (optional) or NULL.
func (renderer *GutterRendererPixbuf) Paintable() *gdk.Paintable {
	var _arg0 *C.GtkSourceGutterRendererPixbuf // out
	var _cret *C.GdkPaintable                  // in

	_arg0 = (*C.GtkSourceGutterRendererPixbuf)(unsafe.Pointer(coreglib.InternObject(renderer).Native()))

	_cret = C.gtk_source_gutter_renderer_pixbuf_get_paintable(_arg0)
	runtime.KeepAlive(renderer)

	var _paintable *gdk.Paintable // out

	if _cret != nil {
		{
			obj := coreglib.Take(unsafe.Pointer(_cret))
			_paintable = &gdk.Paintable{
				Object: obj,
			}
		}
	}

	return _paintable
}

// Pixbuf: get the pixbuf of the renderer.
//
// The function returns the following values:
//
//   - pixbuf: Pixbuf.
func (renderer *GutterRendererPixbuf) Pixbuf() *gdkpixbuf.Pixbuf {
	var _arg0 *C.GtkSourceGutterRendererPixbuf // out
	var _cret *C.GdkPixbuf                     // in

	_arg0 = (*C.GtkSourceGutterRendererPixbuf)(unsafe.Pointer(coreglib.InternObject(renderer).Native()))

	_cret = C.gtk_source_gutter_renderer_pixbuf_get_pixbuf(_arg0)
	runtime.KeepAlive(renderer)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	{
		obj := coreglib.Take(unsafe.Pointer(_cret))
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
//   - paintable: Paintable.
func (renderer *GutterRendererPixbuf) OverlayPaintable(paintable gdk.Paintabler) {
	var _arg0 *C.GtkSourceGutterRendererPixbuf // out
	var _arg1 *C.GdkPaintable                  // out

	_arg0 = (*C.GtkSourceGutterRendererPixbuf)(unsafe.Pointer(coreglib.InternObject(renderer).Native()))
	_arg1 = (*C.GdkPaintable)(unsafe.Pointer(coreglib.InternObject(paintable).Native()))

	C.gtk_source_gutter_renderer_pixbuf_overlay_paintable(_arg0, _arg1)
	runtime.KeepAlive(renderer)
	runtime.KeepAlive(paintable)
}

// The function takes the following parameters:
//
//   - icon (optional): icon, or NULL.
func (renderer *GutterRendererPixbuf) SetGIcon(icon gio.Iconner) {
	var _arg0 *C.GtkSourceGutterRendererPixbuf // out
	var _arg1 *C.GIcon                         // out

	_arg0 = (*C.GtkSourceGutterRendererPixbuf)(unsafe.Pointer(coreglib.InternObject(renderer).Native()))
	if icon != nil {
		_arg1 = (*C.GIcon)(unsafe.Pointer(coreglib.InternObject(icon).Native()))
	}

	C.gtk_source_gutter_renderer_pixbuf_set_gicon(_arg0, _arg1)
	runtime.KeepAlive(renderer)
	runtime.KeepAlive(icon)
}

// The function takes the following parameters:
//
//   - iconName (optional): icon name, or NULL.
func (renderer *GutterRendererPixbuf) SetIconName(iconName string) {
	var _arg0 *C.GtkSourceGutterRendererPixbuf // out
	var _arg1 *C.gchar                         // out

	_arg0 = (*C.GtkSourceGutterRendererPixbuf)(unsafe.Pointer(coreglib.InternObject(renderer).Native()))
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
//   - paintable (optional): paintable, or NULL.
func (renderer *GutterRendererPixbuf) SetPaintable(paintable gdk.Paintabler) {
	var _arg0 *C.GtkSourceGutterRendererPixbuf // out
	var _arg1 *C.GdkPaintable                  // out

	_arg0 = (*C.GtkSourceGutterRendererPixbuf)(unsafe.Pointer(coreglib.InternObject(renderer).Native()))
	if paintable != nil {
		_arg1 = (*C.GdkPaintable)(unsafe.Pointer(coreglib.InternObject(paintable).Native()))
	}

	C.gtk_source_gutter_renderer_pixbuf_set_paintable(_arg0, _arg1)
	runtime.KeepAlive(renderer)
	runtime.KeepAlive(paintable)
}

// The function takes the following parameters:
//
//   - pixbuf (optional): pixbuf, or NULL.
func (renderer *GutterRendererPixbuf) SetPixbuf(pixbuf *gdkpixbuf.Pixbuf) {
	var _arg0 *C.GtkSourceGutterRendererPixbuf // out
	var _arg1 *C.GdkPixbuf                     // out

	_arg0 = (*C.GtkSourceGutterRendererPixbuf)(unsafe.Pointer(coreglib.InternObject(renderer).Native()))
	if pixbuf != nil {
		_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(coreglib.InternObject(pixbuf).Native()))
	}

	C.gtk_source_gutter_renderer_pixbuf_set_pixbuf(_arg0, _arg1)
	runtime.KeepAlive(renderer)
	runtime.KeepAlive(pixbuf)
}

// GutterRendererPixbufClass: instance of this type is always passed by
// reference.
type GutterRendererPixbufClass struct {
	*gutterRendererPixbufClass
}

// gutterRendererPixbufClass is the struct that's finalized.
type gutterRendererPixbufClass struct {
	native *C.GtkSourceGutterRendererPixbufClass
}

func (g *GutterRendererPixbufClass) ParentClass() *GutterRendererClass {
	valptr := &g.native.parent_class
	var _v *GutterRendererClass // out
	_v = (*GutterRendererClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
