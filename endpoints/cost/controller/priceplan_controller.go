package controller

import (
	"fmt"
	"joi-energy-golang/endpoints/cost/usage"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PricePlanController struct {
	MeterReadingService MeterReadingService
	PricePlanService    PricePlanService
}

type MeterReading struct {
	Time  time.Time
	Value float64
}

type MeterReadingService interface {
	GetReadings(smartMeterID string) []usage.ElectricityReading
}

type PricePlanService interface {
	GetPricePlan(smartMeterID string) *usage.PricePlan
}

func (pc *PricePlanController) GetLastWeekCost(c *gin.Context) {
	smartMeterID := c.Param("smartMeterId")
	readings := pc.MeterReadingService.GetReadings(smartMeterID)
	plan := pc.PricePlanService.GetPricePlan(smartMeterID)

	cost, err := usage.CalculateCostOfLastWeek(readings, plan)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"smartMeterId": smartMeterID,
		"pricePlanId":  plan.PlanName,
		"totalCost":    fmt.Sprintf("%.2f", cost),
	})
}
