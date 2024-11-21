package view

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"parking-lot-cli/entity"
	"parking-lot-cli/parking"
	"parking-lot-cli/variable"
	"parking-lot-cli/view/components"
)

func Park(newAttendant *parking.Attendant, listTickets map[string]*entity.Ticket) {
	scanner := bufio.NewScanner(os.Stdin)
	finishSession := false
	if len(newAttendant.Lot) == 0 {
		fmt.Println(variable.ErrorLotIsEmpty)
		finishSession = true
	}
	for !finishSession {
		components.Header("PARK A CAR")
		plateNumber := promptInput(scanner, "Plate Number: ")
		plateNumber = strings.Trim(plateNumber, " ")
		if len(plateNumber) < 1 {
			fmt.Println(variable.ErrorInputInvalid)
			fmt.Println()
			continue
		}
		ticket, err := newAttendant.Park(entity.NewCarToPark(plateNumber))
		if err != nil {
			finishSession = true
			break
		}
		listTickets[ticket.ID] = ticket
		finishSession = true
	}
}
