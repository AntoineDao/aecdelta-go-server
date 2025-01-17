/*
 * aecdeltas
 *
 * The AEC Deltas OpenAPI Spec
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi


// Project - A Project
type Project struct {

	Id string `json:"id"`

	Name string `json:"name"`

	Description string `json:"description,omitempty"`

	Streams []string `json:"streams"`

	Roles []ProjectRole `json:"roles"`
}
