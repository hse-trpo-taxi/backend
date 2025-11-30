package models

import "time"

// Driver represents a taxi driver in the service.
// It contains personal information, licensing details, and performance metrics
// for drivers who provide taxi services.
type Driver struct {
	// ID is the unique identifier for the driver
	ID int `json:"id" db:"id"`
	// Name is the full name of the driver
	Name string `json:"name" db:"name"`
	// Phone is the driver's contact phone number
	Phone string `json:"phone" db:"phone"`
	// LicenseNumber is the driver's license number for verification
	LicenseNumber string `json:"license_number" db:"license_number"`
	// Rating is the driver's average rating from clients (0.0 to 5.0)
	Rating float64 `json:"rating" db:"rating"`
	// CreatedAt is the timestamp when the driver record was created
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	// UpdatedAt is the timestamp when the driver record was last modified
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	Status string `json:"status" db:"status"`

	TimeWork int `json:"time_work" db:"time_work"`

	X float64 `json:"x" db:"x"`

	Y float64 `json:"y" db:"y"`

	ScoreDriver float64 `json:"scoreDriver" db:"score_driver"`

	Passport int `json:"passport" db:"passport"`

	Snils int `json:"snils" db:"snils"`

	Inn int `json:"inn" db:"inn"`
}

type CreateDriverModel struct {
	Name          string  `json:"name" db:"name"`
	Phone         string  `json:"phone" db:"phone"`
	LicenseNumber string  `json:"license_number" db:"license_number"`
	Rating        float64 `json:"rating" db:"rating"`
}

func (model *CreateDriverModel) Validate() bool {
	return model.Rating >= 0
}

type UpdateDriverModel struct {
	Name          string  `json:"name" db:"name"`
	Phone         string  `json:"phone" db:"phone"`
	LicenseNumber string  `json:"license_number" db:"license_number"`
	Rating        float64 `json:"rating" db:"rating"`
}

type DriversMeanScore struct {
	Score       float64 `json:"score"`
	TimeWaiting float64 `json:"timeWaiting"`
	TimeFree    float64 `json:"timeFree"`
}

type DriverCoords struct {
	ID   int     `json:"id"`
	Name string  `json:"name"`
	X    float64 `json:"x"`
	Y    float64 `json:"y"`
}

type DriverStat struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Status   string `json:"status"`
	Phone    string `json:"phone"`
	TimeWork int    `json:"time_work"`
}

type DriverStatRequestModel struct {
	Skip   uint64 `json:"skip"`
	Limit  uint64 `json:"limit"`
	Status string `json:"status"`
}

type DriverShortModel struct {
	Name        string  `json:"name"`
	Phone       string  `json:"phone"`
	Status      string  `json:"status"`
	ScoreDriver float64 `json:"scoreDriver"`
	Passport    int     `json:"passport"`
	Snils       int     `json:"snils"`
	Inn         int     `json:"inn"`
}

type DriverResponseModel struct {
	Car    *CarDriverModel   `json:"car"`
	Driver *DriverShortModel `json:"driver"`
}

type DriverShortInfo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// DriverSchedule represents a single work day for a driver.
// Date stores the day the driver works (DATE in DB).
type DriverSchedule struct {
	Date time.Time `json:"date" db:"date"`
}

// DriverScheduleRequestModel contains optional start and end dates
// to request a driver's schedule. If zero, defaults are applied
// (start = today, end = start + 1 month) by the repository.
type DriverScheduleRequestModel struct {
	Start time.Time `json:"start" db:"start"`
	End   time.Time `json:"end" db:"end"`
}
