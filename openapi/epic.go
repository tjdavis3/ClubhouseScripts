package openapi

import (
	"time"
)

func (epic Epic) GetStartDate() *time.Time {
	now := time.Now()
	startDate := &now
	if epic.Started {
		startDate = epic.StartedAt
		if epic.StartedAtOverride != nil {
			startDate = epic.StartedAtOverride
		}
	} else {
		startDate = epic.CreatedAt
	}
	return startDate
}
