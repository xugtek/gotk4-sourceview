// Code generated by girgen. DO NOT EDIT.

package gtksource

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtksourceview/gtksource.h>
import "C"

// GType values.
var (
	GTypeStyleSchemeChooserWidget = coreglib.Type(C.gtk_source_style_scheme_chooser_widget_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeStyleSchemeChooserWidget, F: marshalStyleSchemeChooserWidget},
	})
}

// StyleSchemeChooserWidgetOverrides contains methods that are overridable.
type StyleSchemeChooserWidgetOverrides struct {
}

func defaultStyleSchemeChooserWidgetOverrides(v *StyleSchemeChooserWidget) StyleSchemeChooserWidgetOverrides {
	return StyleSchemeChooserWidgetOverrides{}
}

// StyleSchemeChooserWidget: widget for choosing style schemes.
//
// The GtkSourceStyleSchemeChooserWidget widget lets the user select a style
// scheme. By default, the chooser presents a predefined list of style schemes.
//
// To change the initially selected style scheme, use
// styleschemechooser.SetStyleScheme. To get the selected style scheme use
// styleschemechooser.GetStyleScheme.
type StyleSchemeChooserWidget struct {
	_ [0]func() // equal guard
	gtk.Widget

	*coreglib.Object
	StyleSchemeChooser
}

var (
	_ gtk.Widgetter     = (*StyleSchemeChooserWidget)(nil)
	_ coreglib.Objector = (*StyleSchemeChooserWidget)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*StyleSchemeChooserWidget, *StyleSchemeChooserWidgetClass, StyleSchemeChooserWidgetOverrides](
		GTypeStyleSchemeChooserWidget,
		initStyleSchemeChooserWidgetClass,
		wrapStyleSchemeChooserWidget,
		defaultStyleSchemeChooserWidgetOverrides,
	)
}

func initStyleSchemeChooserWidgetClass(gclass unsafe.Pointer, overrides StyleSchemeChooserWidgetOverrides, classInitFunc func(*StyleSchemeChooserWidgetClass)) {
	if classInitFunc != nil {
		class := (*StyleSchemeChooserWidgetClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapStyleSchemeChooserWidget(obj *coreglib.Object) *StyleSchemeChooserWidget {
	return &StyleSchemeChooserWidget{
		Widget: gtk.Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Accessible: gtk.Accessible{
				Object: obj,
			},
			Buildable: gtk.Buildable{
				Object: obj,
			},
			ConstraintTarget: gtk.ConstraintTarget{
				Object: obj,
			},
		},
		Object: obj,
		StyleSchemeChooser: StyleSchemeChooser{
			Object: obj,
		},
	}
}

func marshalStyleSchemeChooserWidget(p uintptr) (interface{}, error) {
	return wrapStyleSchemeChooserWidget(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewStyleSchemeChooserWidget creates a new SourceStyleSchemeChooserWidget.
//
// The function returns the following values:
//
//   - styleSchemeChooserWidget: new SourceStyleSchemeChooserWidget.
//
func NewStyleSchemeChooserWidget() *StyleSchemeChooserWidget {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_source_style_scheme_chooser_widget_new()

	var _styleSchemeChooserWidget *StyleSchemeChooserWidget // out

	_styleSchemeChooserWidget = wrapStyleSchemeChooserWidget(coreglib.Take(unsafe.Pointer(_cret)))

	return _styleSchemeChooserWidget
}

// StyleSchemeChooserWidgetClass: instance of this type is always passed by
// reference.
type StyleSchemeChooserWidgetClass struct {
	*styleSchemeChooserWidgetClass
}

// styleSchemeChooserWidgetClass is the struct that's finalized.
type styleSchemeChooserWidgetClass struct {
	native *C.GtkSourceStyleSchemeChooserWidgetClass
}

func (s *StyleSchemeChooserWidgetClass) Parent() *gtk.WidgetClass {
	valptr := &s.native.parent
	var _v *gtk.WidgetClass // out
	_v = (*gtk.WidgetClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
