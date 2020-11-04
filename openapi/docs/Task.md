# Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Complete** | **bool** | True/false boolean indicating whether the Task has been completed. | 
**CompletedAt** | Pointer to [**time.Time**](time.Time.md) | The time/date the Task was completed. | 
**CreatedAt** | [**time.Time**](time.Time.md) | The time/date the Task was created. | 
**Description** | **string** | Full text of the Task. | 
**EntityType** | **string** | A string description of this resource. | 
**ExternalId** | Pointer to **string** | This field can be set to another unique ID. In the case that the Task has been imported from another tool, the ID in the other tool can be indicated here. | 
**GroupMentionIds** | **[]string** | An array of UUIDs of Groups mentioned in this Task. | 
**Id** | **int64** | The unique ID of the Task. | 
**MemberMentionIds** | **[]string** | An array of UUIDs of Members mentioned in this Task. | 
**MentionIds** | **[]string** | Deprecated: use member_mention_ids. | 
**OwnerIds** | **[]string** | An array of UUIDs of the Owners of this Task. | 
**Position** | **int64** | The number corresponding to the Task&#39;s position within a list of Tasks on a Story. | 
**StoryId** | **int64** | The unique identifier of the parent Story. | 
**UpdatedAt** | Pointer to [**time.Time**](time.Time.md) | The time/date the Task was updated. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


