/*
 * Clubhouse API
 *
 * Clubhouse API
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// Profile A group of Member profile details.
type Profile struct {
	// A true/false boolean indicating whether the Member has been deactivated within Clubhouse.
	Deactivated bool `json:"deactivated"`
	DisplayIcon Icon `json:"display_icon"`
	// The primary email address of the Member with the Organization.
	EmailAddress *string `json:"email_address"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// This is the gravatar hash associated with email_address.
	GravatarHash *string `json:"gravatar_hash"`
	// The unique identifier of the profile.
	Id string `json:"id"`
	// The Member's username within the Organization.
	MentionName string `json:"mention_name"`
	// The Member's name within the Organization.
	Name *string `json:"name"`
	// If Two Factor Authentication is activated for this User.
	TwoFactorAuthActivated bool `json:"two_factor_auth_activated,omitempty"`
}
