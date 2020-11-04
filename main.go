package main

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"time"

	"./plantuml"

	"./openapi"
	"github.com/go-chi/chi"
)

const apiKey = "5dc0d638-3eeb-4b5b-91de-ebd6f1617d87"

func main() {
	r := chi.NewRouter()
	r.Get("/epics/{epicId}", handler)
	http.ListenAndServe(":8000", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	outputType := "text"
	epicId, err := strconv.Atoi(chi.URLParam(r, "epicId"))
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	timeline, err := buildTimeline(int64(epicId))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if outputType == "image" {
		puml := plantuml.NewServer("")
		diagram, err := puml.GetDiagramImage(timeline)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		w.Header().Set("content-type", "image/png")
		w.Write(diagram)
	} else {
		w.Header().Set("content-type", "text/plain")
		w.Write(timeline)
	}
}

func buildTimeline(epicId int64) ([]byte, error) {
	var buf bytes.Buffer
	config := openapi.NewConfiguration()
	config.AddDefaultHeader("Clubhouse-Token", apiKey)
	config.BasePath = "https://api.clubhouse.io"
	config.Debug = false
	client := openapi.NewAPIClient(config)

	epic, _, err := client.DefaultApi.ApiV3EpicsEpicPublicIdGet(context.Background(), epicId)
	if err != nil {
		return nil, err
	}
	stories, _, err := client.DefaultApi.ApiV3EpicsEpicPublicIdStoriesGet(context.Background(), epicId, openapi.GetEpicStories{IncludesDescription: false})
	if err != nil {
		return nil, err
	}
	fmt.Fprintln(&buf, "@startgantt")
	fmt.Fprintf(&buf, "title %s Stories\n", epic.Name)
	startDate := epic.GetStartDate().AddDate(0, 0, -30)
	fmt.Fprintln(&buf, "Project starts "+startDate.Format("2006/01/02"))
	fmt.Fprintln(&buf, "caption Project starts "+epic.GetStartDate().Format("01/02/2006"))
	fmt.Fprintln(&buf, `printscale daily
saturday are closed
sunday are closed`)
	fmt.Fprintln(&buf, "today is "+time.Now().Format("2006-01-02")+" and is colored in #AAF")
	epicStart := epic.GetStartDate().Format("2006/01/02")
	fmt.Fprintf(&buf, "%s to %s are colored Blue\n", epicStart, epicStart)
	fmt.Fprintf(&buf, "%s to %s are named [Epic Start]\n", epicStart, epicStart)
	if epic.Deadline != nil {
		color := "Blue"
		if epic.Deadline.Before(time.Now()) {
			color = "Red"
		}
		deadline := epic.Deadline.Format("2006/01/02")
		fmt.Fprintf(&buf, "%s to %s are colored %s\n", deadline, deadline, color)
		fmt.Fprintf(&buf, "%s to %s are named [Epic Due]\n", deadline, deadline)
	}
	var constraints []string
	sort.Sort(openapi.ByStart(stories))
	for _, story := range stories {
		fmt.Fprint(&buf, "[", story.Name, "] as [story_", story.Id, "] ends ", story.GetEndDate(epic.GetStartDate()).Format("2006-01-02"), "\n")
		fmt.Fprintf(&buf, "[story_%d] lasts %d days\n", story.Id, story.GetEstimatedDays(epic.GetStartDate()))
		fmt.Fprintln(&buf, getLink(story))
		fmt.Fprintln(&buf, setColor(story, epic))
		if story.Deadline != nil {
			fmt.Fprintf(&buf, "[%d Due] happens %s\n", story.Id, story.Deadline.Format("2006/01/02"))
			fmt.Fprintf(&buf, "[%d Due] displays on same row as [story_%d]\n", story.Id, story.Id)
			if story.Deadline.Before(time.Now()) {
				fmt.Fprintf(&buf, "[%d Due] is colored in Black\n", story.Id)
			}
		}
		for _, link := range story.StoryLinks {
			if link.Type == "subject" {
				if link.Verb == "blocks" {
					constraints = append(constraints, fmt.Sprintf("[story_%d] starts at [story_%d]'s end", link.ObjectId, link.SubjectId))
				}
			}
		}
	}
	for _, constraint := range constraints {
		fmt.Fprintln(&buf, constraint)
	}
	fmt.Fprintln(&buf, buildLegend())
	fmt.Fprintln(&buf, "@endgantt")
	return buf.Bytes(), nil
}

func getLink(story openapi.StorySlim) string {
	return fmt.Sprintf("[story_%d] links to [[%s]]", story.Id, story.AppUrl)
}
func setColor(story openapi.StorySlim, epic openapi.Epic) string {
	color := "LightGreen"
	if story.Deadline != nil && story.GetEndDate(epic.GetStartDate()).After(*story.Deadline) {
		color = "Red"
	}
	return fmt.Sprintf("[story_%d] is colored in %s", story.Id, color)
}

// buildLegend returns a string that will create a PlantUML legend based on the colors being used
func buildLegend() string {
	legend := `legend right
	=== Legend
	|<#AAF>            | {today} |
	|<#LightBlue>      | Future |
	|<#LightGreen>     | Current |
	|<#Red>            | Past Due |
	|<#Blue>           | Epic Start / End |
	end legend`
	return legend
}
