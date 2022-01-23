// Code generated by girgen. DO NOT EDIT.

package gtksource

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtksourceview/gtksource.h>
// void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_source_search_context_get_type()), F: marshalSearchContexter},
	})
}

type SearchContext struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*SearchContext)(nil)
)

func wrapSearchContext(obj *externglib.Object) *SearchContext {
	return &SearchContext{
		Object: obj,
	}
}

func marshalSearchContexter(p uintptr) (interface{}, error) {
	return wrapSearchContext(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewSearchContext creates a new search context, associated with buffer, and
// customized with settings. If settings is NULL, a new SourceSearchSettings
// object will be created, that you can retrieve with
// gtk_source_search_context_get_settings().
//
// The function takes the following parameters:
//
//    - buffer: SourceBuffer.
//    - settings (optional) or NULL.
//
// The function returns the following values:
//
//    - searchContext: new search context.
//
func NewSearchContext(buffer *Buffer, settings *SearchSettings) *SearchContext {
	var _arg1 *C.GtkSourceBuffer         // out
	var _arg2 *C.GtkSourceSearchSettings // out
	var _cret *C.GtkSourceSearchContext  // in

	_arg1 = (*C.GtkSourceBuffer)(unsafe.Pointer(buffer.Native()))
	if settings != nil {
		_arg2 = (*C.GtkSourceSearchSettings)(unsafe.Pointer(settings.Native()))
	}

	_cret = C.gtk_source_search_context_new(_arg1, _arg2)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(settings)

	var _searchContext *SearchContext // out

	_searchContext = wrapSearchContext(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _searchContext
}

// Backward synchronous backward search. It is recommended to use the
// asynchronous functions instead, to not block the user interface. However, if
// you are sure that the buffer is small, this function is more convenient to
// use.
//
// If the SourceSearchSettings:wrap-around property is FALSE, this function
// doesn't try to wrap around.
//
// The has_wrapped_around out parameter is set independently of whether a match
// is found. So if this function returns FALSE, has_wrapped_around will have the
// same value as the SourceSearchSettings:wrap-around property.
//
// The function takes the following parameters:
//
//    - iter: start of search.
//
// The function returns the following values:
//
//    - matchStart (optional): return location for start of match, or NULL.
//    - matchEnd (optional): return location for end of match, or NULL.
//    - hasWrappedAround (optional): return location to know whether the search
//      has wrapped around, or NULL.
//    - ok: whether a match was found.
//
func (search *SearchContext) Backward(iter *gtk.TextIter) (matchStart *gtk.TextIter, matchEnd *gtk.TextIter, hasWrappedAround bool, ok bool) {
	var _arg0 *C.GtkSourceSearchContext // out
	var _arg1 *C.GtkTextIter            // out
	var _arg2 C.GtkTextIter             // in
	var _arg3 C.GtkTextIter             // in
	var _arg4 C.gboolean                // in
	var _cret C.gboolean                // in

	_arg0 = (*C.GtkSourceSearchContext)(unsafe.Pointer(search.Native()))
	_arg1 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(iter)))

	_cret = C.gtk_source_search_context_backward(_arg0, _arg1, &_arg2, &_arg3, &_arg4)
	runtime.KeepAlive(search)
	runtime.KeepAlive(iter)

	var _matchStart *gtk.TextIter // out
	var _matchEnd *gtk.TextIter   // out
	var _hasWrappedAround bool    // out
	var _ok bool                  // out

	_matchStart = (*gtk.TextIter)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))
	_matchEnd = (*gtk.TextIter)(gextras.NewStructNative(unsafe.Pointer((&_arg3))))
	if _arg4 != 0 {
		_hasWrappedAround = true
	}
	if _cret != 0 {
		_ok = true
	}

	return _matchStart, _matchEnd, _hasWrappedAround, _ok
}

