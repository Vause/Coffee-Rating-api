package models

type RatingSummary struct {
	RatingId        uint    `json:"id" gorm:"primary_key"`
	CoffeeAmount    float32 `json:"coffee_amount"`
	CoffeeBrand     string  `json:"coffee_brand"`
	CoffeeRoastType string  `json:"coffee_roast_type"`
	BrewMethod      string  `json:"brew_method"`
	GrindSize       string  `json:"grind_size"`
	WaterAmount     float32 `json:"water_amount"`
	WaterTemp       float32 `json:"water_temp"`
	SteepTime       float32 `json:"steep_time"`
	MilkAmount      float32 `json:"milk_amount"`
	MilkHeatTime    float32 `json:"milk_heat_time"`
	CoffeeMadeDate  string  `json:"coffee_made_date"`
}

// CREATE TABLE rating_summary
// (
//     rating_id INTEGER PRIMARY KEY AUTOINCREMENT,
//     coffee_amount REAL,
//     coffee_brand TEXT,
//     coffee_roast_type TEXT,
//     brew_method TEXT,
//     grind_size TEXT,
//     water_amount REAL,
//     water_temp REAL,
//     steep_time REAL,
//     milk_amount REAL,
//     milk_heat_time REAL,
//     coffee_made_date NUMERIC
// );
