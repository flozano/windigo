//go:build windows

package windxgi

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/x/codxgi"
)

// [IDXGISurface1] COM interface.
//
// [IDXGISurface1]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nn-dxgi-idxgisurface1
type IDXGISurface1 struct{ IDXGISurface }

type _IDXGISurface1Vt struct {
	_IDXGISurfaceVt
	GetDC     uintptr
	ReleaseDC uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IDXGISurface1) IID() *co.IID {
	return &codxgi.IID_IDXGISurface1
}

// [GetDC] method.
//
// [GetDC]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgisurface1-getdc
func (me *IDXGISurface1) GetDC(discard bool) (win.HDC, error) {
	var hdc win.HDC
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGISurface1Vt](me.Ppvt()).GetDC,
		me.Ppvt(),
		utl.BoolToUintptr(discard),
		uintptr(unsafe.Pointer(&hdc)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return win.HDC(0), hr
	}
	return hdc, nil
}

// [ReleaseDC] method.
//
// [ReleaseDC]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgisurface1-releasedc
func (me *IDXGISurface1) ReleaseDC(pDirtyRect *win.RECT) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGISurface1Vt](me.Ppvt()).ReleaseDC,
		me.Ppvt(),
		uintptr(unsafe.Pointer(pDirtyRect)))
	return utl.HresultToError(ret)
}
