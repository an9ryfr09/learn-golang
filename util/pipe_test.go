package pipe

import (
	"testing"

	"github.com/go-cmd/cmd"
)

func TestPipe(t *testing.T) {
	c := cmd.NewCmd("bash", "-c", "ps aux|grep go")
	<-c.Start()
	t.Error(c.Status().Stdout)
}
