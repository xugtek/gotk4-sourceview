// Code generated by girgen. DO NOT EDIT.

package gtksource

// #include <stdlib.h>
// #include <gtksourceview/gtksource.h>
import "C"

// Finalize: free the resources allocated by GtkSourceView. For example it
// unrefs the singleton objects.
//
// It is not mandatory to call this function, it's just to be friendlier to
// memory debugging tools. This function is meant to be called at the end of
// main(). It can be called several times.
func Finalize() {
	C.gtk_source_finalize()
}

// Init initializes the GtkSourceView library (e.g. for the
// internationalization).
//
// This function can be called several times, but is meant to be called at the
// beginning of main(), before any other GtkSourceView function call.
func Init() {
	C.gtk_source_init()
}
