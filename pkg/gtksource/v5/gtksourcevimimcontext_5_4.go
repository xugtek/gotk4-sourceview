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
// extern void _gotk4_gtksource5_VimIMContext_ConnectWrite(gpointer, GtkSourceView*, gchar*, guintptr);
// extern void _gotk4_gtksource5_VimIMContext_ConnectFormatText(gpointer, GtkTextIter*, GtkTextIter*, guintptr);
// extern void _gotk4_gtksource5_VimIMContext_ConnectEdit(gpointer, GtkSourceView*, gchar*, guintptr);
// extern gboolean _gotk4_gtksource5_VimIMContext_ConnectExecuteCommand(gpointer, gchar*, guintptr);
import "C"

// GType values.
var (
	GTypeVimIMContext = coreglib.Type(C.gtk_source_vim_im_context_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeVimIMContext, F: marshalVimIMContext},
	})
}

// VimIMContextOverrides contains methods that are overridable.
type VimIMContextOverrides struct {
}

func defaultVimIMContextOverrides(v *VimIMContext) VimIMContextOverrides {
	return VimIMContextOverrides{}
}

// VimIMContext: vim emulation.
//
// The GtkSourceVimIMContext is a gtk.IMContext implementation that can be used
// to provide Vim-like editing controls within a view.
//
// The GtkSourceViMIMContext will process incoming gdk.KeyEvent as the user
// types. It should be used in conjunction with a gtk.EventControllerKey.
//
// Various features supported by GtkSourceVimIMContext include:
//
//   - Normal, Insert, Replace, Visual, and Visual Line modes
//   - Support for an integrated command bar and current command preview
//   - Search and replace
//   - Motions and Text Objects
//   - History replay
//   - Jumplists within the current file
//   - Registers including the system and primary clipboards
//   - Creation and motion to marks
//   - Some commonly used Vim commands
//
// It is recommended that applications display the contents of
// vimimcontext:command-bar-text and vimimcontext:command-text to the user as
// they represent the command-bar and current command preview found in Vim.
//
// GtkSourceVimIMContext attempts to work with additional gtk.IMContext
// implementations such as IBus by querying the gtk.TextView before processing
// the command in states which support it (notably Insert and Replace modes).
//
//	GtkEventController *key;
//	GtkIMContext *im_context;
//	GtkWidget *view;
//
//	view = gtk_source_view_new ();
//	im_context = gtk_source_vim_im_context_new ();
//	key = gtk_event_controller_key_new ();
//
//	gtk_event_controller_key_set_im_context (GTK_EVENT_CONTROLLER_KEY (key), im_context);
//	gtk_event_controller_set_propagation_phase (key, GTK_PHASE_CAPTURE);
//	gtk_widget_add_controller (view, key);
//	gtk_im_context_set_client_widget (im_context, view);
//
//	g_object_bind_property (im_context, "command-bar-text", command_bar_label, "label", 0);
//	g_object_bind_property (im_context, "command-text", command_label, "label", 0);.
type VimIMContext struct {
	_ [0]func() // equal guard
	gtk.IMContext
}

var (
	_ gtk.IMContexter = (*VimIMContext)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*VimIMContext, *VimIMContextClass, VimIMContextOverrides](
		GTypeVimIMContext,
		initVimIMContextClass,
		wrapVimIMContext,
		defaultVimIMContextOverrides,
	)
}

func initVimIMContextClass(gclass unsafe.Pointer, overrides VimIMContextOverrides, classInitFunc func(*VimIMContextClass)) {
	if classInitFunc != nil {
		class := (*VimIMContextClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapVimIMContext(obj *coreglib.Object) *VimIMContext {
	return &VimIMContext{
		IMContext: gtk.IMContext{
			Object: obj,
		},
	}
}

func marshalVimIMContext(p uintptr) (interface{}, error) {
	return wrapVimIMContext(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectEdit requests the application open the file found at path.
//
// If path is NULL, then the current file should be reloaded from storage.
//
// This may be executed in relation to the user running the :edit or :e
// commands.
func (self *VimIMContext) ConnectEdit(f func(view *View, path string)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "edit", false, unsafe.Pointer(C._gotk4_gtksource5_VimIMContext_ConnectEdit), f)
}

// ConnectExecuteCommand: signal is emitted when a command should be executed.
// This might be something like :wq or :e <path>.
//
// If the application chooses to implement this, it should return TRUE from this
// signal to indicate the command has been handled.
func (self *VimIMContext) ConnectExecuteCommand(f func(command string) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "execute-command", false, unsafe.Pointer(C._gotk4_gtksource5_VimIMContext_ConnectExecuteCommand), f)
}

// ConnectFormatText requests that the application format the text between begin
// and end.
func (self *VimIMContext) ConnectFormatText(f func(begin, end *gtk.TextIter)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "format-text", false, unsafe.Pointer(C._gotk4_gtksource5_VimIMContext_ConnectFormatText), f)
}

// ConnectWrite requests the application save the file.
//
// If a filename was provided, it will be available to the signal handler as
// path. This may be executed in relation to the user running the :write or :w
// commands.
func (self *VimIMContext) ConnectWrite(f func(view *View, path string)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "write", false, unsafe.Pointer(C._gotk4_gtksource5_VimIMContext_ConnectWrite), f)
}

func NewVimIMContext() *VimIMContext {
	var _cret *C.GtkIMContext // in

	_cret = C.gtk_source_vim_im_context_new()

	var _vimIMContext *VimIMContext // out

	_vimIMContext = wrapVimIMContext(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _vimIMContext
}

// ExecuteCommand executes command as if it was typed into the command bar by
// the user except that this does not emit the vimimcontext::execute-command
// signal.
//
// The function takes the following parameters:
//
//   - command text.
func (self *VimIMContext) ExecuteCommand(command string) {
	var _arg0 *C.GtkSourceVimIMContext // out
	var _arg1 *C.char                  // out

	_arg0 = (*C.GtkSourceVimIMContext)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(command)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_source_vim_im_context_execute_command(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(command)
}

// CommandBarText gets the current command-bar text as it is entered by the
// user.
//
// The function returns the following values:
//
//   - utf8: string containing the command-bar text.
func (self *VimIMContext) CommandBarText() string {
	var _arg0 *C.GtkSourceVimIMContext // out
	var _cret *C.char                  // in

	_arg0 = (*C.GtkSourceVimIMContext)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_source_vim_im_context_get_command_bar_text(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// CommandText gets the current command text as it is entered by the user.
//
// The function returns the following values:
//
//   - utf8: string containing the command text.
func (self *VimIMContext) CommandText() string {
	var _arg0 *C.GtkSourceVimIMContext // out
	var _cret *C.char                  // in

	_arg0 = (*C.GtkSourceVimIMContext)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_source_vim_im_context_get_command_text(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}
