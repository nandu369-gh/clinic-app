package models

import "time"

type Appointment struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	PatientName     string    `json:"patient_name"binding:"required"`
	DoctorName      string    `json:"doctor_name"binding:"required"`
	AppointmentDate time.Time `json:"appointment_date"binding:"required"`
	Symptoms        string    `json:"symptoms"`
	CreatedAt       time.Time `json:"created_at"`
}
