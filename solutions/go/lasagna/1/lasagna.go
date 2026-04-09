package lasagna

// TODO: define the 'OvenTime' constant

const OvenTime int = 40
const TimeforLayer int = 2

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	var remaining_minutes int;
    remaining_minutes = OvenTime - actualMinutesInOven
	return remaining_minutes
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
    var PreparationTime int = TimeforLayer * numberOfLayers
    return PreparationTime
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	var ElapsedTime int = (TimeforLayer * numberOfLayers) + actualMinutesInOven
    return ElapsedTime
}
