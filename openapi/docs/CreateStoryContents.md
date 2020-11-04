# CreateStoryContents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deadline** | Pointer to [**time.Time**](time.Time.md) | The due date of the story. | [optional] 
**Description** | **string** | The description of the story. | [optional] 
**EntityType** | **string** | A string description of this resource. | [optional] 
**EpicId** | Pointer to **int64** | The ID of the epic the to be populated. | [optional] 
**Estimate** | Pointer to **int64** | The numeric point estimate to be populated. | [optional] 
**ExternalTickets** | [**[]CreateEntityTemplateExternalTicket**](CreateEntityTemplateExternalTicket.md) | An array of the external ticket IDs to be populated. | [optional] 
**FileIds** | **[]int64** | An array of the attached file IDs to be populated. | [optional] 
**Files** | [**[]File**](File.md) | An array of files attached to the story. | [optional] 
**FollowerIds** | **[]string** | An array of UUIDs for any Members listed as Followers. | [optional] 
**IterationId** | Pointer to **int64** | The ID of the iteration the to be populated. | [optional] 
**Labels** | [**[]CreateLabelParams**](CreateLabelParams.md) | An array of labels to be populated by the template. | [optional] 
**LinkedFileIds** | **[]int64** | An array of the linked file IDs to be populated. | [optional] 
**LinkedFiles** | [**[]LinkedFile**](LinkedFile.md) | An array of linked files attached to the story. | [optional] 
**Name** | **string** | The name of the story. | [optional] 
**OwnerIds** | **[]string** | An array of UUIDs of the owners of this story. | [optional] 
**ProjectId** | **int64** | The ID of the project the story belongs to. | [optional] 
**StoryType** | **string** | The type of story (feature, bug, chore). | [optional] 
**Tasks** | [**[]EntityTemplateTask**](EntityTemplateTask.md) | An array of tasks to be populated by the template. | [optional] 
**WorkflowStateId** | **int64** | The ID of the workflow state the story is currently in. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


