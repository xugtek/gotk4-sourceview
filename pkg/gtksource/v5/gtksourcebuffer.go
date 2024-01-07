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
// extern void _gotk4_gtksource5_Buffer_ConnectSourceMarkUpdated(gpointer, GtkTextMark*, guintptr);
// extern void _gotk4_gtksource5_Buffer_ConnectHighlightUpdated(gpointer, GtkTextIter*, GtkTextIter*, guintptr);
// extern void _gotk4_gtksource5_Buffer_ConnectCursorMoved(gpointer, guintptr);
// extern void _gotk4_gtksource5_Buffer_ConnectBracketMatched(gpointer, GtkTextIter*, GtkSourceBracketMatchType, guintptr);
// extern void _gotk4_gtksource5_BufferClass_bracket_matched(GtkSourceBuffer*, GtkTextIter*, GtkSourceBracketMatchType);
// void _gotk4_gtksource5_Buffer_virtual_bracket_matched(void* fnptr, GtkSourceBuffer* arg0, GtkTextIter* arg1, GtkSourceBracketMatchType arg2) {
//   ((void (*)(GtkSourceBuffer*, GtkTextIter*, GtkSourceBracketMatchType))(fnptr))(arg0, arg1, arg2);
// };
import "C"

// GType values.
var (
	GTypeBuffer = coreglib.Type(C.gtk_source_buffer_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeBuffer, F: marshalBuffer},
	})
}

// BufferOverrides contains methods that are overridable.
type BufferOverrides struct {
	// The function takes the following parameters:
	//
	//   - iter
	//   - state
	//
	BracketMatched func(iter *gtk.TextIter, state BracketMatchType)
}

func defaultBufferOverrides(v *Buffer) BufferOverrides {
	return BufferOverrides{
		BracketMatched: v.bracketMatched,
	}
}

// Buffer subclass of gtk.TextBuffer.
//
// A GtkSourceBuffer object is the model for view widgets. It extends the
// gtk.TextBuffer class by adding features useful to display and edit source
// code such as syntax highlighting and bracket matching.
//
// To create a GtkSourceBuffer use gtksource.Buffer.New or
// gtksource.Buffer.NewWithLanguage. The second form is just a convenience
// function which allows you to initially set a language. You can also directly
// create a view and get its buffer with gtk.TextView.GetBuffer().
//
// The highlighting is enabled by default, but you can disable it with
// buffer.SetHighlightSyntax.
//
// Context Classes:
//
// It is possible to retrieve some information from the syntax highlighting
// engine. The default context classes that are applied to regions of a
// GtkSourceBuffer:
//
//   - **comment**: the region delimits a comment;
//   - **no-spell-check**: the region should not be spell checked;
//   - **path**: the region delimits a path to a file;
//   - **string**: the region delimits a string.
//
// Custom language definition files can create their own context classes,
// since the functions like buffer.IterHasContextClass take a string parameter
// as the context class.
//
// GtkSourceBuffer provides an API to access the context classes:
// buffer.IterHasContextClass, buffer.GetContextClassesAtIter,
// buffer.IterForwardToContextClassToggle and
// buffer.IterBackwardToContextClassToggle.
//
// And the gtksource.Buffer::highlight-updated signal permits to be notified
// when a context class region changes.
//
// Each context class has also an associated gtk.TextTag with the name
// gtksourceview:context-classes:<name>. For example to retrieve the gtk.TextTag
// for the string context class, one can write:
//
//    GtkTextTagTable *tag_table;
//    GtkTextTag *tag;
//
//    tag_table = gtk_text_buffer_get_tag_table (buffer);
//    tag = gtk_text_tag_table_lookup (tag_table, "gtksourceview:context-classes:string");
//
// The tag must be used for read-only purposes.
//
// Accessing a context class via the associated gtk.TextTag is less convenient
// than the GtkSourceBuffer API, because:
//
//   - The tag doesn't always exist, you need to listen to the
//     gtk.TextTagTable::tag-added and gtk.TextTagTable::tag-removed signals.
//   - Instead of the gtksource.Buffer::highlight-updated signal, you can listen
//     to the gtk.TextBuffer::apply-tag and gtk.TextBuffer::remove-tag signals.
//
// A possible use-case for accessing a context class via the associated
// gtk.TextTag is to read the region but without adding a hard dependency on the
// GtkSourceView library (for example for a spell-checking library that wants to
// read the no-spell-check region).
type Buffer struct {
	_ [0]func() // equal guard
	gtk.TextBuffer
}

