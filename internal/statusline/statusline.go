package statusline

import (
	"encoding/json"
	"fmt"
	"log"
	"swaystats/internal/block"
)

type StatusLine struct {
	Blocks  map[string]block.Props
	Order   []string
	Updates chan block.Update
}

func New() StatusLine {
	return StatusLine{
		Blocks:  make(map[string]block.Props),
		Updates: make(chan block.Update),
	}
}

func (sl *StatusLine) Register(block block.Block, n string) {
	_, exists := sl.Blocks[n]
	if !exists {
		go block.Update(sl.Updates, n)
		sl.Order = append(sl.Order, n)
	}
}

func (sl *StatusLine) Render() {
	for {
		update := <-sl.Updates
		existing, exists := sl.Blocks[update.Name]

		if exists && existing.FullText == update.Props.FullText {
			return
		}

		sl.Blocks[update.Name] = update.Props

		blocks := []block.Props{}

		for _, n := range sl.Order {
			blocks = append(blocks, sl.Blocks[n])
		}

		r, err := json.Marshal(blocks)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(string(r), ",")
	}
}