// BackwardAsync asynchronous version of gtk_source_search_context_backward().
//
// See the documentation of gtk_source_search_context_backward() for more
// details.
//
// See the Result documentation to know how to use this function.
//
// If the operation is cancelled, the callback will only be called if
// cancellable was not NULL. gtk_source_search_context_backward_async() takes
// ownership of cancellable, so you can unref it after calling this function.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//    - iter: start of search.
//    - callback (optional) to call when the operation is finished.
//
func (search *SearchContext) BackwardAsync(ctx context.Context, iter *gtk.TextIter, callback gio.AsyncReadyCallback) {
	var _arg0 *C.GtkSourceSearchContext // out
	var _arg2 *C.GCancellable           // out
	var _arg1 *C.GtkTextIter            // out
	var _arg3 C.GAsyncReadyCallback     // out
	var _arg4 C.gpointer

	_arg0 = (*C.GtkSourceSearchContext)(unsafe.Pointer(search.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(iter)))
	if callback != nil {
		_arg3 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg4 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.gtk_source_search_context_backward_async(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(search)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(iter)
	runtime.KeepAlive(callback)
}

// BackwardFinish finishes a backward search started with
// gtk_source_search_context_backward_async().
//
// See the documentation of gtk_source_search_context_backward() for more
// details.
//
// The function takes the following parameters:
//
//    - result: Result.
//
// The function returns the following values:
//
//    - matchStart (optional): return location for start of match, or NULL.
//    - matchEnd (optional): return location for end of match, or NULL.
//    - hasWrappedAround (optional): return location to know whether the search
//      has wrapped around, or NULL.
//
func (search *SearchContext) BackwardFinish(result gio.AsyncResulter) (matchStart *gtk.TextIter, matchEnd *gtk.TextIter, hasWrappedAround bool, goerr error) {
	var _arg0 *C.GtkSourceSearchContext // out
	var _arg1 *C.GAsyncResult           // out
	var _arg2 C.GtkTextIter             // in
	var _arg3 C.GtkTextIter             // in
	var _arg4 C.gboolean                // in
	var _cerr *C.GError                 // in

	_arg0 = (*C.GtkSourceSearchContext)(unsafe.Pointer(search.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	C.gtk_source_search_context_backward_finish(_arg0, _arg1, &_arg2, &_arg3, &_arg4, &_cerr)
	runtime.KeepAlive(search)
	runtime.KeepAlive(result)

	var _matchStart *gtk.TextIter // out
	var _matchEnd *gtk.TextIter   // out
	var _hasWrappedAround bool    // out
	var _goerr error              // out

	_matchStart = (*gtk.TextIter)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))
	_matchEnd = (*gtk.TextIter)(gextras.NewStructNative(unsafe.Pointer((&_arg3))))
	if _arg4 != 0 {
		_hasWrappedAround = true
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _matchStart, _matchEnd, _hasWrappedAround, _goerr
}

// Forward synchronous forward search. It is recommended to use the asynchronous
// functions instead, to not block the user interface. However, if you are sure
// that the buffer is small, this function is more convenient to use.
//
// If the SourceSearchSettings:wrap-around property is FALSE, this function
// doesn't try to wrap around.
//
// The has_wrapped_around out parameter is set independently of whether a match
// is found. So if this function returns FALSE, has_wrapped_around will have the
// same value as the SourceSearchSettings:wrap-around property.
//
// The function takes the following parameters:
//
//    - iter: start of search.
//
// The function returns the following values:
//
//    - matchStart (optional): return location for start of match, or NULL.
//    - matchEnd (optional): return location for end of match, or NULL.
//    - hasWrappedAround (optional): return location to know whether the search
//      has wrapped around, or NULL.
//    - ok: whether a match was found.
//
func (search *SearchContext) Forward(iter *gtk.TextIter) (matchStart *gtk.TextIter, matchEnd *gtk.TextIter, hasWrappedAround bool, ok bool) {
	var _arg0 *C.GtkSourceSearchContext // out
	var _arg1 *C.GtkTextIter            // out
	var _arg2 C.GtkTextIter             // in
	var _arg3 C.GtkTextIter             // in
	var _arg4 C.gboolean                // in
	var _cret C.gboolean                // in

	_arg0 = (*C.GtkSourceSearchContext)(unsafe.Pointer(search.Native()))
	_arg1 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(iter)))

	_cret = C.gtk_source_search_context_forward(_arg0, _arg1, &_arg2, &_arg3, &_arg4)
	runtime.KeepAlive(search)
	runtime.KeepAlive(iter)

	var _matchStart *gtk.TextIter // out
	var _matchEnd *gtk.TextIter   // out
	var _hasWrappedAround bool    // out
	var _ok bool                  // out

	_matchStart = (*gtk.TextIter)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))
	_matchEnd = (*gtk.TextIter)(gextras.NewStructNative(unsafe.Pointer((&_arg3))))
	if _arg4 != 0 {
		_hasWrappedAround = true
	}
	if _cret != 0 {
		_ok = true
	}

	return _matchStart, _matchEnd, _hasWrappedAround, _ok
}

