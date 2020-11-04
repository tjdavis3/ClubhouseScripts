# LinkedFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | Pointer to **string** | The content type of the image (e.g. txt/plain). | 
**CreatedAt** | [**time.Time**](time.Time.md) | The time/date the LinkedFile was created. | 
**Description** | Pointer to **string** | The description of the file. | 
**EntityType** | **string** | A string description of this resource. | 
**GroupMentionIds** | **[]string** | The groups that are mentioned in the description of the file. | 
**Id** | **int64** | The unique identifier for the file. | 
**MemberMentionIds** | **[]string** | The members that are mentioned in the description of the file. | 
**MentionIds** | **[]string** | Deprecated: use member_mention_ids. | 
**Name** | **string** | The name of the linked file. | 
**Size** | Pointer to **int64** | The filesize, if the integration provided it. | 
**StoryIds** | **[]int64** | The IDs of the stories this file is attached to. | 
**ThumbnailUrl** | Pointer to **string** | The URL of the file thumbnail, if the integration provided it. | 
**Type** | **string** | The integration type (e.g. google, dropbox, box). | 
**UpdatedAt** | [**time.Time**](time.Time.md) | The time/date the LinkedFile was updated. | 
**UploaderId** | **string** | The UUID of the member that uploaded the file. | 
**Url** | **string** | The URL of the file. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


