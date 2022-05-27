package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(sliceOfLayers []string, average int) int {
	if average == 0 {
		average = 2
	}
	return len(sliceOfLayers) * average
}

// TODO: define the 'Quantities()' function
func Quantities(sliceOfLayers []string) (int, float64) {
	var noodles int
	var sauce float64
	for i := range sliceOfLayers {
		if sliceOfLayers[i] == "noodles" {
			noodles += 50
		}
		if sliceOfLayers[i] == "sauce" {
			sauce += 0.2
		}
	}
	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(frindsIngredient, myIngredient []string) {
	myIngredient[len(myIngredient)-1] = frindsIngredient[len(frindsIngredient)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amountsNeededForTwo []float64, numberOfPortions int) []float64 {
	var amoutsNeeded []float64
	for _, v := range amountsNeededForTwo {
		amoutsNeeded = append(amoutsNeeded, v)
	}
	for i := range amoutsNeeded {
		amoutsNeeded[i] *= float64(numberOfPortions) / 2
	}
	return amoutsNeeded
}
