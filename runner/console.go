package runner

import (
	"github.com/dop251/goja"
	"github.com/sirupsen/logrus"
)

type Console struct{}

func (c Console) Log(msg goja.Value) {
	logrus.Info(msg)
}

func (c Console) ToJS(rt *goja.Runtime) goja.Value {
	obj := rt.NewObject()

	if err := obj.Set("log", c.Log); err != nil {
		panic(err)
	}

	return obj
}
