package lasagna

func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}

func Quantities(layers []string) (noddlesG int, sauceL float64) {
	for _, layer := range layers {
		if layer == "noodles" {
			noddlesG += 50
		}
		if layer == "sauce" {
			sauceL += 0.2
		}
	}
	return
}

func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities []float64, scale int) []float64 {
	result := make([]float64, len(quantities))
	for i, quantity := range quantities {
		result[i] = (quantity / 2.0) * float64(scale)
	}
	return result
}
