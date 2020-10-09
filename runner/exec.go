package runner

import (
	"io"
	"os"
	"os/exec"

	"github.com/dop251/goja"
	"github.com/sirupsen/logrus"
)

type Exec struct{}

func (e Exec) Run(rt *goja.Runtime) func(goja.Value) {
	return func(cmd goja.Value) {
		var args []string

		logger := logrus.New()
		logger.SetLevel(logrus.InfoLevel)
		logger.SetOutput(os.Stdout)

		if err := rt.ExportTo(cmd, &args); err != nil {
			logrus.Error(err)

			return
		}

		//nolint:gosec
		command := exec.Command(args[0], args[1:]...)

		pipe, err := command.StdoutPipe()
		if err != nil {
			return
		}

		if err := command.Start(); err != nil {
			logrus.Error(err)
		}

		w := logger.WithField("command", args[0]).Writer()
		if _, err := io.Copy(w, pipe); err != nil {
			logrus.Error(err)
		}

		if err := w.Close(); err != nil {
			logrus.Error(err)
		}

		if err := command.Wait(); err != nil {
			logrus.Error(err)
		}
	}
}

func (e Exec) ToJS(rt *goja.Runtime) goja.Value {
	obj := rt.NewObject()

	if err := obj.Set("run", e.Run(rt)); err != nil {
		panic(err)
	}

	return obj
}
