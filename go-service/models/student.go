package models

type Student struct {
	ID                 int     `json:"id"`
	Name               string  `json:"name"`
	Email              string  `json:"email"`
	SystemAccess       bool    `json:"systemAccess"`
	Phone              *string `json:"phone"`
	Gender             *string `json:"gender"`
	DOB                *string `json:"dob"`
	Class              *string `json:"class"`
	Section            *string `json:"section"`
	Roll               *int    `json:"roll"`
	FatherName         *string `json:"fatherName"`
	FatherPhone        *string `json:"fatherPhone"`
	MotherName         *string `json:"motherName"`
	MotherPhone        *string `json:"motherPhone"`
	GuardianName       *string `json:"guardianName"`
	GuardianPhone      *string `json:"guardianPhone"`
	RelationOfGuardian *string `json:"relationOfGuardian"`
	CurrentAddress     *string `json:"currentAddress"`
	PermanentAddress   *string `json:"permanentAddress"`
	AdmissionDate      *string `json:"admissionDate"`
	ReporterName       *string `json:"reporterName"`
}

type StudentResponse struct {
	Success bool    `json:"success"`
	Data    Student `json:"data"`
	Message string  `json:"message,omitempty"`
}
