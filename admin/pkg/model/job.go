package model

import (
	"gorm.io/gorm"
)

type SqlJob struct {
	gorm.Model
	ProjectTag	string	`json:"projecttag" gorm:"not null;size:255"`
	ApiName       string `json:"apiname" gorm:"not null;unique;size:255"`
	URL           string `json:"url" gorm:"not null"`
	TestOrNot     string `json:"testornot" gorm:"not null"`
	TestFrequency int    `json:"testfrequency" gorm:"not null"`
}

type ScriptJob struct {
	gorm.Model
	ProjectTag	string	`json:"projecttag" gorm:"not null;size:255"`
	ApiName       string `json:"apiname" gorm:"not null;unique;size:255"`
	URL           string `json:"url" gorm:"not null"`
	TestOrNot     string `json:"testornot" gorm:"not null"`
	TestFrequency int    `json:"testfrequency" gorm:"not null"`
}