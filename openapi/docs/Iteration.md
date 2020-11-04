# Iteration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppUrl** | **string** | The Clubhouse application url for the Iteration. | 
**CreatedAt** | [**time.Time**](time.Time.md) | The instant when this iteration was created. | 
**Description** | **string** | The description of the iteration. | 
**EndDate** | [**time.Time**](time.Time.md) | The date this iteration begins. | 
**EntityType** | **string** | A string description of this resource | 
**FollowerIds** | **[]string** | An array of UUIDs for any Members listed as Followers. | 
**GroupIds** | **[]string** | An array of UUIDs for any Groups you want to add as Followers. Currently, only one Group association is presented in our web UI. | 
**GroupMentionIds** | **[]string** | An array of Group IDs that have been mentioned in the Story description. | 
**Id** | **int64** | The ID of the iteration. | 
**Labels** | [**[]Label**](Label.md) | An array of labels attached to the iteration. | 
**MemberMentionIds** | **[]string** | An array of Member IDs that have been mentioned in the Story description. | 
**MentionIds** | **[]string** | Deprecated: use member_mention_ids. | 
**Name** | **string** | The name of the iteration. | 
**StartDate** | [**time.Time**](time.Time.md) | The date this iteration begins. | 
**Stats** | [**IterationStats**](IterationStats.md) |  | 
**Status** | **string** | The status of the iteration. Values are either \&quot;unstarted\&quot;, \&quot;started\&quot;, or \&quot;done\&quot;. | 
**UpdatedAt** | [**time.Time**](time.Time.md) | The instant when this iteration was last updated. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


