package main

import (
	"flag"
	"log"

	"github.com/tracker-cli-go/cli"
)

func main() {
	game := flag.String("game", "valorant", "insert game (valorant)")
	player := flag.String("name", "", "insert player name")
	tagline := flag.String("tagline", "EUW", "insert tagline")
	region := flag.String("region", "eu", "insert region (eu, na, ap, kr)")

	flag.Parse()

	if *game == "" {
		log.Fatal("Missing flag --game")
		return
	}

	if *player == "" {
		log.Fatal("Missing flag --name")
		return
	}

	if *tagline == "" {
		log.Fatal("Missing flag --tagline")
		return
	}

	if *region == "" {
		log.Fatal("Missing flag --region")
		return
	}

	switch *game {
	case "valorant":
		cli.RunCommand(cli.CommadArgs{
			Player:  *player,
			Tagline: *tagline,
			Region:  *region,
		})
	default:
		log.Fatal("Oopsies, something isn't really working")
	}
}
