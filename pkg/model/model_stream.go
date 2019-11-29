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

// Stream - A Stream
type Stream struct {

	Id string `json:"id"`

	Name string `json:"name"`

	Description string `json:"description,omitempty"`

	ProjectId string `json:"project_id"`

	Payload []map[string]interface{} `json:"payload"`

	Roles []StreamRole `json:"roles"`

	// The names of the schemas or compute libraries that can write to this stream. For example `[ \"bhom\", \"speckle\", \"3drepo\" ]`
	Senders []string `json:"senders,omitempty"`
}