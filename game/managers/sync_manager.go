package managers

import (
	"flux/game/util"
	"time"
)

type SyncManager struct {
	Playing bool
	Speed   float64

	lastTime float64
	RealTime float64
	EndTime  float64
}

func CreateSyncManager() *SyncManager {
	return &SyncManager{
		Playing:  false,
		lastTime: 0,
		RealTime: 0,
		EndTime:  util.SelectedMap.Notes[len(util.SelectedMap.Notes)-1].T + 1,
		Speed:    1,
	}
}

func (manager *SyncManager) Update(dt float64) {
	if !manager.Playing {
		return
	}

	now := float64(time.Now().UnixMicro())
	time := manager.Speed * (now - manager.lastTime) * 0.000001
	manager.lastTime = now
	manager.RealTime += time
}

func (manager *SyncManager) Start(from float64) {
	manager.lastTime = float64(time.Now().UnixMicro())
	manager.RealTime = from
	manager.Playing = true
}
