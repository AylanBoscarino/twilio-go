/*
 * Twilio - Ip_messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ListUserResponse struct for ListUserResponse
type ListUserResponse struct {
	Meta ListCredentialResponseMeta `json:"Meta,omitempty"`
	Users []IpMessagingV1ServiceUser `json:"Users,omitempty"`
}
