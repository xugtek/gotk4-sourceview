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
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
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
		{T: externglib.Type(C.gtk_source_completion_provider_get_type()), F: marshalCompletionProviderer},
	})
}

// CompletionProviderOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type CompletionProviderOverrider interface {
	// Activate: this function requests proposal to be activated by the
	// SourceCompletionProvider.
	//
	// What the provider does to activate the proposal is specific to that
	// provider. Many providers may choose to insert a SourceSnippet with edit
	// points the user may cycle through.
	//
	// See also: SourceSnippet, SourceSnippetChunk,
	// gtk_source_view_push_snippet().
	//
	// The function takes the following parameters:
	//
	//    - context: SourceCompletionContext.
	//    - proposal: SourceCompletionProposal.
	//
	Activate(context *CompletionContext, proposal CompletionProposaller)
	// Display: this function requests that the SourceCompletionProvider
	// prepares cell to display the contents of proposal. Based on cells column
	// type, you may want to display different information.
	//
	// This allows for columns of information among completion proposals
	// resulting in better alignment of similar content (icons, return types,
	// method names, and parameter lists).
	//
	// The function takes the following parameters:
	//
	//    - context: SourceCompletionContext.
	//    - proposal: SourceCompletionProposal.
	//    - cell: SourceCompletionCell.
	//
	Display(context *CompletionContext, proposal CompletionProposaller, cell *CompletionCell)
	// Priority: this function should return the priority of self in context.
	//
	// The priority is used to sort groups of completion proposals by provider
	// so that higher priority providers results are shown above lower priority
	// providers.
	//
	// Lower value indicates higher priority.
	//
	// The function takes the following parameters:
	//
	//    - context: SourceCompletionContext.
	//
	// The function returns the following values:
	//
	Priority(context *CompletionContext) int
	// Title gets the title of the completion provider, if any.
	//
	// Currently, titles are not displayed in the completion results, but may be
	// at some point in the future when non-NULL.
	//
	// The function returns the following values:
	//
	//    - utf8 (optional): title for the provider or NULL.
	//
	Title() string
	// IsTrigger: this function is used to determine of a character inserted
	// into the text editor should cause a new completion request to be
	// triggered.
	//
	// An example would be period '.' which might indicate that the user wants
	// to complete method or field names of an object.
	//
	// The function takes the following parameters:
	//
	//    - iter: TextIter.
	//    - ch of the character inserted.
	//
	// The function returns the following values:
	//
	IsTrigger(iter *gtk.TextIter, ch uint32) bool
	// KeyActivates: this function is used to determine if a key typed by the
	// user should activate proposal (resulting in committing the text to the
	// editor).
	//
	// This is useful when using languages where convention may lead to less
	// typing by the user. One example may be the use of "." or "-" to expand a
	// field access in the C programming language.
	//
	// The function takes the following parameters:
	//
	//    - context: SourceCompletionContext.
	//    - proposal: SourceCompletionProposal.
	//    - keyval such as GDK_KEY_period.
	//    - state or 0.
	//
	// The function returns the following values:
	//
	KeyActivates(context *CompletionContext, proposal CompletionProposaller, keyval uint, state gdk.ModifierType) bool
	// PopulateAsync: asynchronously requests that the provider populates the
	// completion results for context.
	//
	// For providers that would like to populate a Model while those results are
	// displayed to the user,
	// gtk_source_completion_context_set_proposals_for_provider() may be used to
	// reduce latency until the user sees results.
	//
	// The function takes the following parameters:
	//
	//    - ctx (optional) or NULL.
	//    - context: SourceCompletionContext.
	//    - callback (optional) to execute upon completion.
	//
	PopulateAsync(ctx context.Context, context *CompletionContext, callback gio.AsyncReadyCallback)
	// PopulateFinish completes an asynchronous operation to populate a
	// completion provider.
	//
	// The function takes the following parameters:
	//
	//    - result provided to callback.
	//
	// The function returns the following values:
	//
	//    - listModel of SourceCompletionProposal.
	//
	PopulateFinish(result gio.AsyncResulter) (gio.ListModeller, error)
	// Refilter: this function can be used to filter results previously provided
	// to the SourceCompletionContext by the SourceCompletionProvider.
	//
	// This can happen as the user types additionl text onto the word so that
	// previously matched items may be removed from the list instead of
	// generating new Model of results.
	//
	// The function takes the following parameters:
	//
	//    - context: SourceCompletionContext.
	//    - model: Model.
	//
	Refilter(context *CompletionContext, model gio.ListModeller)
}

type CompletionProvider struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*CompletionProvider)(nil)
)

