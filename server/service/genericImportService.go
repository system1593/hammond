package service

import (
	"hammond/db"
	"hammond/models"
	"time"
)

func GenericParseRefuelings(content []models.ImportFillup, user *db.User, vehicle *db.Vehicle, timezone string) ([]db.Fillup, []string) {
	var errors []string
	var fillups []db.Fillup
	dateLayout := "2023-04-16T04:41:25.682Z"
	loc, _ := time.LoadLocation(timezone)
	for _, record := range content {
		date, err := time.ParseInLocation(record.Date, dateLayout, loc)
		if err != nil {
			date = time.Date(2000, time.December, 0, 0, 0, 0, 0, loc)
		}

		var missedFillup bool
		if record.HasMissedFillup == nil {
			missedFillup = false
		} else {
			missedFillup = *record.HasMissedFillup
		}

	fillups = append(fillups, db.Fillup{
		VehicleID: 				vehicle.ID,
		UserID: 					user.ID,
		Date: 						date,
		IsTankFull: 			record.IsTankFull,
		HasMissedFillup: 	&missedFillup,
		FuelQuantity: 		float32(record.FuelQuantity),
		PerUnitPrice: 		float32(record.PerUnitPrice),
		FillingStation: 	record.FillingStation,
		OdoReading: 			record.OdoReading,
		TotalAmount: 			float32(record.TotalAmount),
		FuelUnit: 				vehicle.FuelUnit,
		Currency: 				user.Currency,
		DistanceUnit: 		user.DistanceUnit,
		Comments: 				record.Comments,
		Source: 					"Generic Import",
	})
	}

	return fillups, errors
}