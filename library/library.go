// library is full of constant data that is referred to everywhere.
package library

var (
	instance *Library
)

type Library struct {
}

func Instance() *Library {
	if instance == nil {
		instance = &Library{}
	}
	return instance
}
