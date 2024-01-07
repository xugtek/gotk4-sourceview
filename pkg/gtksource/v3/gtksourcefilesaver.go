// Code generated by girgen. DO NOT EDIT.

package gtksource

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
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
	GTypeFileSaver = coreglib.Type(C.gtk_source_file_saver_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeFileSaver, F: marshalFileSaver},
	})
}

// FileSaverOverrides contains methods that are overridable.
type FileSaverOverrides struct {
}

func defaultFileSaverOverrides(v *FileSaver) FileSaverOverrides {
	return FileSaverOverrides{}
}

type FileSaver struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*FileSaver)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*FileSaver, *FileSaverClass, FileSaverOverrides](
		GTypeFileSaver,
		initFileSaverClass,
		wrapFileSaver,
		defaultFileSaverOverrides,
	)
}

func initFileSaverClass(gclass unsafe.Pointer, overrides FileSaverOverrides, classInitFunc func(*FileSaverClass)) {
	if classInitFunc != nil {
		class := (*FileSaverClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapFileSaver(obj *coreglib.Object) *FileSaver {
	return &FileSaver{
		Object: obj,
	}
}

func marshalFileSaver(p uintptr) (interface{}, error) {
	return wrapFileSaver(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewFileSaver creates a new SourceFileSaver object. The buffer will be saved
// to the SourceFile's location.
//
// This constructor is suitable for a simple "save" operation, when the file
// already contains a non-NULL SourceFile:location.
//
// The function takes the following parameters:
//
//   - buffer to save.
//   - file: SourceFile.
//
// The function returns the following values:
//
//   - fileSaver: new SourceFileSaver object.
//
func NewFileSaver(buffer *Buffer, file *File) *FileSaver {
	var _arg1 *C.GtkSourceBuffer    // out
	var _arg2 *C.GtkSourceFile      // out
	var _cret *C.GtkSourceFileSaver // in

	_arg1 = (*C.GtkSourceBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))
	_arg2 = (*C.GtkSourceFile)(unsafe.Pointer(coreglib.InternObject(file).Native()))

	_cret = C.gtk_source_file_saver_new(_arg1, _arg2)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(file)

	var _fileSaver *FileSaver // out

	_fileSaver = wrapFileSaver(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _fileSaver
}

// NewFileSaverWithTarget creates a new SourceFileSaver object with a target
// location. When the file saving is finished successfully, target_location
// is set to the file's SourceFile:location property. If an error occurs,
// the previous valid location is still available in SourceFile.
//
// This constructor is suitable for a "save as" operation, or for saving a new
// buffer for the first time.
//
// The function takes the following parameters:
//
//   - buffer to save.
//   - file: SourceFile.
//   - targetLocation where to save the buffer to.
//
// The function returns the following values:
//
//   - fileSaver: new SourceFileSaver object.
//
func NewFileSaverWithTarget(buffer *Buffer, file *File, targetLocation gio.Filer) *FileSaver {
	var _arg1 *C.GtkSourceBuffer    // out
	var _arg2 *C.GtkSourceFile      // out
	var _arg3 *C.GFile              // out
	var _cret *C.GtkSourceFileSaver // in

	_arg1 = (*C.GtkSourceBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))
	_arg2 = (*C.GtkSourceFile)(unsafe.Pointer(coreglib.InternObject(file).Native()))
	_arg3 = (*C.GFile)(unsafe.Pointer(coreglib.InternObject(targetLocation).Native()))

	_cret = C.gtk_source_file_saver_new_with_target(_arg1, _arg2, _arg3)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(file)
	runtime.KeepAlive(targetLocation)

	var _fileSaver *FileSaver // out

	_fileSaver = wrapFileSaver(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _fileSaver
}

// The function returns the following values:
//
//   - buffer to save.
//
func (saver *FileSaver) Buffer() *Buffer {
	var _arg0 *C.GtkSourceFileSaver // out
	var _cret *C.GtkSourceBuffer    // in

	_arg0 = (*C.GtkSourceFileSaver)(unsafe.Pointer(coreglib.InternObject(saver).Native()))

	_cret = C.gtk_source_file_saver_get_buffer(_arg0)
	runtime.KeepAlive(saver)

	var _buffer *Buffer // out

	_buffer = wrapBuffer(coreglib.Take(unsafe.Pointer(_cret)))

	return _buffer
}

// The function returns the following values:
//
//   - compressionType: compression type.
//
func (saver *FileSaver) CompressionType() CompressionType {
	var _arg0 *C.GtkSourceFileSaver      // out
	var _cret C.GtkSourceCompressionType // in

	_arg0 = (*C.GtkSourceFileSaver)(unsafe.Pointer(coreglib.InternObject(saver).Native()))

	_cret = C.gtk_source_file_saver_get_compression_type(_arg0)
	runtime.KeepAlive(saver)

	var _compressionType CompressionType // out

	_compressionType = CompressionType(_cret)

	return _compressionType
}

// The function returns the following values:
//
//   - encoding: encoding.
//
func (saver *FileSaver) Encoding() *Encoding {
	var _arg0 *C.GtkSourceFileSaver // out
	var _cret *C.GtkSourceEncoding  // in

	_arg0 = (*C.GtkSourceFileSaver)(unsafe.Pointer(coreglib.InternObject(saver).Native()))

	_cret = C.gtk_source_file_saver_get_encoding(_arg0)
	runtime.KeepAlive(saver)

	var _encoding *Encoding // out

	_encoding = (*Encoding)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _encoding
}

// The function returns the following values:
//
//   - file: SourceFile.
//
func (saver *FileSaver) File() *File {
	var _arg0 *C.GtkSourceFileSaver // out
	var _cret *C.GtkSourceFile      // in

	_arg0 = (*C.GtkSourceFileSaver)(unsafe.Pointer(coreglib.InternObject(saver).Native()))

	_cret = C.gtk_source_file_saver_get_file(_arg0)
	runtime.KeepAlive(saver)

	var _file *File // out

	_file = wrapFile(coreglib.Take(unsafe.Pointer(_cret)))

	return _file
}

// The function returns the following values:
//
//   - fileSaverFlags: flags.
//
func (saver *FileSaver) Flags() FileSaverFlags {
	var _arg0 *C.GtkSourceFileSaver     // out
	var _cret C.GtkSourceFileSaverFlags // in

	_arg0 = (*C.GtkSourceFileSaver)(unsafe.Pointer(coreglib.InternObject(saver).Native()))

	_cret = C.gtk_source_file_saver_get_flags(_arg0)
	runtime.KeepAlive(saver)

	var _fileSaverFlags FileSaverFlags // out

	_fileSaverFlags = FileSaverFlags(_cret)

	return _fileSaverFlags
}

// The function returns the following values:
//
//   - file where to save the buffer to.
//
func (saver *FileSaver) Location() *gio.File {
	var _arg0 *C.GtkSourceFileSaver // out
	var _cret *C.GFile              // in

	_arg0 = (*C.GtkSourceFileSaver)(unsafe.Pointer(coreglib.InternObject(saver).Native()))

	_cret = C.gtk_source_file_saver_get_location(_arg0)
	runtime.KeepAlive(saver)

	var _file *gio.File // out

	{
		obj := coreglib.Take(unsafe.Pointer(_cret))
		_file = &gio.File{
			Object: obj,
		}
	}

	return _file
}

// The function returns the following values:
//
//   - newlineType: newline type.
//
func (saver *FileSaver) NewlineType() NewlineType {
	var _arg0 *C.GtkSourceFileSaver  // out
	var _cret C.GtkSourceNewlineType // in

	_arg0 = (*C.GtkSourceFileSaver)(unsafe.Pointer(coreglib.InternObject(saver).Native()))

	_cret = C.gtk_source_file_saver_get_newline_type(_arg0)
	runtime.KeepAlive(saver)

	var _newlineType NewlineType // out

	_newlineType = NewlineType(_cret)

	return _newlineType
}

// SaveFinish finishes a file saving started with
// gtk_source_file_saver_save_async().
//
// If the file has been saved successfully, the following SourceFile properties
// will be updated: the location, the encoding, the newline type and the
// compression type.
//
// Since the 3.20 version, gtk_text_buffer_set_modified() is called with FALSE
// if the file has been saved successfully.
//
// The function takes the following parameters:
//
//   - result: Result.
//
func (saver *FileSaver) SaveFinish(result gio.AsyncResulter) error {
	var _arg0 *C.GtkSourceFileSaver // out
	var _arg1 *C.GAsyncResult       // out
	var _cerr *C.GError             // in

	_arg0 = (*C.GtkSourceFileSaver)(unsafe.Pointer(coreglib.InternObject(saver).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	C.gtk_source_file_saver_save_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(saver)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// SetCompressionType sets the compression type. By default the compression type
// is taken from the SourceFile.
//
// The function takes the following parameters:
//
//   - compressionType: new compression type.
//
func (saver *FileSaver) SetCompressionType(compressionType CompressionType) {
	var _arg0 *C.GtkSourceFileSaver      // out
	var _arg1 C.GtkSourceCompressionType // out

	_arg0 = (*C.GtkSourceFileSaver)(unsafe.Pointer(coreglib.InternObject(saver).Native()))
	_arg1 = C.GtkSourceCompressionType(compressionType)

	C.gtk_source_file_saver_set_compression_type(_arg0, _arg1)
	runtime.KeepAlive(saver)
	runtime.KeepAlive(compressionType)
}

// SetEncoding sets the encoding. If encoding is NULL, the UTF-8 encoding will
// be set. By default the encoding is taken from the SourceFile.
//
// The function takes the following parameters:
//
//   - encoding (optional): new encoding, or NULL for UTF-8.
//
func (saver *FileSaver) SetEncoding(encoding *Encoding) {
	var _arg0 *C.GtkSourceFileSaver // out
	var _arg1 *C.GtkSourceEncoding  // out

	_arg0 = (*C.GtkSourceFileSaver)(unsafe.Pointer(coreglib.InternObject(saver).Native()))
	if encoding != nil {
		_arg1 = (*C.GtkSourceEncoding)(gextras.StructNative(unsafe.Pointer(encoding)))
	}

	C.gtk_source_file_saver_set_encoding(_arg0, _arg1)
	runtime.KeepAlive(saver)
	runtime.KeepAlive(encoding)
}

// The function takes the following parameters:
//
//   - flags: new flags.
//
func (saver *FileSaver) SetFlags(flags FileSaverFlags) {
	var _arg0 *C.GtkSourceFileSaver     // out
	var _arg1 C.GtkSourceFileSaverFlags // out

	_arg0 = (*C.GtkSourceFileSaver)(unsafe.Pointer(coreglib.InternObject(saver).Native()))
	_arg1 = C.GtkSourceFileSaverFlags(flags)

	C.gtk_source_file_saver_set_flags(_arg0, _arg1)
	runtime.KeepAlive(saver)
	runtime.KeepAlive(flags)
}

// SetNewlineType sets the newline type. By default the newline type is taken
// from the SourceFile.
//
// The function takes the following parameters:
//
//   - newlineType: new newline type.
//
func (saver *FileSaver) SetNewlineType(newlineType NewlineType) {
	var _arg0 *C.GtkSourceFileSaver  // out
	var _arg1 C.GtkSourceNewlineType // out

	_arg0 = (*C.GtkSourceFileSaver)(unsafe.Pointer(coreglib.InternObject(saver).Native()))
	_arg1 = C.GtkSourceNewlineType(newlineType)

	C.gtk_source_file_saver_set_newline_type(_arg0, _arg1)
	runtime.KeepAlive(saver)
	runtime.KeepAlive(newlineType)
}

// FileSaverClass: instance of this type is always passed by reference.
type FileSaverClass struct {
	*fileSaverClass
}

// fileSaverClass is the struct that's finalized.
type fileSaverClass struct {
	native *C.GtkSourceFileSaverClass
}

func (f *FileSaverClass) Padding() [10]unsafe.Pointer {
	valptr := &f.native.padding
	var _v [10]unsafe.Pointer // out
	{
		src := &*valptr
		for i := 0; i < 10; i++ {
			_v[i] = (unsafe.Pointer)(unsafe.Pointer(src[i]))
		}
	}
	return _v
}
