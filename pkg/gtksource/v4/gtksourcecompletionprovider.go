// Code generated by girgen. DO NOT EDIT.

package gtksource

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v3"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtksourceview/gtksource.h>
// extern GIcon* _gotk4_gtksource4_CompletionProviderIface_get_gicon(GtkSourceCompletionProvider*);
// extern GdkPixbuf* _gotk4_gtksource4_CompletionProviderIface_get_icon(GtkSourceCompletionProvider*);
// extern GtkSourceCompletionActivation _gotk4_gtksource4_CompletionProviderIface_get_activation(GtkSourceCompletionProvider*);
// extern GtkWidget* _gotk4_gtksource4_CompletionProviderIface_get_info_widget(GtkSourceCompletionProvider*, GtkSourceCompletionProposal*);
// extern gboolean _gotk4_gtksource4_CompletionProviderIface_activate_proposal(GtkSourceCompletionProvider*, GtkSourceCompletionProposal*, GtkTextIter*);
// extern gboolean _gotk4_gtksource4_CompletionProviderIface_get_start_iter(GtkSourceCompletionProvider*, GtkSourceCompletionContext*, GtkSourceCompletionProposal*, GtkTextIter*);
// extern gboolean _gotk4_gtksource4_CompletionProviderIface_match(GtkSourceCompletionProvider*, GtkSourceCompletionContext*);
// extern gchar* _gotk4_gtksource4_CompletionProviderIface_get_icon_name(GtkSourceCompletionProvider*);
// extern gchar* _gotk4_gtksource4_CompletionProviderIface_get_name(GtkSourceCompletionProvider*);
// extern gint _gotk4_gtksource4_CompletionProviderIface_get_interactive_delay(GtkSourceCompletionProvider*);
// extern gint _gotk4_gtksource4_CompletionProviderIface_get_priority(GtkSourceCompletionProvider*);
// extern void _gotk4_gtksource4_CompletionProviderIface_populate(GtkSourceCompletionProvider*, GtkSourceCompletionContext*);
// extern void _gotk4_gtksource4_CompletionProviderIface_update_info(GtkSourceCompletionProvider*, GtkSourceCompletionProposal*, GtkSourceCompletionInfo*);
import "C"

// glib.Type values for gtksourcecompletionprovider.go.
var GTypeCompletionProvider = externglib.Type(C.gtk_source_completion_provider_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeCompletionProvider, F: marshalCompletionProvider},
	})
}

