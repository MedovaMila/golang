package main

import (
	gadget "learning/gadget/Page_356"
)

type Player interface {
	Play(string)
	Stop()
}

func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	recorder := player.(gadget.TapeRecorder)
	recorder.Record()
}

func main() {
	TryOut(gadget.TapeRecorder{})
}
