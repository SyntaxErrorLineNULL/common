package cmd

import "syscall"

type Process struct {
	Pid int
}

// AbortProcess terminates the process associated with the Process instance.
// It sends a SIGTERM signal to the process group identified by the process ID (Pid).
// This method is used to request graceful termination of the process.
func (p *Process) AbortProcess() error {
	// Retrieve the process ID (Pid) from the Process instance.
	// This ID represents the specific process that will be terminated.
	pid := p.Pid

	// Send a SIGTERM signal to the process group that the process belongs to.
	// The '-' before the pid indicates that the signal should be sent to the entire process group
	// instead of just the single process. This is useful when the process has spawned child processes,
	// allowing all related processes to be terminated together.
	return syscall.Kill(-pid, syscall.SIGTERM)
}
