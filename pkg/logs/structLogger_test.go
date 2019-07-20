package logs

import (
	"testing"

	"github.com/franela/goblin"
)

func TestStructLoggerIsLogger(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("StructLogger", func() {
		g.It("implements Logger interface", func() {
			var _ Logger = NewStructLogger(StructLoggerParams{})
		})
	})
}
