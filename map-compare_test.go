package gotools_test

import (
	"testing"

	"github.com/cdvelop/gotools"
)

func TestAreMapsIdentical(t *testing.T) {

	// Caso 1: Mapas idénticos
	map1 := map[string]string{"a": "1", "b": "2"}
	map2 := map[string]string{"a": "1", "b": "2"}
	if !gotools.AreMapsIdentical(map1, map2) {
		t.Errorf("Los mapas deberían ser idénticos, pero no lo son.")
	}

	// Caso 2: Mapas con diferentes valores
	map3 := map[string]string{"a": "1", "b": "2"}
	map4 := map[string]string{"a": "1", "b": "3"}
	if gotools.AreMapsIdentical(map3, map4) {
		t.Errorf("Los mapas deberían tener diferentes valores, pero son idénticos.")
	}

	// Caso 3: Mapas con diferentes claves
	map5 := map[string]string{"a": "1", "b": "2"}
	map6 := map[string]string{"a": "1", "c": "2"}
	if gotools.AreMapsIdentical(map5, map6) {
		t.Errorf("Los mapas deberían tener diferentes claves, pero son idénticos.")
	}

}

func TestAreSliceMapsIdentical(t *testing.T) {
	// Caso 1: Slices de mapas idénticos
	slice1 := []map[string]string{
		{"a": "1", "b": "2"},
		{"c": "3", "d": "4"},
	}
	slice2 := []map[string]string{
		{"a": "1", "b": "2"},
		{"c": "3", "d": "4"},
	}
	if !gotools.AreSliceMapsIdentical(slice1, slice2) {
		t.Errorf("Los slices de mapas deberían ser idénticos, pero no lo son.")
	}

	// Caso 2: Slices de mapas con diferentes valores
	slice3 := []map[string]string{
		{"a": "1", "b": "2"},
		{"c": "3", "d": "4"},
	}
	slice4 := []map[string]string{
		{"a": "1", "b": "2"},
		{"c": "3", "d": "5"},
	}
	if gotools.AreSliceMapsIdentical(slice3, slice4) {
		t.Errorf("Los slices de mapas deberían tener diferentes valores, pero son idénticos.")
	}

	// Caso 3: Slices de mapas con diferentes claves
	slice5 := []map[string]string{
		{"a": "1", "b": "2"},
		{"c": "3", "d": "4"},
	}
	slice6 := []map[string]string{
		{"a": "1", "b": "2"},
		{"e": "3", "d": "4"},
	}
	if gotools.AreSliceMapsIdentical(slice5, slice6) {
		t.Errorf("Los slices de mapas deberían tener diferentes claves, pero son idénticos.")
	}
}
