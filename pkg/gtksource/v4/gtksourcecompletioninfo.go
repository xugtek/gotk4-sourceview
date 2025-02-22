// Code generated by girgen. DO NOT EDIT.

package gtksource

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v3"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtksourceview/gtksource.h>
import "C"

// GType values.
var (
	GTypeCompletionInfo = coreglib.Type(C.gtk_source_completion_info_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeCompletionInfo, F: marshalCompletionInfo},
	})
}

// CompletionInfoOverrides contains methods that are overridable.
type CompletionInfoOverrides struct {
}

func defaultCompletionInfoOverrides(v *CompletionInfo) CompletionInfoOverrides {
	return CompletionInfoOverrides{}
}

type CompletionInfo struct {
	_ [0]func() // equal guard
	gtk.Window
}

var (
	_ gtk.Binner = (*CompletionInfo)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*CompletionInfo, *CompletionInfoClass, CompletionInfoOverrides](
		GTypeCompletionInfo,
		initCompletionInfoClass,
		wrapCompletionInfo,
		defaultCompletionInfoOverrides,
	)
}

func initCompletionInfoClass(gclass unsafe.Pointer, overrides CompletionInfoOverrides, classInitFunc func(*CompletionInfoClass)) {
	if classInitFunc != nil {
		class := (*CompletionInfoClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapCompletionInfo(obj *coreglib.Object) *CompletionInfo {
	return &CompletionInfo{
		Window: gtk.Window{
			Bin: gtk.Bin{
				Container: gtk.Container{
					Widget: gtk.Widget{
						InitiallyUnowned: coreglib.InitiallyUnowned{
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

func marshalCompletionInfo(p uintptr) (interface{}, error) {
	return wrapCompletionInfo(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// The function returns the following values:
//
//   - completionInfo: new GtkSourceCompletionInfo.
func NewCompletionInfo() *CompletionInfo {
	var _cret *C.GtkSourceCompletionInfo // in

	_cret = C.gtk_source_completion_info_new()

	var _completionInfo *CompletionInfo // out

	_completionInfo = wrapCompletionInfo(coreglib.Take(unsafe.Pointer(_cret)))

	return _completionInfo
}

// MoveToIter moves the SourceCompletionInfo to iter. If iter is NULL info is
// moved to the cursor position. Moving will respect the Gravity setting of the
// info window and will ensure the line at iter is not occluded by the window.
//
// The function takes the following parameters:
//
//   - view on which the info window should be positioned.
//   - iter (optional): TextIter.
func (info *CompletionInfo) MoveToIter(view *gtk.TextView, iter *gtk.TextIter) {
	var _arg0 *C.GtkSourceCompletionInfo // out
	var _arg1 *C.GtkTextView             // out
	var _arg2 *C.GtkTextIter             // out

	_arg0 = (*C.GtkSourceCompletionInfo)(unsafe.Pointer(coreglib.InternObject(info).Native()))
	_arg1 = (*C.GtkTextView)(unsafe.Pointer(coreglib.InternObject(view).Native()))
	if iter != nil {
		_arg2 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(iter)))
	}

	C.gtk_source_completion_info_move_to_iter(_arg0, _arg1, _arg2)
	runtime.KeepAlive(info)
	runtime.KeepAlive(view)
	runtime.KeepAlive(iter)
}

// CompletionInfoClass: instance of this type is always passed by reference.
type CompletionInfoClass struct {
	*completionInfoClass
}

// completionInfoClass is the struct that's finalized.
type completionInfoClass struct {
	native *C.GtkSourceCompletionInfoClass
}

func (c *CompletionInfoClass) ParentClass() *gtk.WindowClass {
	valptr := &c.native.parent_class
	var _v *gtk.WindowClass // out
	_v = (*gtk.WindowClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}

func (c *CompletionInfoClass) Padding() [10]unsafe.Pointer {
	valptr := &c.native.padding
	var _v [10]unsafe.Pointer // out
	{
		src := &*valptr
		for i := 0; i < 10; i++ {
			_v[i] = (unsafe.Pointer)(unsafe.Pointer(src[i]))
		}
	}
	return _v
}
