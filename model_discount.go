/*
 * MedHQ API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Discount struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	Value string `json:"value"`
	Comment string `json:"comment,omitempty"`
}
