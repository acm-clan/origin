package main

import (
	"origin/src/origin"
)

func main() {
	or := origin.Default()
	or.Handler(&Echo{})
	or.Run()
}
