package chain

type BaseHandler[A any] interface {
	SetNext(handler BaseHandler[A])
	Execute(id string) A
}
