// Code generated by girgen. DO NOT EDIT.

package gtksource

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtksourceview/gtksource.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_source_view_get_type()), F: marshalViewer},
	})
}

// ViewOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ViewOverrider interface {
	// The function takes the following parameters:
	//
	//    - iter
	//    - button
	//    - state
	//    - nPresses
	//
	LineMarkActivated(iter *gtk.TextIter, button uint, state gdk.ModifierType, nPresses int)
	// The function takes the following parameters:
	//
	MoveLines(down bool)
	// The function takes the following parameters:
	//
	MoveWords(step int)
	// PushSnippet inserts a new snippet at location
	//
	// If another snippet was already active, it will be paused and the new
	// snippet will become active. Once the focus positions of snippet have been
	// exhausted, editing will return to the previous snippet.
	//
	// The function takes the following parameters:
	//
	//    - snippet: SourceSnippet.
	//    - location (optional) or NULL for the cursor position.
	//
	PushSnippet(snippet *Snippet, location *gtk.TextIter)
	ShowCompletion()
}

type View struct {
	_ [0]func() // equal guard
	gtk.TextView
}

var (
	_ gtk.Widgetter       = (*View)(nil)
	_ externglib.Objector = (*View)(nil)
)

func wrapView(obj *externglib.Object) *View {
	return &View{
		TextView: gtk.TextView{
			Widget: gtk.Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				Object: obj,
				Accessible: gtk.Accessible{
					Object: obj,
				},
				Buildable: gtk.Buildable{
					Object: obj,
				},
				ConstraintTarget: gtk.ConstraintTarget{
					Object: obj,
				},
			},
			Object: obj,
			Scrollable: gtk.Scrollable{
				Object: obj,
			},
		},
	}
}

