# Workflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoAssignOwner** | **bool** | Indicates if an owner is automatically assigned when an unowned story is started. | 
**CreatedAt** | [**time.Time**](time.Time.md) | The date the Workflow was created. | 
**DefaultStateId** | **int64** | The unique ID of the default state that new Stories are entered into. | 
**Description** | **string** | A description of the workflow. | 
**EntityType** | **string** | A string description of this resource. | 
**Id** | **int64** | The unique ID of the Workflow. | 
**Name** | **string** | The name of the workflow. | 
**ProjectIds** | **[]float64** | An array of IDs of projects within the Workflow. | 
**States** | [**[]WorkflowState**](WorkflowState.md) | A map of the states in this Workflow. | 
**TeamId** | **int64** | The ID of the team the workflow belongs to. | 
**UpdatedAt** | [**time.Time**](time.Time.md) | The date the Workflow was updated. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


