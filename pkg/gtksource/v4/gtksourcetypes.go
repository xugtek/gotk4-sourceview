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
	GTypeEncoding = coreglib.Type(C.gtk_source_encoding_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeEncoding, F: marshalEncoding},
	})
}

// Encoding: instance of this type is always passed by reference.
type Encoding struct {
	*encoding
}

// encoding is the struct that's finalized.
type encoding struct {
	native *C.GtkSourceEncoding
}

func marshalEncoding(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &Encoding{&encoding{(*C.GtkSourceEncoding)(b)}}, nil
}

// Copy: used by language bindings.
//
// The function returns the following values:
//
//   - encoding: copy of enc.
func (enc *Encoding) Copy() *Encoding {
	var _arg0 *C.GtkSourceEncoding // out
	var _cret *C.GtkSourceEncoding // in

	_arg0 = (*C.GtkSourceEncoding)(gextras.StructNative(unsafe.Pointer(enc)))

	_cret = C.gtk_source_encoding_copy(_arg0)
	runtime.KeepAlive(enc)

	var _encoding *Encoding // out

	_encoding = (*Encoding)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_encoding)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_source_encoding_free((*C.GtkSourceEncoding)(intern.C))
		},
	)

	return _encoding
}

// Charset gets the character set of the SourceEncoding, such as "UTF-8" or
// "ISO-8859-1".
//
// The function returns the following values:
//
//   - utf8: character set of the SourceEncoding.
func (enc *Encoding) Charset() string {
	var _arg0 *C.GtkSourceEncoding // out
	var _cret *C.gchar             // in

	_arg0 = (*C.GtkSourceEncoding)(gextras.StructNative(unsafe.Pointer(enc)))

	_cret = C.gtk_source_encoding_get_charset(_arg0)
	runtime.KeepAlive(enc)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Name gets the name of the SourceEncoding such as "Unicode" or "Western".
//
// The function returns the following values:
//
//   - utf8: name of the SourceEncoding.
func (enc *Encoding) Name() string {
	var _arg0 *C.GtkSourceEncoding // out
	var _cret *C.gchar             // in

	_arg0 = (*C.GtkSourceEncoding)(gextras.StructNative(unsafe.Pointer(enc)))

	_cret = C.gtk_source_encoding_get_name(_arg0)
	runtime.KeepAlive(enc)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// The function returns the following values:
//
//   - utf8: string representation. Free with g_free() when no longer needed.
func (enc *Encoding) String() string {
	var _arg0 *C.GtkSourceEncoding // out
	var _cret *C.gchar             // in

	_arg0 = (*C.GtkSourceEncoding)(gextras.StructNative(unsafe.Pointer(enc)))

	_cret = C.gtk_source_encoding_to_string(_arg0)
	runtime.KeepAlive(enc)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
