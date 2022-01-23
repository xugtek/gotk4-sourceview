// Code generated by girgen. DO NOT EDIT.

package gtksource

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v3"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtksourceview/gtksource.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_source_completion_info_get_type()), F: marshalCompletionInfor},
	})
}

type CompletionInfo struct {
	_ [0]func() // equal guard
	gtk.Window
}

var (
	_ gtk.Binner = (*CompletionInfo)(nil)
)

func wrapCompletionInfo(obj *externglib.Object) *CompletionInfo {
	return &CompletionInfo{
		Window: gtk.Window{
			Bin: gtk.Bin{
				Container: gtk.Container{
					Widget: gtk.Widget{
						InitiallyUnowned: externglib.InitiallyUnowned{
							Object: obj,
						},
						Object: obj,
						ImplementorIface: atk.ImplementorIface{
							Object: obj,
						},
						Buildable: gtk.Buildable{
							Object: obj,
						},
					},
				},
			},
		},
	}
}

func marshalCompletionInfor(p uintptr) (interface{}, error) {
	return wrapCompletionInfo(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// The function returns the following values:
//
//    - completionInfo: new GtkSourceCompletionInfo.
//
func NewCompletionInfo() *CompletionInfo {
	var _cret *C.GtkSourceCompletionInfo // in

	_cret = C.gtk_source_completion_info_new()

	var _completionInfo *CompletionInfo // out

	_completionInfo = wrapCompletionInfo(externglib.Take(unsafe.Pointer(_cret)))

	return _completionInfo
}

// MoveToIter moves the SourceCompletionInfo to iter. If iter is NULL info is
// moved to the cursor position. Moving will respect the Gravity setting of the
// info window and will ensure the line at iter is not occluded by the window.
//
// The function takes the following parameters:
//
//    - view on which the info window should be positioned.
//    - iter (optional): TextIter.
//
func (info *CompletionInfo) MoveToIter(view *gtk.TextView, iter *gtk.TextIter) {
	var _arg0 *C.GtkSourceCompletionInfo // out
	var _arg1 *C.GtkTextView             // out
	var _arg2 *C.GtkTextIter             // out

	_arg0 = (*C.GtkSourceCompletionInfo)(unsafe.Pointer(info.Native()))
	_arg1 = (*C.GtkTextView)(unsafe.Pointer(view.Native()))
	if iter != nil {
		_arg2 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(iter)))
	}

	C.gtk_source_completion_info_move_to_iter(_arg0, _arg1, _arg2)
	runtime.KeepAlive(info)
	runtime.KeepAlive(view)
	runtime.KeepAlive(iter)
}
