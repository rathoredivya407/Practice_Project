package models

//this is the struct for the patient
type Patient struct {
	ID                uint32 `json:"Id"`
	FullName          string `json:"full_name"`
	Age               int    `json:"age"`
	Gender            string `json:"gender"`
	InsuranceProvider string `json:"insurance_provider"`
	PolicyNumber      string `json:"policy_number"`
}
