package commands

import (
	"github.com/sykesm/asciinema/asciicast"
	"github.com/sykesm/asciinema/util"
)

type PlayCommand struct {
	Player asciicast.Player
}

func NewPlayCommand() *PlayCommand {
	return &PlayCommand{
		Player: asciicast.NewPlayer(),
	}
}

func (c *PlayCommand) Execute(url string, maxWait float64) error {
	var cast *asciicast.Asciicast
	var err error

	util.WithSpinner(500, func() {
		cast, err = asciicast.Load(url)
	})

	if err != nil {
		return err
	}

	return c.Player.Play(cast, maxWait)
}
