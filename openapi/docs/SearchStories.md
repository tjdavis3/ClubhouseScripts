# SearchStories

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | **bool** | A true/false boolean indicating whether the Story is in archived state. | [optional] 
**CompletedAtEnd** | [**time.Time**](time.Time.md) | Stories should have been completed before this date. | [optional] 
**CompletedAtStart** | [**time.Time**](time.Time.md) | Stories should have been competed after this date. | [optional] 
**CreatedAtEnd** | [**time.Time**](time.Time.md) | Stories should have been created before this date. | [optional] 
**CreatedAtStart** | [**time.Time**](time.Time.md) | Stories should have been created after this date. | [optional] 
**DeadlineEnd** | [**time.Time**](time.Time.md) | Stories should have a deadline before this date. | [optional] 
**DeadlineStart** | [**time.Time**](time.Time.md) | Stories should have a deadline after this date. | [optional] 
**EpicId** | Pointer to **int64** | The Epic IDs that may be associated with the Stories. | [optional] 
**EpicIds** | **[]int64** | The Epic IDs that may be associated with the Stories. | [optional] 
**Estimate** | **int64** | The number of estimate points associate with the Stories. | [optional] 
**ExternalId** | **string** | An ID or URL that references an external resource. Useful during imports. | [optional] 
**GroupId** | Pointer to **string** | The Group ID that is associated with the Stories | [optional] 
**GroupIds** | **[]string** | The Group IDs that are associated with the Stories | [optional] 
**IncludesDescription** | **bool** | Whether to include the story description in the response. | [optional] 
**IterationId** | Pointer to **int64** | The Iteration ID that may be associated with the Stories. | [optional] 
**IterationIds** | **[]int64** | The Iteration IDs that may be associated with the Stories. | [optional] 
**LabelIds** | **[]int64** | The Label IDs that may be associated with the Stories. | [optional] 
**LabelName** | **string** | The name of any associated Labels. | [optional] 
**OwnerId** | Pointer to **string** | An array of UUIDs for any Users who may be Owners of the Stories. | [optional] 
**OwnerIds** | **[]string** | An array of UUIDs for any Users who may be Owners of the Stories. | [optional] 
**ProjectId** | **int64** | The IDs for the Projects the Stories may be assigned to. | [optional] 
**ProjectIds** | **[]int64** | The IDs for the Projects the Stories may be assigned to. | [optional] 
**RequestedById** | **string** | The UUID of any Users who may have requested the Stories. | [optional] 
**StoryType** | **string** | The type of Stories that you want returned. | [optional] 
**UpdatedAtEnd** | [**time.Time**](time.Time.md) | Stories should have been updated before this date. | [optional] 
**UpdatedAtStart** | [**time.Time**](time.Time.md) | Stories should have been updated after this date. | [optional] 
**WorkflowStateId** | **int64** | The unique IDs of the specific Workflow States that the Stories should be in. | [optional] 
**WorkflowStateTypes** | **[]string** | The type of Workflow State the Stories may be in. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


