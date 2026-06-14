package controllers

import (
	"clinic-app/config"
	"clinic-app/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateAppointment handles POST /appointments
func CreateAppointment(c *gin.Context) {
	var input models.Appointment
	// Validate JSON input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Raw SQL Insert returning the generated ID and Timestamp
	query := `INSERT INTO appointments (patient_name, doctor_name,appointment_date, symptoms)
            VALUES ($1, $2, $3, $4)RETURNING id, created_at`
	err := config.DB.QueryRow(context.Background(), query,
		input.PatientName,
		input.DoctorName,
		input.AppointmentDate,
		input.Symptoms).Scan(&input.ID, &input.CreatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error: " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Appointment scheduled successfully",
		"data": gin.H{
			"id":               input.ID,
			"patient_name":     input.PatientName,
			"doctor_name":      input.DoctorName,
			"appointment_date": input.AppointmentDate.Format("2006-01-02 15:04:05"), // Clean format
			"symptoms":         input.Symptoms,
			"created_at":       input.CreatedAt.Format("2006-01-02 15:04:05"), // Clean format
		},
	})
}