var (
	_ coreglib.Objector = (*Buffer)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*Buffer, *BufferClass, BufferOverrides](
		GTypeBuffer,
		initBufferClass,
		wrapBuffer,
		defaultBufferOverrides,
	)
}

func initBufferClass(gclass unsafe.Pointer, overrides BufferOverrides, classInitFunc func(*BufferClass)) {
	pclass := (*C.GtkSourceBufferClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeBuffer))))

	if overrides.BracketMatched != nil {
		pclass.bracket_matched = (*[0]byte)(C._gotk4_gtksource5_BufferClass_bracket_matched)
	}

	if classInitFunc != nil {
		class := (*BufferClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapBuffer(obj *coreglib.Object) *Buffer {
	return &Buffer{
		TextBuffer: gtk.TextBuffer{
			Object: obj,
		},
	}
}

func marshalBuffer(p uintptr) (interface{}, error) {
	return wrapBuffer(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectBracketMatched: iter is set to a valid iterator pointing to the
// matching bracket if state is GTK_SOURCE_BRACKET_MATCH_FOUND. Otherwise iter
// is meaningless.
//
// The signal is emitted only when the state changes, typically when the cursor
// moves.
//
// A use-case for this signal is to show messages in a gtk.Statusbar.
func (buffer *Buffer) ConnectBracketMatched(f func(iter *gtk.TextIter, state BracketMatchType)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(buffer, "bracket-matched", false, unsafe.Pointer(C._gotk4_gtksource5_Buffer_ConnectBracketMatched), f)
}

// ConnectCursorMoved: "cursor-moved" signal is emitted when then insertion mark
// has moved.
func (buffer *Buffer) ConnectCursorMoved(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(buffer, "cursor-moved", false, unsafe.Pointer(C._gotk4_gtksource5_Buffer_ConnectCursorMoved), f)
}

// ConnectHighlightUpdated signal is emitted when the syntax highlighting
// and context classes (./class.Buffer.html#context-classes) are updated in a
// certain region of the buffer.
func (buffer *Buffer) ConnectHighlightUpdated(f func(start, end *gtk.TextIter)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(buffer, "highlight-updated", false, unsafe.Pointer(C._gotk4_gtksource5_Buffer_ConnectHighlightUpdated), f)
}

// ConnectSourceMarkUpdated signal is emitted each time a mark is added to,
// moved or removed from the buffer.
func (buffer *Buffer) ConnectSourceMarkUpdated(f func(mark *gtk.TextMark)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(buffer, "source-mark-updated", false, unsafe.Pointer(C._gotk4_gtksource5_Buffer_ConnectSourceMarkUpdated), f)
}

// NewBuffer creates a new source buffer.
//
// The function takes the following parameters:
//
//   - table (optional) or NULL to create a new one.
//
// The function returns the following values:
//
//   - buffer: new source buffer.
//
func NewBuffer(table *gtk.TextTagTable) *Buffer {
	var _arg1 *C.GtkTextTagTable // out
	var _cret *C.GtkSourceBuffer // in

	if table != nil {
		_arg1 = (*C.GtkTextTagTable)(unsafe.Pointer(coreglib.InternObject(table).Native()))
	}

	_cret = C.gtk_source_buffer_new(_arg1)
	runtime.KeepAlive(table)

	var _buffer *Buffer // out

	_buffer = wrapBuffer(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _buffer
}

// NewBufferWithLanguage creates a new source buffer using the highlighting
// patterns in language.
//
// This is equivalent to creating a new source buffer with a new tag table and
// then calling buffer.SetLanguage.
//
// The function takes the following parameters:
//
//   - language: SourceLanguage.
//
// The function returns the following values:
//
//   - buffer: new source buffer which will highlight text according to the
//     highlighting patterns in language.
//
func NewBufferWithLanguage(language *Language) *Buffer {
	var _arg1 *C.GtkSourceLanguage // out
	var _cret *C.GtkSourceBuffer   // in

	_arg1 = (*C.GtkSourceLanguage)(unsafe.Pointer(coreglib.InternObject(language).Native()))

	_cret = C.gtk_source_buffer_new_with_language(_arg1)
	runtime.KeepAlive(language)

	var _buffer *Buffer // out

	_buffer = wrapBuffer(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _buffer
}

// ChangeCase changes the case of the text between the specified iterators.
//
// Since 5.4, this function will update the position of start and end to
// surround the modified text.
//
// The function takes the following parameters:
//
//   - caseType: how to change the case.
//   - start: TextIter.
//   - end: TextIter.
//
func (buffer *Buffer) ChangeCase(caseType ChangeCaseType, start, end *gtk.TextIter) {
	var _arg0 *C.GtkSourceBuffer        // out
	var _arg1 C.GtkSourceChangeCaseType // out
	var _arg2 *C.GtkTextIter            // out
	var _arg3 *C.GtkTextIter            // out

	_arg0 = (*C.GtkSourceBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))
	_arg1 = C.GtkSourceChangeCaseType(caseType)
	_arg2 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(start)))
	_arg3 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(end)))

	C.gtk_source_buffer_change_case(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(caseType)
	runtime.KeepAlive(start)
	runtime.KeepAlive(end)
}

// CreateSourceMark creates a source mark in the buffer of category category.
//
// A source mark is a gtk.TextMark but organized into categories. Depending on
// the category a pixbuf can be specified that will be displayed along the line
// of the mark.
//
// Like a gtk.TextMark, a mark can be anonymous if the passed name is NULL.
// Also, the buffer owns the marks so you shouldn't unreference it.
//
// Marks always have left gravity and are moved to the beginning of the line
// when the user deletes the line they were in.
//
// Typical uses for a source mark are bookmarks, breakpoints, current executing
// instruction indication in a source file, etc..
//
// The function takes the following parameters:
//
//   - name (optional) of the mark, or NULL.
//   - category: string defining the mark category.
//   - where: location to place the mark.
//
// The function returns the following values:
//
//   - mark: new mark, owned by the buffer.
//
func (buffer *Buffer) CreateSourceMark(name, category string, where *gtk.TextIter) *Mark {
	var _arg0 *C.GtkSourceBuffer // out
	var _arg1 *C.gchar           // out
	var _arg2 *C.gchar           // out
	var _arg3 *C.GtkTextIter     // out
	var _cret *C.GtkSourceMark   // in

	_arg0 = (*C.GtkSourceBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))
	if name != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(category)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(where)))

	_cret = C.gtk_source_buffer_create_source_mark(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(name)
	runtime.KeepAlive(category)
	runtime.KeepAlive(where)

	var _mark *Mark // out

	_mark = wrapMark(coreglib.Take(unsafe.Pointer(_cret)))

	return _mark
}

