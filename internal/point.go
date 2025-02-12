package internal

import (
	"math"
	"strconv"
	"strings"

	"github.com/robertjshirts/fetch-takehome/gen"
)

func isAlphanumeric(c int) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func GetPoints(receipt *gen.Receipt) (int, error) {
	points := 0

	// 1 point for each alphanumeric character in the retailer name
	for c := range receipt.Retailer {
		if isAlphanumeric(c) {
			points += 1
		}
	}

	// 50 points for a total ending in .00
	// 25 points for a total ending in a multiple of .25
	_, stringCents, _ := strings.Cut(receipt.Total, ".")
	cents, err := strconv.Atoi(stringCents)
	if err != nil {
		return 0, err
	}
	if cents == 0 {
		points += 50
	} else if cents%25 == 0 {
		points += 25
	}

	// 5 points for every 2 items
	points += int(len(receipt.Items)/2) * 5

	// If item desc is multiple of 3, then add item price * 0.2 rounded up
	for _, item := range receipt.Items {
		descLength := len(strings.TrimSpace(item.ShortDescription))
		if descLength%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				return 0, err
			}
			points += int(math.Ceil(price * 0.2))
		}
	}

	// 6 points if purchase date is odd
	if receipt.PurchaseDate.Day()%2 == 1 {
		points += 6
	}

	// 10 points if purchase hour is 14 or 15 (2:00 - 3:59)
	stringHour, _, _ := strings.Cut(receipt.PurchaseTime, ":")
	hour, err := strconv.Atoi(stringHour)
	if err != nil {
		return 0, err
	}
	if hour == 14 || hour == 15 {
		points += 10
	}

	return points, nil
}
