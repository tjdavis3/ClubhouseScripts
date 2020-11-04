package openapi

import (
	"time"
)

func (story StorySlim) GetStartDate() *time.Time {
	defaultStart := time.Now().AddDate(0, 0, 1) // Set the default start date to tomorrow
	startTime := &defaultStart
	if story.Started {
		startTime = story.StartedAt
		if story.StartedAtOverride != nil {
			startTime = story.StartedAtOverride
		}
	}
	return startTime
}

func (story StorySlim) GetEndDate(epicStart *time.Time) *time.Time {
	now := time.Now()
	endTime := &now

	if story.Completed {
		endTime = story.CompletedAt
		if story.CompletedAtOverride != nil {
			endTime = story.CompletedAt
		}
	} else {
		if story.Deadline != nil && story.Deadline.After(*endTime) {
			endTime = story.Deadline
		} else {
			estEnd := story.GetStartDate().AddDate(0, 0, story.GetEstimatedDays(epicStart))
			if estEnd.Before(now) {
				estEnd = now.AddDate(0, 0, 10)
			}
			endTime = &estEnd
		}
	}
	return endTime
}

func (story StorySlim) GetEstimatedDays(epicStart *time.Time) int {
	var estimate int
	if story.Estimate != nil && *story.Estimate > 0 {
		estimate = int(*story.Estimate)
	} else {
		var endDate *time.Time
		if story.Deadline != nil {
			endDate = story.Deadline
		} else {
			estEnd := time.Now().AddDate(0, 0, 10)
			endDate = &estEnd
		}
		startDate := story.GetStartDate()
		if startDate.Before(*epicStart) {
			startDate = epicStart
		}
		// If the story is past due and the start date is in the future, force estimate to 10 days
		if startDate.After(time.Now()) && endDate.Before(time.Now()) {
			newStart := startDate.AddDate(0, 0, 10)
			endDate = &newStart
		}
		estimate = int(endDate.Sub(*startDate).Hours() / 24)
	}
	return estimate
}

type ByStart []StorySlim

func (a ByStart) Len() int           { return len(a) }
func (a ByStart) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByStart) Less(i, j int) bool { return a[i].GetStartDate().Before(*a[j].GetStartDate()) }
