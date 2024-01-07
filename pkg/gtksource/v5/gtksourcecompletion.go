// Code generated by girgen. DO NOT EDIT.

package gtksource

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtksourceview/gtksource.h>
// extern void _gotk4_gtksource5_Completion_ConnectShow(gpointer, guintptr);
// extern void _gotk4_gtksource5_Completion_ConnectProviderRemoved(gpointer, GtkSourceCompletionProvider*, guintptr);
// extern void _gotk4_gtksource5_Completion_ConnectProviderAdded(gpointer, GtkSourceCompletionProvider*, guintptr);
// extern void _gotk4_gtksource5_Completion_ConnectHide(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeCompletion = coreglib.Type(C.gtk_source_completion_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeCompletion, F: marshalCompletion},
	})
}

// CompletionOverrides contains methods that are overridable.
type CompletionOverrides struct {
}

func defaultCompletionOverrides(v *Completion) CompletionOverrides {
	return CompletionOverrides{}
}

// Completion: main Completion Object.
//
// The completion system helps the user when they writes some text, such as
// words, command names, functions, and suchlike. Proposals can be shown,
// to complete the text the user is writing. Each proposal can contain an
// additional piece of information (for example documentation), that is
// displayed when the "Details" button is clicked.
//
// Proposals are created via a completionprovider. There can be for example
// a provider to complete words (see completionwords), another provider
// for the completion of function names, etc. To add a provider, call
// completion.AddProvider.
//
// When several providers match, they are all shown in the completion window,
// but one can switch between providers: see the SourceCompletion::move-page
// signal. It is also possible to activate the first proposals with key
// bindings, see the SourceCompletion:accelerators property.
//
// The completionproposal interface represents a proposal.
//
// If a proposal contains extra information (see
// GTK_SOURCE_COMPLETION_COLUMN_DETAILS), it will be displayed in a supplemental
// details window, which appears when the "Details" button is clicked.
//
// Each view object is associated with a completion instance. This instance
// can be obtained with view.GetCompletion. The view class contains also the
// view::show-completion signal.
//
// A same completionprovider object can be used for several
// GtkSourceCompletion's.
type Completion struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Completion)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*Completion, *CompletionClass, CompletionOverrides](
		GTypeCompletion,
		initCompletionClass,
		wrapCompletion,
		defaultCompletionOverrides,
	)
}