// EnsureHighlight forces buffer to analyze and highlight the given area
// synchronously.
//
// **Note**:
//
// This is a potentially slow operation and should be used only when you need to
// make sure that some text not currently visible is highlighted, for instance
// before printing.
//
// The function takes the following parameters:
//
//   - start of the area to highlight.
//   - end of the area to highlight.
//
func (buffer *Buffer) EnsureHighlight(start, end *gtk.TextIter) {
	var _arg0 *C.GtkSourceBuffer // out
	var _arg1 *C.GtkTextIter     // out
	var _arg2 *C.GtkTextIter     // out

	_arg0 = (*C.GtkSourceBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))
	_arg1 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(start)))
	_arg2 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(end)))

	C.gtk_source_buffer_ensure_highlight(_arg0, _arg1, _arg2)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(start)
	runtime.KeepAlive(end)
}

// ContextClassesAtIter: get all defined context classes at iter.
//
// See the buffer description for the list of default context classes.
//
// The function takes the following parameters:
//
//   - iter: TextIter.
//
// The function returns the following values:
//
//   - utf8s: new NULL terminated array of context class names. Use g_strfreev()
//     to free the array if it is no longer needed.
//
func (buffer *Buffer) ContextClassesAtIter(iter *gtk.TextIter) []string {
	var _arg0 *C.GtkSourceBuffer // out
	var _arg1 *C.GtkTextIter     // out
	var _cret **C.gchar          // in

	_arg0 = (*C.GtkSourceBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))
	_arg1 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(iter)))

	_cret = C.gtk_source_buffer_get_context_classes_at_iter(_arg0, _arg1)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(iter)

	var _utf8s []string // out

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

	return _utf8s
}

