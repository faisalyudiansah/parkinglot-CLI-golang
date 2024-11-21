package view

import (
	"bufio"
	"fmt"
	"os"

	"parking-lot-cli/parking"
	"parking-lot-cli/variable"
	"parking-lot-cli/view/components"
)

func ChangeStyle(newAttendant *parking.Attendant) {
	scanner := bufio.NewScanner(os.Stdin)
	finishSession := false
	if len(newAttendant.Lot) == 0 {
		fmt.Println(variable.ErrorLotIsEmpty)
		finishSession = true
	}
	for !finishSession {
		components.Header("CHANGE PARKING STYLE")
		menu := "choose the parking style you want \n" +
			"1. default\n" +
			"2. most capacity\n" +
			"3. highest number\n"
		fmt.Println(menu)
		input := promptInput(scanner, "options: ")
		switch input {
		case "1":
			newAttendant.ChangeStylePark(newAttendant.OptionStyle("default"))
			fmt.Println("success to set the park style to default!")
			finishSession = true
		case "2":
			newAttendant.ChangeStylePark(newAttendant.OptionStyle("mostCapacity"))
			fmt.Println("success to set the park style to most capacity!")
			finishSession = true
		case "3":
			newAttendant.ChangeStylePark(newAttendant.OptionStyle("highestNumber"))
			fmt.Println("success to set the park style to highest number!")
			finishSession = true
		default:
			fmt.Println("invalid options park style!")
			fmt.Println()
		}
	}
}
