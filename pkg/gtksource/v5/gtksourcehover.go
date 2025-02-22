// Code generated by girgen. DO NOT EDIT.

package gtksource

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtksourceview/gtksource.h>
import "C"

// GType values.
var (
	GTypeHover = coreglib.Type(C.gtk_source_hover_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeHover, F: marshalHover},
	})
}

// HoverOverrides contains methods that are overridable.
type HoverOverrides struct {
}

func defaultHoverOverrides(v *Hover) HoverOverrides {
	return HoverOverrides{}
}

// Hover: interactive tooltips.
//
// GtkSourceHover allows a view to provide contextual information. When enabled,
// if the user hovers over a word in the text editor, a series of registered
// hoverprovider can populate a hoverdisplay with useful information.
//
// To enable call view.GetHover and add hoverprovider using hover.AddProvider.
// To disable, remove all registered providers with hover.RemoveProvider.
//
// You can change how long to wait to display the interactive tooltip by setting
// the hover:hover-delay property in milliseconds.
type Hover struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Hover)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*Hover, *HoverClass, HoverOverrides](
		GTypeHover,
		initHoverClass,
		wrapHover,
		defaultHoverOverrides,
	)
}

func initHoverClass(gclass unsafe.Pointer, overrides HoverOverrides, classInitFunc func(*HoverClass)) {
	if classInitFunc != nil {
		class := (*HoverClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapHover(obj *coreglib.Object) *Hover {
	return &Hover{
		Object: obj,
	}
}

func marshalHover(p uintptr) (interface{}, error) {
	return wrapHover(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (self *Hover) AddProvider(provider HoverProviderer) {
	var _arg0 *C.GtkSourceHover         // out
	var _arg1 *C.GtkSourceHoverProvider // out

	_arg0 = (*C.GtkSourceHover)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GtkSourceHoverProvider)(unsafe.Pointer(coreglib.InternObject(provider).Native()))

	C.gtk_source_hover_add_provider(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(provider)
}

func (self *Hover) RemoveProvider(provider HoverProviderer) {
	var _arg0 *C.GtkSourceHover         // out
	var _arg1 *C.GtkSourceHoverProvider // out

	_arg0 = (*C.GtkSourceHover)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GtkSourceHoverProvider)(unsafe.Pointer(coreglib.InternObject(provider).Native()))

	C.gtk_source_hover_remove_provider(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(provider)
}

// HoverClass: instance of this type is always passed by reference.
type HoverClass struct {
	*hoverClass
}

// hoverClass is the struct that's finalized.
type hoverClass struct {
	native *C.GtkSourceHoverClass
}
