// Code generated by girgen. DO NOT EDIT.

package gtksource

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtksourceview/gtksource.h>
// extern void _gotk4_gtksource3_UndoManager_ConnectCanUndoChanged(gpointer, guintptr);
// extern void _gotk4_gtksource3_UndoManager_ConnectCanRedoChanged(gpointer, guintptr);
// gboolean _gotk4_gtksource3_UndoManager_virtual_can_redo(void* fnptr, GtkSourceUndoManager* arg0) {
//   return ((gboolean (*)(GtkSourceUndoManager*))(fnptr))(arg0);
// };
// gboolean _gotk4_gtksource3_UndoManager_virtual_can_undo(void* fnptr, GtkSourceUndoManager* arg0) {
//   return ((gboolean (*)(GtkSourceUndoManager*))(fnptr))(arg0);
// };
// void _gotk4_gtksource3_UndoManager_virtual_begin_not_undoable_action(void* fnptr, GtkSourceUndoManager* arg0) {
//   ((void (*)(GtkSourceUndoManager*))(fnptr))(arg0);
// };
// void _gotk4_gtksource3_UndoManager_virtual_can_redo_changed(void* fnptr, GtkSourceUndoManager* arg0) {
//   ((void (*)(GtkSourceUndoManager*))(fnptr))(arg0);
// };
// void _gotk4_gtksource3_UndoManager_virtual_can_undo_changed(void* fnptr, GtkSourceUndoManager* arg0) {
//   ((void (*)(GtkSourceUndoManager*))(fnptr))(arg0);
// };
// void _gotk4_gtksource3_UndoManager_virtual_end_not_undoable_action(void* fnptr, GtkSourceUndoManager* arg0) {
//   ((void (*)(GtkSourceUndoManager*))(fnptr))(arg0);
// };
// void _gotk4_gtksource3_UndoManager_virtual_redo(void* fnptr, GtkSourceUndoManager* arg0) {
//   ((void (*)(GtkSourceUndoManager*))(fnptr))(arg0);
// };
// void _gotk4_gtksource3_UndoManager_virtual_undo(void* fnptr, GtkSourceUndoManager* arg0) {
//   ((void (*)(GtkSourceUndoManager*))(fnptr))(arg0);
// };
import "C"

// GType values.
var (
	GTypeUndoManager = coreglib.Type(C.gtk_source_undo_manager_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeUndoManager, F: marshalUndoManager},
	})
}

//
// UndoManager wraps an interface. This means the user can get the
// underlying type by calling Cast().
type UndoManager struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*UndoManager)(nil)
)

// UndoManagerer describes UndoManager's interface methods.
type UndoManagerer interface {
	coreglib.Objector

	// BeginNotUndoableAction: begin a not undoable action on the buffer.
	BeginNotUndoableAction()
	// CanRedo: get whether there are redo operations available.
	CanRedo() bool
	// CanRedoChanged emits the SourceUndoManager::can-redo-changed signal.
	CanRedoChanged()
	// CanUndo: get whether there are undo operations available.
	CanUndo() bool
	// CanUndoChanged emits the SourceUndoManager::can-undo-changed signal.
	CanUndoChanged()
	// EndNotUndoableAction ends a not undoable action on the buffer.
	EndNotUndoableAction()
	// Redo: perform a single redo.
	Redo()
	// Undo: perform a single undo.
	Undo()

	// Can-redo-changed is emitted when the ability to redo has changed.
	ConnectCanRedoChanged(func()) coreglib.SignalHandle
	// Can-undo-changed is emitted when the ability to undo has changed.
	ConnectCanUndoChanged(func()) coreglib.SignalHandle
}

var _ UndoManagerer = (*UndoManager)(nil)

func wrapUndoManager(obj *coreglib.Object) *UndoManager {
	return &UndoManager{
		Object: obj,
	}
}

func marshalUndoManager(p uintptr) (interface{}, error) {
	return wrapUndoManager(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectCanRedoChanged is emitted when the ability to redo has changed.
func (manager *UndoManager) ConnectCanRedoChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(manager, "can-redo-changed", false, unsafe.Pointer(C._gotk4_gtksource3_UndoManager_ConnectCanRedoChanged), f)
}

// ConnectCanUndoChanged is emitted when the ability to undo has changed.
func (manager *UndoManager) ConnectCanUndoChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(manager, "can-undo-changed", false, unsafe.Pointer(C._gotk4_gtksource3_UndoManager_ConnectCanUndoChanged), f)
}

