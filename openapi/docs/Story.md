# Story

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppUrl** | **string** | The Clubhouse application url for the Story. | 
**Archived** | **bool** | True if the story has been archived or not. | 
**Blocked** | **bool** | A true/false boolean indicating if the Story is currently blocked. | 
**Blocker** | **bool** | A true/false boolean indicating if the Story is currently a blocker of another story. | 
**Branches** | [**[]Branch**](Branch.md) | An array of Git branches attached to the story. | 
**Comments** | [**[]Comment**](Comment.md) | An array of comments attached to the story. | 
**Commits** | [**[]Commit**](Commit.md) | An array of commits attached to the story. | 
**Completed** | **bool** | A true/false boolean indicating if the Story has been completed. | 
**CompletedAt** | Pointer to [**time.Time**](time.Time.md) | The time/date the Story was completed. | 
**CompletedAtOverride** | Pointer to [**time.Time**](time.Time.md) | A manual override for the time/date the Story was completed. | 
**CreatedAt** | [**time.Time**](time.Time.md) | The time/date the Story was created. | 
**CycleTime** | **int64** | The cycle time (in seconds) of this story when complete. | [optional] 
**Deadline** | Pointer to [**time.Time**](time.Time.md) | The due date of the story. | 
**Description** | **string** | The description of the story. | 
**EntityType** | **string** | A string description of this resource. | 
**EpicId** | Pointer to **int64** | The ID of the epic the story belongs to. | 
**Estimate** | Pointer to **int64** | The numeric point estimate of the story. Can also be null, which means unestimated. | 
**ExternalId** | Pointer to **string** | This field can be set to another unique ID. In the case that the Story has been imported from another tool, the ID in the other tool can be indicated here. | 
**ExternalLinks** | **[]string** | An array of external link (strings) associated with a Story | 
**ExternalTickets** | [**[]ExternalTicket**](ExternalTicket.md) | An array of External Tickets associated with a Story | 
**Files** | [**[]File**](File.md) | An array of files attached to the story. | 
**FollowerIds** | **[]string** | An array of UUIDs for any Members listed as Followers. | 
**GroupId** | Pointer to **string** | The ID of the group associated with the story. | 
**GroupMentionIds** | **[]string** | An array of Group IDs that have been mentioned in the Story description. | 
**Id** | **int64** | The unique ID of the Story. | 
**IterationId** | Pointer to **int64** | The ID of the iteration the story belongs to. | 
**Labels** | [**[]Label**](Label.md) | An array of labels attached to the story. | 
**LeadTime** | **int64** | The lead time (in seconds) of this story when complete. | [optional] 
**LinkedFiles** | [**[]LinkedFile**](LinkedFile.md) | An array of linked files attached to the story. | 
**MemberMentionIds** | **[]string** | An array of Member IDs that have been mentioned in the Story description. | 
**MentionIds** | **[]string** | Deprecated: use member_mention_ids. | 
**MovedAt** | Pointer to [**time.Time**](time.Time.md) | The time/date the Story was last changed workflow-state. | 
**Name** | **string** | The name of the story. | 
**OwnerIds** | **[]string** | An array of UUIDs of the owners of this story. | 
**Position** | **int64** | A number representing the position of the story in relation to every other story in the current project. | 
**PreviousIterationIds** | **[]int64** | The IDs of the iteration the story belongs to. | 
**ProjectId** | **int64** | The ID of the project the story belongs to. | 
**PullRequests** | [**[]PullRequest**](PullRequest.md) | An array of Pull/Merge Requests attached to the story. | 
**RequestedById** | **string** | The ID of the Member that requested the story. | 
**Started** | **bool** | A true/false boolean indicating if the Story has been started. | 
**StartedAt** | Pointer to [**time.Time**](time.Time.md) | The time/date the Story was started. | 
**StartedAtOverride** | Pointer to [**time.Time**](time.Time.md) | A manual override for the time/date the Story was started. | 
**Stats** | [**StoryStats**](StoryStats.md) |  | 
**StoryLinks** | [**[]TypedStoryLink**](TypedStoryLink.md) | An array of story links attached to the Story. | 
**StoryType** | **string** | The type of story (feature, bug, chore). | 
**SupportTickets** | [**[]SupportTicket**](SupportTicket.md) |  | 
**Tasks** | [**[]Task**](Task.md) | An array of tasks connected to the story. | 
**UpdatedAt** | Pointer to [**time.Time**](time.Time.md) | The time/date the Story was updated. | 
**WorkflowStateId** | **int64** | The ID of the workflow state the story is currently in. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


