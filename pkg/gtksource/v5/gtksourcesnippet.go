// Code generated by girgen. DO NOT EDIT.

package gtksource

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtksourceview/gtksource.h>
import "C"

// GType values.
var (
	GTypeSnippet = coreglib.Type(C.gtk_source_snippet_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSnippet, F: marshalSnippet},
	})
}

// SnippetOverrides contains methods that are overridable.
type SnippetOverrides struct {
}

func defaultSnippetOverrides(v *Snippet) SnippetOverrides {
	return SnippetOverrides{}
}

// Snippet: quick insertion code snippets.
//
// The GtkSourceSnippet represents a series of chunks that can quickly be
// inserted into the view.
//
// Snippets are defined in XML files which are loaded by the snippetmanager.
// Alternatively, applications can create snippets on demand and insert them
// into the view using view.PushSnippet.
//
// Snippet chunks can reference other snippet chunks as well as post-process the
// values from other chunks such as capitalization.
type Snippet struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Snippet)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*Snippet, *SnippetClass, SnippetOverrides](
		GTypeSnippet,
		initSnippetClass,
		wrapSnippet,
		defaultSnippetOverrides,
	)
}

func initSnippetClass(gclass unsafe.Pointer, overrides SnippetOverrides, classInitFunc func(*SnippetClass)) {
	if classInitFunc != nil {
		class := (*SnippetClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapSnippet(obj *coreglib.Object) *Snippet {
	return &Snippet{
		Object: obj,
	}
}

func marshalSnippet(p uintptr) (interface{}, error) {
	return wrapSnippet(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewSnippet creates a new SourceSnippet.
//
// The function takes the following parameters:
//
//   - trigger (optional) word.
//   - languageId (optional): source language.
//
// The function returns the following values:
//
//   - snippet: new SourceSnippet.
//
func NewSnippet(trigger, languageId string) *Snippet {
	var _arg1 *C.gchar            // out
	var _arg2 *C.gchar            // out
	var _cret *C.GtkSourceSnippet // in

	if trigger != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(trigger)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	if languageId != "" {
		_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(languageId)))
		defer C.free(unsafe.Pointer(_arg2))
	}

	_cret = C.gtk_source_snippet_new(_arg1, _arg2)
	runtime.KeepAlive(trigger)
	runtime.KeepAlive(languageId)

	var _snippet *Snippet // out

	_snippet = wrapSnippet(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _snippet
}

// NewSnippetParsed parses the snippet formatted text into a series of chunks
// and adds them to a new SourceSnippet.
//
// The function takes the following parameters:
//
//   - text: formatted snippet text to parse.
//
// The function returns the following values:
//
//   - snippet: newly parsed SourceSnippet, or NULL upon failure and error is
//     set.
//
func NewSnippetParsed(text string) (*Snippet, error) {
	var _arg1 *C.char             // out
	var _cret *C.GtkSourceSnippet // in
	var _cerr *C.GError           // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_source_snippet_new_parsed(_arg1, &_cerr)
	runtime.KeepAlive(text)

	var _snippet *Snippet // out
	var _goerr error      // out

	_snippet = wrapSnippet(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _snippet, _goerr
}

// AddChunk appends chunk to the snippet.
//
// This may only be called before the snippet has been expanded.
//
// The function takes the following parameters:
//
//   - chunk: SourceSnippetChunk.
//
func (snippet *Snippet) AddChunk(chunk *SnippetChunk) {
	var _arg0 *C.GtkSourceSnippet      // out
	var _arg1 *C.GtkSourceSnippetChunk // out

	_arg0 = (*C.GtkSourceSnippet)(unsafe.Pointer(coreglib.InternObject(snippet).Native()))
	_arg1 = (*C.GtkSourceSnippetChunk)(unsafe.Pointer(coreglib.InternObject(chunk).Native()))

	C.gtk_source_snippet_add_chunk(_arg0, _arg1)
	runtime.KeepAlive(snippet)
	runtime.KeepAlive(chunk)
}

// Copy does a deep copy of the snippet.
//
// The function returns the following values:
//
//   - ret: new SourceSnippet.
//
func (snippet *Snippet) Copy() *Snippet {
	var _arg0 *C.GtkSourceSnippet // out
	var _cret *C.GtkSourceSnippet // in

	_arg0 = (*C.GtkSourceSnippet)(unsafe.Pointer(coreglib.InternObject(snippet).Native()))

	_cret = C.gtk_source_snippet_copy(_arg0)
	runtime.KeepAlive(snippet)

	var _ret *Snippet // out

	_ret = wrapSnippet(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _ret
}

// Context gets the context used for expanding the snippet.
//
// The function returns the following values:
//
//   - snippetContext (optional): SourceSnippetContext.
//
func (snippet *Snippet) Context() *SnippetContext {
	var _arg0 *C.GtkSourceSnippet        // out
	var _cret *C.GtkSourceSnippetContext // in

	_arg0 = (*C.GtkSourceSnippet)(unsafe.Pointer(coreglib.InternObject(snippet).Native()))

	_cret = C.gtk_source_snippet_get_context(_arg0)
	runtime.KeepAlive(snippet)

	var _snippetContext *SnippetContext // out

	if _cret != nil {
		_snippetContext = wrapSnippetContext(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _snippetContext
}

// Description gets the description for the snippet.
//
// The function returns the following values:
//
func (snippet *Snippet) Description() string {
	var _arg0 *C.GtkSourceSnippet // out
	var _cret *C.gchar            // in

	_arg0 = (*C.GtkSourceSnippet)(unsafe.Pointer(coreglib.InternObject(snippet).Native()))

	_cret = C.gtk_source_snippet_get_description(_arg0)
	runtime.KeepAlive(snippet)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// FocusPosition gets the current focus for the snippet.
//
// This is changed as the user tabs through focus locations.
//
// The function returns the following values:
//
//   - gint focus position, or -1 if unset.
//
func (snippet *Snippet) FocusPosition() int {
	var _arg0 *C.GtkSourceSnippet // out
	var _cret C.gint              // in

	_arg0 = (*C.GtkSourceSnippet)(unsafe.Pointer(coreglib.InternObject(snippet).Native()))

	_cret = C.gtk_source_snippet_get_focus_position(_arg0)
	runtime.KeepAlive(snippet)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// LanguageID gets the language-id used for the source snippet.
//
// The language identifier should be one that matches a source language
// language:id property.
//
// The function returns the following values:
//
//   - utf8: language identifier.
//
func (snippet *Snippet) LanguageID() string {
	var _arg0 *C.GtkSourceSnippet // out
	var _cret *C.gchar            // in

	_arg0 = (*C.GtkSourceSnippet)(unsafe.Pointer(coreglib.InternObject(snippet).Native()))

	_cret = C.gtk_source_snippet_get_language_id(_arg0)
	runtime.KeepAlive(snippet)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// NChunks gets the number of chunks in the snippet.
//
// Note that not all chunks are editable.
//
// The function returns the following values:
//
//   - guint: number of chunks.
//
func (snippet *Snippet) NChunks() uint {
	var _arg0 *C.GtkSourceSnippet // out
	var _cret C.guint             // in

	_arg0 = (*C.GtkSourceSnippet)(unsafe.Pointer(coreglib.InternObject(snippet).Native()))

	_cret = C.gtk_source_snippet_get_n_chunks(_arg0)
	runtime.KeepAlive(snippet)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Name gets the name for the snippet.
//
// The function returns the following values:
//
func (snippet *Snippet) Name() string {
	var _arg0 *C.GtkSourceSnippet // out
	var _cret *C.gchar            // in

	_arg0 = (*C.GtkSourceSnippet)(unsafe.Pointer(coreglib.InternObject(snippet).Native()))

	_cret = C.gtk_source_snippet_get_name(_arg0)
	runtime.KeepAlive(snippet)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// NthChunk gets the chunk at nth.
//
// The function takes the following parameters:
//
//   - nth chunk to get.
//
// The function returns the following values:
//
//   - snippetChunk: SourceSnippetChunk.
//
func (snippet *Snippet) NthChunk(nth uint) *SnippetChunk {
	var _arg0 *C.GtkSourceSnippet      // out
	var _arg1 C.guint                  // out
	var _cret *C.GtkSourceSnippetChunk // in

	_arg0 = (*C.GtkSourceSnippet)(unsafe.Pointer(coreglib.InternObject(snippet).Native()))
	_arg1 = C.guint(nth)

	_cret = C.gtk_source_snippet_get_nth_chunk(_arg0, _arg1)
	runtime.KeepAlive(snippet)
	runtime.KeepAlive(nth)

	var _snippetChunk *SnippetChunk // out

	_snippetChunk = wrapSnippetChunk(coreglib.Take(unsafe.Pointer(_cret)))

	return _snippetChunk
}

// Trigger gets the trigger for the source snippet.
//
// A trigger is a word that can be expanded into the full snippet when the user
// presses Tab.
//
// The function returns the following values:
//
//   - utf8 (optional): string or NULL.
//
func (snippet *Snippet) Trigger() string {
	var _arg0 *C.GtkSourceSnippet // out
	var _cret *C.gchar            // in

	_arg0 = (*C.GtkSourceSnippet)(unsafe.Pointer(coreglib.InternObject(snippet).Native()))

	_cret = C.gtk_source_snippet_get_trigger(_arg0)
	runtime.KeepAlive(snippet)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// SetDescription sets the description for the snippet.
//
// The function takes the following parameters:
//
//   - description: snippet description.
//
func (snippet *Snippet) SetDescription(description string) {
	var _arg0 *C.GtkSourceSnippet // out
	var _arg1 *C.gchar            // out

	_arg0 = (*C.GtkSourceSnippet)(unsafe.Pointer(coreglib.InternObject(snippet).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(description)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_source_snippet_set_description(_arg0, _arg1)
	runtime.KeepAlive(snippet)
	runtime.KeepAlive(description)
}

// SetLanguageID sets the language identifier for the snippet.
//
// This should match the language:id identifier.
//
// The function takes the following parameters:
//
//   - languageId: language identifier for the snippet.
//
func (snippet *Snippet) SetLanguageID(languageId string) {
	var _arg0 *C.GtkSourceSnippet // out
	var _arg1 *C.gchar            // out

	_arg0 = (*C.GtkSourceSnippet)(unsafe.Pointer(coreglib.InternObject(snippet).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(languageId)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_source_snippet_set_language_id(_arg0, _arg1)
	runtime.KeepAlive(snippet)
	runtime.KeepAlive(languageId)
}

// SetName sets the name for the snippet.
//
// The function takes the following parameters:
//
//   - name: snippet name.
//
func (snippet *Snippet) SetName(name string) {
	var _arg0 *C.GtkSourceSnippet // out
	var _arg1 *C.gchar            // out

	_arg0 = (*C.GtkSourceSnippet)(unsafe.Pointer(coreglib.InternObject(snippet).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_source_snippet_set_name(_arg0, _arg1)
	runtime.KeepAlive(snippet)
	runtime.KeepAlive(name)
}

// SetTrigger sets the trigger for the snippet.
//
// The function takes the following parameters:
//
//   - trigger word.
//
func (snippet *Snippet) SetTrigger(trigger string) {
	var _arg0 *C.GtkSourceSnippet // out
	var _arg1 *C.gchar            // out

	_arg0 = (*C.GtkSourceSnippet)(unsafe.Pointer(coreglib.InternObject(snippet).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(trigger)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_source_snippet_set_trigger(_arg0, _arg1)
	runtime.KeepAlive(snippet)
	runtime.KeepAlive(trigger)
}

// SnippetClass: instance of this type is always passed by reference.
type SnippetClass struct {
	*snippetClass
}

// snippetClass is the struct that's finalized.
type snippetClass struct {
	native *C.GtkSourceSnippetClass
}
