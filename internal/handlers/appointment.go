package handlers

import (
	"net/http"
	"encoding/json"
)

// Appointment represents the appointment data.
type Appointment struct {
	ID          string `json:"id"`
	PatientName string `json:"patient_name"`
	DateTime    string `json:"date_time"`
}

// AppointmentHandler handles appointment-related endpoints.
type AppointmentHandler struct {
	appointments []Appointment
}

// NewAppointmentHandler creates a new AppointmentHandler.
func NewAppointmentHandler() *AppointmentHandler {
	return &AppointmentHandler{}
}

// CreateAppointment handles the creation of a new appointment.
func (h *AppointmentHandler) CreateAppointment(w http.ResponseWriter, r *http.Request) {
	var appointment Appointment
	if err := json.NewDecoder(r.Body).Decode(&appointment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.appointments = append(h.appointments, appointment)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(appointment)
}

// GetAppointments handles retrieving all appointments.
func (h *AppointmentHandler) GetAppointments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.appointments)
}