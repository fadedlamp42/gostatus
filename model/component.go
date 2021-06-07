package model

// Component specifies a simple interface for renderable elements
type Component interface {
	Render() string
}
