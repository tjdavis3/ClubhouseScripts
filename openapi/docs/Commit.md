# Commit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorEmail** | **string** | The email address of the VCS user that authored the Commit. | 
**AuthorId** | Pointer to **string** | The ID of the Member that authored the Commit, if known. | 
**AuthorIdentity** | [**Identity**](Identity.md) |  | 
**CreatedAt** | [**time.Time**](time.Time.md) | The time/date the Commit was created. | 
**EntityType** | **string** | A string description of this resource. | 
**Hash** | **string** | The Commit hash. | 
**Id** | Pointer to **int64** | The unique ID of the Commit. | 
**MergedBranchIds** | **[]int64** | The IDs of the Branches the Commit has been merged into. | 
**Message** | **string** | The Commit message. | 
**RepositoryId** | Pointer to **int64** | The ID of the Repository that contains the Commit. | 
**Timestamp** | [**time.Time**](time.Time.md) | The time/date the Commit was pushed. | 
**UpdatedAt** | Pointer to [**time.Time**](time.Time.md) | The time/date the Commit was updated. | 
**Url** | **string** | The URL of the Commit. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