// CompletionProviderOverrider contains methods that are overridable.
type CompletionProviderOverrider interface {
	// ActivateProposal: activate proposal at iter. When this functions returns
	// FALSE, the default activation of proposal will take place which replaces
	// the word at iter with the text of proposal (see
	// gtk_source_completion_proposal_get_text()).
	//
	// Here is how the default activation selects the boundaries of the word to
	// replace. The end of the word is iter. For the start of the word, it
	// depends on whether a start iter is defined for proposal (see
	// gtk_source_completion_provider_get_start_iter()). If a start iter is
	// defined, the start of the word is the start iter. Else, the word (as long
	// as possible) will contain only alphanumerical and the "_" characters.
	//
	// The function takes the following parameters:
	//
	//    - proposal: SourceCompletionProposal.
	//    - iter: TextIter.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE to indicate that the proposal activation has been handled,
	//      FALSE otherwise.
	//
	ActivateProposal(proposal CompletionProposaller, iter *gtk.TextIter) bool
	// Activation: get with what kind of activation the provider should be
	// activated.
	//
	// The function returns the following values:
	//
	//    - completionActivation: combination of SourceCompletionActivation.
	//
	Activation() CompletionActivation
	// GIcon gets the #GIcon for the icon of provider.
	//
	// The function returns the following values:
	//
	//    - icon (optional) to be used for the provider, or NULL if the provider
	//      does not have a special icon.
	//
	GIcon() gio.Iconner
	// Icon: get the Pixbuf for the icon of the provider.
	//
	// The function returns the following values:
	//
	//    - pixbuf (optional): icon to be used for the provider, or NULL if the
	//      provider does not have a special icon.
	//
	Icon() *gdkpixbuf.Pixbuf
	// IconName gets the icon name of provider.
	//
	// The function returns the following values:
	//
	//    - utf8 (optional): icon name to be used for the provider, or NULL if
	//      the provider does not have a special icon.
	//
	IconName() string
	// InfoWidget: get a customized info widget to show extra information of a
	// proposal. This allows for customized widgets on a proposal basis,
	// although in general providers will have the same custom widget for all
	// their proposals and proposal can be ignored. The implementation of this
	// function is optional.
	//
	// If this function is not implemented, the default widget is a Label. The
	// return value of gtk_source_completion_proposal_get_info() is used as the
	// content of the Label.
	//
	// <note> <para> If implemented,
	// gtk_source_completion_provider_update_info() <emphasis>must</emphasis>
	// also be implemented. </para> </note>.
	//
	// The function takes the following parameters:
	//
	//    - proposal: currently selected SourceCompletionProposal.
	//
	// The function returns the following values:
	//
	//    - widget (optional): custom Widget to show extra information about
	//      proposal, or NULL if the provider does not have a special info
	//      widget.
	//
	InfoWidget(proposal CompletionProposaller) gtk.Widgetter
	// InteractiveDelay: get the delay in milliseconds before starting
	// interactive completion for this provider. A value of -1 indicates to use
	// the default value as set by the SourceCompletion:auto-complete-delay
	// property.
	//
	// The function returns the following values:
	//
	//    - gint: interactive delay in milliseconds.
	//
	InteractiveDelay() int
	// Name: get the name of the provider. This should be a translatable name
	// for display to the user. For example: _("Document word completion
	// provider"). The returned string must be freed with g_free().
	//
	// The function returns the following values:
	//
	//    - utf8: new string containing the name of the provider.
	//
	Name() string
	// Priority: get the provider priority. The priority determines the order in
	// which proposals appear in the completion popup. Higher priorities are
	// sorted before lower priorities. The default priority is 0.
	//
	// The function returns the following values:
	//
	//    - gint: provider priority.
	//
	Priority() int
	// StartIter: get the TextIter at which the completion for proposal starts.
	// When implemented, this information is used to position the completion
	// window accordingly when a proposal is selected in the completion window.
	// The proposal text inside the completion window is aligned on iter.
	//
	// If this function is not implemented, the word boundary is taken to
	// position the completion window. See
	// gtk_source_completion_provider_activate_proposal() for an explanation on
	// the word boundaries.
	//
	// When the proposal is activated, the default handler uses iter as the
	// start of the word to replace. See
	// gtk_source_completion_provider_activate_proposal() for more information.
	//
	// The function takes the following parameters:
	//
	//    - context: SourceCompletionContext.
	//    - proposal: SourceCompletionProposal.
	//
	// The function returns the following values:
	//
	//    - iter: TextIter.
	//    - ok: TRUE if iter was set for proposal, FALSE otherwise.
	//
	StartIter(context *CompletionContext, proposal CompletionProposaller) (*gtk.TextIter, bool)
	// Match: get whether the provider match the context of completion detailed
	// in context.
	//
	// The function takes the following parameters:
	//
	//    - context: SourceCompletionContext.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if provider matches the completion context, FALSE otherwise.
	//
	Match(context *CompletionContext) bool
	// Populate context with proposals from provider added with the
	// gtk_source_completion_context_add_proposals() function.
	//
	// The function takes the following parameters:
	//
	//    - context: SourceCompletionContext.
	//
	Populate(context *CompletionContext)
	// UpdateInfo: update extra information shown in info for proposal.
	//
	// <note> <para> This function <emphasis>must</emphasis> be implemented when
	// gtk_source_completion_provider_get_info_widget() is implemented. </para>
	// </note>.
	//
	// The function takes the following parameters:
	//
	//    - proposal: SourceCompletionProposal.
	//    - info: SourceCompletionInfo.
	//
	UpdateInfo(proposal CompletionProposaller, info *CompletionInfo)
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

	// ActivateProposal: activate proposal at iter.
	ActivateProposal(proposal CompletionProposaller, iter *gtk.TextIter) bool
	// Activation: get with what kind of activation the provider should be
	// activated.
	Activation() CompletionActivation
	// GIcon gets the #GIcon for the icon of provider.
	GIcon() gio.Iconner
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

func ifaceInitCompletionProviderer(gifacePtr, data C.gpointer) {
	iface := (*C.GtkSourceCompletionProviderIface)(unsafe.Pointer(gifacePtr))
	iface.activate_proposal = (*[0]byte)(C._gotk4_gtksource4_CompletionProviderIface_activate_proposal)
	iface.get_activation = (*[0]byte)(C._gotk4_gtksource4_CompletionProviderIface_get_activation)
	iface.get_gicon = (*[0]byte)(C._gotk4_gtksource4_CompletionProviderIface_get_gicon)
	iface.get_icon = (*[0]byte)(C._gotk4_gtksource4_CompletionProviderIface_get_icon)
	iface.get_icon_name = (*[0]byte)(C._gotk4_gtksource4_CompletionProviderIface_get_icon_name)
	iface.get_info_widget = (*[0]byte)(C._gotk4_gtksource4_CompletionProviderIface_get_info_widget)
	iface.get_interactive_delay = (*[0]byte)(C._gotk4_gtksource4_CompletionProviderIface_get_interactive_delay)
	iface.get_name = (*[0]byte)(C._gotk4_gtksource4_CompletionProviderIface_get_name)
	iface.get_priority = (*[0]byte)(C._gotk4_gtksource4_CompletionProviderIface_get_priority)
	iface.get_start_iter = (*[0]byte)(C._gotk4_gtksource4_CompletionProviderIface_get_start_iter)
	iface.match = (*[0]byte)(C._gotk4_gtksource4_CompletionProviderIface_match)
	iface.populate = (*[0]byte)(C._gotk4_gtksource4_CompletionProviderIface_populate)
	iface.update_info = (*[0]byte)(C._gotk4_gtksource4_CompletionProviderIface_update_info)
}

//export _gotk4_gtksource4_CompletionProviderIface_activate_proposal
func _gotk4_gtksource4_CompletionProviderIface_activate_proposal(arg0 *C.GtkSourceCompletionProvider, arg1 *C.GtkSourceCompletionProposal, arg2 *C.GtkTextIter) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(CompletionProviderOverrider)

	var _proposal CompletionProposaller // out
	var _iter *gtk.TextIter             // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtksource.CompletionProposaller is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(CompletionProposaller)
			return ok
		})
		rv, ok := casted.(CompletionProposaller)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtksource.CompletionProposaller")
		}
		_proposal = rv
	}
	_iter = (*gtk.TextIter)(gextras.NewStructNative(unsafe.Pointer(arg2)))

	ok := iface.ActivateProposal(_proposal, _iter)

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtksource4_CompletionProviderIface_get_activation
func _gotk4_gtksource4_CompletionProviderIface_get_activation(arg0 *C.GtkSourceCompletionProvider) (cret C.GtkSourceCompletionActivation) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(CompletionProviderOverrider)

	completionActivation := iface.Activation()

	cret = C.GtkSourceCompletionActivation(completionActivation)

	return cret
}

