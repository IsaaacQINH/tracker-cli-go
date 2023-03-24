package helper

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func FormatSeasonString(seasonUID string) string {
	seasonUID = strings.Replace(seasonUID, "e", "Episode ", 1)
	seasonUID = strings.Replace(seasonUID, "a", ": Act ", 1)

	return seasonUID
}

func ArrowFromMMRChange(change int64) string {
	if change > 0 {
		return color.GreenString(fmt.Sprintf("⬆ %d", change))
	} else if change < 0 {
		return color.RedString(fmt.Sprintf("⬇ %d", change))
	} else {
		return "-"
	}

}
