package mascot_test

import (
	"testing"

	"github.com/kdlug/mascot"
)

func TestMacot(t *testing.T) {
	if mascot.BestMascot() != "Go Gopher" {
		t.Fatal("Wrong Mascot")
	}
}
