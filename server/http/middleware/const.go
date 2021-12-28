package middleware

type (
	processID struct{}
)

const (
	CurrentProcess = "currentProcess"
)

var (
	ProcessKey = &processID{}
)
