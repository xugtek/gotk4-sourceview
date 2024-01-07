// Code generated by girgen. DO NOT EDIT.

package gtksource

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v3"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtksourceview/gtksource.h>
// GIcon* _gotk4_gtksource3_CompletionProvider_virtual_get_gicon(void* fnptr, GtkSourceCompletionProvider* arg0) {
//   return ((GIcon* (*)(GtkSourceCompletionProvider*))(fnptr))(arg0);
// };
// GdkPixbuf* _gotk4_gtksource3_CompletionProvider_virtual_get_icon(void* fnptr, GtkSourceCompletionProvider* arg0) {
//   return ((GdkPixbuf* (*)(GtkSourceCompletionProvider*))(fnptr))(arg0);
// };
// GtkSourceCompletionActivation _gotk4_gtksource3_CompletionProvider_virtual_get_activation(void* fnptr, GtkSourceCompletionProvider* arg0) {
//   return ((GtkSourceCompletionActivation (*)(GtkSourceCompletionProvider*))(fnptr))(arg0);
// };
// GtkWidget* _gotk4_gtksource3_CompletionProvider_virtual_get_info_widget(void* fnptr, GtkSourceCompletionProvider* arg0, GtkSourceCompletionProposal* arg1) {
//   return ((GtkWidget* (*)(GtkSourceCompletionProvider*, GtkSourceCompletionProposal*))(fnptr))(arg0, arg1);
// };
// gboolean _gotk4_gtksource3_CompletionProvider_virtual_activate_proposal(void* fnptr, GtkSourceCompletionProvider* arg0, GtkSourceCompletionProposal* arg1, GtkTextIter* arg2) {
//   return ((gboolean (*)(GtkSourceCompletionProvider*, GtkSourceCompletionProposal*, GtkTextIter*))(fnptr))(arg0, arg1, arg2);
// };
// gboolean _gotk4_gtksource3_CompletionProvider_virtual_get_start_iter(void* fnptr, GtkSourceCompletionProvider* arg0, GtkSourceCompletionContext* arg1, GtkSourceCompletionProposal* arg2, GtkTextIter* arg3) {
//   return ((gboolean (*)(GtkSourceCompletionProvider*, GtkSourceCompletionContext*, GtkSourceCompletionProposal*, GtkTextIter*))(fnptr))(arg0, arg1, arg2, arg3);
// };
// gboolean _gotk4_gtksource3_CompletionProvider_virtual_match(void* fnptr, GtkSourceCompletionProvider* arg0, GtkSourceCompletionContext* arg1) {
//   return ((gboolean (*)(GtkSourceCompletionProvider*, GtkSourceCompletionContext*))(fnptr))(arg0, arg1);
// };
// gchar* _gotk4_gtksource3_CompletionProvider_virtual_get_icon_name(void* fnptr, GtkSourceCompletionProvider* arg0) {
//   return ((gchar* (*)(GtkSourceCompletionProvider*))(fnptr))(arg0);
// };
// gchar* _gotk4_gtksource3_CompletionProvider_virtual_get_name(void* fnptr, GtkSourceCompletionProvider* arg0) {
//   return ((gchar* (*)(GtkSourceCompletionProvider*))(fnptr))(arg0);
// };
// gint _gotk4_gtksource3_CompletionProvider_virtual_get_interactive_delay(void* fnptr, GtkSourceCompletionProvider* arg0) {
//   return ((gint (*)(GtkSourceCompletionProvider*))(fnptr))(arg0);
// };
// gint _gotk4_gtksource3_CompletionProvider_virtual_get_priority(void* fnptr, GtkSourceCompletionProvider* arg0) {
//   return ((gint (*)(GtkSourceCompletionProvider*))(fnptr))(arg0);
// };
// void _gotk4_gtksource3_CompletionProvider_virtual_populate(void* fnptr, GtkSourceCompletionProvider* arg0, GtkSourceCompletionContext* arg1) {
//   ((void (*)(GtkSourceCompletionProvider*, GtkSourceCompletionContext*))(fnptr))(arg0, arg1);
// };
// void _gotk4_gtksource3_CompletionProvider_virtual_update_info(void* fnptr, GtkSourceCompletionProvider* arg0, GtkSourceCompletionProposal* arg1, GtkSourceCompletionInfo* arg2) {
//   ((void (*)(GtkSourceCompletionProvider*, GtkSourceCompletionProposal*, GtkSourceCompletionInfo*))(fnptr))(arg0, arg1, arg2);
// };
import "C"

