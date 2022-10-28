package service

func KondisiWater(water int) string {
	var waterStatus string
	if water <= 5 {
		waterStatus = "Normal"
	} else if water >= 6 && water <= 8 {
		waterStatus = "Attention"
	} else {
		waterStatus = "Danger"
	}
	return waterStatus
}

func KondisiWind(wind int) string {
	var windStatus string
	if wind <= 6 {
		windStatus = "Normal"
	} else if wind >= 7 && wind <= 15 {
		windStatus = "Attention"
	} else {
		windStatus = "Danger"
	}
	return windStatus
}
