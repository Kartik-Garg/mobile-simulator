package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Generate some random data as a placeholder
	var (
		dates  []float64
		prices []float64
	)
	for i := 0; i < 100; i++ {
		dates = append(dates, float64(i))
		prices = append(prices, float64(rand.Intn(100)))
	}

	// Calculate technical indicators (for demonstration)
	ma := simpleMovingAverage(prices, 10)
	rsi := relativeStrengthIndex(prices, 14)

	// Create a new regression model
	model := newRegressionModel(dates, ma, rsi, prices)

	// Predict the price for a new date
	newDate := 110.0
	predictedPrice := predictPrice(model, newDate, ma[len(ma)-1], rsi[len(rsi)-1])

	fmt.Printf("Predicted price for date %.2f: %.2f\n", newDate, predictedPrice)
}

func simpleMovingAverage(prices []float64, period int) []float64 {
	var sma []float64
	for i := 0; i < len(prices)-period+1; i++ {
		sum := 0.0
		for j := i; j < i+period; j++ {
			sum += prices[j]
		}
		sma = append(sma, sum/float64(period))
	}
	return sma
}

func relativeStrengthIndex(prices []float64, period int) []float64 {
	var rsi []float64
	for i := period; i < len(prices); i++ {
		upSum := 0.0
		downSum := 0.0
		for j := i - period; j < i; j++ {
			change := prices[j+1] - prices[j]
			if change > 0 {
				upSum += change
			} else {
				downSum -= change
			}
		}
		rs := upSum / downSum
		rsi = append(rsi, 100 - (100 / (1 + rs)))
	}
	return rsi
}

func newRegressionModel(dates, ma, rsi, prices []float64) func(float64, float64, float64) float64 {
	return func(date, movingAverage, relativeStrengthIndex float64) float64 {
		sum := 0.0
		for i, d := range dates {
			x := []float64{d, ma[i], rsi[i]}
			prediction := linearRegressionPredict(x, prices[i])
			sum += prediction
		}
		return sum / float64(len(dates))
	}
}

func linearRegressionPredict(features []float64, target float64) float64 {
	return features[0] + features[1] + features[2] + target
}

func predictPrice(model func(float64, float64, float64) float64, date, ma, rsi float64) float64 {
	return model(date, ma, rsi)
}
