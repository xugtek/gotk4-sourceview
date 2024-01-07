// Code generated by girgen. DO NOT EDIT.

package gtksource

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtksourceview/gtksource.h>
// extern void _gotk4_gtksource3_CompletionProposal_ConnectChanged(gpointer, guintptr);
// GIcon* _gotk4_gtksource3_CompletionProposal_virtual_get_gicon(void* fnptr, GtkSourceCompletionProposal* arg0) {
//   return ((GIcon* (*)(GtkSourceCompletionProposal*))(fnptr))(arg0);
// };
// GdkPixbuf* _gotk4_gtksource3_CompletionProposal_virtual_get_icon(void* fnptr, GtkSourceCompletionProposal* arg0) {
//   return ((GdkPixbuf* (*)(GtkSourceCompletionProposal*))(fnptr))(arg0);
// };
// gboolean _gotk4_gtksource3_CompletionProposal_virtual_equal(void* fnptr, GtkSourceCompletionProposal* arg0, GtkSourceCompletionProposal* arg1) {
//   return ((gboolean (*)(GtkSourceCompletionProposal*, GtkSourceCompletionProposal*))(fnptr))(arg0, arg1);
// };
// gchar* _gotk4_gtksource3_CompletionProposal_virtual_get_icon_name(void* fnptr, GtkSourceCompletionProposal* arg0) {
//   return ((gchar* (*)(GtkSourceCompletionProposal*))(fnptr))(arg0);
// };
// gchar* _gotk4_gtksource3_CompletionProposal_virtual_get_info(void* fnptr, GtkSourceCompletionProposal* arg0) {
//   return ((gchar* (*)(GtkSourceCompletionProposal*))(fnptr))(arg0);
// };
// gchar* _gotk4_gtksource3_CompletionProposal_virtual_get_label(void* fnptr, GtkSourceCompletionProposal* arg0) {
//   return ((gchar* (*)(GtkSourceCompletionProposal*))(fnptr))(arg0);
// };
// gchar* _gotk4_gtksource3_CompletionProposal_virtual_get_markup(void* fnptr, GtkSourceCompletionProposal* arg0) {
//   return ((gchar* (*)(GtkSourceCompletionProposal*))(fnptr))(arg0);
// };
// gchar* _gotk4_gtksource3_CompletionProposal_virtual_get_text(void* fnptr, GtkSourceCompletionProposal* arg0) {
//   return ((gchar* (*)(GtkSourceCompletionProposal*))(fnptr))(arg0);
// };
// guint _gotk4_gtksource3_CompletionProposal_virtual_hash(void* fnptr, GtkSourceCompletionProposal* arg0) {
//   return ((guint (*)(GtkSourceCompletionProposal*))(fnptr))(arg0);
// };
// void _gotk4_gtksource3_CompletionProposal_virtual_changed(void* fnptr, GtkSourceCompletionProposal* arg0) {
//   ((void (*)(GtkSourceCompletionProposal*))(fnptr))(arg0);
// };
import "C"

// GType values.
var (
	GTypeCompletionProposal = coreglib.Type(C.gtk_source_completion_proposal_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeCompletionProposal, F: marshalCompletionProposal},
	})
}

//
// CompletionProposal wraps an interface. This means the user can get the
// underlying type by calling Cast().
type CompletionProposal struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*CompletionProposal)(nil)
)

// CompletionProposaller describes CompletionProposal's interface methods.
type CompletionProposaller interface {
	coreglib.Objector

	// Changed emits the "changed" signal on proposal.
	Changed()
	// Equal: get whether two proposal objects are the same.
	Equal(other CompletionProposaller) bool
	// GIcon gets the #GIcon for the icon of proposal.
	GIcon() *gio.Icon
	// Icon gets the Pixbuf for the icon of proposal.
	Icon() *gdkpixbuf.Pixbuf
	// IconName gets the icon name of proposal.
	IconName() string
	// Info gets extra information associated to the proposal.
	Info() string
	// Label gets the label of proposal.
	Label() string
	// Markup gets the label of proposal with markup.
	Markup() string
	// Text gets the text of proposal.
	Text() string
	// Hash: get the hash value of proposal.
	Hash() uint

	// Changed is emitted when the proposal has changed.
	ConnectChanged(func()) coreglib.SignalHandle
}

