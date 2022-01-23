// Code generated by girgen. DO NOT EDIT.

package gtksource

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v3"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtksourceview/gtksource.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_source_mark_get_type()), F: marshalMarker},
	})
}

type Mark struct {
	_ [0]func() // equal guard
	gtk.TextMark
}

var (
	_ externglib.Objector = (*Mark)(nil)
)

func wrapMark(obj *externglib.Object) *Mark {
	return &Mark{
		TextMark: gtk.TextMark{
			Object: obj,
		},
	}
}

func marshalMarker(p uintptr) (interface{}, error) {
	return wrapMark(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewMark creates a text mark. Add it to a buffer using
// gtk_text_buffer_add_mark(). If name is NULL, the mark is anonymous;
// otherwise, the mark can be retrieved by name using
// gtk_text_buffer_get_mark(). Normally marks are created using the utility
// function gtk_source_buffer_create_source_mark().
//
// The function takes the following parameters:
//
//    - name: name of the SourceMark, can be NULL when not using a name.
//    - category is used to classify marks according to common characteristics
//      (e.g. all the marks representing a bookmark could belong to the
//      "bookmark" category, or all the marks representing a compilation error
//      could belong to "error" category).
//
// The function returns the following values:
//
//    - mark: new SourceMark that can be added using gtk_text_buffer_add_mark().
//
func NewMark(name, category string) *Mark {
	var _arg1 *C.gchar         // out
	var _arg2 *C.gchar         // out
	var _cret *C.GtkSourceMark // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(category)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_source_mark_new(_arg1, _arg2)
	runtime.KeepAlive(name)
	runtime.KeepAlive(category)

	var _mark *Mark // out

	_mark = wrapMark(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _mark
}

// Category returns the mark category.
//
// The function returns the following values:
//
//    - utf8: category of the SourceMark.
//
func (mark *Mark) Category() string {
	var _arg0 *C.GtkSourceMark // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GtkSourceMark)(unsafe.Pointer(mark.Native()))

	_cret = C.gtk_source_mark_get_category(_arg0)
	runtime.KeepAlive(mark)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Next returns the next SourceMark in the buffer or NULL if the mark was not
// added to a buffer. If there is no next mark, NULL will be returned.
//
// If category is NULL, looks for marks of any category.
//
// The function takes the following parameters:
//
//    - category (optional): string specifying the mark category, or NULL.
//
// The function returns the following values:
//
//    - ret (optional): next SourceMark, or NULL.
//
func (mark *Mark) Next(category string) *Mark {
	var _arg0 *C.GtkSourceMark // out
	var _arg1 *C.gchar         // out
	var _cret *C.GtkSourceMark // in

	_arg0 = (*C.GtkSourceMark)(unsafe.Pointer(mark.Native()))
	if category != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(category)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.gtk_source_mark_next(_arg0, _arg1)
	runtime.KeepAlive(mark)
	runtime.KeepAlive(category)

	var _ret *Mark // out

	if _cret != nil {
		_ret = wrapMark(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _ret
}

// Prev returns the previous SourceMark in the buffer or NULL if the mark was
// not added to a buffer. If there is no previous mark, NULL is returned.
//
// If category is NULL, looks for marks of any category.
//
// The function takes the following parameters:
//
//    - category: string specifying the mark category, or NULL.
//
// The function returns the following values:
//
//    - ret (optional) previous SourceMark, or NULL.
//
func (mark *Mark) Prev(category string) *Mark {
	var _arg0 *C.GtkSourceMark // out
	var _arg1 *C.gchar         // out
	var _cret *C.GtkSourceMark // in

	_arg0 = (*C.GtkSourceMark)(unsafe.Pointer(mark.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(category)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_source_mark_prev(_arg0, _arg1)
	runtime.KeepAlive(mark)
	runtime.KeepAlive(category)

	var _ret *Mark // out

	if _cret != nil {
		_ret = wrapMark(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _ret
}
