package cmd

import (
	"context"
	"errors"
)

type Options struct {
	ctx    context.Context
	doneCh chan struct{}
}

func (o *Options) SetContext(ctx context.Context) error {
	// Check if the provided context is nil. A nil context is invalid and
	// should not be used. Return an error in this case.
	if ctx == nil {
		return errors.New("context cannot be nil")
	}

	// Assign the provided valid context to the task's parentCtx field.
	o.ctx = ctx

	// Return nil to indicate that the context was successfully set.
	return nil
}

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
