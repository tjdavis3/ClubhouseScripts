# CreateTaskParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Complete** | **bool** | True/false boolean indicating whether the Task is completed. Defaults to false. | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) | Defaults to the time/date the Task is created but can be set to reflect another creation time/date. | [optional] 
**Description** | **string** | The Task description. | 
**ExternalId** | **string** | This field can be set to another unique ID. In the case that the Task has been imported from another tool, the ID in the other tool can be indicated here. | [optional] 
**OwnerIds** | **[]string** | An array of UUIDs for any members you want to add as Owners on this new Task. | [optional] 
**UpdatedAt** | [**time.Time**](time.Time.md) | Defaults to the time/date the Task is created in Clubhouse but can be set to reflect another time/date. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


