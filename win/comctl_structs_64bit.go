//go:build windows && (amd64 || arm64)

package win

import (
	"github.com/rodrigocfd/windigo/co"
)

// [TBBUTTON] struct, with C memory layout.
//
// The padding between FsStyle and DwData is what aligns DwData to a
// pointer, so it is four bytes wider here than in the 32-bit build. See
// the 386 file for the same struct with two bytes.
//
// [TBBUTTON]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-tbbutton
type TBBUTTON struct {
	IBitmap   int32 // With multiple image lists, HIWORD is the image list index.
	IdCommand int32
	FsState   co.TBSTATE
	FsStyle   co.BTNS
	bReserved [6]uint8
	DwData    uintptr
	IString   *uint16 // Convert to/from string with [wstr.DecodePtr] and [wstr.EncodeToPtr]; can also be the index in the string list.
}
