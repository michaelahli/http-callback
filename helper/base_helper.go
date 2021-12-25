package helper

type helper struct{}

type Helper interface {
	SetUp()
}

func New() Helper {
	return &helper{}
}
