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
	GTypeMark = coreglib.Type(C.gtk_source_mark_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeMark, F: marshalMark},
	})
}

// MarkOverrides contains methods that are overridable.
type MarkOverrides struct {
}

func defaultMarkOverrides(v *Mark) MarkOverrides {
	return MarkOverrides{}
}

// Mark object for buffer.
//
// A GtkSourceMark marks a position in the text where you want to display
// additional info. It is based on gtk.TextMark and thus is still valid after
// the text has changed though its position may change.
//
// GtkSourceMarks are organized in categories which you have to set when
// you create the mark. Each category can have a priority, a pixbuf and
// other associated attributes. See view.SetMarkAttributes. The pixbuf will
// be displayed in the margin at the line where the mark residents if the
// view:show-line-marks property is set to TRUE. If there are multiple marks in
// the same line, the pixbufs will be drawn on top of each other. The mark with
// the highest priority will be drawn on top.
type Mark struct {
	_ [0]func() // equal guard
	gtk.TextMark
}

var (
	_ coreglib.Objector = (*Mark)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*Mark, *MarkClass, MarkOverrides](
		GTypeMark,
		initMarkClass,
		wrapMark,
		defaultMarkOverrides,
	)
}

func initMarkClass(gclass unsafe.Pointer, overrides MarkOverrides, classInitFunc func(*MarkClass)) {
	if classInitFunc != nil {
		class := (*MarkClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapMark(obj *coreglib.Object) *Mark {
	return &Mark{
		TextMark: gtk.TextMark{
			Object: obj,
		},
	}
}

func marshalMark(p uintptr) (interface{}, error) {
	return wrapMark(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewMark creates a text mark.
//
// Add it to a buffer using gtk.TextBuffer.AddMark(). If name is NULL,
// the mark is anonymous; otherwise, the mark can be retrieved by name using
// gtk.TextBuffer.GetMark(). Normally marks are created using the utility
// function buffer.CreateSourceMark.
//
// The function takes the following parameters:
//
//   - name (optional): name of the SourceMark or NULL.
//   - category is used to classify marks according to common characteristics
//     (e.g. all the marks representing a bookmark could belong to the
//     "bookmark" category, or all the marks representing a compilation error
//     could belong to "error" category).
//
// The function returns the following values:
//
//   - mark: new SourceMark that can be added using gtk.TextBuffer.AddMark().
func NewMark(name, category string) *Mark {
	var _arg1 *C.gchar         // out
	var _arg2 *C.gchar         // out
	var _cret *C.GtkSourceMark // in

	if name != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(category)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_source_mark_new(_arg1, _arg2)
	runtime.KeepAlive(name)
	runtime.KeepAlive(category)

	var _mark *Mark // out

	_mark = wrapMark(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _mark
}

// Category returns the mark category.
//
// The function returns the following values:
//
//   - utf8: category of the SourceMark.
func (mark *Mark) Category() string {
	var _arg0 *C.GtkSourceMark // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GtkSourceMark)(unsafe.Pointer(coreglib.InternObject(mark).Native()))

	_cret = C.gtk_source_mark_get_category(_arg0)
	runtime.KeepAlive(mark)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Next returns the next GtkSourceMark in the buffer or NULL if the mark was not
// added to a buffer.
//
//	If there is no next mark, NULL will be returned.
//
// If category is NULL, looks for marks of any category.
//
// The function takes the following parameters:
//
//   - category (optional): string specifying the mark category, or NULL.
//
// The function returns the following values:
//
//   - ret (optional): next SourceMark, or NULL.
func (mark *Mark) Next(category string) *Mark {
	var _arg0 *C.GtkSourceMark // out
	var _arg1 *C.gchar         // out
	var _cret *C.GtkSourceMark // in

	_arg0 = (*C.GtkSourceMark)(unsafe.Pointer(coreglib.InternObject(mark).Native()))
	if category != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(category)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.gtk_source_mark_next(_arg0, _arg1)
	runtime.KeepAlive(mark)
	runtime.KeepAlive(category)

	var _ret *Mark // out

	if _cret != nil {
		_ret = wrapMark(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _ret
}

// Prev returns the previous GtkSourceMark in the buffer or NULL if the mark was
// not added to a buffer.
//
// If there is no previous mark, NULL is returned.
//
// If category is NULL, looks for marks of any category.
//
// The function takes the following parameters:
//
//   - category (optional): string specifying the mark category, or NULL.
//
// The function returns the following values:
//
//   - ret (optional) previous SourceMark, or NULL.
func (mark *Mark) Prev(category string) *Mark {
	var _arg0 *C.GtkSourceMark // out
	var _arg1 *C.gchar         // out
	var _cret *C.GtkSourceMark // in

	_arg0 = (*C.GtkSourceMark)(unsafe.Pointer(coreglib.InternObject(mark).Native()))
	if category != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(category)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.gtk_source_mark_prev(_arg0, _arg1)
	runtime.KeepAlive(mark)
	runtime.KeepAlive(category)

	var _ret *Mark // out

	if _cret != nil {
		_ret = wrapMark(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _ret
}

// MarkClass: instance of this type is always passed by reference.
type MarkClass struct {
	*markClass
}

// markClass is the struct that's finalized.
type markClass struct {
	native *C.GtkSourceMarkClass
}

func (m *MarkClass) ParentClass() *gtk.TextMarkClass {
	valptr := &m.native.parent_class
	var _v *gtk.TextMarkClass // out
	_v = (*gtk.TextMarkClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
