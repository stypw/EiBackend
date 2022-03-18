package flow

type Next func()

type Step interface {
	Exec(v interface{}, n Next)
}

type StepFunc func(v interface{}, n Next)

func (f StepFunc) Exec(v interface{}, n Next) {
	f(v, n)
}

type Flow interface {
	Exec(v interface{})
	Use(step Step)
	UseFunc(f StepFunc)
}
