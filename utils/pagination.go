package utils

import (
	"hkn-be/constants"
	"hkn-be/objects"
	"strconv"
)

func CalculateTotalPage(total, perPage int) int {
	if (total % perPage) > 0 {
		return (total / perPage) + 1
	} else {
		return total / perPage
	}
}

func GetOffset(page, limit int) int {
	return (page - 1) * limit
}

func ValidatePageLimit(filter *objects.PageLimit) error {
	if filter.Page == "0" || len(filter.Page) <= 0 {
		filter.Page = "1"
	}
	if filter.Limit == "0" || len(filter.Limit) <= 0 {
		filter.Limit = "20"
	}
	_, err := strconv.Atoi(filter.Page)
	if err != nil {
		return constants.ErrPageLimitWrongFormat
	}
	_, err = strconv.Atoi(filter.Limit)
	if err != nil {
		return constants.ErrPageLimitWrongFormat
	}
	return nil
}

func GetInt(num string) int {
	result, _ := strconv.Atoi(num)
	return result
}
