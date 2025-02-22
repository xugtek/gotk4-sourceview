// Code generated by girgen. DO NOT EDIT.

package gtksource

import (
	"fmt"
	_ "runtime/cgo"
	"strings"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gtksourceview-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtksourceview/gtksource.h>
import "C"

// GType values.
var (
	GTypeBackgroundPatternType = coreglib.Type(C.gtk_source_background_pattern_type_get_type())
	GTypeBracketMatchType      = coreglib.Type(C.gtk_source_bracket_match_type_get_type())
	GTypeChangeCaseType        = coreglib.Type(C.gtk_source_change_case_type_get_type())
	GTypeCompressionType       = coreglib.Type(C.gtk_source_compression_type_get_type())
	GTypeNewlineType           = coreglib.Type(C.gtk_source_newline_type_get_type())
	GTypeSmartHomeEndType      = coreglib.Type(C.gtk_source_smart_home_end_type_get_type())
	GTypeViewGutterPosition    = coreglib.Type(C.gtk_source_view_gutter_position_get_type())
	GTypeCompletionActivation  = coreglib.Type(C.gtk_source_completion_activation_get_type())
	GTypeGutterRendererState   = coreglib.Type(C.gtk_source_gutter_renderer_state_get_type())
	GTypeSortFlags             = coreglib.Type(C.gtk_source_sort_flags_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeBackgroundPatternType, F: marshalBackgroundPatternType},
		coreglib.TypeMarshaler{T: GTypeBracketMatchType, F: marshalBracketMatchType},
		coreglib.TypeMarshaler{T: GTypeChangeCaseType, F: marshalChangeCaseType},
		coreglib.TypeMarshaler{T: GTypeCompressionType, F: marshalCompressionType},
		coreglib.TypeMarshaler{T: GTypeNewlineType, F: marshalNewlineType},
		coreglib.TypeMarshaler{T: GTypeSmartHomeEndType, F: marshalSmartHomeEndType},
		coreglib.TypeMarshaler{T: GTypeViewGutterPosition, F: marshalViewGutterPosition},
		coreglib.TypeMarshaler{T: GTypeCompletionActivation, F: marshalCompletionActivation},
		coreglib.TypeMarshaler{T: GTypeGutterRendererState, F: marshalGutterRendererState},
		coreglib.TypeMarshaler{T: GTypeSortFlags, F: marshalSortFlags},
	})
}

type BackgroundPatternType C.gint

const (
	// SourceBackgroundPatternTypeNone: no pattern.
	SourceBackgroundPatternTypeNone BackgroundPatternType = iota
	// SourceBackgroundPatternTypeGrid: grid pattern.
	SourceBackgroundPatternTypeGrid
)

