package runner_test

import (
	"testing"

	"github.com/DonyaOS/PackageManager/runner"
	"github.com/stretchr/testify/assert"
)

func TestRunner(t *testing.T) {
	r := runner.New()

	script := `
module.exports = function() {
	console.log(this.go)
	console.log('Hello World')
}
`

	assert.NoError(t, r.Run(script))
}
