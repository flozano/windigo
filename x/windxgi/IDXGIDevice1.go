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

// [IDXGIDevice1] COM interface.
//
// [IDXGIDevice1]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nn-dxgi-idxgidevice1
type IDXGIDevice1 struct{ IDXGIDevice }

type _IDXGIDevice1Vt struct {
	_IDXGIDeviceVt
	SetMaximumFrameLatency uintptr
	GetMaximumFrameLatency uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IDXGIDevice1) IID() *co.IID {
	return &codxgi.IID_IDXGIDevice1
}

// [AddRef] method.
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IDXGIDevice1) AddRef(releaser *win.OleReleaser) *IDXGIDevice1 {
	return utl.OleNewFromAddRef[*IDXGIDevice1](me, releaser)
}

// [GetMaximumFrameLatency] method.
//
// [GetMaximumFrameLatency]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgidevice1-getmaximumframelatency
func (me *IDXGIDevice1) GetMaximumFrameLatency() (int, error) {
	var maxLatency uint32
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGIDevice1Vt](me.Ppvt()).GetMaximumFrameLatency,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&maxLatency)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return 0, hr
	}
	return int(maxLatency), nil
}

// [SetMaximumFrameLatency] method.
//
// [SetMaximumFrameLatency]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgidevice1-setmaximumframelatency
func (me *IDXGIDevice1) SetMaximumFrameLatency(maxLatency int) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGIDevice1Vt](me.Ppvt()).SetMaximumFrameLatency,
		me.Ppvt(),
		uintptr(uint32(maxLatency)))
	return utl.HresultToError(ret)
}
