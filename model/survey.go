package model

import "time"

type Survey struct {
	ID     uint      `json:"id"`
	Date   time.Time `json:"date"`
	Mail   string    `json:"mail"`
	Place  string    `json:"place"`
	Age    uint      `json:"age"`
	Gender string    `json:"gender"`
	Option string    `json:"option"`
}

type Surveys []Survey
