# ThreadedComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppUrl** | **string** | The Clubhouse application url for the Comment. | 
**AuthorId** | **string** | The unique ID of the Member that authored the Comment. | 
**Comments** | [**[]ThreadedComment**](ThreadedComment.md) | A nested array of threaded comments. | 
**CreatedAt** | [**time.Time**](time.Time.md) | The time/date the Comment was created. | 
**Deleted** | **bool** | True/false boolean indicating whether the Comment is deleted. | 
**EntityType** | **string** | A string description of this resource. | 
**ExternalId** | Pointer to **string** | This field can be set to another unique ID. In the case that the Comment has been imported from another tool, the ID in the other tool can be indicated here. | 
**GroupMentionIds** | **[]string** | An array of Group IDs that have been mentioned in this Comment. | 
**Id** | **int64** | The unique ID of the Comment. | 
**MemberMentionIds** | **[]string** | An array of Member IDs that have been mentioned in this Comment. | 
**MentionIds** | **[]string** | Deprecated: use member_mention_ids. | 
**Text** | **string** | The text of the Comment. | 
**UpdatedAt** | [**time.Time**](time.Time.md) | The time/date the Comment was updated. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


