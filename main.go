package main

import (
	"os"

  "github.com/iam-naveen/compiler/console"
)

func main() {
  console.Run( os.Stdin, os.Stdout )
}
