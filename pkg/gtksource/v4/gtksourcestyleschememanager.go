// Code generated by girgen. DO NOT EDIT.

package gtksource

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtksourceview/gtksource.h>
import "C"

// GType values.
var (
	GTypeStyleSchemeManager = coreglib.Type(C.gtk_source_style_scheme_manager_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeStyleSchemeManager, F: marshalStyleSchemeManager},
	})
}

// StyleSchemeManagerOverrides contains methods that are overridable.
type StyleSchemeManagerOverrides struct {
}

func defaultStyleSchemeManagerOverrides(v *StyleSchemeManager) StyleSchemeManagerOverrides {
	return StyleSchemeManagerOverrides{}
}

type StyleSchemeManager struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*StyleSchemeManager)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*StyleSchemeManager, *StyleSchemeManagerClass, StyleSchemeManagerOverrides](
		GTypeStyleSchemeManager,
		initStyleSchemeManagerClass,
		wrapStyleSchemeManager,
		defaultStyleSchemeManagerOverrides,
	)
}

func initStyleSchemeManagerClass(gclass unsafe.Pointer, overrides StyleSchemeManagerOverrides, classInitFunc func(*StyleSchemeManagerClass)) {
	if classInitFunc != nil {
		class := (*StyleSchemeManagerClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapStyleSchemeManager(obj *coreglib.Object) *StyleSchemeManager {
	return &StyleSchemeManager{
		Object: obj,
	}
}

func marshalStyleSchemeManager(p uintptr) (interface{}, error) {
	return wrapStyleSchemeManager(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewStyleSchemeManager creates a new style manager. If you do not need more
// than one style manager then use gtk_source_style_scheme_manager_get_default()
// instead.
//
// The function returns the following values:
//
//   - styleSchemeManager: new SourceStyleSchemeManager.
func NewStyleSchemeManager() *StyleSchemeManager {
	var _cret *C.GtkSourceStyleSchemeManager // in

	_cret = C.gtk_source_style_scheme_manager_new()

	var _styleSchemeManager *StyleSchemeManager // out

	_styleSchemeManager = wrapStyleSchemeManager(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _styleSchemeManager
}

// AppendSearchPath appends path to the list of directories
// where the manager looks for style scheme files. See
// gtk_source_style_scheme_manager_set_search_path() for details.
//
// The function takes the following parameters:
//
//   - path: directory or a filename.
func (manager *StyleSchemeManager) AppendSearchPath(path string) {
	var _arg0 *C.GtkSourceStyleSchemeManager // out
	var _arg1 *C.gchar                       // out

	_arg0 = (*C.GtkSourceStyleSchemeManager)(unsafe.Pointer(coreglib.InternObject(manager).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(path)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_source_style_scheme_manager_append_search_path(_arg0, _arg1)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(path)
}

// ForceRescan: mark any currently cached information about the available style
// scehems as invalid. All the available style schemes will be reloaded next
// time the manager is accessed.
func (manager *StyleSchemeManager) ForceRescan() {
	var _arg0 *C.GtkSourceStyleSchemeManager // out

	_arg0 = (*C.GtkSourceStyleSchemeManager)(unsafe.Pointer(coreglib.InternObject(manager).Native()))

	C.gtk_source_style_scheme_manager_force_rescan(_arg0)
	runtime.KeepAlive(manager)
}

// Scheme looks up style scheme by id.
//
// The function takes the following parameters:
//
//   - schemeId: style scheme id to find.
//
// The function returns the following values:
//
//   - styleScheme (optional) object. The returned value is owned by manager and
//     must not be unref'ed.
func (manager *StyleSchemeManager) Scheme(schemeId string) *StyleScheme {
	var _arg0 *C.GtkSourceStyleSchemeManager // out
	var _arg1 *C.gchar                       // out
	var _cret *C.GtkSourceStyleScheme        // in

	_arg0 = (*C.GtkSourceStyleSchemeManager)(unsafe.Pointer(coreglib.InternObject(manager).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(schemeId)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_source_style_scheme_manager_get_scheme(_arg0, _arg1)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(schemeId)

	var _styleScheme *StyleScheme // out

	if _cret != nil {
		_styleScheme = wrapStyleScheme(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _styleScheme
}

// SchemeIDs returns the ids of the available style schemes.
//
// The function returns the following values:
//
//   - utf8s (optional): a NULL-terminated array of strings containing the ids
//     of the available style schemes or NULL if no style scheme is available.
//     The array is sorted alphabetically according to the scheme name.
//     The array is owned by the manager and must not be modified.
func (manager *StyleSchemeManager) SchemeIDs() []string {
	var _arg0 *C.GtkSourceStyleSchemeManager // out
	var _cret **C.gchar                      // in

	_arg0 = (*C.GtkSourceStyleSchemeManager)(unsafe.Pointer(coreglib.InternObject(manager).Native()))

	_cret = C.gtk_source_style_scheme_manager_get_scheme_ids(_arg0)
	runtime.KeepAlive(manager)

	var _utf8s []string // out

	if _cret != nil {
		{
			var i int
			var z *C.gchar
			for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
				i++
			}

			src := unsafe.Slice(_cret, i)
			_utf8s = make([]string, i)
			for i := range src {
				_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			}
		}
	}

	return _utf8s
}

// SearchPath returns the current search path for the manager. See
// gtk_source_style_scheme_manager_set_search_path() for details.
//
// The function returns the following values:
//
//   - utf8s: NULL-terminated array of string containing the search path.
//     The array is owned by the manager and must not be modified.
func (manager *StyleSchemeManager) SearchPath() []string {
	var _arg0 *C.GtkSourceStyleSchemeManager // out
	var _cret **C.gchar                      // in

	_arg0 = (*C.GtkSourceStyleSchemeManager)(unsafe.Pointer(coreglib.InternObject(manager).Native()))

	_cret = C.gtk_source_style_scheme_manager_get_search_path(_arg0)
	runtime.KeepAlive(manager)

	var _utf8s []string // out

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
		}
	}

	return _utf8s
}

// PrependSearchPath prepends path to the list of directories
// where the manager looks for style scheme files. See
// gtk_source_style_scheme_manager_set_search_path() for details.
//
// The function takes the following parameters:
//
//   - path: directory or a filename.
func (manager *StyleSchemeManager) PrependSearchPath(path string) {
	var _arg0 *C.GtkSourceStyleSchemeManager // out
	var _arg1 *C.gchar                       // out

	_arg0 = (*C.GtkSourceStyleSchemeManager)(unsafe.Pointer(coreglib.InternObject(manager).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(path)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_source_style_scheme_manager_prepend_search_path(_arg0, _arg1)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(path)
}

// SetSearchPath sets the list of directories where the manager looks for style
// scheme files. If path is NULL, the search path is reset to default.
//
// The function takes the following parameters:
//
//   - path (optional): a NULL-terminated array of strings or NULL.
func (manager *StyleSchemeManager) SetSearchPath(path []string) {
	var _arg0 *C.GtkSourceStyleSchemeManager // out
	var _arg1 **C.gchar                      // out

	_arg0 = (*C.GtkSourceStyleSchemeManager)(unsafe.Pointer(coreglib.InternObject(manager).Native()))
	{
		_arg1 = (**C.gchar)(C.calloc(C.size_t((len(path) + 1)), C.size_t(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg1))
		{
			out := unsafe.Slice(_arg1, len(path)+1)
			var zero *C.gchar
			out[len(path)] = zero
			for i := range path {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(path[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	C.gtk_source_style_scheme_manager_set_search_path(_arg0, _arg1)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(path)
}

// StyleSchemeManagerGetDefault returns the default SourceStyleSchemeManager
// instance.
//
// The function returns the following values:
//
//   - styleSchemeManager Return value is owned by GtkSourceView library and
//     must not be unref'ed.
func StyleSchemeManagerGetDefault() *StyleSchemeManager {
	var _cret *C.GtkSourceStyleSchemeManager // in

	_cret = C.gtk_source_style_scheme_manager_get_default()

	var _styleSchemeManager *StyleSchemeManager // out

	_styleSchemeManager = wrapStyleSchemeManager(coreglib.Take(unsafe.Pointer(_cret)))

	return _styleSchemeManager
}

// StyleSchemeManagerClass: instance of this type is always passed by reference.
type StyleSchemeManagerClass struct {
	*styleSchemeManagerClass
}

// styleSchemeManagerClass is the struct that's finalized.
type styleSchemeManagerClass struct {
	native *C.GtkSourceStyleSchemeManagerClass
}

func (s *StyleSchemeManagerClass) Padding() [10]unsafe.Pointer {
	valptr := &s.native.padding
	var _v [10]unsafe.Pointer // out
	{
		src := &*valptr
		for i := 0; i < 10; i++ {
			_v[i] = (unsafe.Pointer)(unsafe.Pointer(src[i]))
		}
	}
	return _v
}
