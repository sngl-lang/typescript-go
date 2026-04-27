package testrunner

import (
	"testing"

	"github.com/sngl-lang/typescript-go/internal/testutil/baseline"
)

func TestMain(m *testing.M) {
	defer baseline.Track()()
	m.Run()
}
