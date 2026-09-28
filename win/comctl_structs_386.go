//go:build windows

package win

import (
	"github.com/rodrigocfd/windigo/co"
)

// [TBBUTTON] struct, with C memory layout.
//
// Two bytes of padding here rather than the six of the 64-bit build:
// the padding exists to align DwData to a pointer, and a pointer is
// four bytes wide in this build. With the 64-bit padding, every field
// from DwData on would land four bytes late and the toolbar would read
// a command id where a string pointer is -- which does not fail, it
// draws something else.
//
// [TBBUTTON]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-tbbutton
type TBBUTTON struct {
	IBitmap   int32 // With multiple image lists, HIWORD is the image list index.
	IdCommand int32
	FsState   co.TBSTATE
	FsStyle   co.BTNS
	bReserved [2]uint8
	DwData    uintptr
	IString   *uint16 // Convert to/from string with [wstr.DecodePtr] and [wstr.EncodeToPtr]; can also be the index in the string list.
}
