package runner

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"

	"github.com/dop251/goja"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

type Runner struct {
	runtime *goja.Runtime
	this    *goja.Object
	client  *resty.Client
}

var ErrNoInstall = errors.New("there must be an install function")

func New() *Runner {
	rt := goja.New()

	rt.Set("console", Console{}.ToJS(rt))
	rt.Set("exec", Exec{}.ToJS(rt))

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
		client:  resty.New().SetHeader("User-Agent", "donya-pkg"),
	}
}

// nolint: funlen
func (r *Runner) Run(str string) error {
	_, err := r.runtime.RunString(str)
	if err != nil {
		return err
	}

	arch := r.runtime.Get("arch").String()
	if arch != runtime.GOARCH {
		logrus.Warnf("installation script has arch = %s but your current arch is %s", arch, runtime.GOARCH)
	}

	version := r.runtime.Get("version").String()
	if err := r.this.DefineDataProperty(
		"version",
		r.runtime.ToValue(version),
		goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE,
	); err != nil {
		panic(err)
	}

	var sources []string
	if err := r.runtime.ExportTo(r.runtime.Get("sources"), &sources); err != nil {
		logrus.Errorf("sources: %s", err.Error())
	}

	files := make([]string, len(sources))

	for i, source := range sources {
		logrus.Info(source)

		resp, err := r.client.R().Get(source)
		if err != nil {
			logrus.Error(err)
		}

		content := resp.Body()

		tmpfile, err := ioutil.TempFile("", fmt.Sprintf("s_%d_*", i))
		if err != nil {
			logrus.Error(err)

			continue
		}

		if _, err := tmpfile.Write(content); err != nil {
			logrus.Error(err)
		}

		if err := tmpfile.Close(); err != nil {
			logrus.Error(err)
		}

		files[i] = tmpfile.Name()
	}

	if err := r.this.DefineDataProperty(
		"files",
		r.runtime.ToValue(files),
		goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE,
	); err != nil {
		panic(err)
	}

	f, ok := goja.AssertFunction(r.runtime.Get("install"))
	if !ok {
		return ErrNoInstall
	}

	if _, err := f(r.this); err != nil {
		return err
	}

	for _, file := range files {
		if err := os.Remove(file); err != nil {
			logrus.Error(err)
		}
	}

	return nil
}
