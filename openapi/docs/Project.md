# Project

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Abbreviation** | Pointer to **string** | The Project abbreviation used in Story summaries. Should be kept to 3 characters at most. | 
**AppUrl** | **string** | The Clubhouse application url for the Project. | 
**Archived** | **bool** | True/false boolean indicating whether the Project is in an Archived state. | 
**Color** | Pointer to **string** | The color associated with the Project in the Clubhouse member interface. | 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | The time/date that the Project was created. | 
**DaysToThermometer** | **int64** | The number of days before the thermometer appears in the Story summary. | 
**Description** | Pointer to **string** | The description of the Project. | 
**EntityType** | **string** | A string description of this resource. | 
**ExternalId** | Pointer to **string** | This field can be set to another unique ID. In the case that the Project has been imported from another tool, the ID in the other tool can be indicated here. | 
**FollowerIds** | **[]string** | An array of UUIDs for any Members listed as Followers. | 
**Id** | **int64** | The unique ID of the Project. | 
**IterationLength** | **int64** | The number of weeks per iteration in this Project. | 
**Name** | **string** | The name of the Project | 
**ShowThermometer** | **bool** | Configuration to enable or disable thermometers in the Story summary. | 
**StartTime** | [**time.Time**](time.Time.md) | The date at which the Project was started. | 
**Stats** | [**ProjectStats**](ProjectStats.md) |  | 
**TeamId** | **int64** | The ID of the team the project belongs to. | 
**UpdatedAt** | Pointer to [**time.Time**](time.Time.md) | The time/date that the Project was last updated. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


