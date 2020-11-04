# EpicSlim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppUrl** | **string** | The Clubhouse application url for the Epic. | 
**Archived** | **bool** | True/false boolean that indicates whether the Epic is archived or not. | 
**Completed** | **bool** | A true/false boolean indicating if the Epic has been completed. | 
**CompletedAt** | Pointer to [**time.Time**](time.Time.md) | The time/date the Epic was completed. | 
**CompletedAtOverride** | Pointer to [**time.Time**](time.Time.md) | A manual override for the time/date the Epic was completed. | 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | The time/date the Epic was created. | 
**Deadline** | Pointer to [**time.Time**](time.Time.md) | The Epic&#39;s deadline. | 
**Description** | **string** | The Epic&#39;s description. | [optional] 
**EntityType** | **string** | A string description of this resource. | 
**EpicStateId** | **int64** | The ID of the Epic State. | 
**ExternalId** | Pointer to **string** | This field can be set to another unique ID. In the case that the Epic has been imported from another tool, the ID in the other tool can be indicated here. | 
**ExternalTickets** | [**[]ExternalTicket**](ExternalTicket.md) |  | 
**FollowerIds** | **[]string** | An array of UUIDs for any Members you want to add as Followers on this Epic. | 
**GroupId** | Pointer to **string** |  | 
**GroupMentionIds** | **[]string** | An array of Group IDs that have been mentioned in the Epic description. | 
**Id** | **int64** | The unique ID of the Epic. | 
**Labels** | [**[]Label**](Label.md) | An array of Labels attached to the Epic. | 
**MemberMentionIds** | **[]string** | An array of Member IDs that have been mentioned in the Epic description. | 
**MentionIds** | **[]string** | Deprecated: use member_mention_ids. | 
**MilestoneId** | Pointer to **int64** | The ID of the Milestone this Epic is related to. | 
**Name** | **string** | The name of the Epic. | 
**OwnerIds** | **[]string** | An array of UUIDs for any members you want to add as Owners on this new Epic. | 
**PlannedStartDate** | Pointer to [**time.Time**](time.Time.md) | The Epic&#39;s planned start date. | 
**Position** | **int64** | The Epic&#39;s relative position in the Epic workflow state. | 
**ProjectIds** | **[]int64** | The IDs of Projects related to this Epic. | 
**RequestedById** | **string** | The ID of the Member that requested the epic. | 
**Started** | **bool** | A true/false boolean indicating if the Epic has been started. | 
**StartedAt** | Pointer to [**time.Time**](time.Time.md) | The time/date the Epic was started. | 
**StartedAtOverride** | Pointer to [**time.Time**](time.Time.md) | A manual override for the time/date the Epic was started. | 
**State** | **string** | &#x60;Deprecated&#x60; The workflow state that the Epic is in. | 
**Stats** | [**EpicStats**](EpicStats.md) |  | 
**SupportTickets** | [**[]SupportTicket**](SupportTicket.md) |  | 
**UpdatedAt** | Pointer to [**time.Time**](time.Time.md) | The time/date the Epic was updated. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


