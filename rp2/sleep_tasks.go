//go:build (soc.rp2040 || soc.rp2350) && !scheduler.none

package rp2

func sleepPlatform(ticks uint64) {
	LightSleep(ticks)
}
