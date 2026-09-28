//go:build windows

package dll

import (
	"sync"
	"sync/atomic"
	"syscall"
	"unsafe"
)

// Loads system DLLs and procedures.
type SystemDll struct {
	dll  *syscall.DLL
	name string
}

// System DLL.
var (
	dllMutex sync.Mutex

	Advapi     = SystemDll{nil, "advapi32"}
	Comctl     = SystemDll{nil, "comctl32"}
	Dwmapi     = SystemDll{nil, "dwmapi"}
	Dxgi       = SystemDll{nil, "dxgi"}
	Gdi        = SystemDll{nil, "gdi32"}
	Kernel     = SystemDll{nil, "kernel32"}
	Kernelbase = SystemDll{nil, "kernelbase"}
	Ktmw       = SystemDll{nil, "ktmw32"}
	Ole        = SystemDll{nil, "ole32"}
	Oleaut     = SystemDll{nil, "oleaut32"}
	Psapi      = SystemDll{nil, "psapi"}
	Shcore     = SystemDll{nil, "shcore"}
	Shell      = SystemDll{nil, "shell32"}
	Shlwapi    = SystemDll{nil, "shlwapi"}
	User       = SystemDll{nil, "user32"}
	Userenv    = SystemDll{nil, "userenv"}
	Uxtheme    = SystemDll{nil, "uxtheme"}
	Version    = SystemDll{nil, "version"}
)

// Loads the procName system procedure into pDestProc address, and
// reports whether it was there at all.
//
// Load panics when the procedure is missing, which is right for the
// hundreds of calls that have existed since Windows 95 and wrong for
// the handful that arrived later: GetDpiForWindow is Windows 10 1607,
// and an LTSB 2015 is still a supported machine in the field. A caller
// that has a sensible answer for "this Windows cannot do that" should
// be able to ask rather than recover from a panic.
func (me *SystemDll) TryLoad(pDestProc **syscall.Proc, procName string) (uintptr, bool) {
	if pProc := atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(pDestProc))); pProc != nil {
		return (*syscall.Proc)(pProc).Addr(), true // already cached
	}

	dllMutex.Lock()
	defer dllMutex.Unlock()

	if me.dll == nil {
		loaded, err := syscall.LoadDLL(me.name)
		if err != nil {
			return 0, false
		}
		me.dll = loaded
	}

	proc, err := me.dll.FindProc(procName)
	if err != nil {
		return 0, false
	}
	*pDestProc = proc
	return proc.Addr(), true
}

// Loads the procName system procedure into pDestProc address.
func (me *SystemDll) Load(pDestProc **syscall.Proc, procName string) uintptr {
	if pProc := atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(pDestProc))); pProc != nil {
		return (*syscall.Proc)(pProc).Addr() // already cached
	}

	dllMutex.Lock()
	defer dllMutex.Unlock()

	if me.dll == nil {
		me.dll = syscall.MustLoadDLL(me.name)
	}

	*pDestProc = me.dll.MustFindProc(procName)
	return (*pDestProc).Addr()
}
