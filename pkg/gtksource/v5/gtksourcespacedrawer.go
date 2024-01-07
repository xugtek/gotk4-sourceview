// Code generated by girgen. DO NOT EDIT.

package gtksource

import (
	"fmt"
	"runtime"
	"strings"
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
	GTypeSpaceLocationFlags = coreglib.Type(C.gtk_source_space_location_flags_get_type())
	GTypeSpaceTypeFlags     = coreglib.Type(C.gtk_source_space_type_flags_get_type())
	GTypeSpaceDrawer        = coreglib.Type(C.gtk_source_space_drawer_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSpaceLocationFlags, F: marshalSpaceLocationFlags},
		coreglib.TypeMarshaler{T: GTypeSpaceTypeFlags, F: marshalSpaceTypeFlags},
		coreglib.TypeMarshaler{T: GTypeSpaceDrawer, F: marshalSpaceDrawer},
	})
}

// SpaceLocationFlags contains flags for white space locations.
//
// If a line contains only white spaces (no text), the white spaces match both
// GTK_SOURCE_SPACE_LOCATION_LEADING and GTK_SOURCE_SPACE_LOCATION_TRAILING.
type SpaceLocationFlags C.guint

const (
	// SourceSpaceLocationNone: no flags.
	SourceSpaceLocationNone SpaceLocationFlags = 0b0
	// SourceSpaceLocationLeading: leading white spaces on a line, i.e.
	// the indentation.
	SourceSpaceLocationLeading SpaceLocationFlags = 0b1
	// SourceSpaceLocationInsideText: white spaces inside a line of text.
	SourceSpaceLocationInsideText SpaceLocationFlags = 0b10
	// SourceSpaceLocationTrailing: trailing white spaces on a line.
	SourceSpaceLocationTrailing SpaceLocationFlags = 0b100
	// SourceSpaceLocationAll: white spaces anywhere.
	SourceSpaceLocationAll SpaceLocationFlags = 0b111
)

