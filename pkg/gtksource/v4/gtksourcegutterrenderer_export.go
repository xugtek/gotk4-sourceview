// Code generated by girgen. DO NOT EDIT.

package gtksource

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gtk/v3"
)

// #include <stdlib.h>
// #include <gtksourceview/gtksource.h>
import "C"

//export _gotk4_gtksource4_GutterRenderer_ConnectActivate
func _gotk4_gtksource4_GutterRenderer_ConnectActivate(arg0 C.gpointer, arg1 *C.GtkTextIter, arg2 *C.GdkRectangle, arg3 C.GdkEvent, arg4 C.guintptr) {
	var f func(iter *gtk.TextIter, area *gdk.Rectangle, event *gdk.Event)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg4))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(iter *gtk.TextIter, area *gdk.Rectangle, event *gdk.Event))
	}

	var _iter *gtk.TextIter  // out
	var _area *gdk.Rectangle // out
	var _event *gdk.Event    // out

	_iter = (*gtk.TextIter)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	_area = (*gdk.Rectangle)(gextras.NewStructNative(unsafe.Pointer(arg2)))
	{
		v := (*gdk.Event)(gextras.NewStructNative(unsafe.Pointer((&arg3))))
		v = gdk.CopyEventer(v)
		_event = v
	}

	f(_iter, _area, _event)
}

//export _gotk4_gtksource4_GutterRenderer_ConnectQueryActivatable
func _gotk4_gtksource4_GutterRenderer_ConnectQueryActivatable(arg0 C.gpointer, arg1 *C.GtkTextIter, arg2 *C.GdkRectangle, arg3 C.GdkEvent, arg4 C.guintptr) (cret C.gboolean) {
	var f func(iter *gtk.TextIter, area *gdk.Rectangle, event *gdk.Event) (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg4))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(iter *gtk.TextIter, area *gdk.Rectangle, event *gdk.Event) (ok bool))
	}

	var _iter *gtk.TextIter  // out
	var _area *gdk.Rectangle // out
	var _event *gdk.Event    // out

	_iter = (*gtk.TextIter)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	_area = (*gdk.Rectangle)(gextras.NewStructNative(unsafe.Pointer(arg2)))
	{
		v := (*gdk.Event)(gextras.NewStructNative(unsafe.Pointer((&arg3))))
		v = gdk.CopyEventer(v)
		_event = v
	}

	ok := f(_iter, _area, _event)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtksource4_GutterRenderer_ConnectQueryData
func _gotk4_gtksource4_GutterRenderer_ConnectQueryData(arg0 C.gpointer, arg1 *C.GtkTextIter, arg2 *C.GtkTextIter, arg3 C.GtkSourceGutterRendererState, arg4 C.guintptr) {
	var f func(start, end *gtk.TextIter, state GutterRendererState)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg4))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(start, end *gtk.TextIter, state GutterRendererState))
	}

	var _start *gtk.TextIter       // out
	var _end *gtk.TextIter         // out
	var _state GutterRendererState // out

	_start = (*gtk.TextIter)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	_end = (*gtk.TextIter)(gextras.NewStructNative(unsafe.Pointer(arg2)))
	_state = GutterRendererState(arg3)

	f(_start, _end, _state)
}

//export _gotk4_gtksource4_GutterRenderer_ConnectQueryTooltip
func _gotk4_gtksource4_GutterRenderer_ConnectQueryTooltip(arg0 C.gpointer, arg1 *C.GtkTextIter, arg2 *C.GdkRectangle, arg3 C.gint, arg4 C.gint, arg5 *C.GtkTooltip, arg6 C.guintptr) (cret C.gboolean) {
	var f func(iter *gtk.TextIter, area *gdk.Rectangle, x, y int, tooltip *gtk.Tooltip) (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg6))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(iter *gtk.TextIter, area *gdk.Rectangle, x, y int, tooltip *gtk.Tooltip) (ok bool))
	}

	var _iter *gtk.TextIter   // out
	var _area *gdk.Rectangle  // out
	var _x int                // out
	var _y int                // out
	var _tooltip *gtk.Tooltip // out

	_iter = (*gtk.TextIter)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	_area = (*gdk.Rectangle)(gextras.NewStructNative(unsafe.Pointer(arg2)))
	_x = int(arg3)
	_y = int(arg4)
	{
		obj := coreglib.Take(unsafe.Pointer(arg5))
		_tooltip = &gtk.Tooltip{
			Object: obj,
		}
	}

	ok := f(_iter, _area, _x, _y, _tooltip)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtksource4_GutterRenderer_ConnectQueueDraw
func _gotk4_gtksource4_GutterRenderer_ConnectQueueDraw(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}
