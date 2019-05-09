package db

// Entity basic behavior for model
type Entity interface {
	IsEmpty() bool
	IsEqual(b interface{}) bool
}
