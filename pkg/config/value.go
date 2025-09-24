package config

import (
	"time"

	"github.com/gomeal/config/pkg/opt"
)

var _ Value = &valueImpl{}

type valueImpl struct {
	optBool     opt.Opt[bool]
	optInt      opt.Opt[int]
	optString   opt.Opt[string]
	optDuration opt.Opt[time.Duration]
}

func newValue(v any) *valueImpl {
	switch t := v.(type) {
	case bool:
		return &valueImpl{optBool: opt.Some(t)}
	case int:
		return &valueImpl{optInt: opt.Some(t)}
	case string:
		return &valueImpl{optString: opt.Some(t)}
	case time.Duration:
		return &valueImpl{optDuration: opt.Some(t)}
	default:
		panic("invalid value type in config.Value")
	}
}

func (v *valueImpl) Bool() bool {
	return v.optBool.Get()
}

func (v *valueImpl) Int() int {
	return v.optInt.Get()
}

func (v *valueImpl) String() string {
	return v.optString.Get()
}

func (v *valueImpl) Duration() time.Duration {
	return v.optDuration.Get()
}
