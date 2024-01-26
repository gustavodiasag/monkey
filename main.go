package main

import (
    "os"
    "monkey/repl"
)

func main() {
   repl.Start(os.Stdin, os.Stdout)
}
