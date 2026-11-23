//go:build (soc.rp2040 || soc.rp2350) && soc.defaultplatform

package rp2

import _ "unsafe"

func boardInit0()
func boardInit1()

//go:linkname platformInit0 runtime.platformInit0
func platformInit0() {
	Init()
	boardInit0()
}

//go:linkname platformInit1 runtime.platformInit1
func platformInit1() {
	boardInit1()
}
