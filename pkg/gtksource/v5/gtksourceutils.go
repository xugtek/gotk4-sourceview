// Code generated by girgen. DO NOT EDIT.

package gtksource

import (
	"runtime"
	"unsafe"
)

// #include <stdlib.h>
// #include <gtksourceview/gtksource.h>
import "C"

// UtilsEscapeSearchText: use this function to escape the following characters:
// \n, \r, \t and \.
//
// For a regular expression search, use g_regex_escape_string() instead.
//
// One possible use case is to take the TextBuffer's selection and put it in a
// search entry. The selection can contain tabulations, newlines, etc. So it's
// better to escape those special characters to better fit in the search entry.
//
// See also: utils_unescape_search_text.
//
// <warning> Warning: the escape and unescape functions are not reciprocal! For
// example, escape (unescape (\)) = \\. So avoid cycles such as: search entry ->
// unescape -> search settings -> escape -> search entry. The original search
// entry text may be modified. </warning>.
//
// The function takes the following parameters:
//
//   - text to escape.
//
// The function returns the following values:
//
//   - utf8: escaped text.
func UtilsEscapeSearchText(text string) string {
	var _arg1 *C.gchar // out
	var _cret *C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_source_utils_escape_search_text(_arg1)
	runtime.KeepAlive(text)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// UtilsUnescapeSearchText: use this function before
// searchsettings.SetSearchText, to unescape the following sequences of
// characters: \n, \r, \t and \\. The purpose is to easily write those
// characters in a search entry.
//
// Note that unescaping the search text is not needed for regular expression
// searches.
//
// See also: utils_escape_search_text.
//
// The function takes the following parameters:
//
//   - text to unescape.
//
// The function returns the following values:
//
//   - utf8: unescaped text.
func UtilsUnescapeSearchText(text string) string {
	var _arg1 *C.gchar // out
	var _cret *C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_source_utils_unescape_search_text(_arg1)
	runtime.KeepAlive(text)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
