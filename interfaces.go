package main

type DataRepository interface {
	GetAll() []Tracker
	Create(tracker Tracker) Tracker
	Update(tracker Tracker) Tracker
	Delete(id uint)
}