// HighlightMatchingBrackets determines whether bracket match highlighting is
// activated for the source buffer.
//
// The function returns the following values:
//
//   - ok: TRUE if the source buffer will highlight matching brackets.
//
func (buffer *Buffer) HighlightMatchingBrackets() bool {
	var _arg0 *C.GtkSourceBuffer // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkSourceBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))

	_cret = C.gtk_source_buffer_get_highlight_matching_brackets(_arg0)
	runtime.KeepAlive(buffer)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HighlightSyntax determines whether syntax highlighting is activated in the
// source buffer.
//
// The function returns the following values:
//
//   - ok: TRUE if syntax highlighting is enabled, FALSE otherwise.
//
func (buffer *Buffer) HighlightSyntax() bool {
	var _arg0 *C.GtkSourceBuffer // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkSourceBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))

	_cret = C.gtk_source_buffer_get_highlight_syntax(_arg0)
	runtime.KeepAlive(buffer)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// The function returns the following values:
//
//   - ok: whether the buffer has an implicit trailing newline.
//
func (buffer *Buffer) ImplicitTrailingNewline() bool {
	var _arg0 *C.GtkSourceBuffer // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkSourceBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))

	_cret = C.gtk_source_buffer_get_implicit_trailing_newline(_arg0)
	runtime.KeepAlive(buffer)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Language returns the language associated with the buffer, see
// buffer.SetLanguage.
//
// The returned object should not be unreferenced by the user.
//
// The function returns the following values:
//
//   - language (optional): language associated with the buffer, or NULL.
//
func (buffer *Buffer) Language() *Language {
	var _arg0 *C.GtkSourceBuffer   // out
	var _cret *C.GtkSourceLanguage // in

	_arg0 = (*C.GtkSourceBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))

	_cret = C.gtk_source_buffer_get_language(_arg0)
	runtime.KeepAlive(buffer)

	var _language *Language // out

	if _cret != nil {
		_language = wrapLanguage(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _language
}

// SourceMarksAtIter returns the list of marks of the given category at iter.
//
// If category is NULL it returns all marks at iter.
//
// The function takes the following parameters:
//
//   - iter: iterator.
//   - category (optional) to search for, or NULL.
//
// The function returns the following values:
//
//   - sList: a newly allocated List.
//
func (buffer *Buffer) SourceMarksAtIter(iter *gtk.TextIter, category string) []*Mark {
	var _arg0 *C.GtkSourceBuffer // out
	var _arg1 *C.GtkTextIter     // out
	var _arg2 *C.gchar           // out
	var _cret *C.GSList          // in

	_arg0 = (*C.GtkSourceBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))
	_arg1 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(iter)))
	if category != "" {
		_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(category)))
		defer C.free(unsafe.Pointer(_arg2))
	}

	_cret = C.gtk_source_buffer_get_source_marks_at_iter(_arg0, _arg1, _arg2)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(iter)
	runtime.KeepAlive(category)

	var _sList []*Mark // out

	_sList = make([]*Mark, 0, gextras.SListSize(unsafe.Pointer(_cret)))
	gextras.MoveSList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GtkSourceMark)(v)
		var dst *Mark // out
		dst = wrapMark(coreglib.Take(unsafe.Pointer(src)))
		_sList = append(_sList, dst)
	})

	return _sList
}

