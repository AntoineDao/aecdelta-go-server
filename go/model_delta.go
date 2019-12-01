/*
 * aecdeltas
 *
 * The AEC Deltas OpenAPI Spec
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi


// Delta - An AEC Delta Object
type Delta struct {

	Hash string `json:"hash"`

	// created
	Created []map[string]interface{} `json:"created"`

	// deleted
	Deleted []string `json:"deleted"`

	// updated
	Updated []map[string]interface{} `json:"updated"`

	// from
	From string `json:"from"`

	// to
	To string `json:"to"`

	// timestamp
	Timestamp string `json:"timestamp"`

	// signature
	Signature string `json:"signature"`

	// sender
	Sender string `json:"sender"`
}
