package entity_test

import (
	"testing"

	"parking-lot-cli/entity"

	"github.com/stretchr/testify/assert"
)

func TestCar(t *testing.T) {
	t.Run("should successfully add the car when customer want to park it", func(t *testing.T) {
		car := entity.NewCarToPark("Yamaha-400")
		expectResult := "Yamaha-400"
		assert.Equal(t, expectResult, car.PlateNumber)
	})
}
