package models_test

import (
	"javatransition/samples/javatransition/models"
	"testing"
)

func TestNewCity(t *testing.T) {
	expectedName := "bangalore"
	temperatues := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	city := models.NewCity1(expectedName, temperatues, false, true)

	t.Run("name", func(t *testing.T) {
		got := city.Name1()
		if got != expectedName {
			t.Errorf("Expected %v but got %v", expectedName, got)
		}
	})
}