// GType values.
var (
	GTypeCompletionProvider = coreglib.Type(C.gtk_source_completion_provider_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeCompletionProvider, F: marshalCompletionProvider},
	})
}

//
// CompletionProvider wraps an interface. This means the user can get the
// underlying type by calling Cast().
type CompletionProvider struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*CompletionProvider)(nil)
)

// CompletionProviderer describes CompletionProvider's interface methods.
type CompletionProviderer interface {
	coreglib.Objector

	// ActivateProposal: activate proposal at iter.
	ActivateProposal(proposal CompletionProposaller, iter *gtk.TextIter) bool
	// Activation: get with what kind of activation the provider should be
	// activated.
	Activation() CompletionActivation
	// GIcon gets the #GIcon for the icon of provider.
	GIcon() *gio.Icon
	// Icon: get the Pixbuf for the icon of the provider.
	Icon() *gdkpixbuf.Pixbuf
	// IconName gets the icon name of provider.
	IconName() string
	// InfoWidget: get a customized info widget to show extra information of a
	// proposal.
	InfoWidget(proposal CompletionProposaller) gtk.Widgetter
	// InteractiveDelay: get the delay in milliseconds before starting
	// interactive completion for this provider.
	InteractiveDelay() int
	// Name: get the name of the provider.
	Name() string
	// Priority: get the provider priority.
	Priority() int
	// StartIter: get the TextIter at which the completion for proposal starts.
	StartIter(context *CompletionContext, proposal CompletionProposaller) (*gtk.TextIter, bool)
	// Match: get whether the provider match the context of completion detailed
	// in context.
	Match(context *CompletionContext) bool
	// Populate context with proposals from provider added with the
	// gtk_source_completion_context_add_proposals() function.
	Populate(context *CompletionContext)
	// UpdateInfo: update extra information shown in info for proposal.
	UpdateInfo(proposal CompletionProposaller, info *CompletionInfo)
}

var _ CompletionProviderer = (*CompletionProvider)(nil)

func wrapCompletionProvider(obj *coreglib.Object) *CompletionProvider {
	return &CompletionProvider{
		Object: obj,
	}
}

func marshalCompletionProvider(p uintptr) (interface{}, error) {
	return wrapCompletionProvider(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ActivateProposal: activate proposal at iter. When this functions
// returns FALSE, the default activation of proposal will take place
// which replaces the word at iter with the text of proposal (see
// gtk_source_completion_proposal_get_text()).
//
// Here is how the default activation selects the boundaries of the word
// to replace. The end of the word is iter. For the start of the word,
// it depends on whether a start iter is defined for proposal (see
// gtk_source_completion_provider_get_start_iter()). If a start iter is defined,
// the start of the word is the start iter. Else, the word (as long as possible)
// will contain only alphanumerical and the "_" characters.
//
// The function takes the following parameters:
//
//   - proposal: SourceCompletionProposal.
//   - iter: TextIter.
//
// The function returns the following values:
//
//   - ok: TRUE to indicate that the proposal activation has been handled,
//     FALSE otherwise.
//
func (provider *CompletionProvider) ActivateProposal(proposal CompletionProposaller, iter *gtk.TextIter) bool {
	var _arg0 *C.GtkSourceCompletionProvider // out
	var _arg1 *C.GtkSourceCompletionProposal // out
	var _arg2 *C.GtkTextIter                 // out
	var _cret C.gboolean                     // in

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(coreglib.InternObject(provider).Native()))
	_arg1 = (*C.GtkSourceCompletionProposal)(unsafe.Pointer(coreglib.InternObject(proposal).Native()))
	_arg2 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(iter)))

	_cret = C.gtk_source_completion_provider_activate_proposal(_arg0, _arg1, _arg2)
	runtime.KeepAlive(provider)
	runtime.KeepAlive(proposal)
	runtime.KeepAlive(iter)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Activation: get with what kind of activation the provider should be
