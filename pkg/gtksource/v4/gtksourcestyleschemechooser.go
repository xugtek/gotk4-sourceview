// Code generated by girgen. DO NOT EDIT.

package gtksource

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtksourceview/gtksource.h>
// extern GtkSourceStyleScheme* _gotk4_gtksource4_StyleSchemeChooserInterface_get_style_scheme(GtkSourceStyleSchemeChooser*);
// extern void _gotk4_gtksource4_StyleSchemeChooserInterface_set_style_scheme(GtkSourceStyleSchemeChooser*, GtkSourceStyleScheme*);
import "C"

// glib.Type values for gtksourcestyleschemechooser.go.
var GTypeStyleSchemeChooser = externglib.Type(C.gtk_source_style_scheme_chooser_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeStyleSchemeChooser, F: marshalStyleSchemeChooser},
	})
}

// StyleSchemeChooserOverrider contains methods that are overridable.
type StyleSchemeChooserOverrider interface {
	// StyleScheme gets the currently-selected scheme.
	//
	// The function returns the following values:
	//
	//    - styleScheme: currently-selected scheme.
	//
	StyleScheme() *StyleScheme
	// SetStyleScheme sets the scheme.
	//
	// The function takes the following parameters:
	//
	//    - scheme: SourceStyleScheme.
	//
	SetStyleScheme(scheme *StyleScheme)
}

type StyleSchemeChooser struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*StyleSchemeChooser)(nil)
)

// StyleSchemeChooserer describes StyleSchemeChooser's interface methods.
type StyleSchemeChooserer interface {
	externglib.Objector

	// StyleScheme gets the currently-selected scheme.
	StyleScheme() *StyleScheme
	// SetStyleScheme sets the scheme.
	SetStyleScheme(scheme *StyleScheme)
}

var _ StyleSchemeChooserer = (*StyleSchemeChooser)(nil)

func ifaceInitStyleSchemeChooserer(gifacePtr, data C.gpointer) {
	iface := (*C.GtkSourceStyleSchemeChooserInterface)(unsafe.Pointer(gifacePtr))
	iface.get_style_scheme = (*[0]byte)(C._gotk4_gtksource4_StyleSchemeChooserInterface_get_style_scheme)
	iface.set_style_scheme = (*[0]byte)(C._gotk4_gtksource4_StyleSchemeChooserInterface_set_style_scheme)
}

//export _gotk4_gtksource4_StyleSchemeChooserInterface_get_style_scheme
func _gotk4_gtksource4_StyleSchemeChooserInterface_get_style_scheme(arg0 *C.GtkSourceStyleSchemeChooser) (cret *C.GtkSourceStyleScheme) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(StyleSchemeChooserOverrider)

	styleScheme := iface.StyleScheme()

	cret = (*C.GtkSourceStyleScheme)(unsafe.Pointer(externglib.InternObject(styleScheme).Native()))

	return cret
}

//export _gotk4_gtksource4_StyleSchemeChooserInterface_set_style_scheme
func _gotk4_gtksource4_StyleSchemeChooserInterface_set_style_scheme(arg0 *C.GtkSourceStyleSchemeChooser, arg1 *C.GtkSourceStyleScheme) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(StyleSchemeChooserOverrider)

	var _scheme *StyleScheme // out

	_scheme = wrapStyleScheme(externglib.Take(unsafe.Pointer(arg1)))

	iface.SetStyleScheme(_scheme)
}

func wrapStyleSchemeChooser(obj *externglib.Object) *StyleSchemeChooser {
	return &StyleSchemeChooser{
		Object: obj,
	}
}

func marshalStyleSchemeChooser(p uintptr) (interface{}, error) {
	return wrapStyleSchemeChooser(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// StyleScheme gets the currently-selected scheme.
//
// The function returns the following values:
//
//    - styleScheme: currently-selected scheme.
//
func (chooser *StyleSchemeChooser) StyleScheme() *StyleScheme {
	var _arg0 *C.GtkSourceStyleSchemeChooser // out
	var _cret *C.GtkSourceStyleScheme        // in

	_arg0 = (*C.GtkSourceStyleSchemeChooser)(unsafe.Pointer(externglib.InternObject(chooser).Native()))

	_cret = C.gtk_source_style_scheme_chooser_get_style_scheme(_arg0)
	runtime.KeepAlive(chooser)

	var _styleScheme *StyleScheme // out

	_styleScheme = wrapStyleScheme(externglib.Take(unsafe.Pointer(_cret)))

	return _styleScheme
}

// SetStyleScheme sets the scheme.
//
// The function takes the following parameters:
//
//    - scheme: SourceStyleScheme.
//
func (chooser *StyleSchemeChooser) SetStyleScheme(scheme *StyleScheme) {
	var _arg0 *C.GtkSourceStyleSchemeChooser // out
	var _arg1 *C.GtkSourceStyleScheme        // out

	_arg0 = (*C.GtkSourceStyleSchemeChooser)(unsafe.Pointer(externglib.InternObject(chooser).Native()))
	_arg1 = (*C.GtkSourceStyleScheme)(unsafe.Pointer(externglib.InternObject(scheme).Native()))

	C.gtk_source_style_scheme_chooser_set_style_scheme(_arg0, _arg1)
	runtime.KeepAlive(chooser)
	runtime.KeepAlive(scheme)
}
