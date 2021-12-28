package helper

type helper struct{}

type Helper interface {
	SetUp(string)
}

func New() Helper {
	return &helper{}
}
