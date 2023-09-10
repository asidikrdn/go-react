package helpers

import "go-react/server/models"

func GetLastID() uint {
	var lastID uint

	lastID = 0

	for _, mhs := range models.StudentData {
		if mhs.ID > lastID {
			lastID = mhs.ID
		}
	}

	return lastID
}
