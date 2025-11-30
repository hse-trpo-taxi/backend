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

type Order struct {
	Id            int     `json:"id" db:"id"`
	DriverId      int     `json:"driverId" db:"driver_id"`
	FirstAddress  string  `json:"firstAddress" db:"first_address"`
	SecondAddress string  `json:"secondAddress" db:"second_address"`
	Current       bool    `json:"current" db:"current"`
	Passengers    int     `json:"passengers" db:"passengers"`
	TimeClose     int     `json:"timeClose" db:"time_close"`
	Score         int     `json:"score" db:"score"`
	X             float64 `json:"x" db:"x"`
	Y             float64 `json:"y" db:"y"`
}
type OrderDriverList struct {
	Id            int    `json:"id" db:"id"`
	FirstAddress  string `json:"firstAddress" db:"first_address"`
	SecondAddress string `json:"secondAddress" db:"second_address"`
	Current       bool   `json:"current" db:"current"`
}

type Point struct {
	X float64 `json:"x" db:"x"`
	Y float64 `json:"y" db:"y"`
}

type OrderDriverInfo struct {
	Time       int `json:"time" db:"time"`
	Passengers int `json:"passengers" db:"passengers"`
	TimeClose  int `json:"timeClose" db:"time_close"`
	current    *Point
}

type CurrentOrderInfo struct {
	Info *OrderDriverInfo
	X    float64 `json:"x" db:"x"`
	Y    float64 `json:"y" db:"y"`
}

type OrderClientList struct {
	Id            int    `json:"id" db:"id"`
	FirstAddress  string `json:"firstAddress" db:"first_address"`
	SecondAddress string `json:"secondAddress" db:"second_address"`
	Driver        *DriverShortModel
	Score         int `json:"score" db:"score"`
}

type OrderClientListRequestModel struct {
	Skip  uint64 `json:"skip"`
	Limit uint64 `json:"limit"`
}