// SourceMarksAtLine returns the list of marks of the given category at line.
//
// If category is NULL, all marks at line are returned.
//
// The function takes the following parameters:
//
//   - line number.
//   - category (optional) to search for, or NULL.
//
// The function returns the following values:
//
//   - sList: a newly allocated List.
//
func (buffer *Buffer) SourceMarksAtLine(line int, category string) []*Mark {
	var _arg0 *C.GtkSourceBuffer // out
	var _arg1 C.gint             // out
	var _arg2 *C.gchar           // out
	var _cret *C.GSList          // in

	_arg0 = (*C.GtkSourceBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))
	_arg1 = C.gint(line)
	if category != "" {
		_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(category)))
		defer C.free(unsafe.Pointer(_arg2))
	}

	_cret = C.gtk_source_buffer_get_source_marks_at_line(_arg0, _arg1, _arg2)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(line)
	runtime.KeepAlive(category)

	var _sList []*Mark // out

	_sList = make([]*Mark, 0, gextras.SListSize(unsafe.Pointer(_cret)))
	gextras.MoveSList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GtkSourceMark)(v)
		var dst *Mark // out
		dst = wrapMark(coreglib.Take(unsafe.Pointer(src)))
		_sList = append(_sList, dst)
	})

	return _sList
}

// StyleScheme returns the stylescheme associated with the buffer, see
// buffer.SetStyleScheme.
//
// The returned object should not be unreferenced by the user.
//
// The function returns the following values:
//
//   - styleScheme (optional): stylescheme associated with the buffer, or NULL.
//
func (buffer *Buffer) StyleScheme() *StyleScheme {
	var _arg0 *C.GtkSourceBuffer      // out
	var _cret *C.GtkSourceStyleScheme // in

	_arg0 = (*C.GtkSourceBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))

	_cret = C.gtk_source_buffer_get_style_scheme(_arg0)
	runtime.KeepAlive(buffer)

	var _styleScheme *StyleScheme // out

	if _cret != nil {
		_styleScheme = wrapStyleScheme(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _styleScheme
}

// IterHasContextClass: check if the class context_class is set on iter.
//
// See the buffer description for the list of default context classes.
//
// The function takes the following parameters:
//
//   - iter: TextIter.
//   - contextClass class to search for.
//
// The function returns the following values:
//
//   - ok: whether iter has the context class.
//
func (buffer *Buffer) IterHasContextClass(iter *gtk.TextIter, contextClass string) bool {
	var _arg0 *C.GtkSourceBuffer // out
	var _arg1 *C.GtkTextIter     // out
	var _arg2 *C.gchar           // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkSourceBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))
	_arg1 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(iter)))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(contextClass)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_source_buffer_iter_has_context_class(_arg0, _arg1, _arg2)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(iter)
	runtime.KeepAlive(contextClass)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// JoinLines joins the lines of text between the specified iterators.
//
// The function takes the following parameters:
//
//   - start: TextIter.
//   - end: TextIter.
//
func (buffer *Buffer) JoinLines(start, end *gtk.TextIter) {
	var _arg0 *C.GtkSourceBuffer // out
	var _arg1 *C.GtkTextIter     // out
	var _arg2 *C.GtkTextIter     // out

	_arg0 = (*C.GtkSourceBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))
	_arg1 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(start)))
	_arg2 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(end)))

	C.gtk_source_buffer_join_lines(_arg0, _arg1, _arg2)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(start)
	runtime.KeepAlive(end)
}