//export _gotk4_gtksource4_CompletionProviderIface_get_gicon
func _gotk4_gtksource4_CompletionProviderIface_get_gicon(arg0 *C.GtkSourceCompletionProvider) (cret *C.GIcon) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(CompletionProviderOverrider)

	icon := iface.GIcon()

	if icon != nil {
		cret = (*C.GIcon)(unsafe.Pointer(externglib.InternObject(icon).Native()))
	}

	return cret
}

//export _gotk4_gtksource4_CompletionProviderIface_get_icon
func _gotk4_gtksource4_CompletionProviderIface_get_icon(arg0 *C.GtkSourceCompletionProvider) (cret *C.GdkPixbuf) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(CompletionProviderOverrider)

	pixbuf := iface.Icon()

	if pixbuf != nil {
		cret = (*C.GdkPixbuf)(unsafe.Pointer(externglib.InternObject(pixbuf).Native()))
	}

	return cret
}

//export _gotk4_gtksource4_CompletionProviderIface_get_icon_name
func _gotk4_gtksource4_CompletionProviderIface_get_icon_name(arg0 *C.GtkSourceCompletionProvider) (cret *C.gchar) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(CompletionProviderOverrider)

	utf8 := iface.IconName()

	if utf8 != "" {
		cret = (*C.gchar)(unsafe.Pointer(C.CString(utf8)))
		defer C.free(unsafe.Pointer(cret))
	}

	return cret
}

