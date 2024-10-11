package cmd

import (
	"context"
	"errors"
)

type Options struct {
	parentCtx context.Context
	doneCh    chan struct{}
}

// SetContext assigns the provided context to the Options instance. This allows the context to be
// used for controlling operations within the command execution, such as handling cancellations
// or timeouts. The method ensures that the context is valid (non-nil) before assigning it,
// returning an error if an invalid context is provided.
func (o *Options) SetContext(ctx context.Context) error {
	// Check if the provided context is nil. A nil context is invalid and
	// should not be used. Return an error in this case.
	if ctx == nil {
		return errors.New("context cannot be nil")
	}

	// Set the valid context to the parentCtx field of the Options instance.
	// This will allow any tasks that rely on the Options to use the provided context
	// for operations like cancellation or timeouts.
	o.parentCtx = ctx

	// Return nil to indicate that the context was successfully set.
	return nil
}

// SetDoneChannel assigns a done signal channel to the Options instance. This channel is used
// for signaling the completion of certain operations. The method first checks if the provided
// channel is valid (non-nil) and not already closed, returning an error if either condition is
// not met. If valid, the channel is assigned to the Options instance for future use in signaling
// when operations are completed.
func (o *Options) SetDoneChannel(doneCh chan struct{}) error {
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
	o.doneCh = doneCh

	// Return nil to indicate that the channel was successfully set.
	return nil
}
