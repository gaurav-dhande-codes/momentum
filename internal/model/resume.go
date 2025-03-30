package model

import (
	"time"
)

type Resume struct {
	uuid        string
	title       string
	dateCreated time.Time
	dateUpdated time.Time
}
