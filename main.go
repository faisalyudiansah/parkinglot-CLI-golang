package main

import (
	"bufio"
	"fmt"
	"os"

	"parking-lot-cli/entity"
	"parking-lot-cli/parking"
	"parking-lot-cli/variable"
	"parking-lot-cli/view"
	"parking-lot-cli/view/components"
)

func main() {
	RunCli()
}

func promptInput(scanner *bufio.Scanner, text string) string {
	fmt.Print(text)
	scanner.Scan()
	return scanner.Text()
}

func RunCli() {
	listTickets := make(map[string]*entity.Ticket)
	newAttendant := parking.NewAttendant([]*parking.ParkingLot{}, "att-01")
	scanner := bufio.NewScanner(os.Stdin)
	exit := false
	components.ClearTerminal()
	for !exit {
		view.MainMenu(newAttendant)
		input := promptInput(scanner, "input menu: ")
		fmt.Println()
		switch input {
		case "1":
			view.RegisterParkingLot(newAttendant)
		case "2":
			view.Park(newAttendant, listTickets)
		case "3":
			view.Unpark(newAttendant, listTickets)
		case "4":
			view.ChangeStyle(newAttendant)
		case "5":
			fmt.Println("exit...")
			exit = true
		default:
			fmt.Println(variable.ErrorInputInvalid)
		}
	}
}