var _ CompletionProposaller = (*CompletionProposal)(nil)

func wrapCompletionProposal(obj *coreglib.Object) *CompletionProposal {
	return &CompletionProposal{
		Object: obj,
	}
}

func marshalCompletionProposal(p uintptr) (interface{}, error) {
	return wrapCompletionProposal(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectChanged is emitted when the proposal has changed. The completion popup
// will react to this by updating the shown information.
func (proposal *CompletionProposal) ConnectChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(proposal, "changed", false, unsafe.Pointer(C._gotk4_gtksource3_CompletionProposal_ConnectChanged), f)
}

// Changed emits the "changed" signal on proposal. This should be called by
// implementations whenever the name, icon or info of the proposal has changed.
func (proposal *CompletionProposal) Changed() {
	var _arg0 *C.GtkSourceCompletionProposal // out

	_arg0 = (*C.GtkSourceCompletionProposal)(unsafe.Pointer(coreglib.InternObject(proposal).Native()))

	C.gtk_source_completion_proposal_changed(_arg0)
	runtime.KeepAlive(proposal)
}

// Equal: get whether two proposal objects are the same. This is used to
// (together with gtk_source_completion_proposal_hash()) to match proposals in
// the completion model. By default, it uses direct equality (g_direct_equal()).
//
// The function takes the following parameters:
//
//   - other: SourceCompletionProposal.
//
// The function returns the following values:
//
//   - ok: TRUE if proposal and object are the same proposal.
//
func (proposal *CompletionProposal) Equal(other CompletionProposaller) bool {
	var _arg0 *C.GtkSourceCompletionProposal // out
	var _arg1 *C.GtkSourceCompletionProposal // out
	var _cret C.gboolean                     // in

	_arg0 = (*C.GtkSourceCompletionProposal)(unsafe.Pointer(coreglib.InternObject(proposal).Native()))
	_arg1 = (*C.GtkSourceCompletionProposal)(unsafe.Pointer(coreglib.InternObject(other).Native()))

	_cret = C.gtk_source_completion_proposal_equal(_arg0, _arg1)
	runtime.KeepAlive(proposal)
	runtime.KeepAlive(other)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// GIcon gets the #GIcon for the icon of proposal.
//
// The function returns the following values:
//
//   - icon (optional) with the icon of proposal.
//
func (proposal *CompletionProposal) GIcon() *gio.Icon {
	var _arg0 *C.GtkSourceCompletionProposal // out
	var _cret *C.GIcon                       // in

	_arg0 = (*C.GtkSourceCompletionProposal)(unsafe.Pointer(coreglib.InternObject(proposal).Native()))

	_cret = C.gtk_source_completion_proposal_get_gicon(_arg0)
	runtime.KeepAlive(proposal)

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

// Icon gets the Pixbuf for the icon of proposal.
//
// The function returns the following values:
//
//   - pixbuf (optional) with the icon of proposal.
//
func (proposal *CompletionProposal) Icon() *gdkpixbuf.Pixbuf {
	var _arg0 *C.GtkSourceCompletionProposal // out
	var _cret *C.GdkPixbuf                   // in

	_arg0 = (*C.GtkSourceCompletionProposal)(unsafe.Pointer(coreglib.InternObject(proposal).Native()))

	_cret = C.gtk_source_completion_proposal_get_icon(_arg0)
	runtime.KeepAlive(proposal)

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

// IconName gets the icon name of proposal.
//
// The function returns the following values:
//
//   - utf8 (optional): icon name of proposal.
//
func (proposal *CompletionProposal) IconName() string {
	var _arg0 *C.GtkSourceCompletionProposal // out
	var _cret *C.gchar                       // in

	_arg0 = (*C.GtkSourceCompletionProposal)(unsafe.Pointer(coreglib.InternObject(proposal).Native()))

	_cret = C.gtk_source_completion_proposal_get_icon_name(_arg0)
	runtime.KeepAlive(proposal)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Info gets extra information associated to the proposal. This information
// will be used to present the user with extra, detailed information about the
// selected proposal. The returned string must be freed with g_free().
//
// The function returns the following values:
//
//   - utf8 (optional): newly-allocated string containing extra information of
//     proposal or NULL if no extra information is associated to proposal.
//
func (proposal *CompletionProposal) Info() string {
	var _arg0 *C.GtkSourceCompletionProposal // out
	var _cret *C.gchar                       // in

	_arg0 = (*C.GtkSourceCompletionProposal)(unsafe.Pointer(coreglib.InternObject(proposal).Native()))

	_cret = C.gtk_source_completion_proposal_get_info(_arg0)
	runtime.KeepAlive(proposal)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// Label gets the label of proposal. The label is shown in the list of proposals
// as plain text. If you need any markup (such as bold or italic text), you have
// to implement gtk_source_completion_proposal_get_markup(). The returned string
// must be freed with g_free().
//
// The function returns the following values:
//
//   - utf8: new string containing the label of proposal.
//
func (proposal *CompletionProposal) Label() string {
	var _arg0 *C.GtkSourceCompletionProposal // out
	var _cret *C.gchar                       // in

	_arg0 = (*C.GtkSourceCompletionProposal)(unsafe.Pointer(coreglib.InternObject(proposal).Native()))

	_cret = C.gtk_source_completion_proposal_get_label(_arg0)
	runtime.KeepAlive(proposal)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Markup gets the label of proposal with markup. The label is shown in the
// list of proposals and may contain markup. This will be used instead of
// gtk_source_completion_proposal_get_label() if implemented. The returned
// string must be freed with g_free().
//
// The function returns the following values:
//
//   - utf8: new string containing the label of proposal with markup.
//
func (proposal *CompletionProposal) Markup() string {
	var _arg0 *C.GtkSourceCompletionProposal // out
	var _cret *C.gchar                       // in

	_arg0 = (*C.GtkSourceCompletionProposal)(unsafe.Pointer(coreglib.InternObject(proposal).Native()))

	_cret = C.gtk_source_completion_proposal_get_markup(_arg0)
	runtime.KeepAlive(proposal)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Text gets the text of proposal. The text that is inserted into the
// text buffer when the proposal is activated by the default activation.
// You are free to implement a custom activation handler in the
// provider and not implement this function. For more information, see
// gtk_source_completion_provider_activate_proposal(). The returned string must
// be freed with g_free().
//
// The function returns the following values:
//
//   - utf8: new string containing the text of proposal.
//
func (proposal *CompletionProposal) Text() string {
	var _arg0 *C.GtkSourceCompletionProposal // out
	var _cret *C.gchar                       // in

	_arg0 = (*C.GtkSourceCompletionProposal)(unsafe.Pointer(coreglib.InternObject(proposal).Native()))

	_cret = C.gtk_source_completion_proposal_get_text(_arg0)
	runtime.KeepAlive(proposal)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Hash: get the hash value of proposal. This is used to (together with
// gtk_source_completion_proposal_equal()) to match proposals in the completion
// model. By default, it uses a direct hash (g_direct_hash()).
//
// The function returns the following values:
//
//   - guint: hash value of proposal.
//
func (proposal *CompletionProposal) Hash() uint {
	var _arg0 *C.GtkSourceCompletionProposal // out
	var _cret C.guint                        // in

	_arg0 = (*C.GtkSourceCompletionProposal)(unsafe.Pointer(coreglib.InternObject(proposal).Native()))

	_cret = C.gtk_source_completion_proposal_hash(_arg0)
	runtime.KeepAlive(proposal)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Changed emits the "changed" signal on proposal. This should be called by
// implementations whenever the name, icon or info of the proposal has changed.
func (proposal *CompletionProposal) changed() {
	gclass := (*C.GtkSourceCompletionProposalIface)(coreglib.PeekParentClass(proposal))
	fnarg := gclass.changed

	var _arg0 *C.GtkSourceCompletionProposal // out

	_arg0 = (*C.GtkSourceCompletionProposal)(unsafe.Pointer(coreglib.InternObject(proposal).Native()))

	C._gotk4_gtksource3_CompletionProposal_virtual_changed(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(proposal)
}

// Equal: get whether two proposal objects are the same. This is used to
// (together with gtk_source_completion_proposal_hash()) to match proposals in
// the completion model. By default, it uses direct equality (g_direct_equal()).
//
// The function takes the following parameters:
//
//   - other: SourceCompletionProposal.
//
// The function returns the following values:
//
//   - ok: TRUE if proposal and object are the same proposal.
//
func (proposal *CompletionProposal) equal(other CompletionProposaller) bool {
	gclass := (*C.GtkSourceCompletionProposalIface)(coreglib.PeekParentClass(proposal))
	fnarg := gclass.equal

	var _arg0 *C.GtkSourceCompletionProposal // out
	var _arg1 *C.GtkSourceCompletionProposal // out
	var _cret C.gboolean                     // in

	_arg0 = (*C.GtkSourceCompletionProposal)(unsafe.Pointer(coreglib.InternObject(proposal).Native()))
	_arg1 = (*C.GtkSourceCompletionProposal)(unsafe.Pointer(coreglib.InternObject(other).Native()))

	_cret = C._gotk4_gtksource3_CompletionProposal_virtual_equal(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(proposal)
	runtime.KeepAlive(other)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// gIcon gets the #GIcon for the icon of proposal.
//
// The function returns the following values:
//
//   - icon (optional) with the icon of proposal.
//
func (proposal *CompletionProposal) gIcon() *gio.Icon {
	gclass := (*C.GtkSourceCompletionProposalIface)(coreglib.PeekParentClass(proposal))
	fnarg := gclass.get_gicon

	var _arg0 *C.GtkSourceCompletionProposal // out
	var _cret *C.GIcon                       // in

	_arg0 = (*C.GtkSourceCompletionProposal)(unsafe.Pointer(coreglib.InternObject(proposal).Native()))

	_cret = C._gotk4_gtksource3_CompletionProposal_virtual_get_gicon(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(proposal)

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

// Icon gets the Pixbuf for the icon of proposal.
//
// The function returns the following values:
//
//   - pixbuf (optional) with the icon of proposal.
//
func (proposal *CompletionProposal) icon() *gdkpixbuf.Pixbuf {
	gclass := (*C.GtkSourceCompletionProposalIface)(coreglib.PeekParentClass(proposal))
	fnarg := gclass.get_icon

	var _arg0 *C.GtkSourceCompletionProposal // out
	var _cret *C.GdkPixbuf                   // in

	_arg0 = (*C.GtkSourceCompletionProposal)(unsafe.Pointer(coreglib.InternObject(proposal).Native()))

	_cret = C._gotk4_gtksource3_CompletionProposal_virtual_get_icon(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(proposal)

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

// iconName gets the icon name of proposal.
//
// The function returns the following values:
//
//   - utf8 (optional): icon name of proposal.
//
func (proposal *CompletionProposal) iconName() string {
	gclass := (*C.GtkSourceCompletionProposalIface)(coreglib.PeekParentClass(proposal))
	fnarg := gclass.get_icon_name

	var _arg0 *C.GtkSourceCompletionProposal // out
	var _cret *C.gchar                       // in

	_arg0 = (*C.GtkSourceCompletionProposal)(unsafe.Pointer(coreglib.InternObject(proposal).Native()))

	_cret = C._gotk4_gtksource3_CompletionProposal_virtual_get_icon_name(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(proposal)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Info gets extra information associated to the proposal. This information
// will be used to present the user with extra, detailed information about the
// selected proposal. The returned string must be freed with g_free().
//
// The function returns the following values:
//
//   - utf8 (optional): newly-allocated string containing extra information of
//     proposal or NULL if no extra information is associated to proposal.
//
func (proposal *CompletionProposal) info() string {
	gclass := (*C.GtkSourceCompletionProposalIface)(coreglib.PeekParentClass(proposal))
	fnarg := gclass.get_info

	var _arg0 *C.GtkSourceCompletionProposal // out
	var _cret *C.gchar                       // in

	_arg0 = (*C.GtkSourceCompletionProposal)(unsafe.Pointer(coreglib.InternObject(proposal).Native()))

	_cret = C._gotk4_gtksource3_CompletionProposal_virtual_get_info(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(proposal)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// Label gets the label of proposal. The label is shown in the list of proposals
// as plain text. If you need any markup (such as bold or italic text), you have
// to implement gtk_source_completion_proposal_get_markup(). The returned string
// must be freed with g_free().
//
// The function returns the following values:
//
//   - utf8: new string containing the label of proposal.
//
func (proposal *CompletionProposal) label() string {
	gclass := (*C.GtkSourceCompletionProposalIface)(coreglib.PeekParentClass(proposal))
	fnarg := gclass.get_label

	var _arg0 *C.GtkSourceCompletionProposal // out
	var _cret *C.gchar                       // in

	_arg0 = (*C.GtkSourceCompletionProposal)(unsafe.Pointer(coreglib.InternObject(proposal).Native()))

	_cret = C._gotk4_gtksource3_CompletionProposal_virtual_get_label(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(proposal)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Markup gets the label of proposal with markup. The label is shown in the
// list of proposals and may contain markup. This will be used instead of
// gtk_source_completion_proposal_get_label() if implemented. The returned
// string must be freed with g_free().
//
// The function returns the following values:
//
//   - utf8: new string containing the label of proposal with markup.
//
func (proposal *CompletionProposal) markup() string {
	gclass := (*C.GtkSourceCompletionProposalIface)(coreglib.PeekParentClass(proposal))
	fnarg := gclass.get_markup

	var _arg0 *C.GtkSourceCompletionProposal // out
	var _cret *C.gchar                       // in

	_arg0 = (*C.GtkSourceCompletionProposal)(unsafe.Pointer(coreglib.InternObject(proposal).Native()))

	_cret = C._gotk4_gtksource3_CompletionProposal_virtual_get_markup(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(proposal)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Text gets the text of proposal. The text that is inserted into the
// text buffer when the proposal is activated by the default activation.
// You are free to implement a custom activation handler in the
// provider and not implement this function. For more information, see
// gtk_source_completion_provider_activate_proposal(). The returned string must
// be freed with g_free().
//
// The function returns the following values:
//
//   - utf8: new string containing the text of proposal.
//
func (proposal *CompletionProposal) text() string {
	gclass := (*C.GtkSourceCompletionProposalIface)(coreglib.PeekParentClass(proposal))
	fnarg := gclass.get_text

	var _arg0 *C.GtkSourceCompletionProposal // out
	var _cret *C.gchar                       // in

	_arg0 = (*C.GtkSourceCompletionProposal)(unsafe.Pointer(coreglib.InternObject(proposal).Native()))

	_cret = C._gotk4_gtksource3_CompletionProposal_virtual_get_text(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(proposal)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Hash: get the hash value of proposal. This is used to (together with
// gtk_source_completion_proposal_equal()) to match proposals in the completion
// model. By default, it uses a direct hash (g_direct_hash()).
//
// The function returns the following values:
//
//   - guint: hash value of proposal.
//
func (proposal *CompletionProposal) hash() uint {
	gclass := (*C.GtkSourceCompletionProposalIface)(coreglib.PeekParentClass(proposal))
	fnarg := gclass.hash

	var _arg0 *C.GtkSourceCompletionProposal // out
	var _cret C.guint                        // in

	_arg0 = (*C.GtkSourceCompletionProposal)(unsafe.Pointer(coreglib.InternObject(proposal).Native()))

	_cret = C._gotk4_gtksource3_CompletionProposal_virtual_hash(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(proposal)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// CompletionProposalIface: virtual function table for SourceCompletionProposal.
//
// An instance of this type is always passed by reference.
type CompletionProposalIface struct {
	*completionProposalIface
}

// completionProposalIface is the struct that's finalized.
type completionProposalIface struct {
	native *C.GtkSourceCompletionProposalIface
}
