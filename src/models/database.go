package models

import "gorm.io/gorm"

type New struct {
	gorm.Model
	URL      string
	Category string
}

type NewDetail struct {
	gorm.Model
	URL        string
	LastMod    string
	ChangeFreq string
	Title      string
}
