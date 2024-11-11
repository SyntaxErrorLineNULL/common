package common

import "reflect"

// GetRecoverError extracts an error from a recoverable panic.
// It checks if the recovered value is an error type, and if so, returns it.
// If the recovered value is not an error type, it returns nil.
func GetRecoverError(rec any) error {
	// Check if recoverable value is not nil
	if rec != nil {
		// Type switch on the recovered value
		switch e := rec.(type) {
		// If recovered value is of type error
		case error:
			return e

		// If recovered value is of any other type
		default:
			return nil
		}
	} else {
		// If recoverable value is nil
		return nil
	}
}

// GetType takes an interface{} as an argument and returns its reflect.Type.
// This function is useful for obtaining the dynamic type of the provided value,
// even if the value is a pointer or an interface itself.
func GetType(v interface{}) reflect.Type {
	// Check if the provided value is nil.
	// If the input is nil, return nil immediately since there is no type information.
	if v == nil {
		return nil
	}

	// Use reflect.ValueOf to obtain the reflection value of the provided interface.
	// Then use reflect.Indirect to get the value that the interface points to,
	// effectively dereferencing it if it is a pointer. Finally, retrieve the type
	// of the dereferenced value using the Type method.
	return reflect.Indirect(reflect.ValueOf(v)).Type()
}
