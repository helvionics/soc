//go:build soc.rp2040 || soc.rp2350

package rp2

import _ "unsafe"

//go:linkname platformTicks runtime.platformTicks
func platformTicks() uint64 { return Ticks() }

//go:linkname platformTickNS runtime.platformTickNS
func platformTickNS() uint64 { return 1000 }

//go:linkname platformSleep runtime.platformSleep
func platformSleep(ticks uint64) { sleepPlatform(ticks) }
