package entities

import (
	"crypto/rand"
	"time"
)

// Ticket represents a parking ticket issued to a vehicle
type Ticket struct {
	ID           string    // Unique ticket ID
	Vehicle      Vehicle   // Vehicle that parked
	EntryTime    time.Time // When vehicle entered
	ExitTime     time.Time // When vehicle exited (zero if still parked)
	FloorID      int       // Which floor
	SpotID       int       // Which spot on the floor
	VehicleType  VehicleType
	PricePerHour int // Price per hour for this vehicle type
}

// NewTicket creates a new parking ticket
func NewTicket(vehicle Vehicle, floorID, spotID int, pricePerHour int) *Ticket {
	return &Ticket{
		ID:           generateTicketID(),
		Vehicle:      vehicle,
		EntryTime:    time.Now(),
		FloorID:      floorID,
		SpotID:       spotID,
		VehicleType:  vehicle.Type(),
		PricePerHour: pricePerHour,
	}
}

// CalculatePrice calculates the total parking fee based on duration
// Returns price in cents or smallest currency unit
func (t *Ticket) CalculatePrice() int {
	if t.ExitTime.IsZero() {
		// Vehicle is still parked, calculate based on current time
		duration := time.Since(t.EntryTime)
		hours := int(duration.Hours())
		if duration.Minutes() > float64(hours*60) {
			hours++ // Round up to next hour
		}
		if hours == 0 {
			hours = 1 // Minimum 1 hour charge
		}
		return hours * t.PricePerHour
	}

	// Vehicle has exited, calculate based on exit time
	duration := t.ExitTime.Sub(t.EntryTime)
	hours := int(duration.Hours())
	if duration.Minutes() > float64(hours*60) {
		hours++ // Round up to next hour
	}
	if hours == 0 {
		hours = 1 // Minimum 1 hour charge
	}
	return hours * t.PricePerHour
}

// GetDuration returns the parking duration
func (t *Ticket) GetDuration() time.Duration {
	if t.ExitTime.IsZero() {
		return time.Since(t.EntryTime)
	}
	return t.ExitTime.Sub(t.EntryTime)
}

// MarkExit records the exit time
func (t *Ticket) MarkExit() {
	t.ExitTime = time.Now()
}

// IsActive returns true if vehicle is still parked
func (t *Ticket) IsActive() bool {
	return t.ExitTime.IsZero()
}

// generateTicketID generates a unique ticket ID
// In production, use UUID or database sequence
func generateTicketID() string {
	return time.Now().Format("20060102150405") + "-" + randomString(6)
}

// randomString generates a random string using crypto/rand
func randomString(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		// Fallback to time-based if crypto/rand fails
		for i := range b {
			b[i] = charset[time.Now().UnixNano()%int64(len(charset))]
		}
		return string(b)
	}
	for i := range b {
		b[i] = charset[b[i]%byte(len(charset))]
	}
	return string(b)
}
