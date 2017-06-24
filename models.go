package main

type Tracker struct{
	ID uint `gorm:"primary_key"`
	TrackerID uint `gorm:"size:255;unique_index"`
	Name string
	Source string
	LastUpdate string
}
