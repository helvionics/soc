//go:build soc.rp2040 || soc.rp2350

package rp2

import (
	"github.com/helvionics/soc/rp2/reg"
	"unsafe"
)

var resets = (*rp.RESETS_Type)(unsafe.Pointer(rp.RESETS))

// resetBlock resets hardware blocks specified
// by the bit pattern in bits.
func resetBlock(bits uint32) {
	resets.RESET.SetBits(bits)
}

// unresetBlock brings hardware blocks specified by the
// bit pattern in bits out of reset.
func unresetBlock(bits uint32) {
	resets.RESET.ClearBits(bits)
}

// unresetBlockWait brings specified hardware blocks
// specified by the bit pattern in bits
// out of reset and wait for completion.
func unresetBlockWait(bits uint32) {
	unresetBlock(bits)
	for !resets.RESET_DONE.HasBits(bits) {
	}
}
