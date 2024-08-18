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
	GTypeSearchSettings = coreglib.Type(C.gtk_source_search_settings_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSearchSettings, F: marshalSearchSettings},
	})
}

// SearchSettingsOverrides contains methods that are overridable.
type SearchSettingsOverrides struct {
}

func defaultSearchSettingsOverrides(v *SearchSettings) SearchSettingsOverrides {
	return SearchSettingsOverrides{}
}

// SearchSettings: search settings.
//
// A GtkSourceSearchSettings object represents the settings of a search.
// The search settings can be associated with one or several searchcontexts.
type SearchSettings struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*SearchSettings)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*SearchSettings, *SearchSettingsClass, SearchSettingsOverrides](
		GTypeSearchSettings,
		initSearchSettingsClass,
		wrapSearchSettings,
		defaultSearchSettingsOverrides,
	)
}

func initSearchSettingsClass(gclass unsafe.Pointer, overrides SearchSettingsOverrides, classInitFunc func(*SearchSettingsClass)) {
	if classInitFunc != nil {
		class := (*SearchSettingsClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapSearchSettings(obj *coreglib.Object) *SearchSettings {
	return &SearchSettings{
		Object: obj,
	}
}

func marshalSearchSettings(p uintptr) (interface{}, error) {
	return wrapSearchSettings(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewSearchSettings creates a new search settings object.
//
// The function returns the following values:
//
//   - searchSettings: new search settings object.
func NewSearchSettings() *SearchSettings {
	var _cret *C.GtkSourceSearchSettings // in

	_cret = C.gtk_source_search_settings_new()

	var _searchSettings *SearchSettings // out

	_searchSettings = wrapSearchSettings(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _searchSettings
}

// The function returns the following values:
//
//   - ok: whether to search at word boundaries.
func (settings *SearchSettings) AtWordBoundaries() bool {
	var _arg0 *C.GtkSourceSearchSettings // out
	var _cret C.gboolean                 // in

	_arg0 = (*C.GtkSourceSearchSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))

	_cret = C.gtk_source_search_settings_get_at_word_boundaries(_arg0)
	runtime.KeepAlive(settings)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// The function returns the following values:
//
//   - ok: whether the search is case sensitive.
func (settings *SearchSettings) CaseSensitive() bool {
	var _arg0 *C.GtkSourceSearchSettings // out
	var _cret C.gboolean                 // in

	_arg0 = (*C.GtkSourceSearchSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))

	_cret = C.gtk_source_search_settings_get_case_sensitive(_arg0)
	runtime.KeepAlive(settings)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// The function returns the following values:
//
//   - ok: whether to search by regular expressions.
func (settings *SearchSettings) RegexEnabled() bool {
	var _arg0 *C.GtkSourceSearchSettings // out
	var _cret C.gboolean                 // in

	_arg0 = (*C.GtkSourceSearchSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))

	_cret = C.gtk_source_search_settings_get_regex_enabled(_arg0)
	runtime.KeepAlive(settings)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SearchText gets the text to search.
//
// The return value must not be freed.
//
// You may be interested to call utils_escape_search_text after this function.
//
// The function returns the following values:
//
//   - utf8 (optional): text to search, or NULL if the search is disabled.
func (settings *SearchSettings) SearchText() string {
	var _arg0 *C.GtkSourceSearchSettings // out
	var _cret *C.gchar                   // in

	_arg0 = (*C.GtkSourceSearchSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))

	_cret = C.gtk_source_search_settings_get_search_text(_arg0)
	runtime.KeepAlive(settings)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// The function returns the following values:
//
//   - ok: whether to exclude invisible text from the search.
func (settings *SearchSettings) VisibleOnly() bool {
	var _arg0 *C.GtkSourceSearchSettings // out
	var _cret C.gboolean                 // in

	_arg0 = (*C.GtkSourceSearchSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))

	_cret = C.gtk_source_search_settings_get_visible_only(_arg0)
	runtime.KeepAlive(settings)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// The function returns the following values:
//
//   - ok: whether to wrap around the search.
func (settings *SearchSettings) WrapAround() bool {
	var _arg0 *C.GtkSourceSearchSettings // out
	var _cret C.gboolean                 // in

	_arg0 = (*C.GtkSourceSearchSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))

	_cret = C.gtk_source_search_settings_get_wrap_around(_arg0)
	runtime.KeepAlive(settings)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetAtWordBoundaries: change whether the search is done at word boundaries.
//
// If at_word_boundaries is TRUE, a search match must start and end a word.
// The match can span multiple words. See also gtk.TextIter.StartsWord() and
// gtk.TextIter.EndsWord().
//
// The function takes the following parameters:
//
//   - atWordBoundaries: setting.
func (settings *SearchSettings) SetAtWordBoundaries(atWordBoundaries bool) {
	var _arg0 *C.GtkSourceSearchSettings // out
	var _arg1 C.gboolean                 // out

	_arg0 = (*C.GtkSourceSearchSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))
	if atWordBoundaries {
		_arg1 = C.TRUE
	}

	C.gtk_source_search_settings_set_at_word_boundaries(_arg0, _arg1)
	runtime.KeepAlive(settings)
	runtime.KeepAlive(atWordBoundaries)
}

// SetCaseSensitive enables or disables the case sensitivity for the search.
//
// The function takes the following parameters:
//
//   - caseSensitive: setting.
func (settings *SearchSettings) SetCaseSensitive(caseSensitive bool) {
	var _arg0 *C.GtkSourceSearchSettings // out
	var _arg1 C.gboolean                 // out

	_arg0 = (*C.GtkSourceSearchSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))
	if caseSensitive {
		_arg1 = C.TRUE
	}

	C.gtk_source_search_settings_set_case_sensitive(_arg0, _arg1)
	runtime.KeepAlive(settings)
	runtime.KeepAlive(caseSensitive)
}

// SetRegexEnabled enables or disables whether to search by regular expressions.
//
// If enabled, the searchsettings:search-text property contains the pattern of
// the regular expression.
//
// searchcontext uses #GRegex when regex search
// is enabled. See the Regular expression syntax
// (https://developer.gnome.org/glib/stable/glib-regex-syntax.html) page in the
// GLib reference manual.
//
// The function takes the following parameters:
//
//   - regexEnabled: setting.
func (settings *SearchSettings) SetRegexEnabled(regexEnabled bool) {
	var _arg0 *C.GtkSourceSearchSettings // out
	var _arg1 C.gboolean                 // out

	_arg0 = (*C.GtkSourceSearchSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))
	if regexEnabled {
		_arg1 = C.TRUE
	}

	C.gtk_source_search_settings_set_regex_enabled(_arg0, _arg1)
	runtime.KeepAlive(settings)
	runtime.KeepAlive(regexEnabled)
}

// SetSearchText sets the text to search.
//
// If search_text is NULL or is empty, the search will be disabled. A copy of
// search_text will be made, so you can safely free search_text after a call to
// this function.
//
// You may be interested to call utils_unescape_search_text before this
// function.
//
// The function takes the following parameters:
//
//   - searchText (optional): nul-terminated text to search, or NULL to disable
//     the search.
func (settings *SearchSettings) SetSearchText(searchText string) {
	var _arg0 *C.GtkSourceSearchSettings // out
	var _arg1 *C.gchar                   // out

	_arg0 = (*C.GtkSourceSearchSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))
	if searchText != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(searchText)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_source_search_settings_set_search_text(_arg0, _arg1)
	runtime.KeepAlive(settings)
	runtime.KeepAlive(searchText)
}

// SetVisibleOnly enables or disables whether to exclude invisible text from the
// search.
//
// If enabled, only visible text will be searched. A search match may have
// invisible text interspersed.
//
// The function takes the following parameters:
//
//   - visibleOnly: setting.
func (settings *SearchSettings) SetVisibleOnly(visibleOnly bool) {
	var _arg0 *C.GtkSourceSearchSettings // out
	var _arg1 C.gboolean                 // out

	_arg0 = (*C.GtkSourceSearchSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))
	if visibleOnly {
		_arg1 = C.TRUE
	}

	C.gtk_source_search_settings_set_visible_only(_arg0, _arg1)
	runtime.KeepAlive(settings)
	runtime.KeepAlive(visibleOnly)
}

// SetWrapAround enables or disables the wrap around search.
//
// If wrap_around is TRUE, the forward search continues at the beginning of the
// buffer if no search occurrences are found. Similarly, the backward search
// continues to search at the end of the buffer.
//
// The function takes the following parameters:
//
//   - wrapAround: setting.
func (settings *SearchSettings) SetWrapAround(wrapAround bool) {
	var _arg0 *C.GtkSourceSearchSettings // out
	var _arg1 C.gboolean                 // out

	_arg0 = (*C.GtkSourceSearchSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))
	if wrapAround {
		_arg1 = C.TRUE
	}

	C.gtk_source_search_settings_set_wrap_around(_arg0, _arg1)
	runtime.KeepAlive(settings)
	runtime.KeepAlive(wrapAround)
}

// SearchSettingsClass: instance of this type is always passed by reference.
type SearchSettingsClass struct {
	*searchSettingsClass
}

// searchSettingsClass is the struct that's finalized.
type searchSettingsClass struct {
	native *C.GtkSourceSearchSettingsClass
}