// CompletionProviderer describes CompletionProvider's interface methods.
type CompletionProviderer interface {
	externglib.Objector

	// Activate: this function requests proposal to be activated by the
	// SourceCompletionProvider.
	Activate(context *CompletionContext, proposal CompletionProposaller)
	// Display: this function requests that the SourceCompletionProvider
	// prepares cell to display the contents of proposal.
	Display(context *CompletionContext, proposal CompletionProposaller, cell *CompletionCell)
	// Priority: this function should return the priority of self in context.
	Priority(context *CompletionContext) int
	// Title gets the title of the completion provider, if any.
	Title() string
	// IsTrigger: this function is used to determine of a character inserted
	// into the text editor should cause a new completion request to be
	// triggered.
	IsTrigger(iter *gtk.TextIter, ch uint32) bool
	// KeyActivates: this function is used to determine if a key typed by the
	// user should activate proposal (resulting in committing the text to the
	// editor).
	KeyActivates(context *CompletionContext, proposal CompletionProposaller, keyval uint, state gdk.ModifierType) bool
	// PopulateAsync: asynchronously requests that the provider populates the
	// completion results for context.
	PopulateAsync(ctx context.Context, context *CompletionContext, callback gio.AsyncReadyCallback)
	// PopulateFinish completes an asynchronous operation to populate a
	// completion provider.
	PopulateFinish(result gio.AsyncResulter) (gio.ListModeller, error)
	// Refilter: this function can be used to filter results previously provided
	// to the SourceCompletionContext by the SourceCompletionProvider.
	Refilter(context *CompletionContext, model gio.ListModeller)
}

var _ CompletionProviderer = (*CompletionProvider)(nil)

func wrapCompletionProvider(obj *externglib.Object) *CompletionProvider {
	return &CompletionProvider{
		Object: obj,
	}
}

