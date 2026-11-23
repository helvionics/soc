//go:build soc.rp2040 || soc.rp2350

package rp2

import "runtime/interrupt"

func rp2EnterFlashSafeSection() (interrupt.State, bool) {
	return interrupt.Disable(), false
}

func rp2ExitFlashSafeSection(state interrupt.State, _ bool) {
	interrupt.Restore(state)
}
