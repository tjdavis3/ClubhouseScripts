# Branch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | The time/date the Branch was created. | 
**Deleted** | **bool** | A true/false boolean indicating if the Branch has been deleted. | 
**EntityType** | **string** | A string description of this resource. | 
**Id** | Pointer to **int64** | The unique ID of the Branch. | 
**MergedBranchIds** | **[]int64** | The IDs of the Branches the Branch has been merged into. | 
**Name** | **string** | The name of the Branch. | 
**Persistent** | **bool** | A true/false boolean indicating if the Branch is persistent; e.g. master. | 
**PullRequests** | [**[]PullRequest**](PullRequest.md) | An array of PullRequests attached to the Branch (there is usually only one). | 
**RepositoryId** | Pointer to **int64** | The ID of the Repository that contains the Branch. | 
**UpdatedAt** | Pointer to [**time.Time**](time.Time.md) | The time/date the Branch was updated. | 
**Url** | **string** | The URL of the Branch. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


