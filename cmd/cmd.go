package cmd

import (
	"bytes"
	"os/exec"
	"sync"
)

type Commander struct {
	mutex *sync.Mutex
	wg    *sync.WaitGroup
}

func NewCommander() *Commander {
	return &Commander{}
}

func (cm *Commander) Invoke(options *Options) *Process {
	return nil
}

func (cm *Commander) builder(opts *Options) *exec.Cmd {
	command := opts.Command
	args := opts.Args

	var cmd *exec.Cmd

	if opts.parentCtx != nil {
		cmd = exec.CommandContext(opts.parentCtx, command, args...)
	} else {
		cmd = exec.Command(command, args...)
	}

	if opts.stdIn != nil {
		cmd.Stdin = opts.stdIn
	}

	// Initialize a buffer to capture the standard output from the command.
	// This buffer will hold any data written to the standard output stream by the command.
	outBuf := bytes.Buffer{}
	// Initialize a buffer to capture the standard error output from the command.
	// This buffer will hold any error messages written to the standard error stream by the command.
	errBuf := bytes.Buffer{}

	switch {
	// Redirect both standard output and standard error to in-memory buffers (outBuf and errBuf).
	// This is used when StdioBuffer is set to true, enabling separate capture of the command's output and errors.
	// outBuf captures the standard output, and errBuf captures the standard error.
	case opts.StdioBuffer:
		// Assign the address of the outBuf to cmd.Stdout.
		// This setup redirects the standard output of the command to the outBuf buffer.
		// The command's output will be written to outBuf, which can be accessed later for inspection.
		cmd.Stdout = &outBuf
		// Assign the address of the errBuf to cmd.Stderr.
		// This setup redirects the standard error of the command to the errBuf buffer.
		// The command's error messages will be written to errBuf, which can be accessed later for inspection.
		cmd.Stderr = &errBuf

	// Redirect standard output to StdOutBuf and standard error to StdErrBuf.
	// This configuration allows capturing the command's output and errors into different specified buffers.
	// StdOutBuf will contain the standard output, while StdErrBuf will contain the standard error.
	case opts.stdOutBuffer != nil && opts.stdErrBuffer != nil: // buffer combining stderr into stdout
		// Assign the provided StdOutBuf to cmd.Stdout.
		// This setup redirects the standard output of the command to the buffer specified by StdOutBuf.
		// This allows capturing the command's output into the provided buffer for later use.
		cmd.Stdout = opts.stdOutBuffer
		// Assign the provided StdErrBuf to cmd.Stderr.
		// This setup redirects the standard error of the command to the buffer specified by StdErrBuf.
		// This allows capturing the command's error messages into the provided buffer for later use.
		cmd.Stderr = opts.stdErrBuffer

	// If a standard output buffer is provided (`StdOutBuf`) and no separate standard error buffer (`StdErrBuf`)
	// is set, both stdout and stderr are directed to the same output buffer (`StdOutBuf`).
	// This configuration allows for combining the error output and the normal output into a single buffer.
	case opts.stdOutBuffer != nil && opts.stdErrBuffer == nil: // buffer combining stderr into stdout
		// Set the standard output stream to the provided output buffer (`StdOutBuf`).
		// This ensures that the output from the command's execution will be captured in `StdOutBuf`.
		cmd.Stdout = opts.stdOutBuffer
		cmd.Stderr = opts.stdOutBuffer

	// Do not redirect standard output and standard error.
	// This means the command's output and errors will not be captured, effectively discarding them.
	// This configuration is equivalent to redirecting output to (cmd >/dev/null 2>&1).
	default:
		// Set cmd.Stdout to nil
		// This line configures the command's standard output (stdout) to be discarded.
		// By assigning nil to cmd.Stdout, any output produced by the command will be ignored and not captured.
		cmd.Stdout = nil
		// Set cmd.Stderr to nil
		// This line configures the command's standard error (stderr) to be discarded.
		// By assigning nil to cmd.Stderr, any error messages produced by the command will be ignored and not captured.
		cmd.Stderr = nil

	}

	return cmd
}
