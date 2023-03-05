package lasagna

const (
	OvenTime  = 40
	LayerTime = 2
)

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	if actualMinutesInOven < OvenTime {
		return OvenTime - actualMinutesInOven
	} else {
		return 0
	}
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	return LayerTime * numberOfLayers
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time
// and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return PreparationTime(numberOfLayers) + actualMinutesInOven
}
