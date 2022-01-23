// Code generated by girgen. DO NOT EDIT.

package gtksource

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtksourceview/gtksource.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_source_completion_item_get_type()), F: marshalCompletionItemmer},
	})
}

type CompletionItem struct {
	_ [0]func() // equal guard
	*externglib.Object

	CompletionProposal
}

var (
	_ externglib.Objector = (*CompletionItem)(nil)
)

func wrapCompletionItem(obj *externglib.Object) *CompletionItem {
	return &CompletionItem{
		Object: obj,
		CompletionProposal: CompletionProposal{
			Object: obj,
		},
	}
}

func marshalCompletionItemmer(p uintptr) (interface{}, error) {
	return wrapCompletionItem(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewCompletionItem: create a new SourceCompletionItem with label label, icon
// icon and extra information info. Both icon and info can be NULL in which case
// there will be no icon shown and no extra information available.
//
// Deprecated: Use gtk_source_completion_item_new2() instead.
//
// The function takes the following parameters:
//
//    - label: item label.
//    - text: item text.
//    - icon (optional): item icon.
//    - info (optional): item extra information.
//
// The function returns the following values:
//
//    - completionItem: new SourceCompletionItem.
//
func NewCompletionItem(label, text string, icon *gdkpixbuf.Pixbuf, info string) *CompletionItem {
	var _arg1 *C.gchar                   // out
	var _arg2 *C.gchar                   // out
	var _arg3 *C.GdkPixbuf               // out
	var _arg4 *C.gchar                   // out
	var _cret *C.GtkSourceCompletionItem // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg2))
	if icon != nil {
		_arg3 = (*C.GdkPixbuf)(unsafe.Pointer(icon.Native()))
	}
	if info != "" {
		_arg4 = (*C.gchar)(unsafe.Pointer(C.CString(info)))
		defer C.free(unsafe.Pointer(_arg4))
	}

	_cret = C.gtk_source_completion_item_new(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(label)
	runtime.KeepAlive(text)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(info)

	var _completionItem *CompletionItem // out

	_completionItem = wrapCompletionItem(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _completionItem
}

// NewCompletionItemFromStock creates a new SourceCompletionItem from a stock
// item. If label is NULL, the stock label will be used.
//
// Deprecated: Use gtk_source_completion_item_new2() instead.
//
// The function takes the following parameters:
//
//    - label (optional): item label.
//    - text: item text.
//    - stock icon.
//    - info (optional): item extra information.
//
// The function returns the following values:
//
//    - completionItem: new SourceCompletionItem.
//
func NewCompletionItemFromStock(label, text, stock, info string) *CompletionItem {
	var _arg1 *C.gchar                   // out
	var _arg2 *C.gchar                   // out
	var _arg3 *C.gchar                   // out
	var _arg4 *C.gchar                   // out
	var _cret *C.GtkSourceCompletionItem // in

	if label != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(stock)))
	defer C.free(unsafe.Pointer(_arg3))
	if info != "" {
		_arg4 = (*C.gchar)(unsafe.Pointer(C.CString(info)))
		defer C.free(unsafe.Pointer(_arg4))
	}

	_cret = C.gtk_source_completion_item_new_from_stock(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(label)
	runtime.KeepAlive(text)
	runtime.KeepAlive(stock)
	runtime.KeepAlive(info)

	var _completionItem *CompletionItem // out

	_completionItem = wrapCompletionItem(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _completionItem
}

// NewCompletionItemWithMarkup: create a new SourceCompletionItem with markup
// label markup, icon icon and extra information info. Both icon and info can be
// NULL in which case there will be no icon shown and no extra information
// available.
//
// Deprecated: Use gtk_source_completion_item_new2() instead.
//
// The function takes the following parameters:
//
//    - markup: item markup label.
//    - text: item text.
//    - icon (optional): item icon.
//    - info (optional): item extra information.
//
// The function returns the following values:
//
//    - completionItem: new SourceCompletionItem.
//
func NewCompletionItemWithMarkup(markup, text string, icon *gdkpixbuf.Pixbuf, info string) *CompletionItem {
	var _arg1 *C.gchar                   // out
	var _arg2 *C.gchar                   // out
	var _arg3 *C.GdkPixbuf               // out
	var _arg4 *C.gchar                   // out
	var _cret *C.GtkSourceCompletionItem // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(markup)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg2))
	if icon != nil {
		_arg3 = (*C.GdkPixbuf)(unsafe.Pointer(icon.Native()))
	}
	if info != "" {
		_arg4 = (*C.gchar)(unsafe.Pointer(C.CString(info)))
		defer C.free(unsafe.Pointer(_arg4))
	}

	_cret = C.gtk_source_completion_item_new_with_markup(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(markup)
	runtime.KeepAlive(text)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(info)

	var _completionItem *CompletionItem // out

	_completionItem = wrapCompletionItem(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _completionItem
}

// The function takes the following parameters:
//
//    - gicon (optional) or NULL.
//
func (item *CompletionItem) SetGIcon(gicon gio.Iconner) {
	var _arg0 *C.GtkSourceCompletionItem // out
	var _arg1 *C.GIcon                   // out

	_arg0 = (*C.GtkSourceCompletionItem)(unsafe.Pointer(item.Native()))
	if gicon != nil {
		_arg1 = (*C.GIcon)(unsafe.Pointer(gicon.Native()))
	}

	C.gtk_source_completion_item_set_gicon(_arg0, _arg1)
	runtime.KeepAlive(item)
	runtime.KeepAlive(gicon)
}

// The function takes the following parameters:
//
//    - icon (optional) or NULL.
//
func (item *CompletionItem) SetIcon(icon *gdkpixbuf.Pixbuf) {
	var _arg0 *C.GtkSourceCompletionItem // out
	var _arg1 *C.GdkPixbuf               // out

	_arg0 = (*C.GtkSourceCompletionItem)(unsafe.Pointer(item.Native()))
	if icon != nil {
		_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(icon.Native()))
	}

	C.gtk_source_completion_item_set_icon(_arg0, _arg1)
	runtime.KeepAlive(item)
	runtime.KeepAlive(icon)
}

// The function takes the following parameters:
//
//    - iconName (optional): icon name, or NULL.
//
func (item *CompletionItem) SetIconName(iconName string) {
	var _arg0 *C.GtkSourceCompletionItem // out
	var _arg1 *C.gchar                   // out

	_arg0 = (*C.GtkSourceCompletionItem)(unsafe.Pointer(item.Native()))
	if iconName != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(iconName)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_source_completion_item_set_icon_name(_arg0, _arg1)
	runtime.KeepAlive(item)
	runtime.KeepAlive(iconName)
}

// The function takes the following parameters:
//
//    - info (optional): info, or NULL.
//
func (item *CompletionItem) SetInfo(info string) {
	var _arg0 *C.GtkSourceCompletionItem // out
	var _arg1 *C.gchar                   // out

	_arg0 = (*C.GtkSourceCompletionItem)(unsafe.Pointer(item.Native()))
	if info != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(info)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_source_completion_item_set_info(_arg0, _arg1)
	runtime.KeepAlive(item)
	runtime.KeepAlive(info)
}

// The function takes the following parameters:
//
//    - label (optional): label, or NULL.
//
func (item *CompletionItem) SetLabel(label string) {
	var _arg0 *C.GtkSourceCompletionItem // out
	var _arg1 *C.gchar                   // out

	_arg0 = (*C.GtkSourceCompletionItem)(unsafe.Pointer(item.Native()))
	if label != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_source_completion_item_set_label(_arg0, _arg1)
	runtime.KeepAlive(item)
	runtime.KeepAlive(label)
}

// The function takes the following parameters:
//
//    - markup (optional): markup, or NULL.
//
func (item *CompletionItem) SetMarkup(markup string) {
	var _arg0 *C.GtkSourceCompletionItem // out
	var _arg1 *C.gchar                   // out

	_arg0 = (*C.GtkSourceCompletionItem)(unsafe.Pointer(item.Native()))
	if markup != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(markup)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_source_completion_item_set_markup(_arg0, _arg1)
	runtime.KeepAlive(item)
	runtime.KeepAlive(markup)
}

// The function takes the following parameters:
//
//    - text (optional): text, or NULL.
//
func (item *CompletionItem) SetText(text string) {
	var _arg0 *C.GtkSourceCompletionItem // out
	var _arg1 *C.gchar                   // out

	_arg0 = (*C.GtkSourceCompletionItem)(unsafe.Pointer(item.Native()))
	if text != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(text)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_source_completion_item_set_text(_arg0, _arg1)
	runtime.KeepAlive(item)
	runtime.KeepAlive(text)
}

// NewCompletionItem2 creates a new SourceCompletionItem. The desired properties
// need to be set afterwards.
//
// The function returns the following values:
//
//    - completionItem: new SourceCompletionItem.
//
func NewCompletionItem2() *CompletionItem {
	var _cret *C.GtkSourceCompletionItem // in

	_cret = C.gtk_source_completion_item_new2()

	var _completionItem *CompletionItem // out

	_completionItem = wrapCompletionItem(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _completionItem
}
