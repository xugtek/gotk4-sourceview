// Code generated by girgen. DO NOT EDIT.

package gtksource

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtksourceview/gtksource.h>
import "C"

// GType values.
var (
	GTypeFile = coreglib.Type(C.gtk_source_file_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeFile, F: marshalFile},
	})
}

// FileOverrides contains methods that are overridable.
type FileOverrides struct {
}

func defaultFileOverrides(v *File) FileOverrides {
	return FileOverrides{}
}

// File: on-disk representation of a buffer.
//
// A GtkSourceFile object is the on-disk representation of a buffer. With a
// GtkSourceFile, you can create and configure a fileloader and filesaver which
// take by default the values of the GtkSourceFile properties (except for the
// file loader which auto-detect some properties). On a successful load or save
// operation, the GtkSourceFile properties are updated. If an operation fails,
// the GtkSourceFile properties have still the previous valid values.
type File struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*File)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*File, *FileClass, FileOverrides](
		GTypeFile,
		initFileClass,
		wrapFile,
		defaultFileOverrides,
	)
}

func initFileClass(gclass unsafe.Pointer, overrides FileOverrides, classInitFunc func(*FileClass)) {
	if classInitFunc != nil {
		class := (*FileClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapFile(obj *coreglib.Object) *File {
	return &File{
		Object: obj,
	}
}

func marshalFile(p uintptr) (interface{}, error) {
	return wrapFile(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// The function returns the following values:
//
//   - file: new SourceFile object.
//
func NewFile() *File {
	var _cret *C.GtkSourceFile // in

	_cret = C.gtk_source_file_new()

	var _file *File // out

	_file = wrapFile(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _file
}

// CheckFileOnDisk checks synchronously the file on disk, to know whether the
// file is externally modified, or has been deleted, and whether the file is
// read-only.
//
// SourceFile doesn't create a gio.FileMonitor to track those properties, so
// this function needs to be called instead. Creating lots of gio.FileMonitor's
// would take lots of resources.
//
// Since this function is synchronous, it is advised to call it only on local
// files. See file.IsLocal.
func (file *File) CheckFileOnDisk() {
	var _arg0 *C.GtkSourceFile // out

	_arg0 = (*C.GtkSourceFile)(unsafe.Pointer(coreglib.InternObject(file).Native()))

	C.gtk_source_file_check_file_on_disk(_arg0)
	runtime.KeepAlive(file)
}

// The function returns the following values:
//
//   - compressionType: compression type.
//
func (file *File) CompressionType() CompressionType {
	var _arg0 *C.GtkSourceFile           // out
	var _cret C.GtkSourceCompressionType // in

	_arg0 = (*C.GtkSourceFile)(unsafe.Pointer(coreglib.InternObject(file).Native()))

	_cret = C.gtk_source_file_get_compression_type(_arg0)
	runtime.KeepAlive(file)

	var _compressionType CompressionType // out

	_compressionType = CompressionType(_cret)

	return _compressionType
}

// Encoding: encoding is initially NULL. After a successful file loading or
// saving operation, the encoding is non-NULL.
//
// The function returns the following values:
//
//   - encoding: character encoding.
//
func (file *File) Encoding() *Encoding {
	var _arg0 *C.GtkSourceFile     // out
	var _cret *C.GtkSourceEncoding // in

	_arg0 = (*C.GtkSourceFile)(unsafe.Pointer(coreglib.InternObject(file).Native()))

	_cret = C.gtk_source_file_get_encoding(_arg0)
	runtime.KeepAlive(file)

	var _encoding *Encoding // out

	_encoding = (*Encoding)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _encoding
}

// The function returns the following values:
//
//   - ret: #GFile.
//
func (file *File) Location() *gio.File {
	var _arg0 *C.GtkSourceFile // out
	var _cret *C.GFile         // in

	_arg0 = (*C.GtkSourceFile)(unsafe.Pointer(coreglib.InternObject(file).Native()))

	_cret = C.gtk_source_file_get_location(_arg0)
	runtime.KeepAlive(file)

	var _ret *gio.File // out

	{
		obj := coreglib.Take(unsafe.Pointer(_cret))
		_ret = &gio.File{
			Object: obj,
		}
	}

	return _ret
}

// The function returns the following values:
//
//   - newlineType: newline type.
//
func (file *File) NewlineType() NewlineType {
	var _arg0 *C.GtkSourceFile       // out
	var _cret C.GtkSourceNewlineType // in

	_arg0 = (*C.GtkSourceFile)(unsafe.Pointer(coreglib.InternObject(file).Native()))

	_cret = C.gtk_source_file_get_newline_type(_arg0)
	runtime.KeepAlive(file)

	var _newlineType NewlineType // out

	_newlineType = NewlineType(_cret)

	return _newlineType
}

// IsDeleted returns whether the file has been deleted. If the file:location is
// NULL, returns FALSE.
//
// To have an up-to-date value, you must first call file.CheckFileOnDisk.
//
// The function returns the following values:
//
//   - ok: whether the file has been deleted.
//
func (file *File) IsDeleted() bool {
	var _arg0 *C.GtkSourceFile // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkSourceFile)(unsafe.Pointer(coreglib.InternObject(file).Native()))

	_cret = C.gtk_source_file_is_deleted(_arg0)
	runtime.KeepAlive(file)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsExternallyModified returns whether the file is externally modified.
// If the file:location is NULL, returns FALSE.
//
// To have an up-to-date value, you must first call file.CheckFileOnDisk.
//
// The function returns the following values:
//
//   - ok: whether the file is externally modified.
//
func (file *File) IsExternallyModified() bool {
	var _arg0 *C.GtkSourceFile // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkSourceFile)(unsafe.Pointer(coreglib.InternObject(file).Native()))

	_cret = C.gtk_source_file_is_externally_modified(_arg0)
	runtime.KeepAlive(file)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsLocal returns whether the file is local. If the file:location is NULL,
// returns FALSE.
//
// The function returns the following values:
//
//   - ok: whether the file is local.
//
func (file *File) IsLocal() bool {
	var _arg0 *C.GtkSourceFile // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkSourceFile)(unsafe.Pointer(coreglib.InternObject(file).Native()))

	_cret = C.gtk_source_file_is_local(_arg0)
	runtime.KeepAlive(file)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsReadonly returns whether the file is read-only. If the file:location is
// NULL, returns FALSE.
//
// To have an up-to-date value, you must first call file.CheckFileOnDisk.
//
// The function returns the following values:
//
//   - ok: whether the file is read-only.
//
func (file *File) IsReadonly() bool {
	var _arg0 *C.GtkSourceFile // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkSourceFile)(unsafe.Pointer(coreglib.InternObject(file).Native()))

	_cret = C.gtk_source_file_is_readonly(_arg0)
	runtime.KeepAlive(file)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetLocation sets the location.
//
// The function takes the following parameters:
//
//   - location (optional): new #GFile, or NULL.
//
func (file *File) SetLocation(location gio.Filer) {
	var _arg0 *C.GtkSourceFile // out
	var _arg1 *C.GFile         // out

	_arg0 = (*C.GtkSourceFile)(unsafe.Pointer(coreglib.InternObject(file).Native()))
	if location != nil {
		_arg1 = (*C.GFile)(unsafe.Pointer(coreglib.InternObject(location).Native()))
	}

	C.gtk_source_file_set_location(_arg0, _arg1)
	runtime.KeepAlive(file)
	runtime.KeepAlive(location)
}

// FileClass: instance of this type is always passed by reference.
type FileClass struct {
	*fileClass
}

// fileClass is the struct that's finalized.
type fileClass struct {
	native *C.GtkSourceFileClass
}
