/*
 * MedHQ API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type SendingAttempt struct {
	Id int32 `json:"id"`
	DtCreated string `json:"dt_created,omitempty"`
	IsSuccess bool `json:"is_success"`
	Error_ string `json:"error,omitempty"`
	Record int32 `json:"record"`
}
