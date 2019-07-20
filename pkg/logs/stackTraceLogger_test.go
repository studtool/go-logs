package logs

import (
	"testing"

	"github.com/franela/goblin"
)

func TestStackTraceLoggerIsLogger(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("StackTraceLogger", func() {
		g.It("implements Logger interface", func() {
			var _ Logger = NewStackTraceLogger(StackTraceLoggerParams{})
		})
	})
}
