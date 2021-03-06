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
	"os"
)
// CreateFileObj struct for CreateFileObj
type CreateFileObj struct {
	ContentType string `json:"content-type"`
	Filename string `json:"filename"`
	Size int64 `json:"size"`
	Tempfile *os.File `json:"tempfile"`
}
