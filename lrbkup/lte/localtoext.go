package lte

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

// BackupToExt function
func BackupToExt(i int, s string) {

	fmt.Printf("You selected %d) %s, is this correct?", i, s)
	fmt.Println("")

	confirmPrompt := promptui.Prompt{
		Label:     "Is this correct?",
		IsConfirm: true,
	}

	result, err := confirmPrompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed! %v\n", err)
		fmt.Println("I am going to reprompt you in the future")
		return
	}

	fmt.Println("You answered: ", result)
	fmt.Println("Proceeding to do stuff")

}
