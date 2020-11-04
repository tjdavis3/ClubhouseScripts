# Member

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | The time/date the Member was created. | 
**CreatedWithoutInvite** | **bool** | Whether this member was created as a placeholder entity. | 
**Disabled** | **bool** | True/false boolean indicating whether the Member has been disabled within this Organization. | 
**EntityType** | **string** | A string description of this resource. | 
**GroupIds** | **[]string** | The Member&#39;s group ids | 
**Id** | **string** | The Member&#39;s ID in Clubhouse. | 
**Profile** | [**Profile**](Profile.md) |  | 
**ReplacedBy** | **string** | The id of the member that replaces this one when merged. | [optional] 
**Role** | **string** | The Member&#39;s role in the Clubhouse organization. | 
**UpdatedAt** | Pointer to [**time.Time**](time.Time.md) | The time/date the Member was last updated. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


