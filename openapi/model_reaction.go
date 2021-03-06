/*
 * Clubhouse API
 *
 * Clubhouse API
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// Reaction Emoji reaction on a comment.
type Reaction struct {
	// Emoji text of the reaction.
	Emoji string `json:"emoji"`
	// Permissions who have reacted with this.
	PermissionIds []string `json:"permission_ids"`
}
