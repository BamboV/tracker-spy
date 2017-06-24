package main

type Spy struct {
	repository DataRepository
	torrent TorrentClient
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
		if s.update(value) {
			updated = append(updated, value)
		}
	}
	return updated
}

func (s *Spy) update(tracker Tracker) bool {
	remoteTorrent := s.torrent.GetTorrent(tracker.Source, tracker.TrackerID)
	if remoteTorrent.LastUpdate != tracker.LastUpdate {
		if s.torrent.Download(tracker) {
			tracker.LastUpdate = remoteTorrent.LastUpdate
			s.repository.Update(tracker)
			return true
		}
	}
	return false
}