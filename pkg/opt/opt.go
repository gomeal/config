package opt

import "errors"

var errNoValue = errors.New("the value is none in config.opt")

type Opt[T any] struct {
	value   T
	defined bool
}

func Some[T any](v T) Opt[T] {
	return Opt[T]{value: v, defined: true}
}

func (o *Opt[T]) Get() T {
	if !o.defined {
		panic(errNoValue)
	}

	return o.value
}
