# File

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | **string** | Free form string corresponding to a text or image file. | 
**CreatedAt** | [**time.Time**](time.Time.md) | The time/date that the file was created. | 
**Description** | Pointer to **string** | The description of the file. | 
**EntityType** | **string** | A string description of this resource. | 
**ExternalId** | Pointer to **string** | This field can be set to another unique ID. In the case that the File has been imported from another tool, the ID in the other tool can be indicated here. | 
**Filename** | **string** | The name assigned to the file in Clubhouse upon upload. | 
**GroupMentionIds** | **[]string** | The unique IDs of the Groups who are mentioned in the file description. | 
**Id** | **int64** | The unique ID for the file. | 
**MemberMentionIds** | **[]string** | The unique IDs of the Members who are mentioned in the file description. | 
**MentionIds** | **[]string** | Deprecated: use member_mention_ids. | 
**Name** | **string** | The optional User-specified name of the file. | 
**Size** | **int64** | The size of the file. | 
**StoryIds** | **[]int64** | The unique IDs of the Stories associated with this file. | 
**ThumbnailUrl** | Pointer to **string** | The url where the thumbnail of the file can be found in Clubhouse. | 
**UpdatedAt** | Pointer to [**time.Time**](time.Time.md) | The time/date that the file was updated. | 
**UploaderId** | **string** | The unique ID of the Member who uploaded the file. | 
**Url** | Pointer to **string** | The URL for the file. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


