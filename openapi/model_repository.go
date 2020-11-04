/*
 * Clubhouse API
 *
 * Clubhouse API
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)
// Repository Repository refers to a VCS repository.
type Repository struct {
	// The time/date the Repository was created.
	CreatedAt *time.Time `json:"created_at"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// The VCS unique identifier for the Repository.
	ExternalId *string `json:"external_id"`
	// The full name of the VCS repository.
	FullName *string `json:"full_name"`
	// The ID associated to the VCS repository in Clubhouse.
	Id *int64 `json:"id"`
	// The shorthand name of the VCS repository.
	Name *string `json:"name"`
	// The type of Repository. Currently this can only be \"github\".
	Type string `json:"type"`
	// The time/date the Repository was updated.
	UpdatedAt *time.Time `json:"updated_at"`
	// The URL of the Repository.
	Url *string `json:"url"`
}