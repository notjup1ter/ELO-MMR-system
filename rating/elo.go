package rating

import "math"

func getExpectedScore(rating1, rating2 float64) float64 {
  return (1.0/(1.0 + math.pow(10, ((rating1 - rating2)/400)))) 
}

func updateELOs(yourRating, opponentRating float64, score float64) float64 {
  kFactor := 32 //constant required for ELO algorithm
  expectedScore := getExpectedScore(yourRating, opponentRating)
  return yourRating + (kFactor * (score - expectedScore)) 
}