// ForwardAsync asynchronous version of gtk_source_search_context_forward().
//
// See the documentation of gtk_source_search_context_forward() for more
// details.
//
// See the Result documentation to know how to use this function.
//
// If the operation is cancelled, the callback will only be called if
// cancellable was not NULL. gtk_source_search_context_forward_async() takes
// ownership of cancellable, so you can unref it after calling this function.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//    - iter: start of search.
//    - callback (optional) to call when the operation is finished.
//
func (search *SearchContext) ForwardAsync(ctx context.Context, iter *gtk.TextIter, callback gio.AsyncReadyCallback) {
	var _arg0 *C.GtkSourceSearchContext // out
	var _arg2 *C.GCancellable           // out
	var _arg1 *C.GtkTextIter            // out
	var _arg3 C.GAsyncReadyCallback     // out
	var _arg4 C.gpointer

	_arg0 = (*C.GtkSourceSearchContext)(unsafe.Pointer(search.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(iter)))
	if callback != nil {
		_arg3 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg4 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.gtk_source_search_context_forward_async(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(search)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(iter)
	runtime.KeepAlive(callback)
}

// ForwardFinish finishes a forward search started with
// gtk_source_search_context_forward_async().
//
// See the documentation of gtk_source_search_context_forward() for more
// details.
//
// The function takes the following parameters:
//
//    - result: Result.
//
// The function returns the following values:
//
//    - matchStart (optional): return location for start of match, or NULL.
//    - matchEnd (optional): return location for end of match, or NULL.
//    - hasWrappedAround (optional): return location to know whether the search
//      has wrapped around, or NULL.
//
func (search *SearchContext) ForwardFinish(result gio.AsyncResulter) (matchStart *gtk.TextIter, matchEnd *gtk.TextIter, hasWrappedAround bool, goerr error) {
	var _arg0 *C.GtkSourceSearchContext // out
	var _arg1 *C.GAsyncResult           // out
	var _arg2 C.GtkTextIter             // in
	var _arg3 C.GtkTextIter             // in
	var _arg4 C.gboolean                // in
	var _cerr *C.GError                 // in

	_arg0 = (*C.GtkSourceSearchContext)(unsafe.Pointer(search.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	C.gtk_source_search_context_forward_finish(_arg0, _arg1, &_arg2, &_arg3, &_arg4, &_cerr)
	runtime.KeepAlive(search)
	runtime.KeepAlive(result)

	var _matchStart *gtk.TextIter // out
	var _matchEnd *gtk.TextIter   // out
	var _hasWrappedAround bool    // out
	var _goerr error              // out

	_matchStart = (*gtk.TextIter)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))
	_matchEnd = (*gtk.TextIter)(gextras.NewStructNative(unsafe.Pointer((&_arg3))))
	if _arg4 != 0 {
		_hasWrappedAround = true
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _matchStart, _matchEnd, _hasWrappedAround, _goerr
}

// The function returns the following values:
//
//    - buffer: associated buffer.
//
func (search *SearchContext) Buffer() *Buffer {
	var _arg0 *C.GtkSourceSearchContext // out
	var _cret *C.GtkSourceBuffer        // in

	_arg0 = (*C.GtkSourceSearchContext)(unsafe.Pointer(search.Native()))

	_cret = C.gtk_source_search_context_get_buffer(_arg0)
	runtime.KeepAlive(search)

	var _buffer *Buffer // out

	_buffer = wrapBuffer(externglib.Take(unsafe.Pointer(_cret)))

	return _buffer
}

// The function returns the following values:
//
//    - ok: whether to highlight the search occurrences.
//
func (search *SearchContext) Highlight() bool {
	var _arg0 *C.GtkSourceSearchContext // out
	var _cret C.gboolean                // in

	_arg0 = (*C.GtkSourceSearchContext)(unsafe.Pointer(search.Native()))

	_cret = C.gtk_source_search_context_get_highlight(_arg0)
	runtime.KeepAlive(search)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// The function returns the following values:
//
//    - style to apply on search matches.
//
func (search *SearchContext) MatchStyle() *Style {
	var _arg0 *C.GtkSourceSearchContext // out
	var _cret *C.GtkSourceStyle         // in

	_arg0 = (*C.GtkSourceSearchContext)(unsafe.Pointer(search.Native()))

	_cret = C.gtk_source_search_context_get_match_style(_arg0)
	runtime.KeepAlive(search)

	var _style *Style // out

	_style = wrapStyle(externglib.Take(unsafe.Pointer(_cret)))

	return _style
}

// OccurrencePosition gets the position of a search occurrence. If the buffer is
// not already fully scanned, the position may be unknown, and -1 is returned.
// If 0 is returned, it means that this part of the buffer has already been
// scanned, and that match_start and match_end don't delimit an occurrence.
//
// The function takes the following parameters:
//
//    - matchStart: start of the occurrence.
//    - matchEnd: end of the occurrence.
//
// The function returns the following values:
//
//    - gint: position of the search occurrence. The first occurrence has the
//      position 1 (not 0). Returns 0 if match_start and match_end don't delimit
//      an occurrence. Returns -1 if the position is not yet known.
//
func (search *SearchContext) OccurrencePosition(matchStart, matchEnd *gtk.TextIter) int {
	var _arg0 *C.GtkSourceSearchContext // out
	var _arg1 *C.GtkTextIter            // out
	var _arg2 *C.GtkTextIter            // out
	var _cret C.gint                    // in

	_arg0 = (*C.GtkSourceSearchContext)(unsafe.Pointer(search.Native()))
	_arg1 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(matchStart)))
	_arg2 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(matchEnd)))

	_cret = C.gtk_source_search_context_get_occurrence_position(_arg0, _arg1, _arg2)
	runtime.KeepAlive(search)
	runtime.KeepAlive(matchStart)
	runtime.KeepAlive(matchEnd)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// OccurrencesCount gets the total number of search occurrences. If the buffer
// is not already fully scanned, the total number of occurrences is unknown, and
// -1 is returned.
//
// The function returns the following values:
//
//    - gint: total number of search occurrences, or -1 if unknown.
//
func (search *SearchContext) OccurrencesCount() int {
	var _arg0 *C.GtkSourceSearchContext // out
	var _cret C.gint                    // in

	_arg0 = (*C.GtkSourceSearchContext)(unsafe.Pointer(search.Native()))

	_cret = C.gtk_source_search_context_get_occurrences_count(_arg0)
	runtime.KeepAlive(search)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// RegexError: regular expression patterns must follow certain rules. If
// SourceSearchSettings:search-text breaks a rule, the error can be retrieved
// with this function. The error domain is REGEX_ERROR.
//
// Free the return value with g_error_free().
//
// The function returns the following values:
//
//    - err (optional) or NULL if the pattern is valid.
//
func (search *SearchContext) RegexError() error {
	var _arg0 *C.GtkSourceSearchContext // out
	var _cret *C.GError                 // in

	_arg0 = (*C.GtkSourceSearchContext)(unsafe.Pointer(search.Native()))

	_cret = C.gtk_source_search_context_get_regex_error(_arg0)
	runtime.KeepAlive(search)

	var _err error // out

	if _cret != nil {
		_err = gerror.Take(unsafe.Pointer(_cret))
	}

	return _err
}

// The function returns the following values:
//
//    - searchSettings: search settings.
//
func (search *SearchContext) Settings() *SearchSettings {
	var _arg0 *C.GtkSourceSearchContext  // out
	var _cret *C.GtkSourceSearchSettings // in

	_arg0 = (*C.GtkSourceSearchContext)(unsafe.Pointer(search.Native()))

	_cret = C.gtk_source_search_context_get_settings(_arg0)
	runtime.KeepAlive(search)

	var _searchSettings *SearchSettings // out

	_searchSettings = wrapSearchSettings(externglib.Take(unsafe.Pointer(_cret)))

	return _searchSettings
}

// Replace replaces a search match by another text. If match_start and match_end
// doesn't correspond to a search match, FALSE is returned.
//
// match_start and match_end iters are revalidated to point to the replacement
// text boundaries.
//
// For a regular expression replacement, you can check if replace is valid by
// calling g_regex_check_replacement(). The replace text can contain
// backreferences; read the g_regex_replace() documentation for more details.
//
// The function takes the following parameters:
//
//    - matchStart: start of the match to replace.
//    - matchEnd: end of the match to replace.
//    - replace: replacement text.
//    - replaceLength: length of replace in bytes, or -1.
//
func (search *SearchContext) Replace(matchStart, matchEnd *gtk.TextIter, replace string, replaceLength int) error {
	var _arg0 *C.GtkSourceSearchContext // out
	var _arg1 *C.GtkTextIter            // out
	var _arg2 *C.GtkTextIter            // out
	var _arg3 *C.gchar                  // out
	var _arg4 C.gint                    // out
	var _cerr *C.GError                 // in

	_arg0 = (*C.GtkSourceSearchContext)(unsafe.Pointer(search.Native()))
	_arg1 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(matchStart)))
	_arg2 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(matchEnd)))
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(replace)))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = C.gint(replaceLength)

	C.gtk_source_search_context_replace(_arg0, _arg1, _arg2, _arg3, _arg4, &_cerr)
	runtime.KeepAlive(search)
	runtime.KeepAlive(matchStart)
	runtime.KeepAlive(matchEnd)
	runtime.KeepAlive(replace)
	runtime.KeepAlive(replaceLength)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// ReplaceAll replaces all search matches by another text. It is a synchronous
