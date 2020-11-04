# PullRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BranchId** | **int64** | The ID of the branch for the particular pull request. | 
**BranchName** | **string** | The name of the branch for the particular pull request. | 
**BuildStatus** | **string** | The status of the Continuous Integration workflow for the pull request. | [optional] 
**Closed** | **bool** | True/False boolean indicating whether the VCS pull request has been closed. | 
**CreatedAt** | [**time.Time**](time.Time.md) | The time/date the pull request was created. | 
**Draft** | **bool** | True/False boolean indicating whether the VCS pull request is in the draft state. | 
**EntityType** | **string** | A string description of this resource. | 
**Id** | **int64** | The unique ID associated with the pull request in Clubhouse. | 
**NumAdded** | **int64** | Number of lines added in the pull request, according to VCS. | 
**NumCommits** | Pointer to **int64** | The number of commits on the pull request. | 
**NumModified** | Pointer to **int64** | Number of lines modified in the pull request, according to VCS. | 
**NumRemoved** | **int64** | Number of lines removed in the pull request, according to VCS. | 
**Number** | **int64** | The pull requests unique number ID in VCS. | 
**ReviewStatus** | **string** | The status of the review for the pull request. | [optional] 
**TargetBranchId** | **int64** | The ID of the target branch for the particular pull request. | 
**TargetBranchName** | **string** | The name of the target branch for the particular pull request. | 
**Title** | **string** | The title of the pull request. | 
**UpdatedAt** | [**time.Time**](time.Time.md) | The time/date the pull request was created. | 
**Url** | **string** | The URL for the pull request. | 
**VcsLabels** | Pointer to [**[]PullRequestLabel**](PullRequestLabel.md) | An array of PullRequestLabels attached to the PullRequest. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


