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

// NewTeam - A team contains a list of roles that are attributed to users
type NewTeam struct {

	Name string `json:"name"`

	Description string `json:"description,omitempty"`

	// A list of roles. Should always contain the default and owner roles
	Roles []Role `json:"roles"`
}