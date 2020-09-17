package runner

import (
	"errors"
	"runtime"

	"github.com/dop251/goja"
)

type Runner struct {
	runtime *goja.Runtime
	this    goja.Value
}

var ErrNoFunc = errors.New("manifest must export a function")

func New() *Runner {
	rt := goja.New()

	rt.Set("module", make(map[string]interface{}))
	rt.Set("console", Console{}.ToJS(rt))

	this := rt.NewObject()
	if err := this.DefineDataProperty(
		"go",
		rt.ToValue(runtime.Version()),
		goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE,
	); err != nil {
		panic(err)
	}

	return &Runner{
		runtime: rt,
		this:    this,
	}
}

func (r *Runner) Run(str string) error {
	_, err := r.runtime.RunString(str)
	if err != nil {
		return err
	}

	f, ok := goja.AssertFunction(r.runtime.Get("module").ToObject(r.runtime).Get("exports"))
	if !ok {
		return ErrNoFunc
	}

	if _, err := f(r.this); err != nil {
		return err
	}

	return nil
}