func marshalViewer(p uintptr) (interface{}, error) {
	return wrapView(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectChangeCase: keybinding signal to change case of the text at the
// current cursor position.
func (view *View) ConnectChangeCase(f func(caseType ChangeCaseType)) externglib.SignalHandle {
	return view.Connect("change-case", f)
}

// ConnectChangeNumber: keybinding signal to edit a number at the current cursor
// position.
func (view *View) ConnectChangeNumber(f func(count int)) externglib.SignalHandle {
	return view.Connect("change-number", f)
}

// ConnectJoinLines: keybinding signal to join the lines currently selected.
func (view *View) ConnectJoinLines(f func()) externglib.SignalHandle {
	return view.Connect("join-lines", f)
}

// ConnectLineMarkActivated: emitted when a line mark has been activated (for
// instance when there was a button press in the line marks gutter). You can use
// iter to determine on which line the activation took place.
func (view *View) ConnectLineMarkActivated(f func(iter *gtk.TextIter, button uint, state gdk.ModifierType, nPresses int)) externglib.SignalHandle {
	return view.Connect("line-mark-activated", f)
}

// ConnectMoveLines signal is a keybinding which gets emitted when the user
// initiates moving a line. The default binding key is Alt+Up/Down arrow. And
// moves the currently selected lines, or the current line up or down by one
// line.
func (view *View) ConnectMoveLines(f func(down bool)) externglib.SignalHandle {
	return view.Connect("move-lines", f)
}

// ConnectMoveToMatchingBracket: keybinding signal to move the cursor to the
// matching bracket.
func (view *View) ConnectMoveToMatchingBracket(f func(extendSelection bool)) externglib.SignalHandle {
	return view.Connect("move-to-matching-bracket", f)
}

// ConnectMoveWords signal is a keybinding which gets emitted when the user
// initiates moving a word. The default binding key is Alt+Left/Right Arrow and
// moves the current selection, or the current word by one word.
func (view *View) ConnectMoveWords(f func(count int)) externglib.SignalHandle {
	return view.Connect("move-words", f)
}

// ConnectShowCompletion signal is a key binding signal which gets emitted when
// the user requests a completion, by pressing
// <keycombo><keycap>Control</keycap><keycap>space</keycap></keycombo>.
//
// This will create a SourceCompletionContext with the activation type as
// GTK_SOURCE_COMPLETION_ACTIVATION_USER_REQUESTED.
//
// Applications should not connect to it, but may emit it with
// g_signal_emit_by_name() if they need to activate the completion by another
// means, for example with another key binding or a menu entry.
func (view *View) ConnectShowCompletion(f func()) externglib.SignalHandle {
	return view.Connect("show-completion", f)
}

// ConnectSmartHomeEnd: emitted when a the cursor was moved according to the
// smart home end setting. The signal is emitted after the cursor is moved, but
// during the GtkTextView::move-cursor action. This can be used to find out
// whether the cursor was moved by a normal home/end or by a smart home/end.
func (view *View) ConnectSmartHomeEnd(f func(iter *gtk.TextIter, count int)) externglib.SignalHandle {
	return view.Connect("smart-home-end", f)
}

// NewView creates a new SourceView.
//
// By default, an empty SourceBuffer will be lazily created and can be retrieved
// with gtk_text_view_get_buffer().
//
// If you want to specify your own buffer, either override the TextViewClass
// create_buffer factory method, or use gtk_source_view_new_with_buffer().
//
// The function returns the following values:
//
//    - view: new SourceView.
//
func NewView() *View {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_source_view_new()

	var _view *View // out

	_view = wrapView(externglib.Take(unsafe.Pointer(_cret)))

	return _view
}

// NewViewWithBuffer creates a new SourceView widget displaying the buffer
// buffer. One buffer can be shared among many widgets.
//
// The function takes the following parameters:
//
//    - buffer: SourceBuffer.
//
// The function returns the following values:
//
//    - view: new SourceView.
//
func NewViewWithBuffer(buffer *Buffer) *View {
	var _arg1 *C.GtkSourceBuffer // out
	var _cret *C.GtkWidget       // in

	_arg1 = (*C.GtkSourceBuffer)(unsafe.Pointer(buffer.Native()))

	_cret = C.gtk_source_view_new_with_buffer(_arg1)
	runtime.KeepAlive(buffer)

	var _view *View // out

	_view = wrapView(externglib.Take(unsafe.Pointer(_cret)))

	return _view
}

// AutoIndent returns whether auto-indentation of text is enabled.
//
// The function returns the following values:
//
//    - ok: TRUE if auto indentation is enabled.
//
func (view *View) AutoIndent() bool {
	var _arg0 *C.GtkSourceView // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkSourceView)(unsafe.Pointer(view.Native()))

	_cret = C.gtk_source_view_get_auto_indent(_arg0)
	runtime.KeepAlive(view)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// BackgroundPattern returns the SourceBackgroundPatternType specifying if and
// how the background pattern should be displayed for this view.
//
// The function returns the following values:
//
//    - backgroundPatternType: SourceBackgroundPatternType.
//
func (view *View) BackgroundPattern() BackgroundPatternType {
	var _arg0 *C.GtkSourceView                 // out
	var _cret C.GtkSourceBackgroundPatternType // in

	_arg0 = (*C.GtkSourceView)(unsafe.Pointer(view.Native()))

	_cret = C.gtk_source_view_get_background_pattern(_arg0)
	runtime.KeepAlive(view)

	var _backgroundPatternType BackgroundPatternType // out

	_backgroundPatternType = BackgroundPatternType(_cret)

	return _backgroundPatternType
}

// Completion gets the SourceCompletion associated with view. The returned
// object is guaranteed to be the same for the lifetime of view. Each SourceView
// object has a different SourceCompletion.
//
// The function returns the following values:
//
//    - completion associated with view.
//
func (view *View) Completion() *Completion {
	var _arg0 *C.GtkSourceView       // out
	var _cret *C.GtkSourceCompletion // in

	_arg0 = (*C.GtkSourceView)(unsafe.Pointer(view.Native()))

	_cret = C.gtk_source_view_get_completion(_arg0)
	runtime.KeepAlive(view)

	var _completion *Completion // out

	_completion = wrapCompletion(externglib.Take(unsafe.Pointer(_cret)))

	return _completion
}

// EnableSnippets gets the SourceView:enable-snippets property.
//
// If TRUE, matching snippets found in the SourceSnippetManager may be expanded
// when the user presses Tab after a word in the SourceView.
//
// The function returns the following values:
//
//    - ok: TRUE if enabled.
//
func (view *View) EnableSnippets() bool {
	var _arg0 *C.GtkSourceView // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkSourceView)(unsafe.Pointer(view.Native()))

	_cret = C.gtk_source_view_get_enable_snippets(_arg0)
	runtime.KeepAlive(view)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Gutter returns the SourceGutter object associated with window_type for view.
// Only GTK_TEXT_WINDOW_LEFT and GTK_TEXT_WINDOW_RIGHT are supported,
// respectively corresponding to the left and right gutter. The line numbers and
// mark category icons are rendered in the left gutter.
//
// The function takes the following parameters:
//
//    - windowType: gutter window type.
//
// The function returns the following values:
//
//    - gutter: SourceGutter.
//
func (view *View) Gutter(windowType gtk.TextWindowType) *Gutter {
	var _arg0 *C.GtkSourceView    // out
	var _arg1 C.GtkTextWindowType // out
	var _cret *C.GtkSourceGutter  // in

	_arg0 = (*C.GtkSourceView)(unsafe.Pointer(view.Native()))
	_arg1 = C.GtkTextWindowType(windowType)

	_cret = C.gtk_source_view_get_gutter(_arg0, _arg1)
	runtime.KeepAlive(view)
	runtime.KeepAlive(windowType)

	var _gutter *Gutter // out

	_gutter = wrapGutter(externglib.Take(unsafe.Pointer(_cret)))

	return _gutter
}

// HighlightCurrentLine returns whether the current line is highlighted.
//
// The function returns the following values:
//
//    - ok: TRUE if the current line is highlighted.
//
func (view *View) HighlightCurrentLine() bool {
	var _arg0 *C.GtkSourceView // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkSourceView)(unsafe.Pointer(view.Native()))

	_cret = C.gtk_source_view_get_highlight_current_line(_arg0)
	runtime.KeepAlive(view)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Hover gets the SourceHover associated with view. The returned object is
// guaranteed to be the same for the lifetime of view. Each SourceView object
// has a different SourceHover.
//
// The function returns the following values:
//
//    - hover associated with view.
//
func (view *View) Hover() *Hover {
	var _arg0 *C.GtkSourceView  // out
	var _cret *C.GtkSourceHover // in

	_arg0 = (*C.GtkSourceView)(unsafe.Pointer(view.Native()))

	_cret = C.gtk_source_view_get_hover(_arg0)
	runtime.KeepAlive(view)

	var _hover *Hover // out

	_hover = wrapHover(externglib.Take(unsafe.Pointer(_cret)))

	return _hover
}

// IndentOnTab returns whether when the tab key is pressed the current selection
// should get indented instead of replaced with the \t character.
//
// The function returns the following values:
//
//    - ok: TRUE if the selection is indented when tab is pressed.
//
func (view *View) IndentOnTab() bool {
	var _arg0 *C.GtkSourceView // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkSourceView)(unsafe.Pointer(view.Native()))

	_cret = C.gtk_source_view_get_indent_on_tab(_arg0)
	runtime.KeepAlive(view)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IndentWidth returns the number of spaces to use for each step of indent. See
// gtk_source_view_set_indent_width() for details.
//
// The function returns the following values:
//
//    - gint: indent width.
//
func (view *View) IndentWidth() int {
	var _arg0 *C.GtkSourceView // out
	var _cret C.gint           // in

	_arg0 = (*C.GtkSourceView)(unsafe.Pointer(view.Native()))

	_cret = C.gtk_source_view_get_indent_width(_arg0)
	runtime.KeepAlive(view)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Indenter gets the SourceView:indenter property.
//
// The function returns the following values:
//
//    - indenter (optional) or NULL.
//
func (view *View) Indenter() Indenterer {
	var _arg0 *C.GtkSourceView     // out
	var _cret *C.GtkSourceIndenter // in

	_arg0 = (*C.GtkSourceView)(unsafe.Pointer(view.Native()))

	_cret = C.gtk_source_view_get_indenter(_arg0)
	runtime.KeepAlive(view)

	var _indenter Indenterer // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(Indenterer)
				return ok
			})
			rv, ok := casted.(Indenterer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtksource.Indenterer")
			}
			_indenter = rv
		}
	}

	return _indenter
}

// InsertSpacesInsteadOfTabs returns whether when inserting a tabulator
// character it should be replaced by a group of space characters.
//
// The function returns the following values:
//
//    - ok: TRUE if spaces are inserted instead of tabs.
//
func (view *View) InsertSpacesInsteadOfTabs() bool {
	var _arg0 *C.GtkSourceView // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkSourceView)(unsafe.Pointer(view.Native()))

	_cret = C.gtk_source_view_get_insert_spaces_instead_of_tabs(_arg0)
	runtime.KeepAlive(view)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MarkAttributes gets attributes and priority for the category.
//
// The function takes the following parameters:
//
//    - category: category.
//    - priority: place where priority of the category will be stored.
//
// The function returns the following values:
//
//    - markAttributes for the category. The object belongs to view, so it must
//      not be unreffed.
//
func (view *View) MarkAttributes(category string, priority *int) *MarkAttributes {
	var _arg0 *C.GtkSourceView           // out
	var _arg1 *C.gchar                   // out
	var _arg2 *C.gint                    // out
	var _cret *C.GtkSourceMarkAttributes // in

	_arg0 = (*C.GtkSourceView)(unsafe.Pointer(view.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(category)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gint)(unsafe.Pointer(priority))

	_cret = C.gtk_source_view_get_mark_attributes(_arg0, _arg1, _arg2)
	runtime.KeepAlive(view)
	runtime.KeepAlive(category)
	runtime.KeepAlive(priority)

	var _markAttributes *MarkAttributes // out

	_markAttributes = wrapMarkAttributes(externglib.Take(unsafe.Pointer(_cret)))

	return _markAttributes
}

// RightMarginPosition gets the position of the right margin in the given view.
//
// The function returns the following values:
//
//    - guint: position of the right margin.
//
func (view *View) RightMarginPosition() uint {
	var _arg0 *C.GtkSourceView // out
	var _cret C.guint          // in

	_arg0 = (*C.GtkSourceView)(unsafe.Pointer(view.Native()))

	_cret = C.gtk_source_view_get_right_margin_position(_arg0)
	runtime.KeepAlive(view)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// ShowLineMarks returns whether line marks are displayed beside the text.
//
// The function returns the following values:
//
//    - ok: TRUE if the line marks are displayed.
//
func (view *View) ShowLineMarks() bool {
	var _arg0 *C.GtkSourceView // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkSourceView)(unsafe.Pointer(view.Native()))

	_cret = C.gtk_source_view_get_show_line_marks(_arg0)
	runtime.KeepAlive(view)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowLineNumbers returns whether line numbers are displayed beside the text.
//
// The function returns the following values:
//
//    - ok: TRUE if the line numbers are displayed.
//
func (view *View) ShowLineNumbers() bool {
	var _arg0 *C.GtkSourceView // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkSourceView)(unsafe.Pointer(view.Native()))

	_cret = C.gtk_source_view_get_show_line_numbers(_arg0)
	runtime.KeepAlive(view)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowRightMargin returns whether a right margin is displayed.
//
// The function returns the following values:
//
//    - ok: TRUE if the right margin is shown.
//
func (view *View) ShowRightMargin() bool {
	var _arg0 *C.GtkSourceView // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkSourceView)(unsafe.Pointer(view.Native()))

	_cret = C.gtk_source_view_get_show_right_margin(_arg0)
	runtime.KeepAlive(view)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SmartBackspace returns TRUE if pressing the Backspace key will try to delete
// spaces up to the previous tab stop.
//
// The function returns the following values:
//
//    - ok: TRUE if smart Backspace handling is enabled.
//
func (view *View) SmartBackspace() bool {
	var _arg0 *C.GtkSourceView // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkSourceView)(unsafe.Pointer(view.Native()))

	_cret = C.gtk_source_view_get_smart_backspace(_arg0)
	runtime.KeepAlive(view)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SmartHomeEnd returns a SourceSmartHomeEndType end value specifying how the
// cursor will move when HOME and END keys are pressed.
//
// The function returns the following values:
//
//    - smartHomeEndType: SourceSmartHomeEndType value.
//
func (view *View) SmartHomeEnd() SmartHomeEndType {
	var _arg0 *C.GtkSourceView            // out
	var _cret C.GtkSourceSmartHomeEndType // in

	_arg0 = (*C.GtkSourceView)(unsafe.Pointer(view.Native()))

	_cret = C.gtk_source_view_get_smart_home_end(_arg0)
	runtime.KeepAlive(view)

	var _smartHomeEndType SmartHomeEndType // out

	_smartHomeEndType = SmartHomeEndType(_cret)

	return _smartHomeEndType
}

// SpaceDrawer gets the SourceSpaceDrawer associated with view. The returned
// object is guaranteed to be the same for the lifetime of view. Each SourceView
// object has a different SourceSpaceDrawer.
//
// The function returns the following values:
//
//    - spaceDrawer associated with view.
//
func (view *View) SpaceDrawer() *SpaceDrawer {
	var _arg0 *C.GtkSourceView        // out
	var _cret *C.GtkSourceSpaceDrawer // in

	_arg0 = (*C.GtkSourceView)(unsafe.Pointer(view.Native()))

	_cret = C.gtk_source_view_get_space_drawer(_arg0)
	runtime.KeepAlive(view)

	var _spaceDrawer *SpaceDrawer // out

	_spaceDrawer = wrapSpaceDrawer(externglib.Take(unsafe.Pointer(_cret)))

	return _spaceDrawer
}

// TabWidth returns the width of tabulation in characters.
//
// The function returns the following values:
//
//    - guint: width of tab.
//
func (view *View) TabWidth() uint {
	var _arg0 *C.GtkSourceView // out
	var _cret C.guint          // in

	_arg0 = (*C.GtkSourceView)(unsafe.Pointer(view.Native()))

	_cret = C.gtk_source_view_get_tab_width(_arg0)
	runtime.KeepAlive(view)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// VisualColumn determines the visual column at iter taking into consideration
// the SourceView:tab-width of view.
//
// The function takes the following parameters:
//
//    - iter: position in view.
//
// The function returns the following values:
//
//    - guint: visual column at iter.
//
func (view *View) VisualColumn(iter *gtk.TextIter) uint {
	var _arg0 *C.GtkSourceView // out
	var _arg1 *C.GtkTextIter   // out
	var _cret C.guint          // in

	_arg0 = (*C.GtkSourceView)(unsafe.Pointer(view.Native()))
	_arg1 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(iter)))

	_cret = C.gtk_source_view_get_visual_column(_arg0, _arg1)
	runtime.KeepAlive(view)
	runtime.KeepAlive(iter)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// IndentLines inserts one indentation level at the beginning of the specified
// lines. The empty lines are not indented.
//
// The function takes the following parameters:
//
//    - start of the first line to indent.
//    - end of the last line to indent.
//
func (view *View) IndentLines(start, end *gtk.TextIter) {
	var _arg0 *C.GtkSourceView // out
	var _arg1 *C.GtkTextIter   // out
	var _arg2 *C.GtkTextIter   // out

	_arg0 = (*C.GtkSourceView)(unsafe.Pointer(view.Native()))
	_arg1 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(start)))
	_arg2 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(end)))

	C.gtk_source_view_indent_lines(_arg0, _arg1, _arg2)
	runtime.KeepAlive(view)
	runtime.KeepAlive(start)
	runtime.KeepAlive(end)
}

// PushSnippet inserts a new snippet at location
//
// If another snippet was already active, it will be paused and the new snippet
// will become active. Once the focus positions of snippet have been exhausted,
// editing will return to the previous snippet.
//
// The function takes the following parameters:
//
//    - snippet: SourceSnippet.
//    - location (optional) or NULL for the cursor position.
//
func (view *View) PushSnippet(snippet *Snippet, location *gtk.TextIter) {
	var _arg0 *C.GtkSourceView    // out
	var _arg1 *C.GtkSourceSnippet // out
	var _arg2 *C.GtkTextIter      // out

	_arg0 = (*C.GtkSourceView)(unsafe.Pointer(view.Native()))
	_arg1 = (*C.GtkSourceSnippet)(unsafe.Pointer(snippet.Native()))
	if location != nil {
		_arg2 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(location)))
	}

	C.gtk_source_view_push_snippet(_arg0, _arg1, _arg2)
	runtime.KeepAlive(view)
	runtime.KeepAlive(snippet)
	runtime.KeepAlive(location)
}

// SetAutoIndent: if TRUE auto-indentation of text is enabled.
//
// When Enter is pressed to create a new line, the auto-indentation inserts the
// same indentation as the previous line. This is <emphasis>not</emphasis> a
// "smart indentation" where an indentation level is added or removed depending
// on the context.
//
// The function takes the following parameters:
//
//    - enable: whether to enable auto indentation.
//
func (view *View) SetAutoIndent(enable bool) {
	var _arg0 *C.GtkSourceView // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkSourceView)(unsafe.Pointer(view.Native()))
	if enable {
		_arg1 = C.TRUE
	}

	C.gtk_source_view_set_auto_indent(_arg0, _arg1)
	runtime.KeepAlive(view)
	runtime.KeepAlive(enable)
}

// SetBackgroundPattern: set if and how the background pattern should be
// displayed.
//
// The function takes the following parameters:
//
//    - backgroundPattern: SourceBackgroundPatternType.
//
func (view *View) SetBackgroundPattern(backgroundPattern BackgroundPatternType) {
	var _arg0 *C.GtkSourceView                 // out
	var _arg1 C.GtkSourceBackgroundPatternType // out

	_arg0 = (*C.GtkSourceView)(unsafe.Pointer(view.Native()))
	_arg1 = C.GtkSourceBackgroundPatternType(backgroundPattern)

	C.gtk_source_view_set_background_pattern(_arg0, _arg1)
	runtime.KeepAlive(view)
	runtime.KeepAlive(backgroundPattern)
}

// SetEnableSnippets sets the SourceView:enable-snippets property.
//
// If enable_snippets is TRUE, matching snippets found in the
// SourceSnippetManager may be expanded when the user presses Tab after a word
// in the SourceView.
//
// The function takes the following parameters:
//
//    - enableSnippets: if snippets should be enabled.
//
func (view *View) SetEnableSnippets(enableSnippets bool) {
	var _arg0 *C.GtkSourceView // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkSourceView)(unsafe.Pointer(view.Native()))
	if enableSnippets {
		_arg1 = C.TRUE
	}

	C.gtk_source_view_set_enable_snippets(_arg0, _arg1)
	runtime.KeepAlive(view)
	runtime.KeepAlive(enableSnippets)
}

// SetHighlightCurrentLine: if highlight is TRUE the current line will be
// highlighted.
//
// The function takes the following parameters:
//
//    - highlight: whether to highlight the current line.
//
func (view *View) SetHighlightCurrentLine(highlight bool) {
	var _arg0 *C.GtkSourceView // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkSourceView)(unsafe.Pointer(view.Native()))
	if highlight {
		_arg1 = C.TRUE
	}

	C.gtk_source_view_set_highlight_current_line(_arg0, _arg1)
	runtime.KeepAlive(view)
	runtime.KeepAlive(highlight)
}

// SetIndentOnTab: if TRUE, when the tab key is pressed when several lines are
// selected, the selected lines are indented of one level instead of being
// replaced with a \t character. Shift+Tab unindents the selection.
//
// If the first or last line is not selected completely, it is also indented or
// unindented.
//
// When the selection doesn't span several lines, the tab key always replaces
// the selection with a normal \t character.
//
// The function takes the following parameters:
//
//    - enable: whether to indent a block when tab is pressed.
//
func (view *View) SetIndentOnTab(enable bool) {
	var _arg0 *C.GtkSourceView // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkSourceView)(unsafe.Pointer(view.Native()))
	if enable {
		_arg1 = C.TRUE
	}

	C.gtk_source_view_set_indent_on_tab(_arg0, _arg1)
	runtime.KeepAlive(view)
	runtime.KeepAlive(enable)
}

// SetIndentWidth sets the number of spaces to use for each step of indent when
// the tab key is pressed. If width is -1, the value of the SourceView:tab-width
// property will be used.
//
// The SourceView:indent-width interacts with the
// SourceView:insert-spaces-instead-of-tabs property and SourceView:tab-width.
// An example will be clearer: if the SourceView:indent-width is 4 and
// SourceView:tab-width is 8 and SourceView:insert-spaces-instead-of-tabs is
// FALSE, then pressing the tab key at the beginning of a line will insert 4
// spaces. So far so good. Pressing the tab key a second time will remove the 4
// spaces and insert a \t character instead (since SourceView:tab-width is 8).
// On the other hand, if SourceView:insert-spaces-instead-of-tabs is TRUE, the
// second tab key pressed will insert 4 more spaces for a total of 8 spaces in
// the TextBuffer.
//
// The test-widget program (available in the GtkSourceView repository) may be
// useful to better understand the indentation settings (enable the space
// drawing!).
//
// The function takes the following parameters:
//
//    - width: indent width in characters.
//
func (view *View) SetIndentWidth(width int) {
	var _arg0 *C.GtkSourceView // out
	var _arg1 C.gint           // out

	_arg0 = (*C.GtkSourceView)(unsafe.Pointer(view.Native()))
	_arg1 = C.gint(width)

	C.gtk_source_view_set_indent_width(_arg0, _arg1)
	runtime.KeepAlive(view)
	runtime.KeepAlive(width)
}

// SetIndenter sets the indenter for view to indenter.
//
// Note that the indenter will not be used unless SourceView:auto-indent has
// been set to TRUE.
//
// The function takes the following parameters:
//
//    - indenter (optional) or NULL.
//
func (view *View) SetIndenter(indenter Indenterer) {
	var _arg0 *C.GtkSourceView     // out
	var _arg1 *C.GtkSourceIndenter // out

	_arg0 = (*C.GtkSourceView)(unsafe.Pointer(view.Native()))
	if indenter != nil {
		_arg1 = (*C.GtkSourceIndenter)(unsafe.Pointer(indenter.Native()))
	}

	C.gtk_source_view_set_indenter(_arg0, _arg1)
	runtime.KeepAlive(view)
	runtime.KeepAlive(indenter)
}

// SetInsertSpacesInsteadOfTabs: if TRUE a tab key pressed is replaced by a
// group of space characters. Of course it is still possible to insert a real \t
// programmatically with the TextBuffer API.
//
// The function takes the following parameters:
//
//    - enable: whether to insert spaces instead of tabs.
//
func (view *View) SetInsertSpacesInsteadOfTabs(enable bool) {
	var _arg0 *C.GtkSourceView // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkSourceView)(unsafe.Pointer(view.Native()))
	if enable {
		_arg1 = C.TRUE
	}

	C.gtk_source_view_set_insert_spaces_instead_of_tabs(_arg0, _arg1)
	runtime.KeepAlive(view)
	runtime.KeepAlive(enable)
}

// SetMarkAttributes sets attributes and priority for the category.
//
// The function takes the following parameters:
//
//    - category: category.
//    - attributes: mark attributes.
//    - priority of the category.
//
func (view *View) SetMarkAttributes(category string, attributes *MarkAttributes, priority int) {
	var _arg0 *C.GtkSourceView           // out
	var _arg1 *C.gchar                   // out
	var _arg2 *C.GtkSourceMarkAttributes // out
	var _arg3 C.gint                     // out

	_arg0 = (*C.GtkSourceView)(unsafe.Pointer(view.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(category)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GtkSourceMarkAttributes)(unsafe.Pointer(attributes.Native()))
	_arg3 = C.gint(priority)

	C.gtk_source_view_set_mark_attributes(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(view)
	runtime.KeepAlive(category)
	runtime.KeepAlive(attributes)
	runtime.KeepAlive(priority)
}

// SetRightMarginPosition sets the position of the right margin in the given
// view.
//
// The function takes the following parameters:
//
//    - pos: width in characters where to position the right margin.
//
func (view *View) SetRightMarginPosition(pos uint) {
	var _arg0 *C.GtkSourceView // out
	var _arg1 C.guint          // out

	_arg0 = (*C.GtkSourceView)(unsafe.Pointer(view.Native()))
	_arg1 = C.guint(pos)

	C.gtk_source_view_set_right_margin_position(_arg0, _arg1)
	runtime.KeepAlive(view)
	runtime.KeepAlive(pos)
}

// SetShowLineMarks: if TRUE line marks will be displayed beside the text.
//
// The function takes the following parameters:
//
//    - show: whether line marks should be displayed.
//
func (view *View) SetShowLineMarks(show bool) {
	var _arg0 *C.GtkSourceView // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkSourceView)(unsafe.Pointer(view.Native()))
	if show {
		_arg1 = C.TRUE
	}

	C.gtk_source_view_set_show_line_marks(_arg0, _arg1)
	runtime.KeepAlive(view)
	runtime.KeepAlive(show)
}

// SetShowLineNumbers: if TRUE line numbers will be displayed beside the text.
//
// The function takes the following parameters:
//
//    - show: whether line numbers should be displayed.
//
func (view *View) SetShowLineNumbers(show bool) {
	var _arg0 *C.GtkSourceView // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkSourceView)(unsafe.Pointer(view.Native()))
	if show {
		_arg1 = C.TRUE
	}

	C.gtk_source_view_set_show_line_numbers(_arg0, _arg1)
	runtime.KeepAlive(view)
	runtime.KeepAlive(show)
}

// SetShowRightMargin: if TRUE a right margin is displayed.
//
// The function takes the following parameters:
//
//    - show: whether to show a right margin.
//
func (view *View) SetShowRightMargin(show bool) {
	var _arg0 *C.GtkSourceView // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkSourceView)(unsafe.Pointer(view.Native()))
	if show {
		_arg1 = C.TRUE
	}

	C.gtk_source_view_set_show_right_margin(_arg0, _arg1)
	runtime.KeepAlive(view)
	runtime.KeepAlive(show)
}

// SetSmartBackspace: when set to TRUE, pressing the Backspace key will try to
// delete spaces up to the previous tab stop.
//
// The function takes the following parameters:
//
//    - smartBackspace: whether to enable smart Backspace handling.
//
func (view *View) SetSmartBackspace(smartBackspace bool) {
	var _arg0 *C.GtkSourceView // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkSourceView)(unsafe.Pointer(view.Native()))
	if smartBackspace {
		_arg1 = C.TRUE
	}

	C.gtk_source_view_set_smart_backspace(_arg0, _arg1)
	runtime.KeepAlive(view)
	runtime.KeepAlive(smartBackspace)
}

// SetSmartHomeEnd: set the desired movement of the cursor when HOME and END
// keys are pressed.
//
// The function takes the following parameters:
//
//    - smartHomeEnd: desired behavior among SourceSmartHomeEndType.
//
func (view *View) SetSmartHomeEnd(smartHomeEnd SmartHomeEndType) {
	var _arg0 *C.GtkSourceView            // out
	var _arg1 C.GtkSourceSmartHomeEndType // out

	_arg0 = (*C.GtkSourceView)(unsafe.Pointer(view.Native()))
	_arg1 = C.GtkSourceSmartHomeEndType(smartHomeEnd)

	C.gtk_source_view_set_smart_home_end(_arg0, _arg1)
	runtime.KeepAlive(view)
	runtime.KeepAlive(smartHomeEnd)
}

// SetTabWidth sets the width of tabulation in characters. The TextBuffer still
// contains \t characters, but they can take a different visual width in a
// SourceView widget.
//
// The function takes the following parameters:
//
//    - width of tab in characters.
//
func (view *View) SetTabWidth(width uint) {
	var _arg0 *C.GtkSourceView // out
	var _arg1 C.guint          // out

	_arg0 = (*C.GtkSourceView)(unsafe.Pointer(view.Native()))
	_arg1 = C.guint(width)

	C.gtk_source_view_set_tab_width(_arg0, _arg1)
	runtime.KeepAlive(view)
	runtime.KeepAlive(width)
}

// UnindentLines removes one indentation level at the beginning of the specified
// lines.
//
// The function takes the following parameters:
//
//    - start of the first line to indent.
//    - end of the last line to indent.
//
func (view *View) UnindentLines(start, end *gtk.TextIter) {
	var _arg0 *C.GtkSourceView // out
	var _arg1 *C.GtkTextIter   // out
	var _arg2 *C.GtkTextIter   // out

	_arg0 = (*C.GtkSourceView)(unsafe.Pointer(view.Native()))
	_arg1 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(start)))
	_arg2 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(end)))

	C.gtk_source_view_unindent_lines(_arg0, _arg1, _arg2)
	runtime.KeepAlive(view)
	runtime.KeepAlive(start)
	runtime.KeepAlive(end)
}
