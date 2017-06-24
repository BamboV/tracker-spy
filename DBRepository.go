package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DBRepository struct {
	DB *gorm.DB
}

func (d DBRepository) GetAll() []Tracker{
	result := []Tracker{}
	d.DB.Find(&result)

	return result
}

func (d DBRepository) Create(tracker Tracker) Tracker{
	d.DB.Create(&tracker)
	return tracker
}

func (d DBRepository) Update(tracker Tracker) Tracker{
	d.DB.Update(&tracker)
	return tracker
}

func (d DBRepository) Delete(id uint){
	d.DB.Delete(&Tracker{ID:id})
}