func marshalCompletionProviderer(p uintptr) (interface{}, error) {
	return wrapCompletionProvider(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Activate: this function requests proposal to be activated by the
// SourceCompletionProvider.
//
// What the provider does to activate the proposal is specific to that provider.
// Many providers may choose to insert a SourceSnippet with edit points the user
// may cycle through.
//
// See also: SourceSnippet, SourceSnippetChunk, gtk_source_view_push_snippet().
//
// The function takes the following parameters:
//
//    - context: SourceCompletionContext.
//    - proposal: SourceCompletionProposal.
//
func (self *CompletionProvider) Activate(context *CompletionContext, proposal CompletionProposaller) {
	var _arg0 *C.GtkSourceCompletionProvider // out
	var _arg1 *C.GtkSourceCompletionContext  // out
	var _arg2 *C.GtkSourceCompletionProposal // out

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkSourceCompletionContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.GtkSourceCompletionProposal)(unsafe.Pointer(proposal.Native()))

	C.gtk_source_completion_provider_activate(_arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(context)
	runtime.KeepAlive(proposal)
}

// Display: this function requests that the SourceCompletionProvider prepares
// cell to display the contents of proposal. Based on cells column type, you may
// want to display different information.
//
// This allows for columns of information among completion proposals resulting
// in better alignment of similar content (icons, return types, method names,
// and parameter lists).
//
// The function takes the following parameters:
//
//    - context: SourceCompletionContext.
//    - proposal: SourceCompletionProposal.
//    - cell: SourceCompletionCell.
//
func (self *CompletionProvider) Display(context *CompletionContext, proposal CompletionProposaller, cell *CompletionCell) {
	var _arg0 *C.GtkSourceCompletionProvider // out
	var _arg1 *C.GtkSourceCompletionContext  // out
	var _arg2 *C.GtkSourceCompletionProposal // out
	var _arg3 *C.GtkSourceCompletionCell     // out

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkSourceCompletionContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.GtkSourceCompletionProposal)(unsafe.Pointer(proposal.Native()))
	_arg3 = (*C.GtkSourceCompletionCell)(unsafe.Pointer(cell.Native()))

	C.gtk_source_completion_provider_display(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(self)
	runtime.KeepAlive(context)
	runtime.KeepAlive(proposal)
	runtime.KeepAlive(cell)
}

// Priority: this function should return the priority of self in context.
//
// The priority is used to sort groups of completion proposals by provider so
// that higher priority providers results are shown above lower priority
// providers.
//
// Lower value indicates higher priority.
//
// The function takes the following parameters:
//
//    - context: SourceCompletionContext.
//
// The function returns the following values:
//
func (self *CompletionProvider) Priority(context *CompletionContext) int {
	var _arg0 *C.GtkSourceCompletionProvider // out
	var _arg1 *C.GtkSourceCompletionContext  // out
	var _cret C.int                          // in

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkSourceCompletionContext)(unsafe.Pointer(context.Native()))

	_cret = C.gtk_source_completion_provider_get_priority(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(context)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Title gets the title of the completion provider, if any.
//
// Currently, titles are not displayed in the completion results, but may be at
// some point in the future when non-NULL.
//
// The function returns the following values:
//
//    - utf8 (optional): title for the provider or NULL.
//
func (self *CompletionProvider) Title() string {
	var _arg0 *C.GtkSourceCompletionProvider // out
	var _cret *C.char                        // in

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_source_completion_provider_get_title(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// IsTrigger: this function is used to determine of a character inserted into
// the text editor should cause a new completion request to be triggered.
//
// An example would be period '.' which might indicate that the user wants to
// complete method or field names of an object.
//
// The function takes the following parameters:
//
//    - iter: TextIter.
//    - ch of the character inserted.
//
// The function returns the following values:
//
func (self *CompletionProvider) IsTrigger(iter *gtk.TextIter, ch uint32) bool {
	var _arg0 *C.GtkSourceCompletionProvider // out
	var _arg1 *C.GtkTextIter                 // out
	var _arg2 C.gunichar                     // out
	var _cret C.gboolean                     // in

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(iter)))
	_arg2 = C.gunichar(ch)

	_cret = C.gtk_source_completion_provider_is_trigger(_arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(iter)
	runtime.KeepAlive(ch)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// KeyActivates: this function is used to determine if a key typed by the user
// should activate proposal (resulting in committing the text to the editor).
//
// This is useful when using languages where convention may lead to less typing
// by the user. One example may be the use of "." or "-" to expand a field
// access in the C programming language.
//
// The function takes the following parameters:
//
//    - context: SourceCompletionContext.
//    - proposal: SourceCompletionProposal.
//    - keyval such as GDK_KEY_period.
//    - state or 0.
//
// The function returns the following values:
//
func (self *CompletionProvider) KeyActivates(context *CompletionContext, proposal CompletionProposaller, keyval uint, state gdk.ModifierType) bool {
	var _arg0 *C.GtkSourceCompletionProvider // out
	var _arg1 *C.GtkSourceCompletionContext  // out
	var _arg2 *C.GtkSourceCompletionProposal // out
	var _arg3 C.guint                        // out
	var _arg4 C.GdkModifierType              // out
	var _cret C.gboolean                     // in

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkSourceCompletionContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.GtkSourceCompletionProposal)(unsafe.Pointer(proposal.Native()))
	_arg3 = C.guint(keyval)
	_arg4 = C.GdkModifierType(state)

	_cret = C.gtk_source_completion_provider_key_activates(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(self)
	runtime.KeepAlive(context)
	runtime.KeepAlive(proposal)
	runtime.KeepAlive(keyval)
	runtime.KeepAlive(state)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PopulateAsync: asynchronously requests that the provider populates the
// completion results for context.
//
// For providers that would like to populate a Model while those results are
// displayed to the user,
// gtk_source_completion_context_set_proposals_for_provider() may be used to
// reduce latency until the user sees results.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//    - context: SourceCompletionContext.
//    - callback (optional) to execute upon completion.
//
func (self *CompletionProvider) PopulateAsync(ctx context.Context, context *CompletionContext, callback gio.AsyncReadyCallback) {
	var _arg0 *C.GtkSourceCompletionProvider // out
	var _arg2 *C.GCancellable                // out
	var _arg1 *C.GtkSourceCompletionContext  // out
	var _arg3 C.GAsyncReadyCallback          // out
	var _arg4 C.gpointer

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(self.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.GtkSourceCompletionContext)(unsafe.Pointer(context.Native()))
	if callback != nil {
		_arg3 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg4 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.gtk_source_completion_provider_populate_async(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(self)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(context)
	runtime.KeepAlive(callback)
}

// PopulateFinish completes an asynchronous operation to populate a completion
// provider.
//
// The function takes the following parameters:
//
//    - result provided to callback.
//
// The function returns the following values:
//
//    - listModel of SourceCompletionProposal.
//
func (self *CompletionProvider) PopulateFinish(result gio.AsyncResulter) (gio.ListModeller, error) {
	var _arg0 *C.GtkSourceCompletionProvider // out
	var _arg1 *C.GAsyncResult                // out
	var _cret *C.GListModel                  // in
	var _cerr *C.GError                      // in

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.gtk_source_completion_provider_populate_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(result)

	var _listModel gio.ListModeller // out
	var _goerr error                // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.ListModeller is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(gio.ListModeller)
			return ok
		})
		rv, ok := casted.(gio.ListModeller)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.ListModeller")
		}
		_listModel = rv
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _listModel, _goerr
}

// Refilter: this function can be used to filter results previously provided to
// the SourceCompletionContext by the SourceCompletionProvider.
//
// This can happen as the user types additionl text onto the word so that
// previously matched items may be removed from the list instead of generating
// new Model of results.
//
// The function takes the following parameters:
//
//    - context: SourceCompletionContext.
//    - model: Model.
//
func (self *CompletionProvider) Refilter(context *CompletionContext, model gio.ListModeller) {
	var _arg0 *C.GtkSourceCompletionProvider // out
	var _arg1 *C.GtkSourceCompletionContext  // out
	var _arg2 *C.GListModel                  // out

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkSourceCompletionContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.GListModel)(unsafe.Pointer(model.Native()))

	C.gtk_source_completion_provider_refilter(_arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(context)
	runtime.KeepAlive(model)
}
