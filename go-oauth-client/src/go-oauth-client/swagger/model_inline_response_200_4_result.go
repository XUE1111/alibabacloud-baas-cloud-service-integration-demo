/*
 * BaaS API Gateway
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type InlineResponse2004Result struct {

	Response *Response `json:"response"`
	// Object download url
	Url string `json:"url"`
	// Object sha256 checksum in hex format
	Checksum string `json:"checksum"`
}