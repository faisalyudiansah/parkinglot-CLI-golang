package view

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"parking-lot-cli/parking"
	"parking-lot-cli/variable"
	"parking-lot-cli/view/components"
)

func RegisterParkingLot(newAttendant *parking.Attendant) {
	scanner := bufio.NewScanner(os.Stdin)
	finishSession := false
	for !finishSession {
		components.Header("REGISTER PARKING LOT")
		capacity := promptInput(scanner, "Capacity (*numbers): ")
		convertCapacityToNumber, err := strconv.Atoi(capacity)
		if err != nil {
			fmt.Println(variable.ErrorInputNotNumber)
			fmt.Println()
			continue
		}
		if convertCapacityToNumber <= 0 {
			fmt.Println(variable.ErrorInputInvalid)
			fmt.Println()
			continue
		}
		idLot := 1
		if len(newAttendant.Lot) > 0 {
			idLot = len(newAttendant.Lot) + 1
		}
		formatID := fmt.Sprintf("idLot-%v", idLot)
		newLot := parking.NewLot(convertCapacityToNumber, formatID)
		newAttendant.AddLotToTheList(newLot)
		newLot.RegisterObserver(newAttendant)
		fmt.Println(variable.SuccessRegisterLot + " with ID: " + formatID + " | Capacity : " + capacity)
		finishSession = true
	}
}
