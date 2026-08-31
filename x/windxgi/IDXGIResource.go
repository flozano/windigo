//go:build windows

package windxgi

import (
	"syscall"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/x/codxgi"
)

// [IDXGIResource] COM interface.
//
// [IDXGIResource]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nn-dxgi-idxgiresource
type IDXGIResource struct{ IDXGIDeviceSubObject }

type _IDXGIResourceVt struct {
	_IDXGIDeviceSubObjectVt
	GetSharedHandle     uintptr
	GetUsage            uintptr
	SetEvictionPriority uintptr
	GetEvictionPriority uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IDXGIResource) IID() *co.IID {
	return &codxgi.IID_IDXGIResource
}

// [AddRef] method.
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IDXGIResource) AddRef(releaser *win.OleReleaser) *IDXGIResource {
	return utl.OleNewFromAddRef[*IDXGIResource](me, releaser)
}

// [GetEvictionPriority] method.
//
// [GetEvictionPriority]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgiresource-getevictionpriority
func (me *IDXGIResource) GetEvictionPriority() (codxgi.DXGI_RESOURCE_PRIORITY, error) {
	return utl.OleCallReturnStruct[codxgi.DXGI_RESOURCE_PRIORITY](me,
		utl.Vt[_IDXGIResourceVt](me.Ppvt()).GetEvictionPriority)
}

// [GetSharedHandle] method.
//
// [GetSharedHandle]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgiresource-getsharedhandle
func (me *IDXGIResource) GetSharedHandle() (win.HANDLE, error) {
	return utl.OleCallReturnStruct[win.HANDLE](me,
		utl.Vt[_IDXGIResourceVt](me.Ppvt()).GetSharedHandle)
}

// [GetUsage] method.
//
// [GetUsage]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgiresource-getusage
func (me *IDXGIResource) GetUsage() (codxgi.DXGI_USAGE, error) {
	return utl.OleCallReturnStruct[codxgi.DXGI_USAGE](me,
		utl.Vt[_IDXGIResourceVt](me.Ppvt()).GetSharedHandle)
}

// [SetEvictionPriority] method.
//
// [SetEvictionPriority]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/nf-dxgi-idxgiresource-setevictionpriority
func (me *IDXGIResource) SetEvictionPriority(evictionPriority codxgi.DXGI_RESOURCE_PRIORITY) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDXGIResourceVt](me.Ppvt()).SetEvictionPriority,
		me.Ppvt(),
		uintptr(uint32(evictionPriority)))
	return utl.HresultToError(ret)
}