//export _gotk4_gtksource4_CompletionProviderIface_get_info_widget
func _gotk4_gtksource4_CompletionProviderIface_get_info_widget(arg0 *C.GtkSourceCompletionProvider, arg1 *C.GtkSourceCompletionProposal) (cret *C.GtkWidget) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(CompletionProviderOverrider)

	var _proposal CompletionProposaller // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtksource.CompletionProposaller is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(CompletionProposaller)
			return ok
		})
		rv, ok := casted.(CompletionProposaller)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtksource.CompletionProposaller")
		}
		_proposal = rv
	}

	widget := iface.InfoWidget(_proposal)

	if widget != nil {
		cret = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(widget).Native()))
	}

	return cret
}

//export _gotk4_gtksource4_CompletionProviderIface_get_interactive_delay
func _gotk4_gtksource4_CompletionProviderIface_get_interactive_delay(arg0 *C.GtkSourceCompletionProvider) (cret C.gint) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(CompletionProviderOverrider)

	gint := iface.InteractiveDelay()

	cret = C.gint(gint)

	return cret
}

//export _gotk4_gtksource4_CompletionProviderIface_get_name
func _gotk4_gtksource4_CompletionProviderIface_get_name(arg0 *C.GtkSourceCompletionProvider) (cret *C.gchar) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(CompletionProviderOverrider)

	utf8 := iface.Name()

	cret = (*C.gchar)(unsafe.Pointer(C.CString(utf8)))

	return cret
}

//export _gotk4_gtksource4_CompletionProviderIface_get_priority
func _gotk4_gtksource4_CompletionProviderIface_get_priority(arg0 *C.GtkSourceCompletionProvider) (cret C.gint) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(CompletionProviderOverrider)

	gint := iface.Priority()

	cret = C.gint(gint)

	return cret
}

//export _gotk4_gtksource4_CompletionProviderIface_get_start_iter
func _gotk4_gtksource4_CompletionProviderIface_get_start_iter(arg0 *C.GtkSourceCompletionProvider, arg1 *C.GtkSourceCompletionContext, arg2 *C.GtkSourceCompletionProposal, arg3 *C.GtkTextIter) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(CompletionProviderOverrider)

	var _context *CompletionContext     // out
	var _proposal CompletionProposaller // out

	_context = wrapCompletionContext(externglib.Take(unsafe.Pointer(arg1)))
	{
		objptr := unsafe.Pointer(arg2)
		if objptr == nil {
			panic("object of type gtksource.CompletionProposaller is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(CompletionProposaller)
			return ok
		})
		rv, ok := casted.(CompletionProposaller)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtksource.CompletionProposaller")
		}
		_proposal = rv
	}

	iter, ok := iface.StartIter(_context, _proposal)

	*arg3 = *(*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(iter)))
	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtksource4_CompletionProviderIface_match
func _gotk4_gtksource4_CompletionProviderIface_match(arg0 *C.GtkSourceCompletionProvider, arg1 *C.GtkSourceCompletionContext) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(CompletionProviderOverrider)

	var _context *CompletionContext // out

	_context = wrapCompletionContext(externglib.Take(unsafe.Pointer(arg1)))

	ok := iface.Match(_context)

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtksource4_CompletionProviderIface_populate
func _gotk4_gtksource4_CompletionProviderIface_populate(arg0 *C.GtkSourceCompletionProvider, arg1 *C.GtkSourceCompletionContext) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(CompletionProviderOverrider)

	var _context *CompletionContext // out

	_context = wrapCompletionContext(externglib.Take(unsafe.Pointer(arg1)))

	iface.Populate(_context)
}

//export _gotk4_gtksource4_CompletionProviderIface_update_info
func _gotk4_gtksource4_CompletionProviderIface_update_info(arg0 *C.GtkSourceCompletionProvider, arg1 *C.GtkSourceCompletionProposal, arg2 *C.GtkSourceCompletionInfo) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(CompletionProviderOverrider)

	var _proposal CompletionProposaller // out
	var _info *CompletionInfo           // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtksource.CompletionProposaller is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(CompletionProposaller)
			return ok
		})
		rv, ok := casted.(CompletionProposaller)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtksource.CompletionProposaller")
		}
		_proposal = rv
	}
	_info = wrapCompletionInfo(externglib.Take(unsafe.Pointer(arg2)))

	iface.UpdateInfo(_proposal, _info)
}

func wrapCompletionProvider(obj *externglib.Object) *CompletionProvider {
	return &CompletionProvider{
		Object: obj,
	}
}

