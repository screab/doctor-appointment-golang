package models

// Appointment represents an appointment in the system.
type Appointment struct {
    ID          string    `json:"id"`  // Unique identifier for the appointment
    PatientID   string    `json:"patient_id"` // Identifier for the patient
    DoctorID    string    `json:"doctor_id"` // Identifier for the doctor
    StartTime   string    `json:"start_time"` // When the appointment starts
    EndTime     string    `json:"end_time"`   // When the appointment ends
    Status      string    `json:"status"`  // Status of the appointment
}

// CreateAppointmentRequest represents the data needed to create an appointment.
type CreateAppointmentRequest struct {
    PatientID   string    `json:"patient_id"` // Identifier for the patient
    DoctorID    string    `json:"doctor_id"` // Identifier for the doctor
    StartTime   string    `json:"start_time"` // When the appointment should start
    EndTime     string    `json:"end_time"`   // When the appointment should end
}