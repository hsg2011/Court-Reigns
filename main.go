package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)



func main() {
	 // cards just prompts for now 
	 prompts := []string{
		"Star player demands a pay raise. (l = pay, r = refuse)",
        "Offer discounted student tickets this weekend? (l = yes, r = no)",
        "Push players extra hard in practice? (l = hard, r = easy)",
    }

	// Display the prompts
	reader := bufio.NewReader(os.Stdin)

	// loop through each prompt

	for i, text := range prompts {
		fmt.Printf("\n[%d/%d] %s\n", i+1, len(prompts), text)
		fmt.Print("Your choice (l/r): ")

		// read one line of input
		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(strings.ToLower(input))

		// validate input

		if choice != "l" && choice != "r" {
			fmt.Println("Invalid choice. Please enter 'l' or 'r'.")
			i-- // decrement i to repeat this prompt
			continue
		}

		// Process the choice
		if choice == "l" {
			fmt.Println("You chose the left option.")
		} else {
			fmt.Println("You chose the right option.")
		}
	}

	fmt.Println("\nAll prompts completed. Thank you for playing!")
	// End of the game

}
