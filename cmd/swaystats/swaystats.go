package main

import (
	"fmt"
	"swaystats/internal/header"
	"swaystats/internal/statusline"

	"swaystats/internal/modules/time"
	// "swaystats/internal/modules/volume"
)

func main() {
	h := header.New(header.Header{
		Version: 1,
	})

	fmt.Print(string(h))

	sl := statusline.New()

	sl.Register(time.New(), "time")

	sl.Render()
}
