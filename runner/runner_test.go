package runner_test

import (
	"testing"

	"github.com/DonyaOS/PackageManager/runner"
	"github.com/stretchr/testify/assert"
)

func TestRunner(t *testing.T) {
	r := runner.New()

	script := `
version = "13"
arch = "amd64"

sources = [
	"https://httpbin.org/bytes/" + version,
]

function install() {
	console.log(this.go)
	console.log(this.version)
	console.log('Hello World')

	console.log(this.files[0])
	exec.run(['cat', this.files[0]])
}
`

	assert.NoError(t, r.Run(script))
}
