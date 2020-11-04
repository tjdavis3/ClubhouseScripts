# Comment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppUrl** | **string** | The Clubhouse application url for the Comment. | 
**AuthorId** | Pointer to **string** | The unique ID of the Member who is the Comment&#39;s author. | 
**CreatedAt** | [**time.Time**](time.Time.md) | The time/date when the Comment was created. | 
**EntityType** | **string** | A string description of this resource. | 
**ExternalId** | Pointer to **string** | This field can be set to another unique ID. In the case that the Comment has been imported from another tool, the ID in the other tool can be indicated here. | 
**GroupMentionIds** | **[]string** | The unique IDs of the Group who are mentioned in the Comment. | 
**Id** | **int64** | The unique ID of the Comment. | 
**MemberMentionIds** | **[]string** | The unique IDs of the Member who are mentioned in the Comment. | 
**MentionIds** | **[]string** | Deprecated: use member_mention_ids. | 
**Position** | **int64** | The Comments numerical position in the list from oldest to newest. | 
**Reactions** | [**[]Reaction**](Reaction.md) | A set of Reactions to this Comment. | 
**StoryId** | **int64** | The ID of the Story on which the Comment appears. | 
**Text** | **string** | The text of the Comment. | 
**UpdatedAt** | Pointer to [**time.Time**](time.Time.md) | The time/date when the Comment was updated. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


