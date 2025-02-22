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
	GTypeSnippetChunk = coreglib.Type(C.gtk_source_snippet_chunk_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSnippetChunk, F: marshalSnippetChunk},
	})
}

// SnippetChunkOverrides contains methods that are overridable.
type SnippetChunkOverrides struct {
}

func defaultSnippetChunkOverrides(v *SnippetChunk) SnippetChunkOverrides {
	return SnippetChunkOverrides{}
}

// SnippetChunk: chunk of text within the source snippet.
//
// The GtkSourceSnippetChunk represents a single chunk of text that may or may
// not be an edit point within the snippet. Chunks that are an edit point (also
// called a tab stop) have the snippetchunk:focus-position property set.
type SnippetChunk struct {
	_ [0]func() // equal guard
	coreglib.InitiallyUnowned
}

var ()

func init() {
	coreglib.RegisterClassInfo[*SnippetChunk, *SnippetChunkClass, SnippetChunkOverrides](
		GTypeSnippetChunk,
		initSnippetChunkClass,
		wrapSnippetChunk,
		defaultSnippetChunkOverrides,
	)
}

func initSnippetChunkClass(gclass unsafe.Pointer, overrides SnippetChunkOverrides, classInitFunc func(*SnippetChunkClass)) {
	if classInitFunc != nil {
		class := (*SnippetChunkClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapSnippetChunk(obj *coreglib.Object) *SnippetChunk {
	return &SnippetChunk{
		InitiallyUnowned: coreglib.InitiallyUnowned{
			Object: obj,
		},
	}
}

func marshalSnippetChunk(p uintptr) (interface{}, error) {
	return wrapSnippetChunk(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewSnippetChunk: create a new GtkSourceSnippetChunk that can be added to a
// snippet.
func NewSnippetChunk() *SnippetChunk {
	var _cret *C.GtkSourceSnippetChunk // in

	_cret = C.gtk_source_snippet_chunk_new()

	var _snippetChunk *SnippetChunk // out

	_snippetChunk = wrapSnippetChunk(coreglib.Take(unsafe.Pointer(_cret)))

	return _snippetChunk
}

// Copy copies the source snippet.
//
// The function returns the following values:
//
//   - snippetChunk: SourceSnippetChunk.
func (chunk *SnippetChunk) Copy() *SnippetChunk {
	var _arg0 *C.GtkSourceSnippetChunk // out
	var _cret *C.GtkSourceSnippetChunk // in

	_arg0 = (*C.GtkSourceSnippetChunk)(unsafe.Pointer(coreglib.InternObject(chunk).Native()))

	_cret = C.gtk_source_snippet_chunk_copy(_arg0)
	runtime.KeepAlive(chunk)

	var _snippetChunk *SnippetChunk // out

	_snippetChunk = wrapSnippetChunk(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _snippetChunk
}

// Context gets the context for the snippet insertion.
//
// The function returns the following values:
//
//   - snippetContext: SourceSnippetContext.
func (chunk *SnippetChunk) Context() *SnippetContext {
	var _arg0 *C.GtkSourceSnippetChunk   // out
	var _cret *C.GtkSourceSnippetContext // in

	_arg0 = (*C.GtkSourceSnippetChunk)(unsafe.Pointer(coreglib.InternObject(chunk).Native()))

	_cret = C.gtk_source_snippet_chunk_get_context(_arg0)
	runtime.KeepAlive(chunk)

	var _snippetContext *SnippetContext // out

	_snippetContext = wrapSnippetContext(coreglib.Take(unsafe.Pointer(_cret)))

	return _snippetContext
}

// FocusPosition gets the snippetchunk:focus-position.
//
// The focus-position is used to determine how many tabs it takes for the
// snippet to advanced to this chunk.
//
// A focus-position of zero will be the last focus position of the snippet and
// snippet editing ends when it has been reached.
//
// A focus-position of -1 means the chunk cannot be focused by the user.
//
// The function returns the following values:
//
//   - gint: focus-position.
func (chunk *SnippetChunk) FocusPosition() int {
	var _arg0 *C.GtkSourceSnippetChunk // out
	var _cret C.gint                   // in

	_arg0 = (*C.GtkSourceSnippetChunk)(unsafe.Pointer(coreglib.InternObject(chunk).Native()))

	_cret = C.gtk_source_snippet_chunk_get_focus_position(_arg0)
	runtime.KeepAlive(chunk)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Spec gets the specification for the chunk.
//
// The specification is evaluated for variables when other chunks are
// edited within the snippet context. If the user has changed the text,
// the snippetchunk:text and snippetchunk:text-set properties are updated.
//
// The function returns the following values:
//
//   - utf8 (optional): specification, if any.
func (chunk *SnippetChunk) Spec() string {
	var _arg0 *C.GtkSourceSnippetChunk // out
	var _cret *C.gchar                 // in

	_arg0 = (*C.GtkSourceSnippetChunk)(unsafe.Pointer(coreglib.InternObject(chunk).Native()))

	_cret = C.gtk_source_snippet_chunk_get_spec(_arg0)
	runtime.KeepAlive(chunk)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Text gets the snippetchunk:text property.
//
// The text property is updated when the user edits the text of the chunk.
// If it has not been edited, the snippetchunk:spec property is returned.
//
// The function returns the following values:
//
//   - utf8: text of the chunk.
func (chunk *SnippetChunk) Text() string {
	var _arg0 *C.GtkSourceSnippetChunk // out
	var _cret *C.gchar                 // in

	_arg0 = (*C.GtkSourceSnippetChunk)(unsafe.Pointer(coreglib.InternObject(chunk).Native()))

	_cret = C.gtk_source_snippet_chunk_get_text(_arg0)
	runtime.KeepAlive(chunk)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// TextSet gets the snippetchunk:text-set property.
//
// This is typically set when the user has edited a snippet chunk.
func (chunk *SnippetChunk) TextSet() bool {
	var _arg0 *C.GtkSourceSnippetChunk // out
	var _cret C.gboolean               // in

	_arg0 = (*C.GtkSourceSnippetChunk)(unsafe.Pointer(coreglib.InternObject(chunk).Native()))

	_cret = C.gtk_source_snippet_chunk_get_text_set(_arg0)
	runtime.KeepAlive(chunk)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (chunk *SnippetChunk) TooltipText() string {
	var _arg0 *C.GtkSourceSnippetChunk // out
	var _cret *C.char                  // in

	_arg0 = (*C.GtkSourceSnippetChunk)(unsafe.Pointer(coreglib.InternObject(chunk).Native()))

	_cret = C.gtk_source_snippet_chunk_get_tooltip_text(_arg0)
	runtime.KeepAlive(chunk)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

func (chunk *SnippetChunk) SetContext(context *SnippetContext) {
	var _arg0 *C.GtkSourceSnippetChunk   // out
	var _arg1 *C.GtkSourceSnippetContext // out

	_arg0 = (*C.GtkSourceSnippetChunk)(unsafe.Pointer(coreglib.InternObject(chunk).Native()))
	_arg1 = (*C.GtkSourceSnippetContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	C.gtk_source_snippet_chunk_set_context(_arg0, _arg1)
	runtime.KeepAlive(chunk)
	runtime.KeepAlive(context)
}

// SetFocusPosition sets the snippetchunk:focus-position property.
//
// The focus-position is used to determine how many tabs it takes for the
// snippet to advanced to this chunk.
//
// A focus-position of zero will be the last focus position of the snippet and
// snippet editing ends when it has been reached.
//
// A focus-position of -1 means the chunk cannot be focused by the user.
//
// The function takes the following parameters:
//
//   - focusPosition: focus-position.
func (chunk *SnippetChunk) SetFocusPosition(focusPosition int) {
	var _arg0 *C.GtkSourceSnippetChunk // out
	var _arg1 C.gint                   // out

	_arg0 = (*C.GtkSourceSnippetChunk)(unsafe.Pointer(coreglib.InternObject(chunk).Native()))
	_arg1 = C.gint(focusPosition)

	C.gtk_source_snippet_chunk_set_focus_position(_arg0, _arg1)
	runtime.KeepAlive(chunk)
	runtime.KeepAlive(focusPosition)
}

// SetSpec sets the specification for the chunk.
//
// The specification is evaluated for variables when other chunks are
// edited within the snippet context. If the user has changed the text,
// the snippetchunk:text and snippetchunk:text-set properties are updated.
//
// The function takes the following parameters:
//
//   - spec: new specification for the chunk.
func (chunk *SnippetChunk) SetSpec(spec string) {
	var _arg0 *C.GtkSourceSnippetChunk // out
	var _arg1 *C.gchar                 // out

	_arg0 = (*C.GtkSourceSnippetChunk)(unsafe.Pointer(coreglib.InternObject(chunk).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(spec)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_source_snippet_chunk_set_spec(_arg0, _arg1)
	runtime.KeepAlive(chunk)
	runtime.KeepAlive(spec)
}

// SetText sets the text for the snippet chunk.
//
// This is usually used by the snippet engine to update the text, but may be
// useful when creating custom snippets to avoid expansion of any specification.
//
// The function takes the following parameters:
//
//   - text of the property.
func (chunk *SnippetChunk) SetText(text string) {
	var _arg0 *C.GtkSourceSnippetChunk // out
	var _arg1 *C.gchar                 // out

	_arg0 = (*C.GtkSourceSnippetChunk)(unsafe.Pointer(coreglib.InternObject(chunk).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_source_snippet_chunk_set_text(_arg0, _arg1)
	runtime.KeepAlive(chunk)
	runtime.KeepAlive(text)
}

// SetTextSet sets the snippetchunk:text-set property.
//
// This is typically set when the user has edited a snippet chunk by the snippet
// engine.
//
// The function takes the following parameters:
//
//   - textSet: property value.
func (chunk *SnippetChunk) SetTextSet(textSet bool) {
	var _arg0 *C.GtkSourceSnippetChunk // out
	var _arg1 C.gboolean               // out

	_arg0 = (*C.GtkSourceSnippetChunk)(unsafe.Pointer(coreglib.InternObject(chunk).Native()))
	if textSet {
		_arg1 = C.TRUE
	}

	C.gtk_source_snippet_chunk_set_text_set(_arg0, _arg1)
	runtime.KeepAlive(chunk)
	runtime.KeepAlive(textSet)
}

func (chunk *SnippetChunk) SetTooltipText(tooltipText string) {
	var _arg0 *C.GtkSourceSnippetChunk // out
	var _arg1 *C.char                  // out

	_arg0 = (*C.GtkSourceSnippetChunk)(unsafe.Pointer(coreglib.InternObject(chunk).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(tooltipText)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_source_snippet_chunk_set_tooltip_text(_arg0, _arg1)
	runtime.KeepAlive(chunk)
	runtime.KeepAlive(tooltipText)
}

// SnippetChunkClass: instance of this type is always passed by reference.
type SnippetChunkClass struct {
	*snippetChunkClass
}

// snippetChunkClass is the struct that's finalized.
type snippetChunkClass struct {
	native *C.GtkSourceSnippetChunkClass
}
