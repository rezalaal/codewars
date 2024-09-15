package kata8

import (
	"math"
)

/*
#Grasshopper - Terminal game combat function

Create a combat function that takes the player's current health and the amount of damage recieved, and returns the player's new health. Health can't be less than 0.

note:
//Should return the correct positive value
//Should return 0 rather than negative health
*/
func Combat(health, damage float64) float64 {
	if health-damage < 0 {
		return 0
	}
	return health - damage
}


func Combat1(health, damage float64) float64 {
  return math.Max(health-damage, 0.0)
}