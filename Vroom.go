package main

import (
	"fmt"
	"github.com/kumquatexpress/Vroom/helpers"
)

func main() {
	fmt.Printf("%+v\n", helpers.NewVroomOpts("test/test_opts.json"))
}
