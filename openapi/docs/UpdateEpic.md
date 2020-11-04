# UpdateEpic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfterId** | **int64** | The ID of the Epic we want to move this Epic after. | [optional] 
**Archived** | **bool** | A true/false boolean indicating whether the Epic is in archived state. | [optional] 
**BeforeId** | **int64** | The ID of the Epic we want to move this Epic before. | [optional] 
**CompletedAtOverride** | Pointer to [**time.Time**](time.Time.md) | A manual override for the time/date the Epic was completed. | [optional] 
**Deadline** | Pointer to [**time.Time**](time.Time.md) | The Epic&#39;s deadline. | [optional] 
**Description** | **string** | The Epic&#39;s description. | [optional] 
**EpicStateId** | **int64** | The ID of the Epic State. | [optional] 
**FollowerIds** | **[]string** | An array of UUIDs for any Members you want to add as Followers on this Epic. | [optional] 
**GroupId** | Pointer to **string** | The ID of the group to associate with the epic. | [optional] 
**Labels** | [**[]CreateLabelParams**](CreateLabelParams.md) | An array of Labels attached to the Epic. | [optional] 
**MilestoneId** | Pointer to **int64** | The ID of the Milestone this Epic is related to. | [optional] 
**Name** | **string** | The Epic&#39;s name. | [optional] 
**OwnerIds** | **[]string** | An array of UUIDs for any members you want to add as Owners on this Epic. | [optional] 
**PlannedStartDate** | Pointer to [**time.Time**](time.Time.md) | The Epic&#39;s planned start date. | [optional] 
**RequestedById** | **string** | The ID of the member that requested the epic. | [optional] 
**StartedAtOverride** | Pointer to [**time.Time**](time.Time.md) | A manual override for the time/date the Epic was started. | [optional] 
**State** | **string** | &#x60;Deprecated&#x60; The Epic&#39;s state (to do, in progress, or done); will be ignored when &#x60;epic_state_id&#x60; is set. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


