package api

type PlayerData struct {
	Status int64  `json:"status"`
	Data   Player `json:"data"`
}

type MMRData struct {
	Status int64 `json:"status"`
	Data   MMR   `json:"data"`
}

type Player struct {
	Puuid         string `json:"puuid"`
	Region        string `json:"region"`
	AccountLevel  int64  `json:"account_level"`
	Name          string `json:"name"`
	Tag           string `json:"tag"`
	Card          Card   `json:"card"`
	LastUpdate    string `json:"last_update"`
	LastUpdateRaw string `json:"last_update_raw"`
}

type Card struct {
	Id    string `json:"id"`
	Small string `json:"small"`
	Large string `json:"large"`
	Wide  string `json:"wide"`
}

type MMR struct {
	Puuid         string               `json:"puuid"`
	Name          string               `json:"name"`
	Tag           string               `json:"tag"`
	CurrentData   CurrentData          `json:"current_data"`
	HighestRank   HighestRank          `json:"highest_rank"`
	SeasonSummary map[string]SeasonMMR `json:"by_season"`
}

type CurrentData struct {
	CurrentTier       int64             `json:"currenttier"`
	CurrentTierString string            `json:"currenttierpatched"`
	Images            CurrentDataImages `json:"images"`
	RankingTier       int64             `json:"ranking_in_tier"`
	LastMMRChange     int64             `json:"mmr_change_to_last_game"`
	Elo               int64             `json:"elo"`
	UnratedGamesLeft  int64             `json:"games_needed_for_rating"`
	Old               bool              `json:"old"`
}

type CurrentDataImages struct {
	Small        string `json:"small"`
	Large        string `json:"large"`
	TriangleDown string `json:"triangle_down"`
	TriangleUp   string `json:"triangle_up"`
}

type HighestRank struct {
	Old        bool   `json:"old"`
	Tier       int64  `json:"tier"`
	TierString string `json:"patched_tier"`
	Season     string `json:"season"`
}

type SeasonMMR struct {
	Error            string        `json:"error"`
	Wins             int64         `json:"wins"`
	TotalGamesPlayed int64         `json:"number_of_games"`
	FinalTier        int64         `json:"final_rank"`
	FinalTierString  string        `json:"final_rank_patched"`
	ActRankWins      []ActRankWins `json:"act_rank_wins"`
	Old              bool          `json:"old"`
}

type ActRankWins struct {
	Tier       int64  `json:"tier"`
	TierString string `json:"patched_tier"`
}
