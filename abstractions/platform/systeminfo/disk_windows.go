package systeminfo

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

func DiskUsage(path string) (disk DiskStatus) {
	h := windows.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")
	lpFreeBytesAvailable := uint64(0)
	lpTotalNumberOfBytes := uint64(0)
	lpTotalNumberOfFreeBytes := uint64(0)
	_, _, _ = c.Call(uintptr(unsafe.Pointer(windows.StringToUTF16Ptr("C:"))),
		uintptr(unsafe.Pointer(&lpFreeBytesAvailable)),
		uintptr(unsafe.Pointer(&lpTotalNumberOfBytes)),
		uintptr(unsafe.Pointer(&lpTotalNumberOfFreeBytes)))
	disk.All = lpTotalNumberOfBytes
	disk.Free = lpTotalNumberOfFreeBytes
	disk.Used = lpFreeBytesAvailable
	return
}
