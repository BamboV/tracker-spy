package tracker_spy

import "github.com/bamboV/torrent_center/client"

type Spy struct {
	repository DataRepository
	torrent client.CenterClient
}

func NewSpy(repository DataRepository, torrent client.CenterClient) Spy {
	spy := Spy{
		repository: repository,
		torrent: torrent,
	}

	return spy
}

func (s *Spy) GetSpyList() []Tracker{
	return s.repository.GetAll()
}

func (s *Spy) AddToSpyList(tracker Tracker) Tracker {
	return s.repository.Create(tracker)
}

func (s *Spy) Delete(id uint) {
	s.repository.Delete(id)
}

func (s *Spy)UpdateAll() []Tracker{
	updated := []Tracker{}
	for _, value := range s.GetSpyList() {
		if s.update(&value) {
			updated = append(updated, value)
		}
	}
	return updated
}

func (s *Spy) update(tracker *Tracker) bool {
	remoteTorrent, err := s.torrent.GetDistribution(tracker.Source, tracker.TrackerID)

	if err != nil {
		return false
	}
	if remoteTorrent.LastUpdated != tracker.LastUpdate {
		if s.torrent.Download(remoteTorrent.MagnetLink) {
			tracker.LastUpdate = remoteTorrent.LastUpdated
			s.repository.Update(*tracker)
			return true
		}
	}
	return false
}