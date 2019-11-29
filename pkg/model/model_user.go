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

// User - A User
type User struct {

	Id string `json:"id"`

	TeamRoles TeamRole `json:"team_roles"`

	// The user's public PGP signing Key
	PublicKey string `json:"public_key"`
}
