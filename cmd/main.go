package main

import (
	"conway/ui"
	"syscall"
)

var (
	kernel32         = syscall.MustLoadDLL("kernel32.dll")
	procAllocConsole = kernel32.MustFindProc("AllocConsole")
	procFreeConsole  = kernel32.MustFindProc("FreeConsole")
)

func hideConsole() {
	// 尝试释放控制台
	procFreeConsole.Call()
}

func main() {
	// 隐藏控制台
	hideConsole()
	// 开始ui
	ui.ShowGameWindows()
}
