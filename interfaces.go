package tracker_spy

type DataRepository interface {
	GetAll() []Tracker
	Create(tracker Tracker) Tracker
	Update(tracker Tracker) Tracker
	Delete(id uint)
}
