package datepicker_feature

const (
	monthsTotal = 12
	yearsTotal  = 50
	firstYear   = 2015
)

func createMonths() []int {
	result := make([]int, monthsTotal)
	for i := 0; i < monthsTotal; i++ {
		result[i] = i + 1
	}
	return result
}

func createYears() []int {
	result := make([]int, yearsTotal)
	for i := 0; i < yearsTotal; i++ {
		result[i] = i + firstYear
	}
	return result
}