func marshalSpaceLocationFlags(p uintptr) (interface{}, error) {
	return SpaceLocationFlags(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for SpaceLocationFlags.
func (s SpaceLocationFlags) String() string {
	if s == 0 {
		return "SpaceLocationFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(131)

	for s != 0 {
		next := s & (s - 1)
		bit := s - next

		switch bit {
		case SourceSpaceLocationNone:
			builder.WriteString("None|")
		case SourceSpaceLocationLeading:
			builder.WriteString("Leading|")
		case SourceSpaceLocationInsideText:
			builder.WriteString("InsideText|")
		case SourceSpaceLocationTrailing:
			builder.WriteString("Trailing|")
		case SourceSpaceLocationAll:
			builder.WriteString("All|")
		default:
			builder.WriteString(fmt.Sprintf("SpaceLocationFlags(0b%b)|", bit))
		}

		s = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if s contains other.
func (s SpaceLocationFlags) Has(other SpaceLocationFlags) bool {
	return (s & other) == other
}

// SpaceTypeFlags contains flags for white space types.
type SpaceTypeFlags C.guint

const (
	// SourceSpaceTypeNone: no flags.
	SourceSpaceTypeNone SpaceTypeFlags = 0b0
	// SourceSpaceTypeSpace: space character.
	SourceSpaceTypeSpace SpaceTypeFlags = 0b1
	// SourceSpaceTypeTab: tab character.
	SourceSpaceTypeTab SpaceTypeFlags = 0b10
	// SourceSpaceTypeNewline: line break character. If the
	// SourceBuffer:implicit-trailing-newline property is TRUE,
	// SourceSpaceDrawer also draws a line break at the end of the buffer.
	SourceSpaceTypeNewline SpaceTypeFlags = 0b100
	// SourceSpaceTypeNbsp: non-breaking space character.
	SourceSpaceTypeNbsp SpaceTypeFlags = 0b1000
	// SourceSpaceTypeAll: all white spaces.
	SourceSpaceTypeAll SpaceTypeFlags = 0b1111
)

func marshalSpaceTypeFlags(p uintptr) (interface{}, error) {
	return SpaceTypeFlags(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for SpaceTypeFlags.
func (s SpaceTypeFlags) String() string {
	if s == 0 {
		return "SpaceTypeFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(121)

	for s != 0 {
		next := s & (s - 1)
		bit := s - next

		switch bit {
		case SourceSpaceTypeNone:
			builder.WriteString("None|")
		case SourceSpaceTypeSpace:
			builder.WriteString("Space|")
		case SourceSpaceTypeTab:
			builder.WriteString("Tab|")
		case SourceSpaceTypeNewline:
			builder.WriteString("Newline|")
		case SourceSpaceTypeNbsp:
			builder.WriteString("Nbsp|")
		case SourceSpaceTypeAll:
			builder.WriteString("All|")
		default:
			builder.WriteString(fmt.Sprintf("SpaceTypeFlags(0b%b)|", bit))
		}

		s = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if s contains other.
func (s SpaceTypeFlags) Has(other SpaceTypeFlags) bool {
	return (s & other) == other
}

// SpaceDrawerOverrides contains methods that are overridable.
type SpaceDrawerOverrides struct {
}

func defaultSpaceDrawerOverrides(v *SpaceDrawer) SpaceDrawerOverrides {
	return SpaceDrawerOverrides{}
}

// SpaceDrawer: represent white space characters with symbols.
//
// SourceSpaceDrawer provides a way to visualize white spaces, by drawing
// symbols.
//
// Call view.GetSpaceDrawer to get the GtkSourceSpaceDrawer instance of a
// certain view.
//
// By default, no white spaces are drawn because the spacedrawer:enable-matrix
// is FALSE.
//
// To draw white spaces, spacedrawer.SetTypesForLocations can be called to set
// the spacedrawer:matrix property (by default all space types are enabled at
// all locations). Then call spacedrawer.SetEnableMatrix.
//
// For a finer-grained method, there is also the tag's tag:draw-spaces property.
//
// # Example
//
// To draw non-breaking spaces everywhere and draw all types of trailing spaces
// except newlines:
//
//    gtk_source_space_drawer_set_types_for_locations (space_drawer,
//                                                     GTK_SOURCE_SPACE_LOCATION_ALL,
//                                                     GTK_SOURCE_SPACE_TYPE_NBSP);
//
//    gtk_source_space_drawer_set_types_for_locations (space_drawer,
//                                                     GTK_SOURCE_SPACE_LOCATION_TRAILING,
//                                                     GTK_SOURCE_SPACE_TYPE_ALL &
//                                                     ~GTK_SOURCE_SPACE_TYPE_NEWLINE);
//
//    gtk_source_space_drawer_set_enable_matrix (space_drawer, TRUE);
//
// Use-case: draw unwanted white spaces
//
// A possible use-case is to draw only unwanted white spaces. Examples:
//
// - Draw all trailing spaces.
//
// - If the indentation and alignment must be done with spaces, draw tabs.
//
// And non-breaking spaces can always be drawn, everywhere, to distinguish them
// from normal spaces.
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

// NewSpaceDrawer creates a new SourceSpaceDrawer object.
//
// Useful for storing space drawing settings independently of a view.
//
// The function returns the following values:
//
//   - spaceDrawer: new SourceSpaceDrawer.
//
func NewSpaceDrawer() *SpaceDrawer {
	var _cret *C.GtkSourceSpaceDrawer // in

	_cret = C.gtk_source_space_drawer_new()

	var _spaceDrawer *SpaceDrawer // out

	_spaceDrawer = wrapSpaceDrawer(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _spaceDrawer
}

// BindMatrixSetting binds the spacedrawer:matrix property to a gio.Settings
// key.
//
// The gio.Settings key must be of the same type as the spacedrawer:matrix
// property, that is, "au".
//
// The gio.Settings.Bind() function cannot be used, because the default GIO
// mapping functions don't support glib.Variant properties (maybe it will
// be supported by a future GIO version, in which case this function can be
// deprecated).
//
// The function takes the following parameters:
//
//   - settings #GSettings object.
//   - key settings key to bind.
//   - flags for the binding.
//
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
//
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

// Matrix gets the value of the spacedrawer:matrix property, as a glib.Variant.
//
// An empty array can be returned in case the matrix is a zero matrix.
//
// The spacedrawer.GetTypesForLocations function may be more convenient to use.
//
// The function returns the following values:
//
//   - variant value as a new floating #GVariant instance.
//
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
// what kind of white spaces are drawn at that location.
//
// The value is retrieved from the spacedrawer:matrix property.
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
//
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

// SetEnableMatrix sets whether the spacedrawer:matrix property is enabled.
//
// The function takes the following parameters:
//
//   - enableMatrix: new value.
//
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

// SetMatrix sets a new value to the spacedrawer:matrix property, as a
// glib.Variant.
//
// If matrix is NULL, then an empty array is set.
//
// If matrix is floating, it is consumed.
//
// The spacedrawer.SetTypesForLocations function may be more convenient to use.
//
// The function takes the following parameters:
//
//   - matrix (optional): new matrix value, or NULL.
//
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

// SetTypesForLocations modifies the spacedrawer:matrix property at the
// specified locations.
//
// The function takes the following parameters:
//
//   - locations: one or several SourceSpaceLocationFlags.
//   - types: combination of SourceSpaceTypeFlags.
//
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
