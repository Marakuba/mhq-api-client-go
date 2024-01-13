/*
 * MedHQ API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type PatientLabOrder struct {
	Uid string `json:"uid,omitempty"`
	VisitNumber int32 `json:"visit_number"`
	IsCompleted bool `json:"is_completed,omitempty"`
	DtCompleted string `json:"dt_completed"`
	StaffName string `json:"staff_name"`
	Notes string `json:"notes"`
	Analysis []Result `json:"analysis"`
	Equipment string `json:"equipment"`
}