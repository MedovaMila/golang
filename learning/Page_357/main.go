package main

import gadget "learning/gadget/Page_356"

func playList(device gadget.TapePlayer, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	player := gadget.TapePlayer{}
	mixtape := []string{"Jessie,s girl", "Whip It", "9 to 5"}
	playList(player, mixtape)
}
