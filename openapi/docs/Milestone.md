# Milestone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppUrl** | **string** | The Clubhouse application url for the Milestone. | 
**Categories** | [**[]Category**](Category.md) | An array of Categories attached to the Milestone. | 
**Completed** | **bool** | A true/false boolean indicating if the Milestone has been completed. | 
**CompletedAt** | Pointer to [**time.Time**](time.Time.md) | The time/date the Milestone was completed. | 
**CompletedAtOverride** | Pointer to [**time.Time**](time.Time.md) | A manual override for the time/date the Milestone was completed. | 
**CreatedAt** | [**time.Time**](time.Time.md) | The time/date the Milestone was created. | 
**Description** | **string** | The Milestone&#39;s description. | 
**EntityType** | **string** | A string description of this resource. | 
**Id** | **int64** | The unique ID of the Milestone. | 
**Name** | **string** | The name of the Milestone. | 
**Position** | **int64** | A number representing the position of the Milestone in relation to every other Milestone within the Organization. | 
**Started** | **bool** | A true/false boolean indicating if the Milestone has been started. | 
**StartedAt** | Pointer to [**time.Time**](time.Time.md) | The time/date the Milestone was started. | 
**StartedAtOverride** | Pointer to [**time.Time**](time.Time.md) | A manual override for the time/date the Milestone was started. | 
**State** | **string** | The workflow state that the Milestone is in. | 
**Stats** | [**MilestoneStats**](MilestoneStats.md) |  | [optional] 
**UpdatedAt** | [**time.Time**](time.Time.md) | The time/date the Milestone was updated. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