// activated.
//
// The function returns the following values:
//
//   - completionActivation: combination of SourceCompletionActivation.
//
func (provider *CompletionProvider) Activation() CompletionActivation {
	var _arg0 *C.GtkSourceCompletionProvider  // out
	var _cret C.GtkSourceCompletionActivation // in

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(coreglib.InternObject(provider).Native()))

	_cret = C.gtk_source_completion_provider_get_activation(_arg0)
	runtime.KeepAlive(provider)

	var _completionActivation CompletionActivation // out

	_completionActivation = CompletionActivation(_cret)

	return _completionActivation
}

// GIcon gets the #GIcon for the icon of provider.
//
// The function returns the following values:
//
//   - icon (optional) to be used for the provider, or NULL if the provider does
//     not have a special icon.
//
func (provider *CompletionProvider) GIcon() *gio.Icon {
	var _arg0 *C.GtkSourceCompletionProvider // out
	var _cret *C.GIcon                       // in

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(coreglib.InternObject(provider).Native()))

	_cret = C.gtk_source_completion_provider_get_gicon(_arg0)
	runtime.KeepAlive(provider)

	var _icon *gio.Icon // out

	if _cret != nil {
		{
			obj := coreglib.Take(unsafe.Pointer(_cret))
			_icon = &gio.Icon{
				Object: obj,
			}
		}
	}

	return _icon
}

// Icon: get the Pixbuf for the icon of the provider.
//
// The function returns the following values:
//
//   - pixbuf (optional): icon to be used for the provider, or NULL if the
//     provider does not have a special icon.
//
func (provider *CompletionProvider) Icon() *gdkpixbuf.Pixbuf {
	var _arg0 *C.GtkSourceCompletionProvider // out
	var _cret *C.GdkPixbuf                   // in

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(coreglib.InternObject(provider).Native()))

	_cret = C.gtk_source_completion_provider_get_icon(_arg0)
	runtime.KeepAlive(provider)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	if _cret != nil {
		{
			obj := coreglib.Take(unsafe.Pointer(_cret))
			_pixbuf = &gdkpixbuf.Pixbuf{
				Object: obj,
				LoadableIcon: gio.LoadableIcon{
					Icon: gio.Icon{
						Object: obj,
					},
				},
			}
		}
	}

	return _pixbuf
}

// IconName gets the icon name of provider.
//
// The function returns the following values:
//
//   - utf8 (optional): icon name to be used for the provider, or NULL if the
//     provider does not have a special icon.
//
func (provider *CompletionProvider) IconName() string {
	var _arg0 *C.GtkSourceCompletionProvider // out
	var _cret *C.gchar                       // in

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(coreglib.InternObject(provider).Native()))

	_cret = C.gtk_source_completion_provider_get_icon_name(_arg0)
	runtime.KeepAlive(provider)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// InfoWidget: get a customized info widget to show extra information of a
