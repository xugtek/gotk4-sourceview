// Code generated by girgen. DO NOT EDIT.

package gtksource

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtksourceview/gtksource.h>
import "C"

// glib.Type values for gtksourcegutter.go.
var GTypeGutter = externglib.Type(C.gtk_source_gutter_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeGutter, F: marshalGutter},
	})
}

// GutterOverrider contains methods that are overridable.
type GutterOverrider interface {
}

type Gutter struct {
	_ [0]func() // equal guard
	gtk.Widget
}

var (
	_ gtk.Widgetter = (*Gutter)(nil)
)

func classInitGutterer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapGutter(obj *externglib.Object) *Gutter {
	return &Gutter{
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
	}
}

func marshalGutter(p uintptr) (interface{}, error) {
	return wrapGutter(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// The function returns the following values:
//
//    - view: associated SourceView.
//
func (gutter *Gutter) View() *View {
	var _arg0 *C.GtkSourceGutter // out
	var _cret *C.GtkSourceView   // in

	_arg0 = (*C.GtkSourceGutter)(unsafe.Pointer(externglib.InternObject(gutter).Native()))

	_cret = C.gtk_source_gutter_get_view(_arg0)
	runtime.KeepAlive(gutter)

	var _view *View // out

	_view = wrapView(externglib.Take(unsafe.Pointer(_cret)))

	return _view
}

// Insert renderer into the gutter. If renderer is yet unowned then gutter
// claims its ownership. Otherwise just increases renderer's reference count.
// renderer cannot be already inserted to another gutter.
//
// The function takes the following parameters:
//
//    - renderer: gutter renderer (must inherit from SourceGutterRenderer).
//    - position: renderer position.
//
// The function returns the following values:
//
//    - ok: TRUE if operation succeeded. Otherwise FALSE.
//
func (gutter *Gutter) Insert(renderer GutterRendererer, position int) bool {
	var _arg0 *C.GtkSourceGutter         // out
	var _arg1 *C.GtkSourceGutterRenderer // out
	var _arg2 C.gint                     // out
	var _cret C.gboolean                 // in

	_arg0 = (*C.GtkSourceGutter)(unsafe.Pointer(externglib.InternObject(gutter).Native()))
	_arg1 = (*C.GtkSourceGutterRenderer)(unsafe.Pointer(externglib.InternObject(renderer).Native()))
	_arg2 = C.gint(position)

	_cret = C.gtk_source_gutter_insert(_arg0, _arg1, _arg2)
	runtime.KeepAlive(gutter)
	runtime.KeepAlive(renderer)
	runtime.KeepAlive(position)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// The function takes the following parameters:
//
func (gutter *Gutter) Remove(renderer GutterRendererer) {
	var _arg0 *C.GtkSourceGutter         // out
	var _arg1 *C.GtkSourceGutterRenderer // out

	_arg0 = (*C.GtkSourceGutter)(unsafe.Pointer(externglib.InternObject(gutter).Native()))
	_arg1 = (*C.GtkSourceGutterRenderer)(unsafe.Pointer(externglib.InternObject(renderer).Native()))

	C.gtk_source_gutter_remove(_arg0, _arg1)
	runtime.KeepAlive(gutter)
	runtime.KeepAlive(renderer)
}

// Reorder reorders renderer in gutter to new position.
//
// The function takes the following parameters:
//
//    - renderer: CellRenderer.
//    - position: new renderer position.
//
func (gutter *Gutter) Reorder(renderer GutterRendererer, position int) {
	var _arg0 *C.GtkSourceGutter         // out
	var _arg1 *C.GtkSourceGutterRenderer // out
	var _arg2 C.gint                     // out

	_arg0 = (*C.GtkSourceGutter)(unsafe.Pointer(externglib.InternObject(gutter).Native()))
	_arg1 = (*C.GtkSourceGutterRenderer)(unsafe.Pointer(externglib.InternObject(renderer).Native()))
	_arg2 = C.gint(position)

	C.gtk_source_gutter_reorder(_arg0, _arg1, _arg2)
	runtime.KeepAlive(gutter)
	runtime.KeepAlive(renderer)
	runtime.KeepAlive(position)
}
