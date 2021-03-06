/*
 * Clubhouse API
 *
 * Clubhouse API
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)
// CreateStoryParams Used to create multiple stories in a single request.
type CreateStoryParams struct {
	// Controls the story's archived state.
	Archived bool `json:"archived,omitempty"`
	// An array of comments to add to the story.
	Comments []CreateStoryCommentParams `json:"comments,omitempty"`
	// A manual override for the time/date the Story was completed.
	CompletedAtOverride time.Time `json:"completed_at_override,omitempty"`
	// The time/date the Story was created.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// The due date of the story.
	Deadline *time.Time `json:"deadline,omitempty"`
	// The description of the story.
	Description string `json:"description,omitempty"`
	// The ID of the epic the story belongs to.
	EpicId *int64 `json:"epic_id,omitempty"`
	// The numeric point estimate of the story. Can also be null, which means unestimated.
	Estimate *int64 `json:"estimate,omitempty"`
	// This field can be set to another unique ID. In the case that the Story has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalId string `json:"external_id,omitempty"`
	// An array of External Tickets associated with this story. These External Tickets must have unquie external id. Duplicated External Tickets will be removed.
	ExternalTickets []CreateExternalTicketParams `json:"external_tickets,omitempty"`
	// An array of IDs of files attached to the story.
	FileIds []int64 `json:"file_ids,omitempty"`
	// An array of UUIDs of the followers of this story.
	FollowerIds []string `json:"follower_ids,omitempty"`
	// The id of the group to associate with this story.
	GroupId *string `json:"group_id,omitempty"`
	// The ID of the iteration the story belongs to.
	IterationId *int64 `json:"iteration_id,omitempty"`
	// An array of labels attached to the story.
	Labels []CreateLabelParams `json:"labels,omitempty"`
	// An array of IDs of linked files attached to the story.
	LinkedFileIds []int64 `json:"linked_file_ids,omitempty"`
	// The name of the story.
	Name string `json:"name"`
	// An array of UUIDs of the owners of this story.
	OwnerIds []string `json:"owner_ids,omitempty"`
	// The ID of the project the story belongs to.
	ProjectId int64 `json:"project_id"`
	// The ID of the member that requested the story.
	RequestedById string `json:"requested_by_id,omitempty"`
	// A manual override for the time/date the Story was started.
	StartedAtOverride time.Time `json:"started_at_override,omitempty"`
	// An array of story links attached to the story.
	StoryLinks []CreateStoryLinkParams `json:"story_links,omitempty"`
	// The type of story (feature, bug, chore).
	StoryType string `json:"story_type,omitempty"`
	SupportTickets []CreateExternalTicketParams `json:"support_tickets,omitempty"`
	// An array of tasks connected to the story.
	Tasks []CreateTaskParams `json:"tasks,omitempty"`
	// The time/date the Story was updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// The ID of the workflow state the story will be in.
	WorkflowStateId int64 `json:"workflow_state_id,omitempty"`
}
