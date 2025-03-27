package model

import (
	"time"
)

type Jobs struct {
	uuid             string
	title            string
	company_name     string
	location         string
	currentStatus    string
	salaryLowerBound int
	salaryUpperBound int
	datePosted       time.Time
	dateApplied      time.Time
}
