package view

import (
	"fmt"

	"parking-lot-cli/parking"
	"parking-lot-cli/view/components"
)

func MainMenu(newAttendant *parking.Attendant) {
	fmt.Println()
	components.Header("DASHBOARD MENU")
	fmt.Printf("you are logged in as an attendant with ID : %v\n", newAttendant.IdAttendant)
	fmt.Println()
	menu := "parking lot menu\n" +
		"1. register parking lot\n" +
		"2. park\n" +
		"3. unpark\n" +
		"4. change parking style\n" +
		"5. exit\n"
	fmt.Println(menu)
}
