/*
 * Clubhouse API
 *
 * Clubhouse API
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SupportTicket struct for SupportTicket
type SupportTicket struct {
	ExternalId string `json:"external_id"`
	ExternalUrl string `json:"external_url,omitempty"`
	Id string `json:"id"`
	StoryIds []float64 `json:"story_ids"`
}
