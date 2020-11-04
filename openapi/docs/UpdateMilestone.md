# UpdateMilestone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfterId** | **int64** | The ID of the Milestone we want to move this Milestone after. | [optional] 
**BeforeId** | **int64** | The ID of the Milestone we want to move this Milestone before. | [optional] 
**Categories** | [**[]CreateCategoryParams**](CreateCategoryParams.md) | An array of IDs of Categories attached to the Milestone. | [optional] 
**CompletedAtOverride** | Pointer to [**time.Time**](time.Time.md) | A manual override for the time/date the Milestone was completed. | [optional] 
**Description** | **string** | The Milestone&#39;s description. | [optional] 
**Name** | **string** | The name of the Milestone. | [optional] 
**StartedAtOverride** | Pointer to [**time.Time**](time.Time.md) | A manual override for the time/date the Milestone was started. | [optional] 
**State** | **string** | The workflow state that the Milestone is in. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


