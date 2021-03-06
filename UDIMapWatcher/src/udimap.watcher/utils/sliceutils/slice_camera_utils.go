package sliceutils

import (
	"udimap.watcher/models"
)

func ContainCamera(cameras []*models.Camera, camera *models.Camera) bool {
	for _, c := range cameras {
		if c.ID == camera.ID {
			return true
		}
	}
	return false
}

func IndexOfCamera(camera *models.Camera, cameras []*models.Camera) int {
	for i, c := range cameras {
		if c.ID == camera.ID {
			return i
		}
	}
	return -1
}

func ContainTramQuanTracNgap(list []*models.TramQuanTracNgap, item *models.TramQuanTracNgap) bool {
	for _, c := range list {
		if c.ID == item.ID {
			return true
		}
	}
	return false
}

func IndexOfTramQuanTracNgap(list []*models.TramQuanTracNgap, item *models.TramQuanTracNgap) int {
	for i, c := range list {
		if c.ID == item.ID {
			return i
		}
	}
	return -1
}

func ContainFloodPoint(list []*models.FloodPoint, item *models.FloodPoint) bool {
	for _, c := range list {
		if c.ID == item.ID {
			return true
		}
	}
	return false
}

func IndexOfFloodPoint(list []*models.FloodPoint, item *models.FloodPoint) int {
	for i, c := range list {
		if c.ID == item.ID {
			return i
		}
	}
	return -1
}

func ContainTramDoMua(listTram []*models.TramDoMua, tram *models.TramDoMua) bool {
	for _, c := range listTram {
		if c.ID == tram.ID {
			return true
		}
	}
	return false
}

func IndexOfTramDoMua(tram *models.TramDoMua, listTram []*models.TramDoMua) int {
	for i, c := range listTram {
		if c.ID == tram.ID {
			return i
		}
	}
	return -1
}

func ContainTramDoTrieu(listTram []*models.TramDoTrieu, tram *models.TramDoTrieu) bool {
	for _, c := range listTram {
		if c.IDText == tram.IDText {
			return true
		}
	}
	return false
}

func IndexOfTramDoTrieu(tram *models.TramDoTrieu, listTram []*models.TramDoTrieu) int {
	for i, c := range listTram {
		if c.ID == tram.ID {
			return i
		}
	}
	return -1
}