// RemoveSourceMarks: remove all marks of category between start and end from
// the buffer.
//
// If category is NULL, all marks in the range will be removed.
//
// The function takes the following parameters:
//
//   - start: TextIter.
//   - end: TextIter.
//   - category (optional) to search for, or NULL.
//
func (buffer *Buffer) RemoveSourceMarks(start, end *gtk.TextIter, category string) {
	var _arg0 *C.GtkSourceBuffer // out
	var _arg1 *C.GtkTextIter     // out
	var _arg2 *C.GtkTextIter     // out
	var _arg3 *C.gchar           // out

	_arg0 = (*C.GtkSourceBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))
	_arg1 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(start)))
	_arg2 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(end)))
	if category != "" {
		_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(category)))
		defer C.free(unsafe.Pointer(_arg3))
	}

	C.gtk_source_buffer_remove_source_marks(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(start)
	runtime.KeepAlive(end)
	runtime.KeepAlive(category)
}

// SetHighlightMatchingBrackets controls the bracket match highlighting function
// in the buffer.
//
// If activated, when you position your cursor over a bracket character (a
// parenthesis, a square bracket, etc.) the matching opening or closing bracket
// character will be highlighted.
//
// The function takes the following parameters:
//
//   - highlight: TRUE if you want matching brackets highlighted.
//
func (buffer *Buffer) SetHighlightMatchingBrackets(highlight bool) {
	var _arg0 *C.GtkSourceBuffer // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkSourceBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))
	if highlight {
		_arg1 = C.TRUE
	}

	C.gtk_source_buffer_set_highlight_matching_brackets(_arg0, _arg1)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(highlight)
}

// SetHighlightSyntax controls whether syntax is highlighted in the buffer.
//
// If highlight is TRUE, the text will be highlighted according to the syntax
// patterns specified in the language set with buffer.SetLanguage.
//
// If highlight is FALSE, syntax highlighting is disabled and all the
// gtk.TextTag objects that have been added by the syntax highlighting engine
// are removed from the buffer.
//
// The function takes the following parameters:
//
//   - highlight: TRUE to enable syntax highlighting, FALSE to disable it.
//
func (buffer *Buffer) SetHighlightSyntax(highlight bool) {
	var _arg0 *C.GtkSourceBuffer // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkSourceBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))
	if highlight {
		_arg1 = C.TRUE
	}

	C.gtk_source_buffer_set_highlight_syntax(_arg0, _arg1)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(highlight)
}

// SetImplicitTrailingNewline sets whether the buffer has an implicit trailing
// newline.
//
// If an explicit trailing newline is present in a gtk.TextBuffer, gtk.TextView
// shows it as an empty line. This is generally not what the user expects.
//
// If implicit_trailing_newline is TRUE (the default value): - when a fileloader
// loads the content of a file into the buffer, the trailing newline (if present
// in the file) is not inserted into the buffer. - when a filesaver saves the
// content of the buffer into a file, a trailing newline is added to the file.
//
// On the other hand, if implicit_trailing_newline is FALSE, the file's content
// is not modified when loaded into the buffer, and the buffer's content is not
// modified when saved into a file.
//
// The function takes the following parameters:
//
//   - implicitTrailingNewline: new value.
//
func (buffer *Buffer) SetImplicitTrailingNewline(implicitTrailingNewline bool) {
	var _arg0 *C.GtkSourceBuffer // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkSourceBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))
	if implicitTrailingNewline {
		_arg1 = C.TRUE
	}

	C.gtk_source_buffer_set_implicit_trailing_newline(_arg0, _arg1)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(implicitTrailingNewline)
}

