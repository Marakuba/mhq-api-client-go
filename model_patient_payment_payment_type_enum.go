/*
 * MedHQ API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
// PatientPaymentPaymentTypeEnum : * `cash` - Наличные * `card` - Банковская карта * `certificate` - Сертификат * `sbp` - СБП
type PatientPaymentPaymentTypeEnum string

// List of PatientPaymentPaymentTypeEnum
const (
	CASH_PatientPaymentPaymentTypeEnum PatientPaymentPaymentTypeEnum = "cash"
	CARD_PatientPaymentPaymentTypeEnum PatientPaymentPaymentTypeEnum = "card"
	CERTIFICATE_PatientPaymentPaymentTypeEnum PatientPaymentPaymentTypeEnum = "certificate"
	SBP_PatientPaymentPaymentTypeEnum PatientPaymentPaymentTypeEnum = "sbp"
)