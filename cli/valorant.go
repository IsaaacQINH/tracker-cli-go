package cli

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/tracker-cli-go/pkg/api"
)

func RunCommand(args CommadArgs) {
	vapi := api.ValorantApi{
		Client: api.Client{
			Host:    "https://api.henrikdev.xyz",
			Timeout: 30,
		},
	}

	fmt.Print("\n")
	fmt.Printf("Fetching player data for %s#%s", color.CyanString(args.Player), color.GreenString(args.Tagline))

	accountData := vapi.GetPlayer(args.Player, args.Tagline)
	statsData := vapi.GetMMRData(args.Region, args.Player, args.Tagline)

	fmt.Printf("\nPlayer UUID: %s", accountData.Data.Puuid)
	fmt.Printf("\nAccount Level: %d", accountData.Data.AccountLevel)
	fmt.Printf("\nElo: %d", statsData.Data.CurrentData.Elo)
	fmt.Printf("\nCurrent Rank: %s", statsData.Data.CurrentData.CurrentTierString)
	fmt.Print("\n")
}
