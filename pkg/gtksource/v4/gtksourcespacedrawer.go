// Code generated by girgen. DO NOT EDIT.

package gtksource

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtksourceview/gtksource.h>
import "C"

// GType values.
var (
	GTypeSpaceDrawer = coreglib.Type(C.gtk_source_space_drawer_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSpaceDrawer, F: marshalSpaceDrawer},
	})
}

// SpaceDrawerOverrides contains methods that are overridable.
type SpaceDrawerOverrides struct {
}

func defaultSpaceDrawerOverrides(v *SpaceDrawer) SpaceDrawerOverrides {
	return SpaceDrawerOverrides{}
}

type SpaceDrawer struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*SpaceDrawer)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*SpaceDrawer, *SpaceDrawerClass, SpaceDrawerOverrides](
		GTypeSpaceDrawer,
		initSpaceDrawerClass,
		wrapSpaceDrawer,
		defaultSpaceDrawerOverrides,
	)
}

func initSpaceDrawerClass(gclass unsafe.Pointer, overrides SpaceDrawerOverrides, classInitFunc func(*SpaceDrawerClass)) {
	if classInitFunc != nil {
		class := (*SpaceDrawerClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapSpaceDrawer(obj *coreglib.Object) *SpaceDrawer {
	return &SpaceDrawer{
		Object: obj,
	}
}

func marshalSpaceDrawer(p uintptr) (interface{}, error) {
	return wrapSpaceDrawer(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewSpaceDrawer creates a new SourceSpaceDrawer object. Useful for storing
// space drawing settings independently of a SourceView.
//
// The function returns the following values:
//
//   - spaceDrawer: new SourceSpaceDrawer.
func NewSpaceDrawer() *SpaceDrawer {
	var _cret *C.GtkSourceSpaceDrawer // in

	_cret = C.gtk_source_space_drawer_new()

	var _spaceDrawer *SpaceDrawer // out

	_spaceDrawer = wrapSpaceDrawer(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _spaceDrawer
}

// BindMatrixSetting binds the SourceSpaceDrawer:matrix property to a #GSettings
// key.
//
// The #GSettings key must be of the same type as the SourceSpaceDrawer:matrix
// property, that is, "au".
//
// The g_settings_bind() function cannot be used, because the default GIO
// mapping functions don't support #GVariant properties (maybe it will be
// supported by a future GIO version, in which case this function can be
// deprecated).
//
// The function takes the following parameters:
//
//   - settings #GSettings object.
//   - key settings key to bind.
//   - flags for the binding.
func (drawer *SpaceDrawer) BindMatrixSetting(settings *gio.Settings, key string, flags gio.SettingsBindFlags) {
	var _arg0 *C.GtkSourceSpaceDrawer // out
	var _arg1 *C.GSettings            // out
	var _arg2 *C.gchar                // out
	var _arg3 C.GSettingsBindFlags    // out

	_arg0 = (*C.GtkSourceSpaceDrawer)(unsafe.Pointer(coreglib.InternObject(drawer).Native()))
	_arg1 = (*C.GSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.GSettingsBindFlags(flags)

	C.gtk_source_space_drawer_bind_matrix_setting(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(drawer)
	runtime.KeepAlive(settings)
	runtime.KeepAlive(key)
	runtime.KeepAlive(flags)
}

// The function returns the following values:
//
//   - ok: whether the SourceSpaceDrawer:matrix property is enabled.
func (drawer *SpaceDrawer) EnableMatrix() bool {
	var _arg0 *C.GtkSourceSpaceDrawer // out
	var _cret C.gboolean              // in

	_arg0 = (*C.GtkSourceSpaceDrawer)(unsafe.Pointer(coreglib.InternObject(drawer).Native()))

	_cret = C.gtk_source_space_drawer_get_enable_matrix(_arg0)
	runtime.KeepAlive(drawer)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Matrix gets the value of the SourceSpaceDrawer:matrix property, as a
// #GVariant. An empty array can be returned in case the matrix is a zero
// matrix.
//
// The gtk_source_space_drawer_get_types_for_locations() function may be more
// convenient to use.
//
// The function returns the following values:
//
//   - variant value as a new floating #GVariant instance.
func (drawer *SpaceDrawer) Matrix() *glib.Variant {
	var _arg0 *C.GtkSourceSpaceDrawer // out
	var _cret *C.GVariant             // in

	_arg0 = (*C.GtkSourceSpaceDrawer)(unsafe.Pointer(coreglib.InternObject(drawer).Native()))

	_cret = C.gtk_source_space_drawer_get_matrix(_arg0)
	runtime.KeepAlive(drawer)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_variant)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_variant_unref((*C.GVariant)(intern.C))
		},
	)

	return _variant
}

// TypesForLocations: if only one location is specified, this function returns
// what kind of white spaces are drawn at that location. The value is retrieved
// from the SourceSpaceDrawer:matrix property.
//
// If several locations are specified, this function returns the logical AND
// for those locations. Which means that if a certain kind of white space is
// present in the return value, then that kind of white space is drawn at all
// the specified locations.
//
// The function takes the following parameters:
//
//   - locations: one or several SourceSpaceLocationFlags.
//
// The function returns the following values:
//
//   - spaceTypeFlags: combination of SourceSpaceTypeFlags.
func (drawer *SpaceDrawer) TypesForLocations(locations SpaceLocationFlags) SpaceTypeFlags {
	var _arg0 *C.GtkSourceSpaceDrawer       // out
	var _arg1 C.GtkSourceSpaceLocationFlags // out
	var _cret C.GtkSourceSpaceTypeFlags     // in

	_arg0 = (*C.GtkSourceSpaceDrawer)(unsafe.Pointer(coreglib.InternObject(drawer).Native()))
	_arg1 = C.GtkSourceSpaceLocationFlags(locations)

	_cret = C.gtk_source_space_drawer_get_types_for_locations(_arg0, _arg1)
	runtime.KeepAlive(drawer)
	runtime.KeepAlive(locations)

	var _spaceTypeFlags SpaceTypeFlags // out

	_spaceTypeFlags = SpaceTypeFlags(_cret)

	return _spaceTypeFlags
}

// SetEnableMatrix sets whether the SourceSpaceDrawer:matrix property is
// enabled.
//
// The function takes the following parameters:
//
//   - enableMatrix: new value.
func (drawer *SpaceDrawer) SetEnableMatrix(enableMatrix bool) {
	var _arg0 *C.GtkSourceSpaceDrawer // out
	var _arg1 C.gboolean              // out

	_arg0 = (*C.GtkSourceSpaceDrawer)(unsafe.Pointer(coreglib.InternObject(drawer).Native()))
	if enableMatrix {
		_arg1 = C.TRUE
	}

	C.gtk_source_space_drawer_set_enable_matrix(_arg0, _arg1)
	runtime.KeepAlive(drawer)
	runtime.KeepAlive(enableMatrix)
}

// SetMatrix sets a new value to the SourceSpaceDrawer:matrix property, as a
// #GVariant. If matrix is NULL, then an empty array is set.
//
// If matrix is floating, it is consumed.
//
// The gtk_source_space_drawer_set_types_for_locations() function may be more
// convenient to use.
//
// The function takes the following parameters:
//
//   - matrix (optional): new matrix value, or NULL.
func (drawer *SpaceDrawer) SetMatrix(matrix *glib.Variant) {
	var _arg0 *C.GtkSourceSpaceDrawer // out
	var _arg1 *C.GVariant             // out

	_arg0 = (*C.GtkSourceSpaceDrawer)(unsafe.Pointer(coreglib.InternObject(drawer).Native()))
	if matrix != nil {
		_arg1 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(matrix)))
	}

	C.gtk_source_space_drawer_set_matrix(_arg0, _arg1)
	runtime.KeepAlive(drawer)
	runtime.KeepAlive(matrix)
}

// SetTypesForLocations modifies the SourceSpaceDrawer:matrix property at the
// specified locations.
//
// The function takes the following parameters:
//
//   - locations: one or several SourceSpaceLocationFlags.
//   - types: combination of SourceSpaceTypeFlags.
func (drawer *SpaceDrawer) SetTypesForLocations(locations SpaceLocationFlags, types SpaceTypeFlags) {
	var _arg0 *C.GtkSourceSpaceDrawer       // out
	var _arg1 C.GtkSourceSpaceLocationFlags // out
	var _arg2 C.GtkSourceSpaceTypeFlags     // out

	_arg0 = (*C.GtkSourceSpaceDrawer)(unsafe.Pointer(coreglib.InternObject(drawer).Native()))
	_arg1 = C.GtkSourceSpaceLocationFlags(locations)
	_arg2 = C.GtkSourceSpaceTypeFlags(types)

	C.gtk_source_space_drawer_set_types_for_locations(_arg0, _arg1, _arg2)
	runtime.KeepAlive(drawer)
	runtime.KeepAlive(locations)
	runtime.KeepAlive(types)
}

// SpaceDrawerClass: instance of this type is always passed by reference.
type SpaceDrawerClass struct {
	*spaceDrawerClass
}

// spaceDrawerClass is the struct that's finalized.
type spaceDrawerClass struct {
	native *C.GtkSourceSpaceDrawerClass
}

func (s *SpaceDrawerClass) Padding() [20]unsafe.Pointer {
	valptr := &s.native.padding
	var _v [20]unsafe.Pointer // out
	{
		src := &*valptr
		for i := 0; i < 20; i++ {
			_v[i] = (unsafe.Pointer)(unsafe.Pointer(src[i]))
		}
	}
	return _v
}
