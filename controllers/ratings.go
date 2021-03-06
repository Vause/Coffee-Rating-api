package controllers

import (
	"net/http"

	"github.com/Vause/Coffee-Rating-api/models"
	"github.com/Vause/Coffee-Rating-api/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var zapLogger *zap.SugaredLogger

type CreateRatingInput struct {
	CoffeeAmount    *float32 `json:"coffee_amount" binding:"required"`
	CoffeeBrand     string   `json:"coffee_brand" binding:"required"`
	CoffeeRoastType string   `json:"coffee_roast_type" binding:"required"`
	BrewMethod      string   `json:"brew_method" binding:"required"`
	GrindSize       string   `json:"grind_size" binding:"required"`
	WaterAmount     *float32 `json:"water_amount" binding:"required"`
	WaterTemp       *float32 `json:"water_temp" binding:"required"`
	SteepTime       *float32 `json:"steep_time" binding:"required"`
	MilkAmount      *float32 `json:"milk_amount" binding:"required"`
	MilkHeatTime    *float32 `json:"milk_heat_time" binding:"required"`
	TotalRating     *float32 `json:"total_rating"`
	Notes           string   `json:"notes"`
	CoffeeMadeDate  string   `json:"coffee_made_date" binding:"required"`
}

type UpdateRatingInput struct {
	CoffeeAmount    *float32 `json:"coffee_amount"`
	CoffeeBrand     string   `json:"coffee_brand"`
	CoffeeRoastType string   `json:"coffee_roast_type"`
	BrewMethod      string   `json:"brew_method"`
	GrindSize       string   `json:"grind_size"`
	WaterAmount     *float32 `json:"water_amount"`
	WaterTemp       *float32 `json:"water_temp"`
	SteepTime       *float32 `json:"steep_time"`
	MilkAmount      *float32 `json:"milk_amount"`
	MilkHeatTime    *float32 `json:"milk_heat_time"`
	TotalRating     *float32 `json:"total_rating"`
	Notes           string   `json:"notes"`
	CoffeeMadeDate  string   `json:"coffee_made_date"`
}

func InitLog(logger *zap.SugaredLogger) {
	zapLogger = logger
}

func GetRatings(c *gin.Context) {
	var ratings []models.RatingSummary

	if err := models.DB.Find(&ratings).Error; err != nil {
		zapLogger.Errorw("Get Ratings Failure",
			"content", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"data": ratings})
}

func GetRatingById(c *gin.Context) {
	var rating models.RatingSummary

	if err := models.DB.Where("rating_id = ?", c.Param("id")).First(&rating).Error; err != nil {
		zapLogger.Errorw("Get Rating Failure",
			"content", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rating})
}

func CreateRating(c *gin.Context) {
	var input CreateRatingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		zapLogger.Errorw("Error Creating Rating",
			"err", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rating := models.RatingSummary{
		CoffeeAmount:    *input.CoffeeAmount,
		CoffeeBrand:     input.CoffeeBrand,
		CoffeeRoastType: input.CoffeeRoastType,
		BrewMethod:      input.BrewMethod,
		GrindSize:       input.GrindSize,
		WaterAmount:     *input.WaterAmount,
		WaterTemp:       *input.WaterTemp,
		SteepTime:       *input.SteepTime,
		MilkAmount:      *input.MilkAmount,
		MilkHeatTime:    *input.MilkHeatTime,
		TotalRating:     *input.TotalRating,
		Notes:           input.Notes,
		CoffeeMadeDate:  input.CoffeeMadeDate,
	}

	models.DB.Create(&rating)

	c.JSON(http.StatusOK, gin.H{"data": rating})
}

func UpdateRating(c *gin.Context) {
	var rating models.RatingSummary
	if err := models.DB.Where("rating_id = ?", c.Param("id")).First(&rating).Error; err != nil {
		zapLogger.Errorw("Update Rating Error",
			"err", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateRatingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		zapLogger.Errorw("Error Binding Record",
			"err", err.Error(),
			"rating", rating)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var interfaceInput, _ = utils.StructToMap(input)
	models.DB.Model(&rating).Updates(interfaceInput)

	c.JSON(http.StatusOK, gin.H{"data": rating})
}

func DeleteRating(c *gin.Context) {
	var rating models.RatingSummary
	if err := models.DB.Where("rating_id = ?", c.Param("id")).First(&rating).Error; err != nil {
		zapLogger.Errorw("Delete Rating Failure",
			"err", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Model(&rating).Delete(&rating)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
