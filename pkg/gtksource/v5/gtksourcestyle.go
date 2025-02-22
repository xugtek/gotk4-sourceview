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
	GTypeStyle = coreglib.Type(C.gtk_source_style_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeStyle, F: marshalStyle},
	})
}

// StyleOverrides contains methods that are overridable.
type StyleOverrides struct {
}

func defaultStyleOverrides(v *Style) StyleOverrides {
	return StyleOverrides{}
}

// Style represents a style.
//
// The GtkSourceStyle structure is used to describe text attributes which are
// set when given style is used.
type Style struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Style)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*Style, *StyleClass, StyleOverrides](
		GTypeStyle,
		initStyleClass,
		wrapStyle,
		defaultStyleOverrides,
	)
}

func initStyleClass(gclass unsafe.Pointer, overrides StyleOverrides, classInitFunc func(*StyleClass)) {
	if classInitFunc != nil {
		class := (*StyleClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapStyle(obj *coreglib.Object) *Style {
	return &Style{
		Object: obj,
	}
}

func marshalStyle(p uintptr) (interface{}, error) {
	return wrapStyle(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Apply: this function modifies the gtk.TextTag properties that are related to
// the SourceStyle properties. Other gtk.TextTag properties are left untouched.
//
// If style is non-NULL, applies style to tag.
//
// If style is NULL, the related *-set properties of gtk.TextTag are set to
// FALSE.
//
// The function takes the following parameters:
//
//   - tag to apply styles to.
func (style *Style) Apply(tag *gtk.TextTag) {
	var _arg0 *C.GtkSourceStyle // out
	var _arg1 *C.GtkTextTag     // out

	if style != nil {
		_arg0 = (*C.GtkSourceStyle)(unsafe.Pointer(coreglib.InternObject(style).Native()))
	}
	_arg1 = (*C.GtkTextTag)(unsafe.Pointer(coreglib.InternObject(tag).Native()))

	C.gtk_source_style_apply(_arg0, _arg1)
	runtime.KeepAlive(style)
	runtime.KeepAlive(tag)
}

// Copy creates a copy of style, that is a new SourceStyle instance which has
// the same attributes set.
//
// The function returns the following values:
//
//   - ret: copy of style, call g_object_unref() when you are done with it.
func (style *Style) Copy() *Style {
	var _arg0 *C.GtkSourceStyle // out
	var _cret *C.GtkSourceStyle // in

	_arg0 = (*C.GtkSourceStyle)(unsafe.Pointer(coreglib.InternObject(style).Native()))

	_cret = C.gtk_source_style_copy(_arg0)
	runtime.KeepAlive(style)

	var _ret *Style // out

	_ret = wrapStyle(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _ret
}

// StyleClass: instance of this type is always passed by reference.
type StyleClass struct {
	*styleClass
}

// styleClass is the struct that's finalized.
type styleClass struct {
	native *C.GtkSourceStyleClass
}
