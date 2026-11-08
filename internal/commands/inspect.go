package commands

import (
	"fmt"
	"net/http"
)

func commandInspect(client *http.Client, input string) error {
	if val, ok := caughtPokemons[input]; ok {
		fmt.Println("- ID: ", val.ID)
		fmt.Println("- Name: ", val.Name)
		fmt.Println("- Order:", val.Order)
		fmt.Println("- Gender Rate: ", val.GenderRate)
		fmt.Println("- Capture Rate: ", val.CaptureRate)
		fmt.Println("- Base, Happiness: ", val.BaseHappiness)
		fmt.Println("- Is Baby: ", val.IsBaby)
		fmt.Println("- Is IsLegendary: ", val.IsLegendary)
		fmt.Println("- Is Mythical: ", val.IsMythical)
		fmt.Println("- Hatch Counter: ", val.HatchCounter)
		fmt.Println("- Has Gender Differences: ", val.HasGenderDifferences)
		fmt.Println("- Forms Switchable: ", val.FormsSwitchable)
	} else {
		fmt.Println("You have not caught this pokemon")
	}
	return nil
}