// function, so it can block the user interface.
//
// For a regular expression replacement, you can check if replace is valid by
// calling g_regex_check_replacement(). The replace text can contain
// backreferences; read the g_regex_replace() documentation for more details.
//
// The function takes the following parameters:
//
//    - replace: replacement text.
//    - replaceLength: length of replace in bytes, or -1.
//
// The function returns the following values:
//
//    - guint: number of replaced matches.
//
func (search *SearchContext) ReplaceAll(replace string, replaceLength int) (uint, error) {
	var _arg0 *C.GtkSourceSearchContext // out
	var _arg1 *C.gchar                  // out
	var _arg2 C.gint                    // out
	var _cret C.guint                   // in
	var _cerr *C.GError                 // in

	_arg0 = (*C.GtkSourceSearchContext)(unsafe.Pointer(search.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(replace)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gint(replaceLength)

	_cret = C.gtk_source_search_context_replace_all(_arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(search)
	runtime.KeepAlive(replace)
	runtime.KeepAlive(replaceLength)

	var _guint uint  // out
	var _goerr error // out

	_guint = uint(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _guint, _goerr
}

// SetHighlight enables or disables the search occurrences highlighting.
//
// The function takes the following parameters:
//
//    - highlight: setting.
//
func (search *SearchContext) SetHighlight(highlight bool) {
	var _arg0 *C.GtkSourceSearchContext // out
	var _arg1 C.gboolean                // out

	_arg0 = (*C.GtkSourceSearchContext)(unsafe.Pointer(search.Native()))
	if highlight {
		_arg1 = C.TRUE
	}

	C.gtk_source_search_context_set_highlight(_arg0, _arg1)
	runtime.KeepAlive(search)
	runtime.KeepAlive(highlight)
}

// SetMatchStyle: set the style to apply on search matches. If match_style is
// NULL, default theme's scheme 'match-style' will be used. To enable or disable
// the search highlighting, use gtk_source_search_context_set_highlight().
//
// The function takes the following parameters:
//
//    - matchStyle (optional) or NULL.
//
func (search *SearchContext) SetMatchStyle(matchStyle *Style) {
	var _arg0 *C.GtkSourceSearchContext // out
	var _arg1 *C.GtkSourceStyle         // out

	_arg0 = (*C.GtkSourceSearchContext)(unsafe.Pointer(search.Native()))
	if matchStyle != nil {
		_arg1 = (*C.GtkSourceStyle)(unsafe.Pointer(matchStyle.Native()))
	}

	C.gtk_source_search_context_set_match_style(_arg0, _arg1)
	runtime.KeepAlive(search)
	runtime.KeepAlive(matchStyle)
}