// BeginNotUndoableAction: begin a not undoable action on
// the buffer. All changes between this call and the call to
// gtk_source_undo_manager_end_not_undoable_action() cannot be undone. This
// function should be re-entrant.
func (manager *UndoManager) BeginNotUndoableAction() {
	var _arg0 *C.GtkSourceUndoManager // out

	_arg0 = (*C.GtkSourceUndoManager)(unsafe.Pointer(coreglib.InternObject(manager).Native()))

	C.gtk_source_undo_manager_begin_not_undoable_action(_arg0)
	runtime.KeepAlive(manager)
}

// CanRedo: get whether there are redo operations available.
//
// The function returns the following values:
//
//   - ok: TRUE if there are redo operations available, FALSE otherwise.
func (manager *UndoManager) CanRedo() bool {
	var _arg0 *C.GtkSourceUndoManager // out
	var _cret C.gboolean              // in

	_arg0 = (*C.GtkSourceUndoManager)(unsafe.Pointer(coreglib.InternObject(manager).Native()))

	_cret = C.gtk_source_undo_manager_can_redo(_arg0)
	runtime.KeepAlive(manager)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CanRedoChanged emits the SourceUndoManager::can-redo-changed signal.
func (manager *UndoManager) CanRedoChanged() {
	var _arg0 *C.GtkSourceUndoManager // out

	_arg0 = (*C.GtkSourceUndoManager)(unsafe.Pointer(coreglib.InternObject(manager).Native()))

	C.gtk_source_undo_manager_can_redo_changed(_arg0)
	runtime.KeepAlive(manager)
}

// CanUndo: get whether there are undo operations available.
//
// The function returns the following values:
//
//   - ok: TRUE if there are undo operations available, FALSE otherwise.
func (manager *UndoManager) CanUndo() bool {
	var _arg0 *C.GtkSourceUndoManager // out
	var _cret C.gboolean              // in

	_arg0 = (*C.GtkSourceUndoManager)(unsafe.Pointer(coreglib.InternObject(manager).Native()))

	_cret = C.gtk_source_undo_manager_can_undo(_arg0)
	runtime.KeepAlive(manager)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CanUndoChanged emits the SourceUndoManager::can-undo-changed signal.
func (manager *UndoManager) CanUndoChanged() {
	var _arg0 *C.GtkSourceUndoManager // out

	_arg0 = (*C.GtkSourceUndoManager)(unsafe.Pointer(coreglib.InternObject(manager).Native()))

	C.gtk_source_undo_manager_can_undo_changed(_arg0)
	runtime.KeepAlive(manager)
}

// EndNotUndoableAction ends a not undoable action on the buffer.
func (manager *UndoManager) EndNotUndoableAction() {
	var _arg0 *C.GtkSourceUndoManager // out

	_arg0 = (*C.GtkSourceUndoManager)(unsafe.Pointer(coreglib.InternObject(manager).Native()))

	C.gtk_source_undo_manager_end_not_undoable_action(_arg0)
	runtime.KeepAlive(manager)
}

// Redo: perform a single redo. Calling this function when there are no redo
// operations available is an error. Use gtk_source_undo_manager_can_redo() to
// find out if there are redo operations available.
func (manager *UndoManager) Redo() {
	var _arg0 *C.GtkSourceUndoManager // out

	_arg0 = (*C.GtkSourceUndoManager)(unsafe.Pointer(coreglib.InternObject(manager).Native()))

	C.gtk_source_undo_manager_redo(_arg0)
	runtime.KeepAlive(manager)
}

// Undo: perform a single undo. Calling this function when there are no undo
// operations available is an error. Use gtk_source_undo_manager_can_undo() to
// find out if there are undo operations available.
func (manager *UndoManager) Undo() {
	var _arg0 *C.GtkSourceUndoManager // out

	_arg0 = (*C.GtkSourceUndoManager)(unsafe.Pointer(coreglib.InternObject(manager).Native()))

	C.gtk_source_undo_manager_undo(_arg0)
	runtime.KeepAlive(manager)
}

// beginNotUndoableAction: begin a not undoable action on
// the buffer. All changes between this call and the call to
// gtk_source_undo_manager_end_not_undoable_action() cannot be undone. This
// function should be re-entrant.
func (manager *UndoManager) beginNotUndoableAction() {
	gclass := (*C.GtkSourceUndoManagerIface)(coreglib.PeekParentClass(manager))
	fnarg := gclass.begin_not_undoable_action

	var _arg0 *C.GtkSourceUndoManager // out

	_arg0 = (*C.GtkSourceUndoManager)(unsafe.Pointer(coreglib.InternObject(manager).Native()))

	C._gotk4_gtksource3_UndoManager_virtual_begin_not_undoable_action(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(manager)
}

// canRedo: get whether there are redo operations available.
//
// The function returns the following values:
//
//   - ok: TRUE if there are redo operations available, FALSE otherwise.
func (manager *UndoManager) canRedo() bool {
	gclass := (*C.GtkSourceUndoManagerIface)(coreglib.PeekParentClass(manager))
	fnarg := gclass.can_redo

	var _arg0 *C.GtkSourceUndoManager // out
	var _cret C.gboolean              // in

	_arg0 = (*C.GtkSourceUndoManager)(unsafe.Pointer(coreglib.InternObject(manager).Native()))

	_cret = C._gotk4_gtksource3_UndoManager_virtual_can_redo(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(manager)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// canRedoChanged emits the SourceUndoManager::can-redo-changed signal.
func (manager *UndoManager) canRedoChanged() {
	gclass := (*C.GtkSourceUndoManagerIface)(coreglib.PeekParentClass(manager))
	fnarg := gclass.can_redo_changed

	var _arg0 *C.GtkSourceUndoManager // out

	_arg0 = (*C.GtkSourceUndoManager)(unsafe.Pointer(coreglib.InternObject(manager).Native()))

	C._gotk4_gtksource3_UndoManager_virtual_can_redo_changed(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(manager)
}

// canUndo: get whether there are undo operations available.
//
// The function returns the following values:
//
//   - ok: TRUE if there are undo operations available, FALSE otherwise.
func (manager *UndoManager) canUndo() bool {
	gclass := (*C.GtkSourceUndoManagerIface)(coreglib.PeekParentClass(manager))
	fnarg := gclass.can_undo

	var _arg0 *C.GtkSourceUndoManager // out
	var _cret C.gboolean              // in

	_arg0 = (*C.GtkSourceUndoManager)(unsafe.Pointer(coreglib.InternObject(manager).Native()))

	_cret = C._gotk4_gtksource3_UndoManager_virtual_can_undo(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(manager)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// canUndoChanged emits the SourceUndoManager::can-undo-changed signal.
func (manager *UndoManager) canUndoChanged() {
	gclass := (*C.GtkSourceUndoManagerIface)(coreglib.PeekParentClass(manager))
	fnarg := gclass.can_undo_changed

	var _arg0 *C.GtkSourceUndoManager // out

	_arg0 = (*C.GtkSourceUndoManager)(unsafe.Pointer(coreglib.InternObject(manager).Native()))

	C._gotk4_gtksource3_UndoManager_virtual_can_undo_changed(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(manager)
}

// endNotUndoableAction ends a not undoable action on the buffer.
func (manager *UndoManager) endNotUndoableAction() {
	gclass := (*C.GtkSourceUndoManagerIface)(coreglib.PeekParentClass(manager))
	fnarg := gclass.end_not_undoable_action

	var _arg0 *C.GtkSourceUndoManager // out

	_arg0 = (*C.GtkSourceUndoManager)(unsafe.Pointer(coreglib.InternObject(manager).Native()))

	C._gotk4_gtksource3_UndoManager_virtual_end_not_undoable_action(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(manager)
}

// Redo: perform a single redo. Calling this function when there are no redo
// operations available is an error. Use gtk_source_undo_manager_can_redo() to
// find out if there are redo operations available.
func (manager *UndoManager) redo() {
	gclass := (*C.GtkSourceUndoManagerIface)(coreglib.PeekParentClass(manager))
	fnarg := gclass.redo

	var _arg0 *C.GtkSourceUndoManager // out

	_arg0 = (*C.GtkSourceUndoManager)(unsafe.Pointer(coreglib.InternObject(manager).Native()))

	C._gotk4_gtksource3_UndoManager_virtual_redo(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(manager)
}

// Undo: perform a single undo. Calling this function when there are no undo
// operations available is an error. Use gtk_source_undo_manager_can_undo() to
// find out if there are undo operations available.
func (manager *UndoManager) undo() {
	gclass := (*C.GtkSourceUndoManagerIface)(coreglib.PeekParentClass(manager))
	fnarg := gclass.undo

	var _arg0 *C.GtkSourceUndoManager // out

	_arg0 = (*C.GtkSourceUndoManager)(unsafe.Pointer(coreglib.InternObject(manager).Native()))

	C._gotk4_gtksource3_UndoManager_virtual_undo(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(manager)
}

// UndoManagerIface: instance of this type is always passed by reference.
type UndoManagerIface struct {
	*undoManagerIface
}

// undoManagerIface is the struct that's finalized.
type undoManagerIface struct {
	native *C.GtkSourceUndoManagerIface
}