func marshalCompletionProvider(p uintptr) (interface{}, error) {
	return wrapCompletionProvider(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ActivateProposal: activate proposal at iter. When this functions returns
// FALSE, the default activation of proposal will take place which replaces the
// word at iter with the text of proposal (see
// gtk_source_completion_proposal_get_text()).
//
// Here is how the default activation selects the boundaries of the word to
// replace. The end of the word is iter. For the start of the word, it depends
// on whether a start iter is defined for proposal (see
// gtk_source_completion_provider_get_start_iter()). If a start iter is defined,
// the start of the word is the start iter. Else, the word (as long as possible)
// will contain only alphanumerical and the "_" characters.
//
// The function takes the following parameters:
//
//    - proposal: SourceCompletionProposal.
//    - iter: TextIter.
//
// The function returns the following values:
//
//    - ok: TRUE to indicate that the proposal activation has been handled, FALSE
//      otherwise.
//
func (provider *CompletionProvider) ActivateProposal(proposal CompletionProposaller, iter *gtk.TextIter) bool {
	var _arg0 *C.GtkSourceCompletionProvider // out
	var _arg1 *C.GtkSourceCompletionProposal // out
	var _arg2 *C.GtkTextIter                 // out
	var _cret C.gboolean                     // in

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(externglib.InternObject(provider).Native()))
	_arg1 = (*C.GtkSourceCompletionProposal)(unsafe.Pointer(externglib.InternObject(proposal).Native()))
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
//    - completionActivation: combination of SourceCompletionActivation.
//
func (provider *CompletionProvider) Activation() CompletionActivation {
	var _arg0 *C.GtkSourceCompletionProvider  // out
	var _cret C.GtkSourceCompletionActivation // in

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(externglib.InternObject(provider).Native()))

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
//    - icon (optional) to be used for the provider, or NULL if the provider does
//      not have a special icon.
//
func (provider *CompletionProvider) GIcon() gio.Iconner {
	var _arg0 *C.GtkSourceCompletionProvider // out
	var _cret *C.GIcon                       // in

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(externglib.InternObject(provider).Native()))

	_cret = C.gtk_source_completion_provider_get_gicon(_arg0)
	runtime.KeepAlive(provider)

	var _icon gio.Iconner // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(gio.Iconner)
				return ok
			})
			rv, ok := casted.(gio.Iconner)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.Iconner")
			}
			_icon = rv
		}
	}

	return _icon
}