func marshalBackgroundPatternType(p uintptr) (interface{}, error) {
	return BackgroundPatternType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for BackgroundPatternType.
func (b BackgroundPatternType) String() string {
	switch b {
	case SourceBackgroundPatternTypeNone:
		return "None"
	case SourceBackgroundPatternTypeGrid:
		return "Grid"
	default:
		return fmt.Sprintf("BackgroundPatternType(%d)", b)
	}
}

type BracketMatchType C.gint

const (
	// SourceBracketMatchNone: there is no bracket to match.
	SourceBracketMatchNone BracketMatchType = iota
	// SourceBracketMatchOutOfRange: matching a bracket failed because the
	// maximum range was reached.
	SourceBracketMatchOutOfRange
	// SourceBracketMatchNotFound: matching bracket was not found.
	SourceBracketMatchNotFound
	// SourceBracketMatchFound: matching bracket was found.
	SourceBracketMatchFound
)

func marshalBracketMatchType(p uintptr) (interface{}, error) {
	return BracketMatchType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for BracketMatchType.
func (b BracketMatchType) String() string {
	switch b {
	case SourceBracketMatchNone:
		return "None"
	case SourceBracketMatchOutOfRange:
		return "OutOfRange"
	case SourceBracketMatchNotFound:
		return "NotFound"
	case SourceBracketMatchFound:
		return "Found"
	default:
		return fmt.Sprintf("BracketMatchType(%d)", b)
	}
}

type ChangeCaseType C.gint

const (
	// SourceChangeCaseLower: change case to lowercase.
	SourceChangeCaseLower ChangeCaseType = iota
	// SourceChangeCaseUpper: change case to uppercase.
	SourceChangeCaseUpper
	// SourceChangeCaseToggle: toggle case of each character.
	SourceChangeCaseToggle
	// SourceChangeCaseTitle: capitalize each word.
	SourceChangeCaseTitle
)

func marshalChangeCaseType(p uintptr) (interface{}, error) {
	return ChangeCaseType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for ChangeCaseType.
func (c ChangeCaseType) String() string {
	switch c {
	case SourceChangeCaseLower:
		return "Lower"
	case SourceChangeCaseUpper:
		return "Upper"
	case SourceChangeCaseToggle:
		return "Toggle"
	case SourceChangeCaseTitle:
		return "Title"
	default:
		return fmt.Sprintf("ChangeCaseType(%d)", c)
	}
}

func CompletionErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.gtk_source_completion_error_quark()

	var _quark glib.Quark // out

	_quark = glib.Quark(_cret)

	return _quark
}

type CompressionType C.gint

const (
	// SourceCompressionTypeNone: plain text.
	SourceCompressionTypeNone CompressionType = iota
	// SourceCompressionTypeGzip: gzip compression.
	SourceCompressionTypeGzip
)

func marshalCompressionType(p uintptr) (interface{}, error) {
	return CompressionType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for CompressionType.
func (c CompressionType) String() string {
	switch c {
	case SourceCompressionTypeNone:
		return "None"
	case SourceCompressionTypeGzip:
		return "Gzip"
	default:
		return fmt.Sprintf("CompressionType(%d)", c)
	}
}

func FileLoaderErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.gtk_source_file_loader_error_quark()

	var _quark glib.Quark // out

	_quark = glib.Quark(_cret)

	return _quark
}

func FileSaverErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.gtk_source_file_saver_error_quark()

	var _quark glib.Quark // out

	_quark = glib.Quark(_cret)

	return _quark
}

type NewlineType C.gint

const (
	// SourceNewlineTypeLf: line feed, used on UNIX.
	SourceNewlineTypeLf NewlineType = iota
	// SourceNewlineTypeCr: carriage return, used on Mac.
	SourceNewlineTypeCr
	// SourceNewlineTypeCrLf: carriage return followed by a line feed, used on
	// Windows.
	SourceNewlineTypeCrLf
)

func marshalNewlineType(p uintptr) (interface{}, error) {
	return NewlineType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for NewlineType.
func (n NewlineType) String() string {
	switch n {
	case SourceNewlineTypeLf:
		return "Lf"
	case SourceNewlineTypeCr:
		return "Cr"
	case SourceNewlineTypeCrLf:
		return "CrLf"
	default:
		return fmt.Sprintf("NewlineType(%d)", n)
	}
}

type SmartHomeEndType C.gint

const (
	// SourceSmartHomeEndDisabled: smart-home-end disabled.
	SourceSmartHomeEndDisabled SmartHomeEndType = iota
	// SourceSmartHomeEndBefore: move to the first/last non-whitespace character
	// on the first press of the HOME/END keys and to the beginning/end of the
	// line on the second press.
	SourceSmartHomeEndBefore
	// SourceSmartHomeEndAfter: move to the beginning/end of the line on the
	// first press of the HOME/END keys and to the first/last non-whitespace
	// character on the second press.
	SourceSmartHomeEndAfter
	// SourceSmartHomeEndAlways always move to the first/last non-whitespace
	// character when the HOME/END keys are pressed.
	SourceSmartHomeEndAlways
)

func marshalSmartHomeEndType(p uintptr) (interface{}, error) {
	return SmartHomeEndType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for SmartHomeEndType.
func (s SmartHomeEndType) String() string {
	switch s {
	case SourceSmartHomeEndDisabled:
		return "Disabled"
	case SourceSmartHomeEndBefore:
		return "Before"
	case SourceSmartHomeEndAfter:
		return "After"
	case SourceSmartHomeEndAlways:
		return "Always"
	default:
		return fmt.Sprintf("SmartHomeEndType(%d)", s)
	}
}

type ViewGutterPosition C.gint

const (
	// SourceViewGutterPositionLines: gutter position of the lines renderer.
	SourceViewGutterPositionLines ViewGutterPosition = -30
	// SourceViewGutterPositionMarks: gutter position of the marks renderer.
	SourceViewGutterPositionMarks ViewGutterPosition = -20
)

func marshalViewGutterPosition(p uintptr) (interface{}, error) {
	return ViewGutterPosition(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for ViewGutterPosition.
func (v ViewGutterPosition) String() string {
	switch v {
	case SourceViewGutterPositionLines:
		return "Lines"
	case SourceViewGutterPositionMarks:
		return "Marks"
	default:
		return fmt.Sprintf("ViewGutterPosition(%d)", v)
	}
}

type CompletionActivation C.guint

const (
	// SourceCompletionActivationNone: none.
	SourceCompletionActivationNone CompletionActivation = 0b0
	// SourceCompletionActivationInteractive: interactive activation.
	// By default, it occurs on each insertion in the TextBuffer. This can be
	// blocked temporarily with gtk_source_completion_block_interactive().
	SourceCompletionActivationInteractive CompletionActivation = 0b1
	// SourceCompletionActivationUserRequested: user requested
	// activation. By default, it occurs when the user presses
	// <keycombo><keycap>Control</keycap><keycap>space</keycap></keycombo>.
	SourceCompletionActivationUserRequested CompletionActivation = 0b10
)

func marshalCompletionActivation(p uintptr) (interface{}, error) {
	return CompletionActivation(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for CompletionActivation.
func (c CompletionActivation) String() string {
	if c == 0 {
		return "CompletionActivation(0)"
	}

	var builder strings.Builder
	builder.Grow(108)

	for c != 0 {
		next := c & (c - 1)
		bit := c - next

		switch bit {
		case SourceCompletionActivationNone:
			builder.WriteString("None|")
		case SourceCompletionActivationInteractive:
			builder.WriteString("Interactive|")
		case SourceCompletionActivationUserRequested:
			builder.WriteString("UserRequested|")
		default:
			builder.WriteString(fmt.Sprintf("CompletionActivation(0b%b)|", bit))
		}

		c = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if c contains other.
func (c CompletionActivation) Has(other CompletionActivation) bool {
	return (c & other) == other
}

type GutterRendererState C.guint

const (
	// SourceGutterRendererStateNormal: normal state.
	SourceGutterRendererStateNormal GutterRendererState = 0b0
	// SourceGutterRendererStateCursor: area in the renderer represents the line
	// on which the insert cursor is currently positioned.
	SourceGutterRendererStateCursor GutterRendererState = 0b1
	// SourceGutterRendererStatePrelit: mouse pointer is currently over the
	// activatable area of the renderer.
	SourceGutterRendererStatePrelit GutterRendererState = 0b10
	// SourceGutterRendererStateSelected: area in the renderer represents a line
	// in the buffer which contains part of the selection.
	SourceGutterRendererStateSelected GutterRendererState = 0b100
)

func marshalGutterRendererState(p uintptr) (interface{}, error) {
	return GutterRendererState(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for GutterRendererState.
func (g GutterRendererState) String() string {
	if g == 0 {
		return "GutterRendererState(0)"
	}

	var builder strings.Builder
	builder.Grow(129)

	for g != 0 {
		next := g & (g - 1)
		bit := g - next

		switch bit {
		case SourceGutterRendererStateNormal:
			builder.WriteString("Normal|")
		case SourceGutterRendererStateCursor:
			builder.WriteString("Cursor|")
		case SourceGutterRendererStatePrelit:
			builder.WriteString("Prelit|")
		case SourceGutterRendererStateSelected:
			builder.WriteString("Selected|")
		default:
			builder.WriteString(fmt.Sprintf("GutterRendererState(0b%b)|", bit))
		}

		g = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if g contains other.
func (g GutterRendererState) Has(other GutterRendererState) bool {
	return (g & other) == other
}

type SortFlags C.guint

const (
	// SourceSortFlagsNone: no flags specified.
	SourceSortFlagsNone SortFlags = 0b0
	// SourceSortFlagsCaseSensitive: case sensitive sort.
	SourceSortFlagsCaseSensitive SortFlags = 0b1
	// SourceSortFlagsReverseOrder: sort in reverse order.
	SourceSortFlagsReverseOrder SortFlags = 0b10
	// SourceSortFlagsRemoveDuplicates: remove duplicates.
	SourceSortFlagsRemoveDuplicates SortFlags = 0b100
)

func marshalSortFlags(p uintptr) (interface{}, error) {
	return SortFlags(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for SortFlags.
func (s SortFlags) String() string {
	if s == 0 {
		return "SortFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(108)

	for s != 0 {
		next := s & (s - 1)
		bit := s - next

		switch bit {
		case SourceSortFlagsNone:
			builder.WriteString("None|")
		case SourceSortFlagsCaseSensitive:
			builder.WriteString("CaseSensitive|")
		case SourceSortFlagsReverseOrder:
			builder.WriteString("ReverseOrder|")
		case SourceSortFlagsRemoveDuplicates:
			builder.WriteString("RemoveDuplicates|")
		default:
			builder.WriteString(fmt.Sprintf("SortFlags(0b%b)|", bit))
		}

		s = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if s contains other.
func (s SortFlags) Has(other SortFlags) bool {
	return (s & other) == other
}
