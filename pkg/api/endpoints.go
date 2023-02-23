package api

import (
	"encoding/json"
	"fmt"
)

type ValorantApi struct {
	Client Client
}

func (api *ValorantApi) GetPlayer(player string, tagline string) PlayerData {
	res := api.Client.Request(fmt.Sprintf("/valorant/v1/account/%s/%s", player, tagline))

	pData := PlayerData{}
	json.Unmarshal(res, &pData)

	return pData
}

func (api *ValorantApi) GetMMRData(region string, player string, tagline string) MMRData {
	res := api.Client.Request(fmt.Sprintf("/valorant/v2/mmr/%s/%s/%s", region, player, tagline))

	mmrData := MMRData{}
	json.Unmarshal(res, &mmrData)

	return mmrData
}