// proposal. This allows for customized widgets on a proposal basis, although in
// general providers will have the same custom widget for all their proposals
// and proposal can be ignored. The implementation of this function is optional.
//
// If this function is not implemented, the default widget is a Label.
// The return value of gtk_source_completion_proposal_get_info() is used as the
// content of the Label.
//
// <note> <para> If implemented, gtk_source_completion_provider_update_info()
// <emphasis>must</emphasis> also be implemented. </para> </note>.
//
// The function takes the following parameters:
//
//   - proposal: currently selected SourceCompletionProposal.
//
// The function returns the following values:
//
//   - widget (optional): custom Widget to show extra information about
//     proposal, or NULL if the provider does not have a special info widget.
//
func (provider *CompletionProvider) InfoWidget(proposal CompletionProposaller) gtk.Widgetter {
	var _arg0 *C.GtkSourceCompletionProvider // out
	var _arg1 *C.GtkSourceCompletionProposal // out
	var _cret *C.GtkWidget                   // in

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(coreglib.InternObject(provider).Native()))
	_arg1 = (*C.GtkSourceCompletionProposal)(unsafe.Pointer(coreglib.InternObject(proposal).Native()))

	_cret = C.gtk_source_completion_provider_get_info_widget(_arg0, _arg1)
	runtime.KeepAlive(provider)
	runtime.KeepAlive(proposal)

	var _widget gtk.Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(gtk.Widgetter)
				return ok
			})
			rv, ok := casted.(gtk.Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// InteractiveDelay: get the delay in milliseconds before starting interactive
// completion for this provider. A value of -1 indicates to use the default
// value as set by the SourceCompletion:auto-complete-delay property.
//
// The function returns the following values:
//
//   - gint: interactive delay in milliseconds.
//
func (provider *CompletionProvider) InteractiveDelay() int {
	var _arg0 *C.GtkSourceCompletionProvider // out
	var _cret C.gint                         // in

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(coreglib.InternObject(provider).Native()))

	_cret = C.gtk_source_completion_provider_get_interactive_delay(_arg0)
	runtime.KeepAlive(provider)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Name: get the name of the provider. This should be a translatable name for
// display to the user. For example: _("Document word completion provider").
// The returned string must be freed with g_free().
//
// The function returns the following values:
//
//   - utf8: new string containing the name of the provider.
//
func (provider *CompletionProvider) Name() string {
	var _arg0 *C.GtkSourceCompletionProvider // out
	var _cret *C.gchar                       // in

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(coreglib.InternObject(provider).Native()))

	_cret = C.gtk_source_completion_provider_get_name(_arg0)
	runtime.KeepAlive(provider)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Priority: get the provider priority. The priority determines the order in
// which proposals appear in the completion popup. Higher priorities are sorted
// before lower priorities. The default priority is 0.
//
// The function returns the following values:
//
//   - gint: provider priority.
//
func (provider *CompletionProvider) Priority() int {
	var _arg0 *C.GtkSourceCompletionProvider // out
	var _cret C.gint                         // in

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(coreglib.InternObject(provider).Native()))

	_cret = C.gtk_source_completion_provider_get_priority(_arg0)
	runtime.KeepAlive(provider)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// StartIter: get the TextIter at which the completion for proposal starts.
// When implemented, this information is used to position the completion
// window accordingly when a proposal is selected in the completion window.
// The proposal text inside the completion window is aligned on iter.
//
// If this function is not implemented, the word boundary is taken to position
// the completion window. See gtk_source_completion_provider_activate_proposal()
// for an explanation on the word boundaries.
//
// When the proposal is activated, the default handler uses iter as the start of
// the word to replace. See gtk_source_completion_provider_activate_proposal()
// for more information.
//
// The function takes the following parameters:
//
//   - context: SourceCompletionContext.
//   - proposal: SourceCompletionProposal.
//
// The function returns the following values:
//
//   - iter: TextIter.
//   - ok: TRUE if iter was set for proposal, FALSE otherwise.
//
func (provider *CompletionProvider) StartIter(context *CompletionContext, proposal CompletionProposaller) (*gtk.TextIter, bool) {
	var _arg0 *C.GtkSourceCompletionProvider // out
	var _arg1 *C.GtkSourceCompletionContext  // out
	var _arg2 *C.GtkSourceCompletionProposal // out
	var _arg3 C.GtkTextIter                  // in
	var _cret C.gboolean                     // in

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(coreglib.InternObject(provider).Native()))
	_arg1 = (*C.GtkSourceCompletionContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg2 = (*C.GtkSourceCompletionProposal)(unsafe.Pointer(coreglib.InternObject(proposal).Native()))

	_cret = C.gtk_source_completion_provider_get_start_iter(_arg0, _arg1, _arg2, &_arg3)
	runtime.KeepAlive(provider)
	runtime.KeepAlive(context)
	runtime.KeepAlive(proposal)

	var _iter *gtk.TextIter // out
	var _ok bool            // out

	_iter = (*gtk.TextIter)(gextras.NewStructNative(unsafe.Pointer((&_arg3))))
	if _cret != 0 {
		_ok = true
	}

	return _iter, _ok
}

// Match: get whether the provider match the context of completion detailed in
// context.
//
// The function takes the following parameters:
//
//   - context: SourceCompletionContext.
//
// The function returns the following values:
//
//   - ok: TRUE if provider matches the completion context, FALSE otherwise.
//
func (provider *CompletionProvider) Match(context *CompletionContext) bool {
	var _arg0 *C.GtkSourceCompletionProvider // out
	var _arg1 *C.GtkSourceCompletionContext  // out
	var _cret C.gboolean                     // in

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(coreglib.InternObject(provider).Native()))
	_arg1 = (*C.GtkSourceCompletionContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_cret = C.gtk_source_completion_provider_match(_arg0, _arg1)
	runtime.KeepAlive(provider)
	runtime.KeepAlive(context)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Populate context with proposals from provider added with the
// gtk_source_completion_context_add_proposals() function.
//
// The function takes the following parameters:
//
//   - context: SourceCompletionContext.
//
func (provider *CompletionProvider) Populate(context *CompletionContext) {
	var _arg0 *C.GtkSourceCompletionProvider // out
	var _arg1 *C.GtkSourceCompletionContext  // out

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(coreglib.InternObject(provider).Native()))
	_arg1 = (*C.GtkSourceCompletionContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	C.gtk_source_completion_provider_populate(_arg0, _arg1)
	runtime.KeepAlive(provider)
	runtime.KeepAlive(context)
}

// UpdateInfo: update extra information shown in info for proposal.
//
// <note> <para> This function <emphasis>must</emphasis> be implemented when
// gtk_source_completion_provider_get_info_widget() is implemented. </para>
// </note>.
//
// The function takes the following parameters:
//
//   - proposal: SourceCompletionProposal.
//   - info: SourceCompletionInfo.
//
func (provider *CompletionProvider) UpdateInfo(proposal CompletionProposaller, info *CompletionInfo) {
	var _arg0 *C.GtkSourceCompletionProvider // out
	var _arg1 *C.GtkSourceCompletionProposal // out
	var _arg2 *C.GtkSourceCompletionInfo     // out

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(coreglib.InternObject(provider).Native()))
	_arg1 = (*C.GtkSourceCompletionProposal)(unsafe.Pointer(coreglib.InternObject(proposal).Native()))
	_arg2 = (*C.GtkSourceCompletionInfo)(unsafe.Pointer(coreglib.InternObject(info).Native()))

	C.gtk_source_completion_provider_update_info(_arg0, _arg1, _arg2)
	runtime.KeepAlive(provider)
	runtime.KeepAlive(proposal)
	runtime.KeepAlive(info)
}

// activateProposal: activate proposal at iter. When this functions
// returns FALSE, the default activation of proposal will take place
// which replaces the word at iter with the text of proposal (see
// gtk_source_completion_proposal_get_text()).
//
// Here is how the default activation selects the boundaries of the word
// to replace. The end of the word is iter. For the start of the word,
// it depends on whether a start iter is defined for proposal (see
// gtk_source_completion_provider_get_start_iter()). If a start iter is defined,
// the start of the word is the start iter. Else, the word (as long as possible)
// will contain only alphanumerical and the "_" characters.
//
// The function takes the following parameters:
//
//   - proposal: SourceCompletionProposal.
//   - iter: TextIter.
//
// The function returns the following values:
//
//   - ok: TRUE to indicate that the proposal activation has been handled,
//     FALSE otherwise.
//
func (provider *CompletionProvider) activateProposal(proposal CompletionProposaller, iter *gtk.TextIter) bool {
	gclass := (*C.GtkSourceCompletionProviderIface)(coreglib.PeekParentClass(provider))
	fnarg := gclass.activate_proposal

	var _arg0 *C.GtkSourceCompletionProvider // out
	var _arg1 *C.GtkSourceCompletionProposal // out
	var _arg2 *C.GtkTextIter                 // out
	var _cret C.gboolean                     // in

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(coreglib.InternObject(provider).Native()))
	_arg1 = (*C.GtkSourceCompletionProposal)(unsafe.Pointer(coreglib.InternObject(proposal).Native()))
	_arg2 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(iter)))

	_cret = C._gotk4_gtksource3_CompletionProvider_virtual_activate_proposal(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2)
	runtime.KeepAlive(provider)
	runtime.KeepAlive(proposal)
	runtime.KeepAlive(iter)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Activation: get with what kind of activation the provider should be
