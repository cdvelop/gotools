package gotools

// Función para comparar dos mapas
func AreMapsIdentical(map_a, map_b map[string]string) bool {
	// Verificar si los mapas tienen la misma cantidad de elementos
	if len(map_a) != len(map_b) {
		return false
	}

	// Verificar si los pares clave-valor en el mapa son idénticos
	for key, value_a := range map_a {
		value_b, found := map_b[key]
		if !found || value_a != value_b {
			return false
		}
	}

	return true
}

// Función para comparar dos slices de mapas usando AreMapsIdentical
func AreSliceMapsIdentical(maps_a, maps_b []map[string]string) bool {
	// Verificar si los slices tienen la misma longitud
	if len(maps_a) != len(maps_b) {
		return false
	}

	// Verificar cada mapa individual usando la función AreMapsIdentical
	for i := 0; i < len(maps_a); i++ {
		if !AreMapsIdentical(maps_a[i], maps_b[i]) {
			return false
		}
	}

	return true
}
