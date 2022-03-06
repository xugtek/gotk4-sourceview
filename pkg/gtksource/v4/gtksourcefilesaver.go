// Code generated by girgen. DO NOT EDIT.

package gtksource

import (
	"context"
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtksourceview/gtksource.h>
// extern void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
// extern void _gotk4_gio2_FileProgressCallback(goffset, goffset, gpointer);
// extern void callbackDelete(gpointer);
import "C"

// glib.Type values for gtksourcefilesaver.go.
var (
	GTypeFileSaverError = externglib.Type(C.gtk_source_file_saver_error_get_type())
	GTypeFileSaverFlags = externglib.Type(C.gtk_source_file_saver_flags_get_type())
	GTypeFileSaver      = externglib.Type(C.gtk_source_file_saver_get_type())
)

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeFileSaverError, F: marshalFileSaverError},
		{T: GTypeFileSaverFlags, F: marshalFileSaverFlags},
		{T: GTypeFileSaver, F: marshalFileSaver},
	})
}

// FileSaverError: error code used with the GTK_SOURCE_FILE_SAVER_ERROR domain.
type FileSaverError C.gint

const (
	// SourceFileSaverErrorInvalidChars: buffer contains invalid characters.
	SourceFileSaverErrorInvalidChars FileSaverError = iota
	// SourceFileSaverErrorExternallyModified: file is externally modified.
	SourceFileSaverErrorExternallyModified
)

