package lasagna

import "strings"

func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}

func Quantities(layers []string) (noddlesG int, sauceL float64) {
	noddlesG = strings.Count(strings.Join(layers, ""), "noodles") * 50
	sauceL = float64(strings.Count(strings.Join(layers, ""), "sauce")) * 0.2
	return
}

func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities []float64, scale int) (result []float64) {
	for _, quantity := range quantities {
		result = append(result, (quantity/2.0)*float64(scale))
	}
	return result
}