// activated.
//
// The function returns the following values:
//
//   - completionActivation: combination of SourceCompletionActivation.
//
func (provider *CompletionProvider) activation() CompletionActivation {
	gclass := (*C.GtkSourceCompletionProviderIface)(coreglib.PeekParentClass(provider))
	fnarg := gclass.get_activation

	var _arg0 *C.GtkSourceCompletionProvider  // out
	var _cret C.GtkSourceCompletionActivation // in

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(coreglib.InternObject(provider).Native()))

	_cret = C._gotk4_gtksource3_CompletionProvider_virtual_get_activation(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(provider)

	var _completionActivation CompletionActivation // out

	_completionActivation = CompletionActivation(_cret)

	return _completionActivation
}

// gIcon gets the #GIcon for the icon of provider.
//
// The function returns the following values:
//
//   - icon (optional) to be used for the provider, or NULL if the provider does
//     not have a special icon.
//
func (provider *CompletionProvider) gIcon() *gio.Icon {
	gclass := (*C.GtkSourceCompletionProviderIface)(coreglib.PeekParentClass(provider))
	fnarg := gclass.get_gicon

	var _arg0 *C.GtkSourceCompletionProvider // out
	var _cret *C.GIcon                       // in

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(coreglib.InternObject(provider).Native()))

	_cret = C._gotk4_gtksource3_CompletionProvider_virtual_get_gicon(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(provider)

	var _icon *gio.Icon // out

	if _cret != nil {
		{
			obj := coreglib.Take(unsafe.Pointer(_cret))
			_icon = &gio.Icon{
				Object: obj,
			}
		}
	}

	return _icon
}

