package types

type IEnvironment interface {
	GetName() string
	GetDebug() string
}

// Serve information
type Environment struct {
	Name  string
	Debug string
}

func (e *Environment) GetName() string {
	return e.Name
}

func (e *Environment) GetDebug() string {
	return e.Debug
}
