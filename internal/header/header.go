package header

import (
	"encoding/json"
	"fmt"
	"log"
)

type Header struct {
	Version      int  `json:"version"`
	Click_events bool `json:"click_events,omitempty"`
	Cont_signal  int  `json:"cont_signal,omitempty"`
	Stop_signal  int  `json:"stop_signal,omitempty"`
}

func New(h Header) string {
	header, err := json.Marshal(h)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s%s[", string(header), string(0x0A))
}
