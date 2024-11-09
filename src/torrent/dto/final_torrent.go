package dto

import (
	"tm/src/transmission/dto"
)

type FinalTorrent struct {
	Torrent             *Torrent                 `json:"db"`
	TransmissionTorrent *dto.TransmissionTorrent `json:"transmission"`
}

type FinalTorrentsList struct {
	FinalTorrentArray []*FinalTorrent
	FinalTorrentCount int
}
