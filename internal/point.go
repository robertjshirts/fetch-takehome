package internal

import (
	"math"
	"strconv"
	"strings"

	"github.com/robertjshirts/fetch-takehome/gen"
)

func isAlphanumeric(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func GetPoints(receipt *gen.Receipt) (int, error) {
	points := 0

	// 1 point for each alphanumeric character in the retailer name
	for _, c := range receipt.Retailer {
		if isAlphanumeric(c) {
			points += 1
		}
	}

	// Get cents from total
	_, stringCents, _ := strings.Cut(receipt.Total, ".")
	cents, err := strconv.Atoi(stringCents)
	if err != nil {
		return 0, err
	}

	// 50 points for a total ending in .00
	if cents == 0 {
		points += 50
	}

	// 25 points for a total ending in a multiple of .25
	if cents%25 == 0 {
		points += 25
	}

	// 5 points for every 2 items
	points += int(len(receipt.Items)/2) * 5

	for _, item := range receipt.Items {
		// Check length of description
		descLength := len(strings.TrimSpace(item.ShortDescription))
		if descLength%3 == 0 {
			// Parse item price
			price, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				return 0, err
			}

			// Add points
			points += int(math.Ceil(price * 0.2))
		}
	}

	// 6 points if purchase date is odd
	if receipt.PurchaseDate.Day()%2 == 1 {
		points += 6
	}

	// Get hour from purchase time
	stringHour, _, _ := strings.Cut(receipt.PurchaseTime, ":")
	hour, err := strconv.Atoi(stringHour)
	if err != nil {
		return 0, err
	}

	// 10 points if between 2:00PM-3:59PM
	if hour == 14 || hour == 15 {
		points += 10
	}

	return points, nil
}
