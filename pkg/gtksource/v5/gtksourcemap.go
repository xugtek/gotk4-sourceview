// Code generated by girgen. DO NOT EDIT.

package gtksource

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtksourceview/gtksource.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_source_map_get_type()), F: marshalMapper},
	})
}

type Map struct {
	_ [0]func() // equal guard
	View
}

var (
	_ gtk.Widgetter       = (*Map)(nil)
	_ externglib.Objector = (*Map)(nil)
)

func wrapMap(obj *externglib.Object) *Map {
	return &Map{
		View: View{
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
		},
	}
}

func marshalMapper(p uintptr) (interface{}, error) {
	return wrapMap(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewMap creates a new SourceMap.
//
// The function returns the following values:
//
//    - _map: new SourceMap.
//
func NewMap() *Map {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_source_map_new()

	var __map *Map // out

	__map = wrapMap(externglib.Take(unsafe.Pointer(_cret)))

	return __map
}

// GetView gets the SourceMap:view property, which is the view this widget is
// mapping.
//
// The function returns the following values:
//
//    - view (optional) or NULL.
//
func (_map *Map) GetView() *View {
	var _arg0 *C.GtkSourceMap  // out
	var _cret *C.GtkSourceView // in

	_arg0 = (*C.GtkSourceMap)(unsafe.Pointer(_map.Native()))

	_cret = C.gtk_source_map_get_view(_arg0)
	runtime.KeepAlive(_map)

	var _view *View // out

	if _cret != nil {
		_view = wrapView(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _view
}

// SetView sets the view that map will be doing the mapping to.
//
// The function takes the following parameters:
//
//    - view: SourceView.
//
func (_map *Map) SetView(view *View) {
	var _arg0 *C.GtkSourceMap  // out
	var _arg1 *C.GtkSourceView // out

	_arg0 = (*C.GtkSourceMap)(unsafe.Pointer(_map.Native()))
	_arg1 = (*C.GtkSourceView)(unsafe.Pointer(view.Native()))

	C.gtk_source_map_set_view(_arg0, _arg1)
	runtime.KeepAlive(_map)
	runtime.KeepAlive(view)
}
