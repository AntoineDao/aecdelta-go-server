/*
 * aecdeltas
 *
 * The AEC Deltas OpenAPI Spec
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

//package openapi
package model

// Role - An individual role for a team. Note that a role with null permissions cannot access the resource.
type Role struct {

	// The name of the role
	Name string `json:"name"`

	// A description of the role
	Description map[string]interface{} `json:"description,omitempty"`

	Permission map[string]interface{} `json:"permission,omitempty"`
}
