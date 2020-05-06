package maps

// GetMap Devuelve un tipo de dat Map
func GetMap() map[string]int {
	//miMap := make(map[string]int)

	miMap := map[string]int{
		"Esteban": 30,
		"Maria":   31,
	}

	miMap["edad1"] = 18
	miMap["edad2"] = 19

	delete(miMap, "edad1")

	miMap["edad2"] = 40

	return miMap
}

func GetKeyMap(key string) int {

	miMap := map[string]int{
		"Esteban": 30,
		"Maria":   31,
	}

	miMap["edad1"] = 18
	miMap["edad2"] = 19

	delete(miMap, "edad1")

	miMap["edad2"] = 40

	return miMap
}
