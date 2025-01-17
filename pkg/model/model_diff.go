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

// Diff - A diff between two revisions
type Diff struct {

	RevisionDatetime string `json:"revision_datetime"`

	Author Author `json:"author"`

	Delta Delta `json:"delta"`
}