// Icon: get the Pixbuf for the icon of the provider.
//
// The function returns the following values:
//
//    - pixbuf (optional): icon to be used for the provider, or NULL if the
//      provider does not have a special icon.
//
func (provider *CompletionProvider) Icon() *gdkpixbuf.Pixbuf {
	var _arg0 *C.GtkSourceCompletionProvider // out
	var _cret *C.GdkPixbuf                   // in

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(externglib.InternObject(provider).Native()))

	_cret = C.gtk_source_completion_provider_get_icon(_arg0)
	runtime.KeepAlive(provider)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	if _cret != nil {
		{
			obj := externglib.Take(unsafe.Pointer(_cret))
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
//    - utf8 (optional): icon name to be used for the provider, or NULL if the
//      provider does not have a special icon.
//
func (provider *CompletionProvider) IconName() string {
	var _arg0 *C.GtkSourceCompletionProvider // out
	var _cret *C.gchar                       // in

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(externglib.InternObject(provider).Native()))

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
// If this function is not implemented, the default widget is a Label. The
// return value of gtk_source_completion_proposal_get_info() is used as the
// content of the Label.
//
// <note> <para> If implemented, gtk_source_completion_provider_update_info()
// <emphasis>must</emphasis> also be implemented. </para> </note>.
//
// The function takes the following parameters:
//
//    - proposal: currently selected SourceCompletionProposal.
//
// The function returns the following values:
//
//    - widget (optional): custom Widget to show extra information about
//      proposal, or NULL if the provider does not have a special info widget.
//
func (provider *CompletionProvider) InfoWidget(proposal CompletionProposaller) gtk.Widgetter {
	var _arg0 *C.GtkSourceCompletionProvider // out
	var _arg1 *C.GtkSourceCompletionProposal // out
	var _cret *C.GtkWidget                   // in

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(externglib.InternObject(provider).Native()))
	_arg1 = (*C.GtkSourceCompletionProposal)(unsafe.Pointer(externglib.InternObject(proposal).Native()))

	_cret = C.gtk_source_completion_provider_get_info_widget(_arg0, _arg1)
	runtime.KeepAlive(provider)
	runtime.KeepAlive(proposal)

	var _widget gtk.Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
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
//    - gint: interactive delay in milliseconds.
//
func (provider *CompletionProvider) InteractiveDelay() int {
	var _arg0 *C.GtkSourceCompletionProvider // out
	var _cret C.gint                         // in

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(externglib.InternObject(provider).Native()))

	_cret = C.gtk_source_completion_provider_get_interactive_delay(_arg0)
	runtime.KeepAlive(provider)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Name: get the name of the provider. This should be a translatable name for
// display to the user. For example: _("Document word completion provider"). The
// returned string must be freed with g_free().
//
// The function returns the following values:
//
//    - utf8: new string containing the name of the provider.
//
func (provider *CompletionProvider) Name() string {
	var _arg0 *C.GtkSourceCompletionProvider // out
	var _cret *C.gchar                       // in

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(externglib.InternObject(provider).Native()))

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
//    - gint: provider priority.
//
func (provider *CompletionProvider) Priority() int {
	var _arg0 *C.GtkSourceCompletionProvider // out
	var _cret C.gint                         // in

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(externglib.InternObject(provider).Native()))

	_cret = C.gtk_source_completion_provider_get_priority(_arg0)
	runtime.KeepAlive(provider)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// StartIter: get the TextIter at which the completion for proposal starts. When
// implemented, this information is used to position the completion window
// accordingly when a proposal is selected in the completion window. The
// proposal text inside the completion window is aligned on iter.
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
//    - context: SourceCompletionContext.
//    - proposal: SourceCompletionProposal.
//
// The function returns the following values:
//
//    - iter: TextIter.
//    - ok: TRUE if iter was set for proposal, FALSE otherwise.
//
func (provider *CompletionProvider) StartIter(context *CompletionContext, proposal CompletionProposaller) (*gtk.TextIter, bool) {
	var _arg0 *C.GtkSourceCompletionProvider // out
	var _arg1 *C.GtkSourceCompletionContext  // out
	var _arg2 *C.GtkSourceCompletionProposal // out
	var _arg3 C.GtkTextIter                  // in
	var _cret C.gboolean                     // in

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(externglib.InternObject(provider).Native()))
	_arg1 = (*C.GtkSourceCompletionContext)(unsafe.Pointer(externglib.InternObject(context).Native()))
	_arg2 = (*C.GtkSourceCompletionProposal)(unsafe.Pointer(externglib.InternObject(proposal).Native()))

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
//    - context: SourceCompletionContext.
//
// The function returns the following values:
//
//    - ok: TRUE if provider matches the completion context, FALSE otherwise.
//
func (provider *CompletionProvider) Match(context *CompletionContext) bool {
	var _arg0 *C.GtkSourceCompletionProvider // out
	var _arg1 *C.GtkSourceCompletionContext  // out
	var _cret C.gboolean                     // in

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(externglib.InternObject(provider).Native()))
	_arg1 = (*C.GtkSourceCompletionContext)(unsafe.Pointer(externglib.InternObject(context).Native()))

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
//    - context: SourceCompletionContext.
//
func (provider *CompletionProvider) Populate(context *CompletionContext) {
	var _arg0 *C.GtkSourceCompletionProvider // out
	var _arg1 *C.GtkSourceCompletionContext  // out

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(externglib.InternObject(provider).Native()))
	_arg1 = (*C.GtkSourceCompletionContext)(unsafe.Pointer(externglib.InternObject(context).Native()))

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
//    - proposal: SourceCompletionProposal.
//    - info: SourceCompletionInfo.
//
func (provider *CompletionProvider) UpdateInfo(proposal CompletionProposaller, info *CompletionInfo) {
	var _arg0 *C.GtkSourceCompletionProvider // out
	var _arg1 *C.GtkSourceCompletionProposal // out
	var _arg2 *C.GtkSourceCompletionInfo     // out

	_arg0 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(externglib.InternObject(provider).Native()))
	_arg1 = (*C.GtkSourceCompletionProposal)(unsafe.Pointer(externglib.InternObject(proposal).Native()))
	_arg2 = (*C.GtkSourceCompletionInfo)(unsafe.Pointer(externglib.InternObject(info).Native()))

	C.gtk_source_completion_provider_update_info(_arg0, _arg1, _arg2)
	runtime.KeepAlive(provider)
	runtime.KeepAlive(proposal)
	runtime.KeepAlive(info)
}
