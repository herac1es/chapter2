package main

import (
	"log"
	"os"

	_ "github.com/iwillbesober/chapter2/matchers"
	"github.com/iwillbesober/chapter2/search"
)

func init() {
	log.SetOutput(os.Stdout)
}
func main() {
	search.Run("coronavirus")
}
