# UpdateStories

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfterId** | **int64** | The ID of the story that the stories are to be moved below. | [optional] 
**Archived** | **bool** | If the Stories should be archived or not. | [optional] 
**BeforeId** | **int64** | The ID of the story that the stories are to be moved before. | [optional] 
**Deadline** | Pointer to [**time.Time**](time.Time.md) | The due date of the story. | [optional] 
**EpicId** | Pointer to **int64** | The ID of the epic the story belongs to. | [optional] 
**Estimate** | Pointer to **int64** | The numeric point estimate of the story. Can also be null, which means unestimated. | [optional] 
**FollowerIdsAdd** | **[]string** | The UUIDs of the new followers to be added. | [optional] 
**FollowerIdsRemove** | **[]string** | The UUIDs of the followers to be removed. | [optional] 
**GroupId** | Pointer to **string** | The Id of the Group the Stories should belong to. | [optional] 
**IterationId** | Pointer to **int64** | The ID of the iteration the story belongs to. | [optional] 
**LabelsAdd** | [**[]CreateLabelParams**](CreateLabelParams.md) | An array of labels to be added. | [optional] 
**LabelsRemove** | [**[]CreateLabelParams**](CreateLabelParams.md) | An array of labels to be removed. | [optional] 
**MoveTo** | **string** |  | [optional] 
**OwnerIdsAdd** | **[]string** | The UUIDs of the new owners to be added. | [optional] 
**OwnerIdsRemove** | **[]string** | The UUIDs of the owners to be removed. | [optional] 
**ProjectId** | **int64** | The ID of the Project the Stories should belong to. | [optional] 
**RequestedById** | **string** | The ID of the member that requested the story. | [optional] 
**StoryIds** | **[]int64** | The unique IDs of the Stories you wish to update. | 
**StoryType** | **string** | The type of story (feature, bug, chore). | [optional] 
**WorkflowStateId** | **int64** | The ID of the workflow state to put the stories in. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


