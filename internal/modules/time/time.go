package time

import (
	"fmt"
	"swaystats/internal/block"
	gotime "time"
)

type Time struct{}

func New() (Time, error) {
	return Time{}, nil
}

func (t Time) Update(c chan block.Update, n string) {
	for {
		c <- block.Update{
			Name:  n,
			Props: data(),
		}
		gotime.Sleep(1 * gotime.Second)
	}
}

func data() block.Props {
	ct := gotime.Now()

	s := ct.Second()
	sep := " "
	if s%2 == 0 {
		sep = ":"
	}

	date := ct.Format("02. Jan")
	weekday := ct.Format("Monday")
	time := ct.Format(fmt.Sprintf("15%s04", sep))

	return block.Props{
		FullText: fmt.Sprintf(`%s, %s, <span color="#83a598" weight="bold">%s</span>`, date, weekday, time),
		Name:     "time",
		Markup:   "pango",
	}
}
