/*
 * aecdeltas
 *
 * The AEC Deltas OpenAPI Spec
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi


// NewProject - A Project
type NewProject struct {

	Name string `json:"name"`

	Description string `json:"description,omitempty"`

	Roles []ProjectRole `json:"roles,omitempty"`
}