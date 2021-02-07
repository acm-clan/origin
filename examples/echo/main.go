package main

import (
	"origin/src/origin"
)

func main() {
	or := origin.Default()
	or.Module(&Echo{})
	or.Run()
}
