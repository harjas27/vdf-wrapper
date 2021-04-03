package wrapper

/*
#cgo LDFLAGS: -L./lib -lvdflib
#include "./lib/vdflib.h"
*/
import "C"
import (
	"unsafe"
)

// VDF is the struct holding necessary state for a hash chain delay function.
type VDF struct {
	difficulty int
	input      [32]byte
	output     [516]byte
	outputChan chan [516]byte
	finished   bool
}

//size of long integers in quadratic function group
const sizeInBits = 2048

// New create a new instance of VDF.
func New(difficulty int, input [32]byte) *VDF {
	return &VDF{
		difficulty: difficulty,
		input:      input,
		outputChan: make(chan [516]byte),
	}
}

// Execute runs the VDF until it's finished and put the result into output channel.
// currently on i7-6700K, it takes about 2 seconds when iteration is set to 10000
func (vdf *VDF) Execute() {
	vdf.finished = false

	inputStr := (*C.char)(unsafe.Pointer(&vdf.input[0]))
	outputStr := (*C.char)(unsafe.Pointer(&vdf.output[0]))
	C.compute(inputStr, C.ulong(len(vdf.input)), outputStr, C.ulong(len(vdf.output)), C.ulong(vdf.difficulty), C.ushort(sizeInBits))

	go func() {
		vdf.outputChan <- vdf.output
	}()

	vdf.finished = true
}

// Verify runs the verification of generated proof
// currently on i7-6700K, verification takes about 700 ms
func (vdf *VDF) Verify(proof [516]byte) bool {
	proofStr := (*C.char)(unsafe.Pointer(&proof[0]))
	inputStr := (*C.char)(unsafe.Pointer(&vdf.input[0]))
	x := C.verify(inputStr, C.ulong(len(vdf.input)), proofStr, C.ulong(len(proof)), C.ulong(vdf.difficulty), C.ushort(sizeInBits))
	if x == 1 {
		return true
	}
	return false
}

// GetOutputChannel returns the vdf output channel.
// VDF output consists of 258 bytes of serialized Y and  258 bytes of serialized Proof
func (vdf *VDF) GetOutputChannel() chan [516]byte {
	return vdf.outputChan
}

// IsFinished returns whether the vdf execution is finished or not.
func (vdf *VDF) IsFinished() bool {
	return vdf.finished
}

// GetOutput returns the vdf output, which can be bytes of 0s is the vdf is not finished.
func (vdf *VDF) GetOutput() [516]byte {
	return vdf.output
}
