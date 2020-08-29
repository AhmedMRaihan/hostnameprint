package main

import (
	"fmt"
	"os"
  "time"
)

func main() {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	fmt.Println("hostname: ", name)

  t := time.Now()
  fmt.Println("\ntimenow: " + t.String())
}