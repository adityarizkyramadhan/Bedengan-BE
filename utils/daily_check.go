package utils

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// DailyCheckKavlingRawSQL checks if the departure date has passed and updates kavling availability
func DailyCheckKavlingRawSQL(db *gorm.DB) error {
	currentTime := time.Now()

	// Query for invoices where the departure date has passed
	rows, err := db.Raw(`
        SELECT invoice_reservasis.user_id, reservasis.kavling_id 
        FROM invoice_reservasis 
        JOIN reservasis ON invoice_reservasis.id = reservasis.invoice_reservasi_id 
        WHERE invoice_reservasis.tanggal_kepulangan <= ?
    `, currentTime).Rows()
	if err != nil {
		return fmt.Errorf("error executing raw SQL: %w", err)
	}
	defer rows.Close()

	var userID string
	var kavlingID sql.NullString

	// Iterate through the rows and process the results
	for rows.Next() {
		if err := rows.Scan(&userID, &kavlingID); err != nil {
			return fmt.Errorf("error scanning row: %w", err)
		}

		// If kavlingID is valid (not null), update its is_available field
		if kavlingID.Valid {
			if err := db.Exec(`UPDATE kavlings SET is_available = true WHERE id = ?`, kavlingID.String).Error; err != nil {
				return fmt.Errorf("error updating kavling availability for kavling ID %s: %w", kavlingID.String, err)
			}
		}
	}

	// Check for errors that may have occurred during the iteration
	if err := rows.Err(); err != nil {
		return fmt.Errorf("error during row iteration: %w", err)
	}

	return nil
}
