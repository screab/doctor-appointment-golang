package services

import "time"

// Appointment represents an appointment in the system.
type Appointment struct {
    ID          string    `json:"id"`
    Description string    `json:"description"`
    Time        time.Time `json:"time"`
}

// AppointmentService provides methods for managing appointments.
type AppointmentService struct {
    appointments map[string]Appointment
}

// NewAppointmentService creates a new instance of AppointmentService.
func NewAppointmentService() *AppointmentService {
    return &AppointmentService{
        appointments: make(map[string]Appointment),
    }
}

// CreateAppointment adds a new appointment.
func (s *AppointmentService) CreateAppointment(appointment Appointment) {
    s.appointments[appointment.ID] = appointment
}

// GetAppointment retrieves an appointment by ID.
func (s *AppointmentService) GetAppointment(id string) (Appointment, bool) {
    appointment, exists := s.appointments[id]
    return appointment, exists
}

// UpdateAppointment modifies an existing appointment.
func (s *AppointmentService) UpdateAppointment(appointment Appointment) bool {
    _, exists := s.appointments[appointment.ID]
    if exists {
        s.appointments[appointment.ID] = appointment
        return true
    }
    return false
}

// DeleteAppointment removes an appointment by ID.
func (s *AppointmentService) DeleteAppointment(id string) {
    delete(s.appointments, id)
}