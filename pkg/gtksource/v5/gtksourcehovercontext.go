// Code generated by girgen. DO NOT EDIT.

package gtksource

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtksourceview/gtksource.h>
import "C"

// GType values.
var (
	GTypeHoverContext = coreglib.Type(C.gtk_source_hover_context_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeHoverContext, F: marshalHoverContext},
	})
}

// HoverContextOverrides contains methods that are overridable.
type HoverContextOverrides struct {
}

func defaultHoverContextOverrides(v *HoverContext) HoverContextOverrides {
	return HoverContextOverrides{}
}

// HoverContext: context for populating hoverdisplay contents.
//
// GtkSourceHoverContext contains information about the request to populate
// contents for a hoverdisplay.
//
// It can be used to retrieve the view, buffer, and gtk.TextIter for the regions
// of text which are being displayed.
//
// Use hovercontext.GetBounds to get the word that was requested.
// hovercontext.GetIter will get you the location of the pointer when the
// request was made.
type HoverContext struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*HoverContext)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*HoverContext, *HoverContextClass, HoverContextOverrides](
		GTypeHoverContext,
		initHoverContextClass,
		wrapHoverContext,
		defaultHoverContextOverrides,
	)
}

func initHoverContextClass(gclass unsafe.Pointer, overrides HoverContextOverrides, classInitFunc func(*HoverContextClass)) {
	if classInitFunc != nil {
		class := (*HoverContextClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapHoverContext(obj *coreglib.Object) *HoverContext {
	return &HoverContext{
		Object: obj,
	}
}

func marshalHoverContext(p uintptr) (interface{}, error) {
	return wrapHoverContext(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Bounds gets the current word bounds of the hover.
//
// If begin is non-NULL, it will be set to the start position of the current
// word being hovered.
//
// If end is non-NULL, it will be set to the end position for the current word
// being hovered.
//
// The function returns the following values:
//
//   - begin (optional): TextIter.
//   - end (optional): TextIter.
//   - ok: TRUE if the marks are still valid and begin or end was set.
func (self *HoverContext) Bounds() (begin, end *gtk.TextIter, ok bool) {
	var _arg0 *C.GtkSourceHoverContext // out
	var _arg1 C.GtkTextIter            // in
	var _arg2 C.GtkTextIter            // in
	var _cret C.gboolean               // in

	_arg0 = (*C.GtkSourceHoverContext)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_source_hover_context_get_bounds(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(self)

	var _begin *gtk.TextIter // out
	var _end *gtk.TextIter   // out
	var _ok bool             // out

	_begin = (*gtk.TextIter)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))
	_end = (*gtk.TextIter)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))
	if _cret != 0 {
		_ok = true
	}

	return _begin, _end, _ok
}

// Buffer: convenience function to get the buffer.
//
// The function returns the following values:
//
//   - buffer for the view.
func (self *HoverContext) Buffer() *Buffer {
	var _arg0 *C.GtkSourceHoverContext // out
	var _cret *C.GtkSourceBuffer       // in

	_arg0 = (*C.GtkSourceHoverContext)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_source_hover_context_get_buffer(_arg0)
	runtime.KeepAlive(self)

	var _buffer *Buffer // out

	_buffer = wrapBuffer(coreglib.Take(unsafe.Pointer(_cret)))

	return _buffer
}

func (self *HoverContext) Iter(iter *gtk.TextIter) bool {
	var _arg0 *C.GtkSourceHoverContext // out
	var _arg1 *C.GtkTextIter           // out
	var _cret C.gboolean               // in

	_arg0 = (*C.GtkSourceHoverContext)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(iter)))

	_cret = C.gtk_source_hover_context_get_iter(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(iter)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// The function returns the following values:
//
//   - view that owns the context.
func (self *HoverContext) View() *View {
	var _arg0 *C.GtkSourceHoverContext // out
	var _cret *C.GtkSourceView         // in

	_arg0 = (*C.GtkSourceHoverContext)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_source_hover_context_get_view(_arg0)
	runtime.KeepAlive(self)

	var _view *View // out

	_view = wrapView(coreglib.Take(unsafe.Pointer(_cret)))

	return _view
}

// HoverContextClass: instance of this type is always passed by reference.
type HoverContextClass struct {
	*hoverContextClass
}

// hoverContextClass is the struct that's finalized.
type hoverContextClass struct {
	native *C.GtkSourceHoverContextClass
}