// Icon: get the Pixbuf for the icon of the provider.
//
// The function returns the following values:
//
//   - pixbuf (optional): icon to be used for the provider, or NULL if the
//     provider does not have a special icon.
//
func (provider *CompletionProvider) icon() *gdkpixbuf.Pixbuf {
	gclass := (*C.GtkSourceCompletionProviderIface)(coreglib.PeekParentClass(provider))
	fnarg := gclass.get_icon

	var _arg0 *C.GtkSourceCompletionProvider // out
	var _cret *C.GdkPixbuf                   // in

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(coreglib.InternObject(provider).Native()))

	_cret = C._gotk4_gtksource3_CompletionProvider_virtual_get_icon(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(provider)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	if _cret != nil {
		{
			obj := coreglib.Take(unsafe.Pointer(_cret))
			_pixbuf = &gdkpixbuf.Pixbuf{
				Object: obj,
				LoadableIcon: gio.LoadableIcon{
					Icon: gio.Icon{
						Object: obj,
					},
				},
			}
		}
	}

	return _pixbuf
}

// iconName gets the icon name of provider.
//
// The function returns the following values:
//
//   - utf8 (optional): icon name to be used for the provider, or NULL if the
//     provider does not have a special icon.
//
func (provider *CompletionProvider) iconName() string {
	gclass := (*C.GtkSourceCompletionProviderIface)(coreglib.PeekParentClass(provider))
	fnarg := gclass.get_icon_name

	var _arg0 *C.GtkSourceCompletionProvider // out
	var _cret *C.gchar                       // in

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(coreglib.InternObject(provider).Native()))

	_cret = C._gotk4_gtksource3_CompletionProvider_virtual_get_icon_name(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(provider)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// infoWidget: get a customized info widget to show extra information of a
// proposal. This allows for customized widgets on a proposal basis, although in
// general providers will have the same custom widget for all their proposals
// and proposal can be ignored. The implementation of this function is optional.
//
// If this function is not implemented, the default widget is a Label.
// The return value of gtk_source_completion_proposal_get_info() is used as the
// content of the Label.
//
// <note> <para> If implemented, gtk_source_completion_provider_update_info()
// <emphasis>must</emphasis> also be implemented. </para> </note>.
//
// The function takes the following parameters:
//
//   - proposal: currently selected SourceCompletionProposal.
//
// The function returns the following values:
//
//   - widget (optional): custom Widget to show extra information about
//     proposal, or NULL if the provider does not have a special info widget.
//
func (provider *CompletionProvider) infoWidget(proposal CompletionProposaller) gtk.Widgetter {
	gclass := (*C.GtkSourceCompletionProviderIface)(coreglib.PeekParentClass(provider))
	fnarg := gclass.get_info_widget

	var _arg0 *C.GtkSourceCompletionProvider // out
	var _arg1 *C.GtkSourceCompletionProposal // out
	var _cret *C.GtkWidget                   // in

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(coreglib.InternObject(provider).Native()))
	_arg1 = (*C.GtkSourceCompletionProposal)(unsafe.Pointer(coreglib.InternObject(proposal).Native()))

	_cret = C._gotk4_gtksource3_CompletionProvider_virtual_get_info_widget(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(provider)
	runtime.KeepAlive(proposal)

	var _widget gtk.Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(gtk.Widgetter)
				return ok
			})
			rv, ok := casted.(gtk.Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// interactiveDelay: get the delay in milliseconds before starting interactive
// completion for this provider. A value of -1 indicates to use the default
// value as set by the SourceCompletion:auto-complete-delay property.
//
// The function returns the following values:
//
//   - gint: interactive delay in milliseconds.
//
func (provider *CompletionProvider) interactiveDelay() int {
	gclass := (*C.GtkSourceCompletionProviderIface)(coreglib.PeekParentClass(provider))
	fnarg := gclass.get_interactive_delay

	var _arg0 *C.GtkSourceCompletionProvider // out
	var _cret C.gint                         // in

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(coreglib.InternObject(provider).Native()))

	_cret = C._gotk4_gtksource3_CompletionProvider_virtual_get_interactive_delay(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(provider)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Name: get the name of the provider. This should be a translatable name for
// display to the user. For example: _("Document word completion provider").
// The returned string must be freed with g_free().
//
// The function returns the following values:
//
//   - utf8: new string containing the name of the provider.
//
func (provider *CompletionProvider) name() string {
	gclass := (*C.GtkSourceCompletionProviderIface)(coreglib.PeekParentClass(provider))
	fnarg := gclass.get_name

	var _arg0 *C.GtkSourceCompletionProvider // out
	var _cret *C.gchar                       // in

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(coreglib.InternObject(provider).Native()))

	_cret = C._gotk4_gtksource3_CompletionProvider_virtual_get_name(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(provider)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Priority: get the provider priority. The priority determines the order in
// which proposals appear in the completion popup. Higher priorities are sorted
// before lower priorities. The default priority is 0.
//
// The function returns the following values:
//
//   - gint: provider priority.
//
func (provider *CompletionProvider) priority() int {
	gclass := (*C.GtkSourceCompletionProviderIface)(coreglib.PeekParentClass(provider))
	fnarg := gclass.get_priority

	var _arg0 *C.GtkSourceCompletionProvider // out
	var _cret C.gint                         // in

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(coreglib.InternObject(provider).Native()))

	_cret = C._gotk4_gtksource3_CompletionProvider_virtual_get_priority(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(provider)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// startIter: get the TextIter at which the completion for proposal starts.
// When implemented, this information is used to position the completion
// window accordingly when a proposal is selected in the completion window.
// The proposal text inside the completion window is aligned on iter.
//
// If this function is not implemented, the word boundary is taken to position
// the completion window. See gtk_source_completion_provider_activate_proposal()
// for an explanation on the word boundaries.
//
// When the proposal is activated, the default handler uses iter as the start of
// the word to replace. See gtk_source_completion_provider_activate_proposal()
// for more information.
//
// The function takes the following parameters:
//
//   - context: SourceCompletionContext.
//   - proposal: SourceCompletionProposal.
//
// The function returns the following values:
//
//   - iter: TextIter.
//   - ok: TRUE if iter was set for proposal, FALSE otherwise.
//
func (provider *CompletionProvider) startIter(context *CompletionContext, proposal CompletionProposaller) (*gtk.TextIter, bool) {
	gclass := (*C.GtkSourceCompletionProviderIface)(coreglib.PeekParentClass(provider))
	fnarg := gclass.get_start_iter

	var _arg0 *C.GtkSourceCompletionProvider // out
	var _arg1 *C.GtkSourceCompletionContext  // out
	var _arg2 *C.GtkSourceCompletionProposal // out
	var _arg3 C.GtkTextIter                  // in
	var _cret C.gboolean                     // in

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(coreglib.InternObject(provider).Native()))
	_arg1 = (*C.GtkSourceCompletionContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg2 = (*C.GtkSourceCompletionProposal)(unsafe.Pointer(coreglib.InternObject(proposal).Native()))

	_cret = C._gotk4_gtksource3_CompletionProvider_virtual_get_start_iter(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, &_arg3)
	runtime.KeepAlive(provider)
	runtime.KeepAlive(context)
	runtime.KeepAlive(proposal)

	var _iter *gtk.TextIter // out
	var _ok bool            // out

	_iter = (*gtk.TextIter)(gextras.NewStructNative(unsafe.Pointer((&_arg3))))
	if _cret != 0 {
		_ok = true
	}

	return _iter, _ok
}

// Match: get whether the provider match the context of completion detailed in
// context.
//
// The function takes the following parameters:
//
//   - context: SourceCompletionContext.
//
// The function returns the following values:
//
//   - ok: TRUE if provider matches the completion context, FALSE otherwise.
//
func (provider *CompletionProvider) match(context *CompletionContext) bool {
	gclass := (*C.GtkSourceCompletionProviderIface)(coreglib.PeekParentClass(provider))
	fnarg := gclass.match

	var _arg0 *C.GtkSourceCompletionProvider // out
	var _arg1 *C.GtkSourceCompletionContext  // out
	var _cret C.gboolean                     // in

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(coreglib.InternObject(provider).Native()))
	_arg1 = (*C.GtkSourceCompletionContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_cret = C._gotk4_gtksource3_CompletionProvider_virtual_match(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(provider)
	runtime.KeepAlive(context)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Populate: populate context with proposals from provider added with the
// gtk_source_completion_context_add_proposals() function.
//
// The function takes the following parameters:
//
//   - context: SourceCompletionContext.
//
func (provider *CompletionProvider) populate(context *CompletionContext) {
	gclass := (*C.GtkSourceCompletionProviderIface)(coreglib.PeekParentClass(provider))
	fnarg := gclass.populate

	var _arg0 *C.GtkSourceCompletionProvider // out
	var _arg1 *C.GtkSourceCompletionContext  // out

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(coreglib.InternObject(provider).Native()))
	_arg1 = (*C.GtkSourceCompletionContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	C._gotk4_gtksource3_CompletionProvider_virtual_populate(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(provider)
	runtime.KeepAlive(context)
}

// updateInfo: update extra information shown in info for proposal.
//
// <note> <para> This function <emphasis>must</emphasis> be implemented when
// gtk_source_completion_provider_get_info_widget() is implemented. </para>
// </note>.
//
// The function takes the following parameters:
//
//   - proposal: SourceCompletionProposal.
//   - info: SourceCompletionInfo.
//
func (provider *CompletionProvider) updateInfo(proposal CompletionProposaller, info *CompletionInfo) {
	gclass := (*C.GtkSourceCompletionProviderIface)(coreglib.PeekParentClass(provider))
	fnarg := gclass.update_info

	var _arg0 *C.GtkSourceCompletionProvider // out
	var _arg1 *C.GtkSourceCompletionProposal // out
	var _arg2 *C.GtkSourceCompletionInfo     // out

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(coreglib.InternObject(provider).Native()))
	_arg1 = (*C.GtkSourceCompletionProposal)(unsafe.Pointer(coreglib.InternObject(proposal).Native()))
	_arg2 = (*C.GtkSourceCompletionInfo)(unsafe.Pointer(coreglib.InternObject(info).Native()))

	C._gotk4_gtksource3_CompletionProvider_virtual_update_info(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2)
	runtime.KeepAlive(provider)
	runtime.KeepAlive(proposal)
	runtime.KeepAlive(info)
}

// CompletionProviderIface: virtual function table for SourceCompletionProvider.
//
// An instance of this type is always passed by reference.
type CompletionProviderIface struct {
	*completionProviderIface
}

// completionProviderIface is the struct that's finalized.
type completionProviderIface struct {
	native *C.GtkSourceCompletionProviderIface
}
