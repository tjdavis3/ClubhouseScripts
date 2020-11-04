# CreateComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorId** | **string** | The Member ID of the Comment&#39;s author. Defaults to the user identified by the API token. | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) | Defaults to the time/date the comment is created, but can be set to reflect another date. | [optional] 
**ExternalId** | **string** | This field can be set to another unique ID. In the case that the comment has been imported from another tool, the ID in the other tool can be indicated here. | [optional] 
**Text** | **string** | The comment text. | 
**UpdatedAt** | [**time.Time**](time.Time.md) | Defaults to the time/date the comment is last updated, but can be set to reflect another date. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


