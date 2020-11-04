# UpdateStory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfterId** | **int64** | The ID of the story we want to move this story after. | [optional] 
**Archived** | **bool** | True if the story is archived, otherwise false. | [optional] 
**BeforeId** | **int64** | The ID of the story we want to move this story before. | [optional] 
**BranchIds** | **[]int64** | An array of IDs of Branches attached to the story. | [optional] 
**CommitIds** | **[]int64** | An array of IDs of Commits attached to the story. | [optional] 
**CompletedAtOverride** | Pointer to [**time.Time**](time.Time.md) | A manual override for the time/date the Story was completed. | [optional] 
**Deadline** | Pointer to [**time.Time**](time.Time.md) | The due date of the story. | [optional] 
**Description** | **string** | The description of the story. | [optional] 
**EpicId** | Pointer to **int64** | The ID of the epic the story belongs to. | [optional] 
**Estimate** | Pointer to **int64** | The numeric point estimate of the story. Can also be null, which means unestimated. | [optional] 
**FileIds** | **[]int64** | An array of IDs of files attached to the story. | [optional] 
**FollowerIds** | **[]string** | An array of UUIDs of the followers of this story. | [optional] 
**GroupId** | Pointer to **string** | The ID of the group to associate with this story | [optional] 
**IterationId** | Pointer to **int64** | The ID of the iteration the story belongs to. | [optional] 
**Labels** | [**[]CreateLabelParams**](CreateLabelParams.md) | An array of labels attached to the story. | [optional] 
**LinkedFileIds** | **[]int64** | An array of IDs of linked files attached to the story. | [optional] 
**MoveTo** | **string** |  | [optional] 
**Name** | **string** | The title of the story. | [optional] 
**OwnerIds** | **[]string** | An array of UUIDs of the owners of this story. | [optional] 
**ProjectId** | **int64** | The ID of the project the story belongs to. | [optional] 
**PullRequestIds** | **[]int64** | An array of IDs of Pull/Merge Requests attached to the story. | [optional] 
**RequestedById** | **string** | The ID of the member that requested the story. | [optional] 
**StartedAtOverride** | Pointer to [**time.Time**](time.Time.md) | A manual override for the time/date the Story was started. | [optional] 
**StoryType** | **string** | The type of story (feature, bug, chore). | [optional] 
**WorkflowStateId** | **int64** | The ID of the workflow state to put the story in. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


