// Code generated by girgen. DO NOT EDIT.

package gtksource

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtksourceview/gtksource.h>
import "C"

// NewCompletionItem2 creates a new SourceCompletionItem. The desired properties
// need to be set afterwards.
//
// The function returns the following values:
//
//   - completionItem: new SourceCompletionItem.
//
func NewCompletionItem2() *CompletionItem {
	var _cret *C.GtkSourceCompletionItem // in

	_cret = C.gtk_source_completion_item_new2()

	var _completionItem *CompletionItem // out

	_completionItem = wrapCompletionItem(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _completionItem
}