func initCompletionClass(gclass unsafe.Pointer, overrides CompletionOverrides, classInitFunc func(*CompletionClass)) {
	if classInitFunc != nil {
		class := (*CompletionClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapCompletion(obj *coreglib.Object) *Completion {
	return &Completion{
		Object: obj,
	}
}

func marshalCompletion(p uintptr) (interface{}, error) {
	return wrapCompletion(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectHide: "hide" signal is emitted when the completion window should be
// hidden.
func (self *Completion) ConnectHide(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "hide", false, unsafe.Pointer(C._gotk4_gtksource5_Completion_ConnectHide), f)
}

// ConnectProviderAdded: "provided-added" signal is emitted when a new provider
// is added to the completion.
func (self *Completion) ConnectProviderAdded(f func(provider CompletionProviderer)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "provider-added", false, unsafe.Pointer(C._gotk4_gtksource5_Completion_ConnectProviderAdded), f)
}

// ConnectProviderRemoved: "provided-removed" signal is emitted when a provider
// has been removed from the completion.
func (self *Completion) ConnectProviderRemoved(f func(provider CompletionProviderer)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "provider-removed", false, unsafe.Pointer(C._gotk4_gtksource5_Completion_ConnectProviderRemoved), f)
}

// ConnectShow: "show" signal is emitted when the completion window should be
// shown.
func (self *Completion) ConnectShow(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "show", false, unsafe.Pointer(C._gotk4_gtksource5_Completion_ConnectShow), f)
}

// AddProvider adds a completionprovider to the list of providers to be queried
// for completion results.
//
// The function takes the following parameters:
//
//   - provider: SourceCompletionProvider.
//
func (self *Completion) AddProvider(provider CompletionProviderer) {
	var _arg0 *C.GtkSourceCompletion         // out
	var _arg1 *C.GtkSourceCompletionProvider // out

	_arg0 = (*C.GtkSourceCompletion)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(coreglib.InternObject(provider).Native()))

	C.gtk_source_completion_add_provider(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(provider)
}

func (self *Completion) BlockInteractive() {
	var _arg0 *C.GtkSourceCompletion // out

	_arg0 = (*C.GtkSourceCompletion)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	C.gtk_source_completion_block_interactive(_arg0)
	runtime.KeepAlive(self)
}

// Buffer gets the connected view's buffer.
//
// The function returns the following values:
//
//   - buffer: SourceBuffer.
//
func (self *Completion) Buffer() *Buffer {
	var _arg0 *C.GtkSourceCompletion // out
	var _cret *C.GtkSourceBuffer     // in

	_arg0 = (*C.GtkSourceCompletion)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_source_completion_get_buffer(_arg0)
	runtime.KeepAlive(self)

	var _buffer *Buffer // out

	_buffer = wrapBuffer(coreglib.Take(unsafe.Pointer(_cret)))

	return _buffer
}

// The function returns the following values:
//
func (self *Completion) PageSize() uint {
	var _arg0 *C.GtkSourceCompletion // out
	var _cret C.guint                // in

	_arg0 = (*C.GtkSourceCompletion)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_source_completion_get_page_size(_arg0)
	runtime.KeepAlive(self)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// View gets the view that owns the completion.
//
// The function returns the following values:
//
//   - view: SourceView.
//
func (self *Completion) View() *View {
	var _arg0 *C.GtkSourceCompletion // out
	var _cret *C.GtkSourceView       // in

	_arg0 = (*C.GtkSourceCompletion)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_source_completion_get_view(_arg0)
	runtime.KeepAlive(self)

	var _view *View // out

	_view = wrapView(coreglib.Take(unsafe.Pointer(_cret)))

	return _view
}

// Hide emits the "hide" signal.
//
// When the "hide" signal is emitted, the completion window will be dismissed.
func (self *Completion) Hide() {
	var _arg0 *C.GtkSourceCompletion // out

	_arg0 = (*C.GtkSourceCompletion)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	C.gtk_source_completion_hide(_arg0)
	runtime.KeepAlive(self)
}

// RemoveProvider removes a completionprovider previously added with
// completion.AddProvider.
//
// The function takes the following parameters:
//
//   - provider: SourceCompletionProvider.
//
func (self *Completion) RemoveProvider(provider CompletionProviderer) {
	var _arg0 *C.GtkSourceCompletion         // out
	var _arg1 *C.GtkSourceCompletionProvider // out

	_arg0 = (*C.GtkSourceCompletion)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(coreglib.InternObject(provider).Native()))

	C.gtk_source_completion_remove_provider(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(provider)
}

// The function takes the following parameters:
//
func (self *Completion) SetPageSize(pageSize uint) {
	var _arg0 *C.GtkSourceCompletion // out
	var _arg1 C.guint                // out

	_arg0 = (*C.GtkSourceCompletion)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.guint(pageSize)

	C.gtk_source_completion_set_page_size(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(pageSize)
}

// Show emits the "show" signal.
//
// When the "show" signal is emitted, the completion window will be displayed if
// there are any results available.
func (self *Completion) Show() {
	var _arg0 *C.GtkSourceCompletion // out

	_arg0 = (*C.GtkSourceCompletion)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	C.gtk_source_completion_show(_arg0)
	runtime.KeepAlive(self)
}

func (self *Completion) UnblockInteractive() {
	var _arg0 *C.GtkSourceCompletion // out

	_arg0 = (*C.GtkSourceCompletion)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	C.gtk_source_completion_unblock_interactive(_arg0)
	runtime.KeepAlive(self)
}

// CompletionFuZZYHighlight: this will add <b> tags around matched characters in
// haystack based on casefold_query.
//
// The function takes the following parameters:
//
//   - haystack: string to be highlighted.
//   - casefoldQuery: typed-text used to highlight haystack.
//
// The function returns the following values:
//
//   - attrList (optional) or NULL.
//
func CompletionFuZZYHighlight(haystack, casefoldQuery string) *pango.AttrList {
	var _arg1 *C.char          // out
	var _arg2 *C.char          // out
	var _cret *C.PangoAttrList // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(haystack)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(casefoldQuery)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_source_completion_fuzzy_highlight(_arg1, _arg2)
	runtime.KeepAlive(haystack)
	runtime.KeepAlive(casefoldQuery)

	var _attrList *pango.AttrList // out

	if _cret != nil {
		_attrList = (*pango.AttrList)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_attrList)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.pango_attr_list_unref((*C.PangoAttrList)(intern.C))
			},
		)
	}

	return _attrList
}

// CompletionFuZZYMatch: this helper function can do a fuzzy match for you
// giving a haystack and casefolded needle.
//
// Casefold your needle using glib.UTF8Casefold() before running the query.
//
// Score will be set with the score of the match upon success. Otherwise,
// it will be set to zero.
//
// The function takes the following parameters:
//
//   - haystack (optional): string to be searched.
//   - casefoldNeedle: g_utf8_casefold() version of the needle.
//
// The function returns the following values:
//
//   - priority (optional): optional location for the score of the match.
//   - ok: TRUE if haystack matched casefold_needle, otherwise FALSE.
//
func CompletionFuZZYMatch(haystack, casefoldNeedle string) (uint, bool) {
	var _arg1 *C.char    // out
	var _arg2 *C.char    // out
	var _arg3 C.guint    // in
	var _cret C.gboolean // in

	if haystack != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(haystack)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(casefoldNeedle)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_source_completion_fuzzy_match(_arg1, _arg2, &_arg3)
	runtime.KeepAlive(haystack)
	runtime.KeepAlive(casefoldNeedle)

	var _priority uint // out
	var _ok bool       // out

	_priority = uint(_arg3)
	if _cret != 0 {
		_ok = true
	}

	return _priority, _ok
}

// CompletionClass: instance of this type is always passed by reference.
type CompletionClass struct {
	*completionClass
}

// completionClass is the struct that's finalized.
type completionClass struct {
	native *C.GtkSourceCompletionClass
}
