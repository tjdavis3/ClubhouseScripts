# StoryContents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deadline** | [**time.Time**](time.Time.md) | The due date of the story. | [optional] 
**Description** | **string** | The description of the story. | [optional] 
**EntityType** | **string** | A string description of this resource. | [optional] 
**EpicId** | **int64** | The ID of the epic the story belongs to. | [optional] 
**Estimate** | **int64** | The numeric point estimate of the story. Can also be null, which means unestimated. | [optional] 
**ExternalTickets** | [**[]ExternalTicket**](ExternalTicket.md) | An array of external tickets connected to the story. | [optional] 
**Files** | [**[]File**](File.md) | An array of files attached to the story. | [optional] 
**FollowerIds** | **[]string** | An array of UUIDs for any Members listed as Followers. | [optional] 
**IterationId** | **int64** | The ID of the iteration the story belongs to. | [optional] 
**Labels** | [**[]Label**](Label.md) | An array of labels attached to the story. | [optional] 
**LinkedFiles** | [**[]LinkedFile**](LinkedFile.md) | An array of linked files attached to the story. | [optional] 
**Name** | **string** | The name of the story. | [optional] 
**OwnerIds** | **[]string** | An array of UUIDs of the owners of this story. | [optional] 
**ProjectId** | **int64** | The ID of the project the story belongs to. | [optional] 
**StoryType** | **string** | The type of story (feature, bug, chore). | [optional] 
**Tasks** | [**[]StoryContentsTask**](StoryContentsTask.md) | An array of tasks connected to the story. | [optional] 
**WorkflowStateId** | **int64** | The ID of the workflow state the story is currently in. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