func marshalFileSaverError(p uintptr) (interface{}, error) {
	return FileSaverError(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for FileSaverError.
func (f FileSaverError) String() string {
	switch f {
	case SourceFileSaverErrorInvalidChars:
		return "InvalidChars"
	case SourceFileSaverErrorExternallyModified:
		return "ExternallyModified"
	default:
		return fmt.Sprintf("FileSaverError(%d)", f)
	}
}

// FileSaverFlags flags to define the behavior of a SourceFileSaver.
type FileSaverFlags C.guint

const (
	// SourceFileSaverFlagsNone: no flags.
	SourceFileSaverFlagsNone FileSaverFlags = 0b0
	// SourceFileSaverFlagsIgnoreInvalidChars: ignore invalid characters.
	SourceFileSaverFlagsIgnoreInvalidChars FileSaverFlags = 0b1
	// SourceFileSaverFlagsIgnoreModificationTime: save file despite external
	// modifications.
	SourceFileSaverFlagsIgnoreModificationTime FileSaverFlags = 0b10
	// SourceFileSaverFlagsCreateBackup: create a backup before saving the file.
	SourceFileSaverFlagsCreateBackup FileSaverFlags = 0b100
)

func marshalFileSaverFlags(p uintptr) (interface{}, error) {
	return FileSaverFlags(externglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for FileSaverFlags.
func (f FileSaverFlags) String() string {
	if f == 0 {
		return "FileSaverFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(139)

	for f != 0 {
		next := f & (f - 1)
		bit := f - next

		switch bit {
		case SourceFileSaverFlagsNone:
			builder.WriteString("None|")
		case SourceFileSaverFlagsIgnoreInvalidChars:
			builder.WriteString("IgnoreInvalidChars|")
		case SourceFileSaverFlagsIgnoreModificationTime:
			builder.WriteString("IgnoreModificationTime|")
		case SourceFileSaverFlagsCreateBackup:
			builder.WriteString("CreateBackup|")
		default:
			builder.WriteString(fmt.Sprintf("FileSaverFlags(0b%b)|", bit))
		}

		f = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if f contains other.
func (f FileSaverFlags) Has(other FileSaverFlags) bool {
	return (f & other) == other
}

// FileSaverOverrider contains methods that are overridable.
type FileSaverOverrider interface {
}

type FileSaver struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*FileSaver)(nil)
)

func classInitFileSaverer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapFileSaver(obj *externglib.Object) *FileSaver {
	return &FileSaver{
		Object: obj,
	}
}

func marshalFileSaver(p uintptr) (interface{}, error) {
	return wrapFileSaver(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewFileSaver creates a new SourceFileSaver object. The buffer will be saved
// to the SourceFile's location.
//
// This constructor is suitable for a simple "save" operation, when the file
// already contains a non-NULL SourceFile:location.
//
// The function takes the following parameters:
//
//    - buffer to save.
//    - file: SourceFile.
//
// The function returns the following values:
//
//    - fileSaver: new SourceFileSaver object.
//
func NewFileSaver(buffer *Buffer, file *File) *FileSaver {
	var _arg1 *C.GtkSourceBuffer    // out
	var _arg2 *C.GtkSourceFile      // out
	var _cret *C.GtkSourceFileSaver // in

	_arg1 = (*C.GtkSourceBuffer)(unsafe.Pointer(externglib.InternObject(buffer).Native()))
	_arg2 = (*C.GtkSourceFile)(unsafe.Pointer(externglib.InternObject(file).Native()))

	_cret = C.gtk_source_file_saver_new(_arg1, _arg2)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(file)

	var _fileSaver *FileSaver // out

	_fileSaver = wrapFileSaver(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _fileSaver
}

// NewFileSaverWithTarget creates a new SourceFileSaver object with a target
// location. When the file saving is finished successfully, target_location is
// set to the file's SourceFile:location property. If an error occurs, the
// previous valid location is still available in SourceFile.
//
// This constructor is suitable for a "save as" operation, or for saving a new
// buffer for the first time.
//
// The function takes the following parameters:
//
//    - buffer to save.
//    - file: SourceFile.
//    - targetLocation where to save the buffer to.
//
// The function returns the following values:
//
//    - fileSaver: new SourceFileSaver object.
//
func NewFileSaverWithTarget(buffer *Buffer, file *File, targetLocation gio.Filer) *FileSaver {
	var _arg1 *C.GtkSourceBuffer    // out
	var _arg2 *C.GtkSourceFile      // out
	var _arg3 *C.GFile              // out
	var _cret *C.GtkSourceFileSaver // in

	_arg1 = (*C.GtkSourceBuffer)(unsafe.Pointer(externglib.InternObject(buffer).Native()))
	_arg2 = (*C.GtkSourceFile)(unsafe.Pointer(externglib.InternObject(file).Native()))
	_arg3 = (*C.GFile)(unsafe.Pointer(externglib.InternObject(targetLocation).Native()))

	_cret = C.gtk_source_file_saver_new_with_target(_arg1, _arg2, _arg3)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(file)
	runtime.KeepAlive(targetLocation)

	var _fileSaver *FileSaver // out

	_fileSaver = wrapFileSaver(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _fileSaver
}

// The function returns the following values:
//
//    - buffer to save.
//
func (saver *FileSaver) Buffer() *Buffer {
	var _arg0 *C.GtkSourceFileSaver // out
	var _cret *C.GtkSourceBuffer    // in

	_arg0 = (*C.GtkSourceFileSaver)(unsafe.Pointer(externglib.InternObject(saver).Native()))

	_cret = C.gtk_source_file_saver_get_buffer(_arg0)
	runtime.KeepAlive(saver)

	var _buffer *Buffer // out

	_buffer = wrapBuffer(externglib.Take(unsafe.Pointer(_cret)))

	return _buffer
}

// The function returns the following values:
//
//    - compressionType: compression type.
//
func (saver *FileSaver) CompressionType() CompressionType {
	var _arg0 *C.GtkSourceFileSaver      // out
	var _cret C.GtkSourceCompressionType // in

	_arg0 = (*C.GtkSourceFileSaver)(unsafe.Pointer(externglib.InternObject(saver).Native()))

	_cret = C.gtk_source_file_saver_get_compression_type(_arg0)
	runtime.KeepAlive(saver)

	var _compressionType CompressionType // out

	_compressionType = CompressionType(_cret)

	return _compressionType
}

// The function returns the following values:
//
//    - encoding: encoding.
//
func (saver *FileSaver) Encoding() *Encoding {
	var _arg0 *C.GtkSourceFileSaver // out
	var _cret *C.GtkSourceEncoding  // in

	_arg0 = (*C.GtkSourceFileSaver)(unsafe.Pointer(externglib.InternObject(saver).Native()))

	_cret = C.gtk_source_file_saver_get_encoding(_arg0)
	runtime.KeepAlive(saver)

	var _encoding *Encoding // out

	_encoding = (*Encoding)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _encoding
}

// The function returns the following values:
//
//    - file: SourceFile.
//
func (saver *FileSaver) File() *File {
	var _arg0 *C.GtkSourceFileSaver // out
	var _cret *C.GtkSourceFile      // in

	_arg0 = (*C.GtkSourceFileSaver)(unsafe.Pointer(externglib.InternObject(saver).Native()))

	_cret = C.gtk_source_file_saver_get_file(_arg0)
	runtime.KeepAlive(saver)

	var _file *File // out

	_file = wrapFile(externglib.Take(unsafe.Pointer(_cret)))

	return _file
}

// The function returns the following values:
//
//    - fileSaverFlags: flags.
//
func (saver *FileSaver) Flags() FileSaverFlags {
	var _arg0 *C.GtkSourceFileSaver     // out
	var _cret C.GtkSourceFileSaverFlags // in

	_arg0 = (*C.GtkSourceFileSaver)(unsafe.Pointer(externglib.InternObject(saver).Native()))

	_cret = C.gtk_source_file_saver_get_flags(_arg0)
	runtime.KeepAlive(saver)

	var _fileSaverFlags FileSaverFlags // out

	_fileSaverFlags = FileSaverFlags(_cret)

	return _fileSaverFlags
}

// The function returns the following values:
//
//    - file where to save the buffer to.
//
func (saver *FileSaver) Location() gio.Filer {
	var _arg0 *C.GtkSourceFileSaver // out
	var _cret *C.GFile              // in

	_arg0 = (*C.GtkSourceFileSaver)(unsafe.Pointer(externglib.InternObject(saver).Native()))

	_cret = C.gtk_source_file_saver_get_location(_arg0)
	runtime.KeepAlive(saver)

	var _file gio.Filer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.Filer is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(gio.Filer)
			return ok
		})
		rv, ok := casted.(gio.Filer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.Filer")
		}
		_file = rv
	}

	return _file
}

// The function returns the following values:
//
//    - newlineType: newline type.
//
func (saver *FileSaver) NewlineType() NewlineType {
	var _arg0 *C.GtkSourceFileSaver  // out
	var _cret C.GtkSourceNewlineType // in

	_arg0 = (*C.GtkSourceFileSaver)(unsafe.Pointer(externglib.InternObject(saver).Native()))

	_cret = C.gtk_source_file_saver_get_newline_type(_arg0)
	runtime.KeepAlive(saver)

	var _newlineType NewlineType // out

	_newlineType = NewlineType(_cret)

	return _newlineType
}

// SaveAsync saves asynchronously the buffer into the file. See the Result
// documentation to know how to use this function.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - ioPriority: i/O priority of the request. E.g. G_PRIORITY_LOW,
//      G_PRIORITY_DEFAULT or G_PRIORITY_HIGH.
//    - progressCallback (optional): function to call back with progress
//      information, or NULL if progress information is not needed.
//    - callback (optional) to call when the request is satisfied.
//
func (saver *FileSaver) SaveAsync(ctx context.Context, ioPriority int, progressCallback gio.FileProgressCallback, callback gio.AsyncReadyCallback) {
	var _arg0 *C.GtkSourceFileSaver   // out
	var _arg2 *C.GCancellable         // out
	var _arg1 C.gint                  // out
	var _arg3 C.GFileProgressCallback // out
	var _arg4 C.gpointer
	var _arg5 C.GDestroyNotify
	var _arg6 C.GAsyncReadyCallback // out
	var _arg7 C.gpointer

	_arg0 = (*C.GtkSourceFileSaver)(unsafe.Pointer(externglib.InternObject(saver).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.gint(ioPriority)
	if progressCallback != nil {
		_arg3 = (*[0]byte)(C._gotk4_gio2_FileProgressCallback)
		_arg4 = C.gpointer(gbox.Assign(progressCallback))
		_arg5 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}
	if callback != nil {
		_arg6 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg7 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.gtk_source_file_saver_save_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)
	runtime.KeepAlive(saver)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(ioPriority)
	runtime.KeepAlive(progressCallback)
	runtime.KeepAlive(callback)
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
//    - result: Result.
//
func (saver *FileSaver) SaveFinish(result gio.AsyncResulter) error {
	var _arg0 *C.GtkSourceFileSaver // out
	var _arg1 *C.GAsyncResult       // out
	var _cerr *C.GError             // in

	_arg0 = (*C.GtkSourceFileSaver)(unsafe.Pointer(externglib.InternObject(saver).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(externglib.InternObject(result).Native()))

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
//    - compressionType: new compression type.
//
func (saver *FileSaver) SetCompressionType(compressionType CompressionType) {
	var _arg0 *C.GtkSourceFileSaver      // out
	var _arg1 C.GtkSourceCompressionType // out

	_arg0 = (*C.GtkSourceFileSaver)(unsafe.Pointer(externglib.InternObject(saver).Native()))
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
//    - encoding (optional): new encoding, or NULL for UTF-8.
//
func (saver *FileSaver) SetEncoding(encoding *Encoding) {
	var _arg0 *C.GtkSourceFileSaver // out
	var _arg1 *C.GtkSourceEncoding  // out

	_arg0 = (*C.GtkSourceFileSaver)(unsafe.Pointer(externglib.InternObject(saver).Native()))
	if encoding != nil {
		_arg1 = (*C.GtkSourceEncoding)(gextras.StructNative(unsafe.Pointer(encoding)))
	}

	C.gtk_source_file_saver_set_encoding(_arg0, _arg1)
	runtime.KeepAlive(saver)
	runtime.KeepAlive(encoding)
}

// The function takes the following parameters:
//
//    - flags: new flags.
//
func (saver *FileSaver) SetFlags(flags FileSaverFlags) {
	var _arg0 *C.GtkSourceFileSaver     // out
	var _arg1 C.GtkSourceFileSaverFlags // out

	_arg0 = (*C.GtkSourceFileSaver)(unsafe.Pointer(externglib.InternObject(saver).Native()))
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
//    - newlineType: new newline type.
//
func (saver *FileSaver) SetNewlineType(newlineType NewlineType) {
	var _arg0 *C.GtkSourceFileSaver  // out
	var _arg1 C.GtkSourceNewlineType // out

	_arg0 = (*C.GtkSourceFileSaver)(unsafe.Pointer(externglib.InternObject(saver).Native()))
	_arg1 = C.GtkSourceNewlineType(newlineType)

	C.gtk_source_file_saver_set_newline_type(_arg0, _arg1)
	runtime.KeepAlive(saver)
	runtime.KeepAlive(newlineType)
}
