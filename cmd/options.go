package cmd

import (
	"context"
	"errors"
)

type Options struct {
	ctx context.Context
}

func (o *Options) WithContext(ctx context.Context) error {
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
