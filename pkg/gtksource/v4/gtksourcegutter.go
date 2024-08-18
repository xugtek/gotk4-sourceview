// Code generated by girgen. DO NOT EDIT.

package gtksource

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v3"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtksourceview/gtksource.h>
import "C"

// GType values.
var (
	GTypeGutter = coreglib.Type(C.gtk_source_gutter_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeGutter, F: marshalGutter},
	})
}

// GutterOverrides contains methods that are overridable.
type GutterOverrides struct {
}

func defaultGutterOverrides(v *Gutter) GutterOverrides {
	return GutterOverrides{}
}

type Gutter struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Gutter)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*Gutter, *GutterClass, GutterOverrides](
		GTypeGutter,
		initGutterClass,
		wrapGutter,
		defaultGutterOverrides,
	)
}

func initGutterClass(gclass unsafe.Pointer, overrides GutterOverrides, classInitFunc func(*GutterClass)) {
	if classInitFunc != nil {
		class := (*GutterClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapGutter(obj *coreglib.Object) *Gutter {
	return &Gutter{
		Object: obj,
	}
}

func marshalGutter(p uintptr) (interface{}, error) {
	return wrapGutter(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// RendererAtPos finds the SourceGutterRenderer at (x, y).
//
// The function takes the following parameters:
//
//   - x position to get identified.
//   - y position to get identified.
//
// The function returns the following values:
//
//   - gutterRenderer (optional): renderer at (x, y) or NULL.
func (gutter *Gutter) RendererAtPos(x, y int) GutterRendererer {
	var _arg0 *C.GtkSourceGutter         // out
	var _arg1 C.gint                     // out
	var _arg2 C.gint                     // out
	var _cret *C.GtkSourceGutterRenderer // in

	_arg0 = (*C.GtkSourceGutter)(unsafe.Pointer(coreglib.InternObject(gutter).Native()))
	_arg1 = C.gint(x)
	_arg2 = C.gint(y)

	_cret = C.gtk_source_gutter_get_renderer_at_pos(_arg0, _arg1, _arg2)
	runtime.KeepAlive(gutter)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)

	var _gutterRenderer GutterRendererer // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(GutterRendererer)
				return ok
			})
			rv, ok := casted.(GutterRendererer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtksource.GutterRendererer")
			}
			_gutterRenderer = rv
		}
	}

	return _gutterRenderer
}

// The function returns the following values:
//
//   - view: associated SourceView.
func (gutter *Gutter) View() *View {
	var _arg0 *C.GtkSourceGutter // out
	var _cret *C.GtkSourceView   // in

	_arg0 = (*C.GtkSourceGutter)(unsafe.Pointer(coreglib.InternObject(gutter).Native()))

	_cret = C.gtk_source_gutter_get_view(_arg0)
	runtime.KeepAlive(gutter)

	var _view *View // out

	_view = wrapView(coreglib.Take(unsafe.Pointer(_cret)))

	return _view
}

// The function returns the following values:
//
//   - textWindowType of gutter.
func (gutter *Gutter) WindowType() gtk.TextWindowType {
	var _arg0 *C.GtkSourceGutter  // out
	var _cret C.GtkTextWindowType // in

	_arg0 = (*C.GtkSourceGutter)(unsafe.Pointer(coreglib.InternObject(gutter).Native()))

	_cret = C.gtk_source_gutter_get_window_type(_arg0)
	runtime.KeepAlive(gutter)

	var _textWindowType gtk.TextWindowType // out

	_textWindowType = gtk.TextWindowType(_cret)

	return _textWindowType
}

// Insert renderer into the gutter. If renderer is yet unowned then gutter
// claims its ownership. Otherwise just increases renderer's reference count.
// renderer cannot be already inserted to another gutter.
//
// The function takes the following parameters:
//
//   - renderer: gutter renderer (must inherit from SourceGutterRenderer).
//   - position: renderer position.
//
// The function returns the following values:
//
//   - ok: TRUE if operation succeeded. Otherwise FALSE.
func (gutter *Gutter) Insert(renderer GutterRendererer, position int) bool {
	var _arg0 *C.GtkSourceGutter         // out
	var _arg1 *C.GtkSourceGutterRenderer // out
	var _arg2 C.gint                     // out
	var _cret C.gboolean                 // in

	_arg0 = (*C.GtkSourceGutter)(unsafe.Pointer(coreglib.InternObject(gutter).Native()))
	_arg1 = (*C.GtkSourceGutterRenderer)(unsafe.Pointer(coreglib.InternObject(renderer).Native()))
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

// QueueDraw invalidates the drawable area of the gutter. You can use this
// to force a redraw of the gutter if something has changed and needs to be
// redrawn.
func (gutter *Gutter) QueueDraw() {
	var _arg0 *C.GtkSourceGutter // out

	_arg0 = (*C.GtkSourceGutter)(unsafe.Pointer(coreglib.InternObject(gutter).Native()))

	C.gtk_source_gutter_queue_draw(_arg0)
	runtime.KeepAlive(gutter)
}

// Remove removes renderer from gutter.
//
// The function takes the following parameters:
//
//   - renderer: SourceGutterRenderer.
func (gutter *Gutter) Remove(renderer GutterRendererer) {
	var _arg0 *C.GtkSourceGutter         // out
	var _arg1 *C.GtkSourceGutterRenderer // out

	_arg0 = (*C.GtkSourceGutter)(unsafe.Pointer(coreglib.InternObject(gutter).Native()))
	_arg1 = (*C.GtkSourceGutterRenderer)(unsafe.Pointer(coreglib.InternObject(renderer).Native()))

	C.gtk_source_gutter_remove(_arg0, _arg1)
	runtime.KeepAlive(gutter)
	runtime.KeepAlive(renderer)
}

// Reorder reorders renderer in gutter to new position.
//
// The function takes the following parameters:
//
//   - renderer: CellRenderer.
//   - position: new renderer position.
func (gutter *Gutter) Reorder(renderer GutterRendererer, position int) {
	var _arg0 *C.GtkSourceGutter         // out
	var _arg1 *C.GtkSourceGutterRenderer // out
	var _arg2 C.gint                     // out

	_arg0 = (*C.GtkSourceGutter)(unsafe.Pointer(coreglib.InternObject(gutter).Native()))
	_arg1 = (*C.GtkSourceGutterRenderer)(unsafe.Pointer(coreglib.InternObject(renderer).Native()))
	_arg2 = C.gint(position)

	C.gtk_source_gutter_reorder(_arg0, _arg1, _arg2)
	runtime.KeepAlive(gutter)
	runtime.KeepAlive(renderer)
	runtime.KeepAlive(position)
}

// GutterClass: instance of this type is always passed by reference.
type GutterClass struct {
	*gutterClass
}

// gutterClass is the struct that's finalized.
type gutterClass struct {
	native *C.GtkSourceGutterClass
}

func (g *GutterClass) Padding() [10]unsafe.Pointer {
	valptr := &g.native.padding
	var _v [10]unsafe.Pointer // out
	{
		src := &*valptr
		for i := 0; i < 10; i++ {
			_v[i] = (unsafe.Pointer)(unsafe.Pointer(src[i]))
		}
	}
	return _v
}
