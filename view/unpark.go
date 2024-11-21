package view

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"parking-lot-cli/entity"
	"parking-lot-cli/parking"
	"parking-lot-cli/variable"
	"parking-lot-cli/view/components"
)

func Unpark(newAttendant *parking.Attendant, listTickets map[string]*entity.Ticket) {
	scanner := bufio.NewScanner(os.Stdin)
	finishSession := false
	if (len(newAttendant.Lot) == 0 && len(newAttendant.LotUnavailable) == 0) || len(listTickets) == 0 {
		fmt.Println(variable.ErrorLotIsEmpty)
		finishSession = true
	}
	for !finishSession {
		components.Header("UNPARK")
		ticket := promptInput(scanner, "Ticket: ")
		_, errValidate := strconv.Atoi(ticket)
		if errValidate != nil {
			fmt.Println(variable.ErrorInputInvalid)
			fmt.Println()
			continue
		}
		_, errUnpark := newAttendant.UnPark(listTickets[ticket])
		if errUnpark != nil {
			fmt.Println()
			continue
		}
		delete(listTickets, ticket)
		finishSession = true
	}
}
