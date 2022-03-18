package flow

type flow struct {
	steps []Step
	idx   int
	v     interface{}
}

func (f *flow) exec() {
	if f.idx >= len(f.steps) {
		return
	}
	step := f.steps[f.idx]
	f.idx++
	step.Exec(f.v, f.exec)
}

func (f *flow) Exec(v interface{}) {
	f.v = v
	f.idx = 0
	f.exec()
}

func (f *flow) Use(step Step) {
	f.steps = append(f.steps, step)
}

func (f *flow) UseFunc(sf StepFunc) {
	f.steps = append(f.steps, sf)
}

func Create() Flow {
	return &flow{steps: make([]Step, 0), idx: 0, v: nil}
}
