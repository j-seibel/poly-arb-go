package client

import (
	"fmt"
	"os"
	"runtime/pprof"
)

type Profiler struct {
	file *os.File
}

// StartCPUProfile starts the CPU profiling and stores the file handle in the Profiler struct.
func (p *Profiler) StartCPUProfile() error {
	var err error
	p.file, err = os.Create("cpu.pprof")
	if err != nil {
		return fmt.Errorf("could not create CPU profile: %v", err)
	}

	if err := pprof.StartCPUProfile(p.file); err != nil {
		return fmt.Errorf("could not start CPU profile: %v", err)
	}

	fmt.Println("CPU profiling started")
	return nil
}

// StopCPUProfile stops the CPU profiling using the file handle from the Profiler struct.
func (p *Profiler) StopCPUProfile() {
	pprof.StopCPUProfile() // Stop profiling.
	if p.file != nil {
		p.file.Close() // Close the file.
	}
	fmt.Println("CPU profiling stopped")
}
