# WorkflowState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | **string** | The hex color for this Workflow State. | 
**CreatedAt** | [**time.Time**](time.Time.md) | The time/date the Workflow State was created. | 
**Description** | **string** | The description of what sort of Stories belong in that Workflow state. | 
**EntityType** | **string** | A string description of this resource. | 
**Id** | **int64** | The unique ID of the Workflow State. | 
**Name** | **string** | The Workflow State&#39;s name. | 
**NumStories** | **int64** | The number of Stories currently in that Workflow State. | 
**NumStoryTemplates** | **int64** | The number of Story Templates associated with that Workflow State. | 
**Position** | **int64** | The position that the Workflow State is in, starting with 0 at the left. | 
**Type** | **string** | The type of Workflow State (Unstarted, Started, or Finished) | 
**UpdatedAt** | [**time.Time**](time.Time.md) | When the Workflow State was last updated. | 
**Verb** | Pointer to **string** | The verb that triggers a move to that Workflow State when making VCS commits. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


