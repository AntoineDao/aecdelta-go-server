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

// NewRevision - A Revision
type NewRevision struct {

	StreamId string `json:"stream_id"`

	Delta map[string]interface{} `json:"delta,omitempty"`
}
