# CreateStoryParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | **bool** | Controls the story&#39;s archived state. | [optional] 
**Comments** | [**[]CreateStoryCommentParams**](CreateStoryCommentParams.md) | An array of comments to add to the story. | [optional] 
**CompletedAtOverride** | [**time.Time**](time.Time.md) | A manual override for the time/date the Story was completed. | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) | The time/date the Story was created. | [optional] 
**Deadline** | Pointer to [**time.Time**](time.Time.md) | The due date of the story. | [optional] 
**Description** | **string** | The description of the story. | [optional] 
**EpicId** | Pointer to **int64** | The ID of the epic the story belongs to. | [optional] 
**Estimate** | Pointer to **int64** | The numeric point estimate of the story. Can also be null, which means unestimated. | [optional] 
**ExternalId** | **string** | This field can be set to another unique ID. In the case that the Story has been imported from another tool, the ID in the other tool can be indicated here. | [optional] 
**ExternalTickets** | [**[]CreateExternalTicketParams**](CreateExternalTicketParams.md) | An array of External Tickets associated with this story. These External Tickets must have unquie external id. Duplicated External Tickets will be removed. | [optional] 
**FileIds** | **[]int64** | An array of IDs of files attached to the story. | [optional] 
**FollowerIds** | **[]string** | An array of UUIDs of the followers of this story. | [optional] 
**GroupId** | Pointer to **string** | The id of the group to associate with this story. | [optional] 
**IterationId** | Pointer to **int64** | The ID of the iteration the story belongs to. | [optional] 
**Labels** | [**[]CreateLabelParams**](CreateLabelParams.md) | An array of labels attached to the story. | [optional] 
**LinkedFileIds** | **[]int64** | An array of IDs of linked files attached to the story. | [optional] 
**Name** | **string** | The name of the story. | 
**OwnerIds** | **[]string** | An array of UUIDs of the owners of this story. | [optional] 
**ProjectId** | **int64** | The ID of the project the story belongs to. | 
**RequestedById** | **string** | The ID of the member that requested the story. | [optional] 
**StartedAtOverride** | [**time.Time**](time.Time.md) | A manual override for the time/date the Story was started. | [optional] 
**StoryLinks** | [**[]CreateStoryLinkParams**](CreateStoryLinkParams.md) | An array of story links attached to the story. | [optional] 
**StoryType** | **string** | The type of story (feature, bug, chore). | [optional] 
**SupportTickets** | [**[]CreateExternalTicketParams**](CreateExternalTicketParams.md) |  | [optional] 
**Tasks** | [**[]CreateTaskParams**](CreateTaskParams.md) | An array of tasks connected to the story. | [optional] 
**UpdatedAt** | [**time.Time**](time.Time.md) | The time/date the Story was updated. | [optional] 
**WorkflowStateId** | **int64** | The ID of the workflow state the story will be in. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


