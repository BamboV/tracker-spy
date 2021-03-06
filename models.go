package tracker_spy

type Tracker struct{
	ID uint `gorm:"primary_key" json:"id"`
	TrackerID int `gorm:"size:255;unique_index" json:"tracker_id"`
	Name string `json:"name"`
	Source string `json:"source"`
	LastUpdate string `json:"last_update"`
}
