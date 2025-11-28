package models

import (
	"time"
)

type SupportRequest struct {
	Id                 int       `json:"id" db:"id"`
	Name               string    `json:"Name" db:"Name"`
	Comment            string    `json:"comment" db:"comment"`
	TimeReactingScore  int       `json:"timeReactingScore" db:"time_reacting_score"`
	QualityAnswerScore int       `json:"qualityAnswerScore" db:"quality_answer_score"`
	Date               time.Time `json:"date" db:"date"`
}

type SupportModel struct {
	Name               string
	Comment            string
	TimeReactingScore  int
	QualityAnswerScore int
	Date               time.Time
}
