package lasagna

import (
	"fmt"
	"strconv"
)

const OvenTime = 40

func main() {
	fmt.Println("Remaining time in oven: " + strconv.Itoa(RemainingOvenTime(30)))
	fmt.Println("Preparation time in minutes: " + strconv.Itoa(PreparationTime(3)))
	fmt.Println("Elapsed time: " + strconv.Itoa(ElapsedTime(20, 20)))

}

func RemainingOvenTime(current int) int {
	return OvenTime - current
}

func PreparationTime(numberOfLayers int) int {
	timePerLayerInMinutes := numberOfLayers * 2
	return timePerLayerInMinutes
}

func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return PreparationTime(numberOfLayers) + actualMinutesInOven
}
