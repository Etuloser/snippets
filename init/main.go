package main

import (
	"fmt"
	"init/say"
)

// mian -.
func main() {
	say.SayHi()
	fmt.Println("main func execute")
	// output:
	// init func execute
	// Hi!
	// main func execute
}
