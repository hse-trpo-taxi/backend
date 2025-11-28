package models

type OrderWeekStat struct {
	Mon int `json:"mon"`
	Tue int `json:"tue"`
	Wed int `json:"wed"`
	Thu int `json:"thu"`
	Fri int `json:"fri"`
	Sat int `json:"sat"`
	Sun int `json:"sun"`
}

type CurrentStat struct {
	Order     int `json:"order"`
	Free      int `json:"free"`
	Technical int `json:"technical"`
}
