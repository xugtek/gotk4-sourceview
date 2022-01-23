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

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_source_gutter_renderer_text_get_type()), F: marshalGutterRendererTexter},
	})
}

type GutterRendererText struct {
	_ [0]func() // equal guard
	GutterRenderer
}

var (
	_ GutterRendererer = (*GutterRendererText)(nil)
)

func wrapGutterRendererText(obj *externglib.Object) *GutterRendererText {
	return &GutterRendererText{
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

func marshalGutterRendererTexter(p uintptr) (interface{}, error) {
	return wrapGutterRendererText(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewGutterRendererText: create a new SourceGutterRendererText.
//
// The function returns the following values:
//
//    - gutterRendererText: SourceGutterRenderer.
//
func NewGutterRendererText() *GutterRendererText {
	var _cret *C.GtkSourceGutterRenderer // in

	_cret = C.gtk_source_gutter_renderer_text_new()

	var _gutterRendererText *GutterRendererText // out

	_gutterRendererText = wrapGutterRendererText(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _gutterRendererText
}

// Measure measures the text provided using the pango layout used by the
// SourceGutterRendererText.
//
// The function takes the following parameters:
//
//    - text to measure.
//
// The function returns the following values:
//
//    - width (optional): location to store the width of the text in pixels, or
//      NULL.
//    - height (optional): location to store the height of the text in pixels, or
//      NULL.
//
func (renderer *GutterRendererText) Measure(text string) (width int, height int) {
	var _arg0 *C.GtkSourceGutterRendererText // out
	var _arg1 *C.gchar                       // out
	var _arg2 C.gint                         // in
	var _arg3 C.gint                         // in

	_arg0 = (*C.GtkSourceGutterRendererText)(unsafe.Pointer(renderer.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_source_gutter_renderer_text_measure(_arg0, _arg1, &_arg2, &_arg3)
	runtime.KeepAlive(renderer)
	runtime.KeepAlive(text)

	var _width int  // out
	var _height int // out

	_width = int(_arg2)
	_height = int(_arg3)

	return _width, _height
}

// MeasureMarkup measures the pango markup provided using the pango layout used
// by the SourceGutterRendererText.
//
// The function takes the following parameters:
//
//    - markup: pango markup to measure.
//
// The function returns the following values:
//
//    - width (optional): location to store the width of the text in pixels, or
//      NULL.
//    - height (optional): location to store the height of the text in pixels, or
//      NULL.
//
func (renderer *GutterRendererText) MeasureMarkup(markup string) (width int, height int) {
	var _arg0 *C.GtkSourceGutterRendererText // out
	var _arg1 *C.gchar                       // out
	var _arg2 C.gint                         // in
	var _arg3 C.gint                         // in

	_arg0 = (*C.GtkSourceGutterRendererText)(unsafe.Pointer(renderer.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(markup)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_source_gutter_renderer_text_measure_markup(_arg0, _arg1, &_arg2, &_arg3)
	runtime.KeepAlive(renderer)
	runtime.KeepAlive(markup)

	var _width int  // out
	var _height int // out

	_width = int(_arg2)
	_height = int(_arg3)

	return _width, _height
}

// The function takes the following parameters:
//
//    - markup
//    - length
//
func (renderer *GutterRendererText) SetMarkup(markup string, length int) {
	var _arg0 *C.GtkSourceGutterRendererText // out
	var _arg1 *C.gchar                       // out
	var _arg2 C.gint                         // out

	_arg0 = (*C.GtkSourceGutterRendererText)(unsafe.Pointer(renderer.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(markup)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gint(length)

	C.gtk_source_gutter_renderer_text_set_markup(_arg0, _arg1, _arg2)
	runtime.KeepAlive(renderer)
	runtime.KeepAlive(markup)
	runtime.KeepAlive(length)
}

// The function takes the following parameters:
//
//    - text
//    - length
//
func (renderer *GutterRendererText) SetText(text string, length int) {
	var _arg0 *C.GtkSourceGutterRendererText // out
	var _arg1 *C.gchar                       // out
	var _arg2 C.gint                         // out

	_arg0 = (*C.GtkSourceGutterRendererText)(unsafe.Pointer(renderer.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gint(length)

	C.gtk_source_gutter_renderer_text_set_text(_arg0, _arg1, _arg2)
	runtime.KeepAlive(renderer)
	runtime.KeepAlive(text)
	runtime.KeepAlive(length)
}
