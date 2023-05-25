package cli

import (
	"fmt"
	"math"

	"github.com/fatih/color"
	"github.com/isaaacqinh/tracker-cli-go/pkg/api"
	"github.com/isaaacqinh/tracker-cli-go/pkg/helper"
)

func RunCommand(args CommadArgs) {
	vapi := api.ValorantApi{
		Client: api.Client{
			Host:    "https://api.henrikdev.xyz",
			Timeout: 30,
		},
	}

	player := color.CyanString(args.Player)

	if args.Player == "Sapphic Dino" {
		player = color.MagentaString(args.Player)
	}

	fmt.Print("\n")
	fmt.Printf("Fetching player data for %s#%s", player, color.GreenString(args.Tagline))

	accountData := vapi.GetPlayer(args.Player, args.Tagline)
	statsData := vapi.GetMMRData(args.Region, args.Player, args.Tagline)

	fmt.Printf("\nAccount Level: %d", accountData.AccountLevel)
	fmt.Printf("\nElo: %d\tRR: %d", statsData.CurrentData.Elo, int64(math.Mod(float64(statsData.CurrentData.Elo), float64(100))))
	fmt.Printf("\nCurrent Rank: %s", statsData.CurrentData.CurrentTierString)
	fmt.Printf("\nLatest Change: %s", helper.ArrowFromMMRChange(statsData.CurrentData.LastMMRChange))
	fmt.Print("\n")
	fmt.Printf("\nHighest Rank: %s (%s)", statsData.HighestRank.TierString, helper.FormatSeasonString(statsData.HighestRank.Season))
	fmt.Print("\n")
}
