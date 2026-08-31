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

// [IDXGIDevice] COM interface.
//
// [IDXGIDevice]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nn-dxgi-idxgidevice
type IDXGIDevice struct{ IDXGIObject }

type _IDXGIDeviceVt struct {
	_IDXGIObjectVt
	GetAdapter             uintptr
	CreateSurface          uintptr
	QueryResourceResidency uintptr
	SetGPUThreadPriority   uintptr
	GetGPUThreadPriority   uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IDXGIDevice) IID() *co.IID {
	return &codxgi.IID_IDXGIDevice
}

// [AddRef] method.
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IDXGIDevice) AddRef(releaser *win.OleReleaser) *IDXGIDevice {
	return utl.OleNewFromAddRef[*IDXGIDevice](me, releaser)
}

// [CreateSurface] method.
//
// [CreateSurface]:https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgidevice-createsurface
func (me *IDXGIDevice) CreateSurface(
	releaser *win.OleReleaser,
	pDest *DXGI_SURFACE_DESC,
	numSurfaces int,
	usage codxgi.DXGI_USAGE,
	pSharedResource *DXGI_SHARED_RESOURCE,
) (*IDXGISurface, error) {
	var ppvtQueried uintptr
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGIDeviceVt](me.Ppvt()).CreateSurface,
		me.Ppvt(),
		uintptr(unsafe.Pointer(pDest)),
		uintptr(uint32(numSurfaces)),
		uintptr(usage),
		uintptr(unsafe.Pointer(pSharedResource)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*IDXGISurface](ret, ppvtQueried, releaser)
}

// [GetAdapter] method.
//
// [GetAdapter]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgidevice-getadapter
func (me *IDXGIDevice) GetAdapter(releaser *win.OleReleaser) (*IDXGIAdapter, error) {
	return utl.OleNewFromCallWithoutParms[*IDXGIAdapter](me, releaser,
		utl.Vt[_IDXGIDeviceVt](me.Ppvt()).GetAdapter)
}

// [GetGPUThreadPriority] method.
//
// [GetGPUThreadPriority]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgidevice-getgputhreadpriority
func (me *IDXGIDevice) GetGPUThreadPriority() (int, error) {
	var priority int32
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGIDeviceVt](me.Ppvt()).GetGPUThreadPriority,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&priority)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return 0, hr
	}
	return int(priority), nil
}

// [QueryResourceResidency] method.
//
// [QueryResourceResidency]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgidevice-queryresourceresidency
func (me *IDXGIDevice) QueryResourceResidency(
	resources ...*IDXGIResource,
) ([]codxgi.DXGI_RESIDENCY, error) {
	ppvts := make([]uintptr, 0, len(resources))
	for _, res := range resources {
		ppvts = append(ppvts, res.Ppvt())
	}
	statuses := make([]codxgi.DXGI_RESIDENCY, len(resources))

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGIDeviceVt](me.Ppvt()).QueryResourceResidency,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&ppvts[0])),
		uintptr(unsafe.Pointer(&statuses[0])),
		uintptr(uint32(len(resources))))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return nil, hr
	}
	return statuses, nil
}

// [SetGPUThreadPriority] method.
//
// [SetGPUThreadPriority]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgidevice-setgputhreadpriority
func (me *IDXGIDevice) SetGPUThreadPriority(priority int) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGIDeviceVt](me.Ppvt()).SetGPUThreadPriority,
		me.Ppvt(),
		uintptr(uint32(priority)))
	return utl.HresultToError(ret)
}
