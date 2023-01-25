package lib

import (
	"fmt"

	display "github.com/djdhm/external-dep/lib"
)

const Version = "v1.0.0"

func TestLib() {
	fmt.Println("I am moab-dep-2 version " + Version)
	fmt.Println("Using external deps version " + display.Version())
}
