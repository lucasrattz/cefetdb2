package types

type ContextKey string

func (key ContextKey) Value() string {
	return string(key)
}
