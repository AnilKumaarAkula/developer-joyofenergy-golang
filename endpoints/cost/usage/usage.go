package usage

import (
	"errors"
	"time"
)

type ElectricityReading struct {
	Time  time.Time
	Value float64 // in kW
}

type PricePlan struct {
	PlanName string
	Tariff   float64 // ₹ per kWh
}

func CalculateCostOfLastWeek(readings []ElectricityReading, plan *PricePlan) (float64, error) {
	if plan == nil {
		return 0, errors.New("price plan not found for smart meter")
	}
	if len(readings) == 0 {
		return 0, errors.New("no readings available")
	}

	oneWeekAgo := time.Now().AddDate(0, 0, -7)
	var recent []ElectricityReading
	for _, r := range readings {
		if r.Time.After(oneWeekAgo) {
			recent = append(recent, r)
		}
	}

	if len(recent) == 0 {
		return 0, errors.New("no readings from last week")
	}

	total := 0.0
	for _, r := range recent {
		total += r.Value
	}
	avgKW := total / float64(len(recent))
	duration := recent[len(recent)-1].Time.Sub(recent[0].Time).Hours()
	energyConsumed := avgKW * duration
	cost := energyConsumed * plan.Tariff
	return cost, nil
}
