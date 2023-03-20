package wrapper

// Wrapper defines a wrap for errors
type Wrapper interface {
	// error returns the string value of the embedded error
	error
	// Unwrap returns the main error
	Unwrap() error
}

// New wraps two errors in a Wrapper instance
func New(err, embedded error) Wrapper {
	return &wrapper{
		err:   err,
		error: embedded,
	}
}

type wrapper struct {
	// error is the embedded error
	error
	// err is the main error returned in Unwrap method
	err error
}

// Unwrap returns the main error
func (w *wrapper) Unwrap() error {
	return w.err
}
