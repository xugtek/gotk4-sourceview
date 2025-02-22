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
	GTypeHoverDisplay = coreglib.Type(C.gtk_source_hover_display_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeHoverDisplay, F: marshalHoverDisplay},
	})
}

// HoverDisplayOverrides contains methods that are overridable.
type HoverDisplayOverrides struct {
}

func defaultHoverDisplayOverrides(v *HoverDisplay) HoverDisplayOverrides {
	return HoverDisplayOverrides{}
}

// HoverDisplay: display for interactive tooltips.
//
// GtkSourceHoverDisplay is a gtk.Widget that may be populated with widgets to
// be displayed to the user in interactive tooltips. The children widgets are
// packed vertically using a gtk.Box.
//
// Implement the hoverprovider interface to be notified of when to populate a
// GtkSourceHoverDisplay on behalf of the user.
type HoverDisplay struct {
	_ [0]func() // equal guard
	gtk.Widget
}

var (
	_ gtk.Widgetter = (*HoverDisplay)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*HoverDisplay, *HoverDisplayClass, HoverDisplayOverrides](
		GTypeHoverDisplay,
		initHoverDisplayClass,
		wrapHoverDisplay,
		defaultHoverDisplayOverrides,
	)
}

func initHoverDisplayClass(gclass unsafe.Pointer, overrides HoverDisplayOverrides, classInitFunc func(*HoverDisplayClass)) {
	if classInitFunc != nil {
		class := (*HoverDisplayClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapHoverDisplay(obj *coreglib.Object) *HoverDisplay {
	return &HoverDisplay{
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
	}
}

func marshalHoverDisplay(p uintptr) (interface{}, error) {
	return wrapHoverDisplay(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (self *HoverDisplay) Append(child gtk.Widgetter) {
	var _arg0 *C.GtkSourceHoverDisplay // out
	var _arg1 *C.GtkWidget             // out

	_arg0 = (*C.GtkSourceHoverDisplay)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	C.gtk_source_hover_display_append(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}

// The function takes the following parameters:
//
//   - child
//   - sibling
func (self *HoverDisplay) InsertAfter(child, sibling gtk.Widgetter) {
	var _arg0 *C.GtkSourceHoverDisplay // out
	var _arg1 *C.GtkWidget             // out
	var _arg2 *C.GtkWidget             // out

	_arg0 = (*C.GtkSourceHoverDisplay)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(sibling).Native()))

	C.gtk_source_hover_display_insert_after(_arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
	runtime.KeepAlive(sibling)
}

func (self *HoverDisplay) Prepend(child gtk.Widgetter) {
	var _arg0 *C.GtkSourceHoverDisplay // out
	var _arg1 *C.GtkWidget             // out

	_arg0 = (*C.GtkSourceHoverDisplay)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	C.gtk_source_hover_display_prepend(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}

func (self *HoverDisplay) Remove(child gtk.Widgetter) {
	var _arg0 *C.GtkSourceHoverDisplay // out
	var _arg1 *C.GtkWidget             // out

	_arg0 = (*C.GtkSourceHoverDisplay)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	C.gtk_source_hover_display_remove(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}

// HoverDisplayClass: instance of this type is always passed by reference.
type HoverDisplayClass struct {
	*hoverDisplayClass
}

// hoverDisplayClass is the struct that's finalized.
type hoverDisplayClass struct {
	native *C.GtkSourceHoverDisplayClass
}

func (h *HoverDisplayClass) ParentClass() *gtk.WidgetClass {
	valptr := &h.native.parent_class
	var _v *gtk.WidgetClass // out
	_v = (*gtk.WidgetClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
