package mutex

import "errors"

var (
	ErrInvalidConnection = errors.New("invalid mutex connection")
	ErrNotReady          = errors.New("mutex is not ready")
)