// SetLanguage associates a language with the buffer.
//
// Note that a language affects not only the syntax highlighting, but also the
// context classes (./class.Buffer.html#context-classes). If you want to disable
// just the syntax highlighting, see buffer.SetHighlightSyntax.
//
// The buffer holds a reference to language.
//
// The function takes the following parameters:
//
//   - language (optional) to set, or NULL.
//
func (buffer *Buffer) SetLanguage(language *Language) {
	var _arg0 *C.GtkSourceBuffer   // out
	var _arg1 *C.GtkSourceLanguage // out

	_arg0 = (*C.GtkSourceBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))
	if language != nil {
		_arg1 = (*C.GtkSourceLanguage)(unsafe.Pointer(coreglib.InternObject(language).Native()))
	}

	C.gtk_source_buffer_set_language(_arg0, _arg1)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(language)
}

// SetStyleScheme sets a stylescheme to be used by the buffer and the view.
//
// Note that a stylescheme affects not only the syntax highlighting, but also
// other view features such as highlighting the current line, matching brackets,
// the line numbers, etc.
//
// Instead of setting a NULL scheme, it is better to disable syntax highlighting
// with buffer.SetHighlightSyntax, and setting the stylescheme with the
// "classic" or "tango" ID, because those two style schemes follow more closely
// the GTK theme (for example for the background color).
//
// The buffer holds a reference to scheme.
//
// The function takes the following parameters:
//
//   - scheme (optional) or NULL.
//
func (buffer *Buffer) SetStyleScheme(scheme *StyleScheme) {
	var _arg0 *C.GtkSourceBuffer      // out
	var _arg1 *C.GtkSourceStyleScheme // out

	_arg0 = (*C.GtkSourceBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))
	if scheme != nil {
		_arg1 = (*C.GtkSourceStyleScheme)(unsafe.Pointer(coreglib.InternObject(scheme).Native()))
	}

	C.gtk_source_buffer_set_style_scheme(_arg0, _arg1)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(scheme)
}

// SortLines: sort the lines of text between the specified iterators.
//
// The function takes the following parameters:
//
//   - start: TextIter.
//   - end: TextIter.
//   - flags specifying how the sort should behave.
//   - column: sort considering the text starting at the given column.
//
func (buffer *Buffer) SortLines(start, end *gtk.TextIter, flags SortFlags, column int) {
	var _arg0 *C.GtkSourceBuffer   // out
	var _arg1 *C.GtkTextIter       // out
	var _arg2 *C.GtkTextIter       // out
	var _arg3 C.GtkSourceSortFlags // out
	var _arg4 C.gint               // out

	_arg0 = (*C.GtkSourceBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))
	_arg1 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(start)))
	_arg2 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(end)))
	_arg3 = C.GtkSourceSortFlags(flags)
	_arg4 = C.gint(column)

	C.gtk_source_buffer_sort_lines(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(start)
	runtime.KeepAlive(end)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(column)
}

// The function takes the following parameters:
//
//   - iter
//   - state
//
func (buffer *Buffer) bracketMatched(iter *gtk.TextIter, state BracketMatchType) {
	gclass := (*C.GtkSourceBufferClass)(coreglib.PeekParentClass(buffer))
	fnarg := gclass.bracket_matched

	var _arg0 *C.GtkSourceBuffer          // out
	var _arg1 *C.GtkTextIter              // out
	var _arg2 C.GtkSourceBracketMatchType // out

	_arg0 = (*C.GtkSourceBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))
	_arg1 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(iter)))
	_arg2 = C.GtkSourceBracketMatchType(state)

	C._gotk4_gtksource5_Buffer_virtual_bracket_matched(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(iter)
	runtime.KeepAlive(state)
}

// BufferClass: instance of this type is always passed by reference.
type BufferClass struct {
	*bufferClass
}

// bufferClass is the struct that's finalized.
type bufferClass struct {
	native *C.GtkSourceBufferClass
}

func (b *BufferClass) ParentClass() *gtk.TextBufferClass {
	valptr := &b.native.parent_class
	var _v *gtk.TextBufferClass // out
	_v = (*gtk.TextBufferClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
