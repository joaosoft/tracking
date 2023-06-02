package main

import (
	"github.com/joaosoft/tracking"
)

func main() {
	m, err := tracking.NewTracking()
	if err != nil {
		panic(err)
	}

	if err := m.Start(); err != nil {
		panic(err)
	}
}
