package cmd

import (
	"context"
	"errors"
	"io"
)

var ErrEmptyCommandArgs = errors.New("name or args is empty")

type Options struct {
	// Command specifies the actual command to execute.
	// This can be a path to an executable or a command string with arguments.
	// If arguments are included in the command string, they are automatically parsed.
	Command string
	// Args holds the command-line arguments for the executable.
	// These arguments are passed to the command when it is executed.
	// If the Command field already includes arguments, this field is ignored.
	Args []string

	parentCtx context.Context
	doneCh    chan struct{}
	// stdOutBuffer is an optional writer where a copy of the command's standard output is sent.
	// This allows real-time processing or logging of the output while the command is running.
	stdOutBuffer io.Writer
	// stdInPipeReader provides a way to connect an input stream to the command's stdin.
	// This allows feeding input to the command during its execution.
	stdInPipeReader io.ReadCloser
	// stdOutPipeWriter captures the command's standard output.
	// This allows the output to be redirected, stored, or processed by the application.
	stdOutPipeWriter io.WriteCloser
}

// SetNameAndArgs sets the command name and its arguments in the Options instance.
// It validates the inputs to ensure that the command name is not empty
// and that at least one argument is provided. If the inputs are valid,
// the method updates the Options instance and returns nil; otherwise,
// it returns an error indicating that the command name or arguments are empty.
func (opts *Options) SetNameAndArgs(commandName string, args []string) error {
	// Check if the command name is empty.
	// An empty command name indicates that no command is specified,
	// which is invalid for setting command options.
	if commandName == "" || len(args) == 0 {
		// Return an error indicating that command arguments are empty.
		// This error will be used by the calling function to handle
		// the scenario where valid command inputs are required.
		return ErrEmptyCommandArgs
	}

	// Set the Command field of the Options instance to the provided command name.
	// This updates the internal state of the Options instance
	// to reflect the command that should be executed later.
	opts.Command = commandName

	// Set the Args field of the Options instance to the provided arguments slice.
	// This stores the arguments that will be passed to the command
	// when it is executed, allowing for flexible command execution.
	opts.Args = args

	// Return nil to indicate that the operation was successful.
	// This signifies to the calling function that both the command name
	// and arguments have been set correctly, allowing further processing.
	return nil
}

// SetContext assigns the provided context to the Options instance. This allows the context to be
// used for controlling operations within the command execution, such as handling cancellations
// or timeouts. The method ensures that the context is valid (non-nil) before assigning it,
// returning an error if an invalid context is provided.
func (opts *Options) SetContext(ctx context.Context) error {
	// Check if the provided context is nil. A nil context is invalid and
	// should not be used. Return an error in this case.
	if ctx == nil {
		return errors.New("context cannot be nil")
	}

	// Set the valid context to the parentCtx field of the Options instance.
	// This will allow any tasks that rely on the Options to use the provided context
	// for operations like cancellation or timeouts.
	opts.parentCtx = ctx

	// Return nil to indicate that the context was successfully set.
	return nil
}

// SetDoneChannel assigns a done signal channel to the Options instance. This channel is used
// for signaling the completion of certain operations. The method first checks if the provided
// channel is valid (non-nil) and not already closed, returning an error if either condition is
// not met. If valid, the channel is assigned to the Options instance for future use in signaling
// when operations are completed.
func (opts *Options) SetDoneChannel(doneCh chan struct{}) error {
	// Check if the provided channel is nil. A nil channel is invalid and
	// cannot be used, so return an error in this case.
	if doneCh == nil {
		return errors.New("chan is empty")
	}

	// Use a non-blocking select statement to check if the channel is closed.
	// This ensures that the method can detect a closed channel without blocking
	// the execution, allowing for safe channel assignment.
	select {
	case <-doneCh:
		// If this case is executed, it means the channel has already been closed.
		// Return an error indicating that a closed channel cannot be set.
		return errors.New("chan is close")
	default:
		// If the channel is not closed (i.e., it's still open), proceed to set it.
		// This case will execute immediately if the channel is open.
	}

	// Assign the provided open channel to the task's doneCh field. This allows
	// the task to later use this channel for signaling that it is done.
	opts.doneCh = doneCh

	// Return nil to indicate that the channel was successfully set.
	return nil
}

// WithPipe sets up pipe connections for standard input and output streams.
// This method returns a writer for the input pipe and a reader for the output pipe, allowing communication between processes
// through pipes. Pipes are used to simulate stdin and stdout streams, often for inter-process communication.
func (opts *Options) WithPipe() (*io.PipeWriter, *io.PipeReader) {
	// Create a new pipe that consists of a reader and a writer for the input stream.
	// `io.Pipe()` sets up an in-memory pipe where one side writes to it and the other reads from it.
	inputPipeReader, inputPipeWriter := io.Pipe()

	// Assign the input pipe reader to the `stdInPipeReader` field of the `Options` struct.
	// This allows the `Options` struct to hold a reference to the input side of the pipe,
	// which can later be used to simulate or manage stdin in a process.
	opts.stdInPipeReader = inputPipeReader

	// Create another pipe that consists of a reader and a writer for the output stream.
	// This will allow output from a process to be captured by reading from the pipe.
	outputPipeReader, outputPipeWriter := io.Pipe()

	// Assign the output pipe writer to the `stdOutPipeWriter` field of the `Options` struct.
	// This will be used to write the standard output of the process to the pipe, where it can be read later.
	opts.stdOutPipeWriter = outputPipeWriter

	// Return the writer for the input pipe and the reader for the output pipe.
	// These will be used for writing data into the input and reading data from the output, respectively,
	// providing a way to send data into the process and capture its output.
	return inputPipeWriter, outputPipeReader
}

// WithStdOutBuffer sets the output buffer for standard output in the `Options` instance.
// This method allows redirection of standard output to the provided `io.Writer` buffer,
// which could be used for various operations such as capturing output or processing data streams.
func (opts *Options) WithStdOutBuffer(buf io.Writer) error {
	// Check if the provided buffer is nil, which would indicate that no valid writer was passed.
	// If `buf` is nil, the function returns an error to prevent setting an invalid output destination.
	if buf == nil {
		// Return an error indicating that the writer is empty.
		// This helps signal to the caller that a valid `io.Writer` must be provided.
		return errors.New("writer is empty")
	}

	// Assign the provided `io.Writer` to the `stdOutBuffer` field of the `Options` instance.
	// This sets the custom output destination for future operations that rely on this buffer.
	opts.stdOutBuffer = buf

	// Return nil to indicate that the operation was successful and no errors were encountered.
	// The `Options` instance is now configured with the specified output buffer.
	return nil
}
