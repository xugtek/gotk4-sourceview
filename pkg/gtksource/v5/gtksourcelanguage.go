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
import "C"

// glib.Type values for gtksourcelanguage.go.
var GTypeLanguage = externglib.Type(C.gtk_source_language_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeLanguage, F: marshalLanguage},
	})
}

// LanguageOverrider contains methods that are overridable.
type LanguageOverrider interface {
}

type Language struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*Language)(nil)
)

func classInitLanguager(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapLanguage(obj *externglib.Object) *Language {
	return &Language{
		Object: obj,
	}
}

func marshalLanguage(p uintptr) (interface{}, error) {
	return wrapLanguage(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Globs returns the globs associated to this language. This is just an utility
// wrapper around gtk_source_language_get_metadata() to retrieve the "globs"
// metadata property and split it into an array.
//
// The function returns the following values:
//
//    - utf8s (optional): a newly-allocated NULL terminated array containing the
//      globs or NULL if no globs are found. The returned array must be freed
//      with g_strfreev().
//
func (language *Language) Globs() []string {
	var _arg0 *C.GtkSourceLanguage // out
	var _cret **C.gchar            // in

	_arg0 = (*C.GtkSourceLanguage)(unsafe.Pointer(externglib.InternObject(language).Native()))

	_cret = C.gtk_source_language_get_globs(_arg0)
	runtime.KeepAlive(language)

	var _utf8s []string // out

	if _cret != nil {
		defer C.free(unsafe.Pointer(_cret))
		{
			var i int
			var z *C.gchar
			for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
				i++
			}

			src := unsafe.Slice(_cret, i)
			_utf8s = make([]string, i)
			for i := range src {
				_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
				defer C.free(unsafe.Pointer(src[i]))
			}
		}
	}

	return _utf8s
}

// Hidden returns whether the language should be hidden from the user.
//
// The function returns the following values:
//
//    - ok: TRUE if the language should be hidden, FALSE otherwise.
//
func (language *Language) Hidden() bool {
	var _arg0 *C.GtkSourceLanguage // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkSourceLanguage)(unsafe.Pointer(externglib.InternObject(language).Native()))

	_cret = C.gtk_source_language_get_hidden(_arg0)
	runtime.KeepAlive(language)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ID returns the ID of the language. The ID is not locale-dependent. The
// returned string is owned by language and should not be freed or modified.
//
// The function returns the following values:
//
//    - utf8: ID of language.
//
func (language *Language) ID() string {
	var _arg0 *C.GtkSourceLanguage // out
	var _cret *C.gchar             // in

	_arg0 = (*C.GtkSourceLanguage)(unsafe.Pointer(externglib.InternObject(language).Native()))

	_cret = C.gtk_source_language_get_id(_arg0)
	runtime.KeepAlive(language)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// The function takes the following parameters:
//
//    - name: metadata property name.
//
// The function returns the following values:
//
//    - utf8 (optional): value of property name stored in the metadata of
//      language or NULL if language does not contain the specified metadata
//      property. The returned string is owned by language and should not be
//      freed or modified.
//
func (language *Language) Metadata(name string) string {
	var _arg0 *C.GtkSourceLanguage // out
	var _arg1 *C.gchar             // out
	var _cret *C.gchar             // in

	_arg0 = (*C.GtkSourceLanguage)(unsafe.Pointer(externglib.InternObject(language).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_source_language_get_metadata(_arg0, _arg1)
	runtime.KeepAlive(language)
	runtime.KeepAlive(name)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// MIMETypes returns the mime types associated to this language. This is just an
// utility wrapper around gtk_source_language_get_metadata() to retrieve the
// "mimetypes" metadata property and split it into an array.
//
// The function returns the following values:
//
//    - utf8s (optional): a newly-allocated NULL terminated array containing the
//      mime types or NULL if no mime types are found. The returned array must be
//      freed with g_strfreev().
//
func (language *Language) MIMETypes() []string {
	var _arg0 *C.GtkSourceLanguage // out
	var _cret **C.gchar            // in

	_arg0 = (*C.GtkSourceLanguage)(unsafe.Pointer(externglib.InternObject(language).Native()))

	_cret = C.gtk_source_language_get_mime_types(_arg0)
	runtime.KeepAlive(language)

	var _utf8s []string // out

	if _cret != nil {
		defer C.free(unsafe.Pointer(_cret))
		{
			var i int
			var z *C.gchar
			for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
				i++
			}

			src := unsafe.Slice(_cret, i)
			_utf8s = make([]string, i)
			for i := range src {
				_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
				defer C.free(unsafe.Pointer(src[i]))
			}
		}
	}

	return _utf8s
}

// Name returns the localized name of the language. The returned string is owned
// by language and should not be freed or modified.
//
// The function returns the following values:
//
//    - utf8: name of language.
//
func (language *Language) Name() string {
	var _arg0 *C.GtkSourceLanguage // out
	var _cret *C.gchar             // in

	_arg0 = (*C.GtkSourceLanguage)(unsafe.Pointer(externglib.InternObject(language).Native()))

	_cret = C.gtk_source_language_get_name(_arg0)
	runtime.KeepAlive(language)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Section returns the localized section of the language. Each language belong
// to a section (ex. HTML belongs to the Markup section). The returned string is
// owned by language and should not be freed or modified.
//
// The function returns the following values:
//
//    - utf8: section of language.
//
func (language *Language) Section() string {
	var _arg0 *C.GtkSourceLanguage // out
	var _cret *C.gchar             // in

	_arg0 = (*C.GtkSourceLanguage)(unsafe.Pointer(externglib.InternObject(language).Native()))

	_cret = C.gtk_source_language_get_section(_arg0)
	runtime.KeepAlive(language)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// StyleFallback returns the ID of the style to use if the specified style_id is
// not present in the current style scheme.
//
// The function takes the following parameters:
//
//    - styleId: style ID.
//
// The function returns the following values:
//
//    - utf8 (optional): ID of the style to use if the specified style_id is not
//      present in the current style scheme or NULL if the style has no fallback
//      defined. The returned string is owned by the language and must not be
//      modified.
//
func (language *Language) StyleFallback(styleId string) string {
	var _arg0 *C.GtkSourceLanguage // out
	var _arg1 *C.gchar             // out
	var _cret *C.gchar             // in

	_arg0 = (*C.GtkSourceLanguage)(unsafe.Pointer(externglib.InternObject(language).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(styleId)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_source_language_get_style_fallback(_arg0, _arg1)
	runtime.KeepAlive(language)
	runtime.KeepAlive(styleId)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// StyleIDs returns the ids of the styles defined by this language.
//
// The function returns the following values:
//
//    - utf8s (optional): a newly-allocated NULL terminated array containing ids
//      of the styles defined by this language or NULL if no style is defined.
//      The returned array must be freed with g_strfreev().
//
func (language *Language) StyleIDs() []string {
	var _arg0 *C.GtkSourceLanguage // out
	var _cret **C.gchar            // in

	_arg0 = (*C.GtkSourceLanguage)(unsafe.Pointer(externglib.InternObject(language).Native()))

	_cret = C.gtk_source_language_get_style_ids(_arg0)
	runtime.KeepAlive(language)

	var _utf8s []string // out

	if _cret != nil {
		defer C.free(unsafe.Pointer(_cret))
		{
			var i int
			var z *C.gchar
			for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
				i++
			}

			src := unsafe.Slice(_cret, i)
			_utf8s = make([]string, i)
			for i := range src {
				_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
				defer C.free(unsafe.Pointer(src[i]))
			}
		}
	}

	return _utf8s
}

// StyleName returns the name of the style with ID style_id defined by this
// language.
//
// The function takes the following parameters:
//
//    - styleId: style ID.
//
// The function returns the following values:
//
//    - utf8 (optional): name of the style with ID style_id defined by this
//      language or NULL if the style has no name or there is no style with ID
//      style_id defined by this language. The returned string is owned by the
//      language and must not be modified.
//
func (language *Language) StyleName(styleId string) string {
	var _arg0 *C.GtkSourceLanguage // out
	var _arg1 *C.gchar             // out
	var _cret *C.gchar             // in

	_arg0 = (*C.GtkSourceLanguage)(unsafe.Pointer(externglib.InternObject(language).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(styleId)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_source_language_get_style_name(_arg0, _arg1)
	runtime.KeepAlive(language)
	runtime.KeepAlive(styleId)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}
