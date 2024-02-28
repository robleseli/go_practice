package main // Declare the package name, which is main for an executable Go program.

import (
	"bufio"   // Import the bufio package to read input from the console.
	"fmt"     // Import the fmt package to print formatted output.
	"log"     // Import the log package to log errors.
	"math/rand" // Import the rand package to generate random numbers.
	"os"      // Import the os package for operating system-specific functionality.
	"strconv"  // Import the strconv package for string conversions.
	"strings"  // Import the strings package for string manipulation.
	"time"    // Import the time package to work with dates and times.
)

func main() {
	seconds := time.Now().Unix() // Get the current Unix timestamp.
	rand.Seed(seconds) // Seed the random number generator with the current timestamp.

	target := rand.Intn(100) + 1 // Generate a random number between 1 and 100.

	// Display instructions to the user.
	fmt.Println("There's a random number between 1 and 100")
	fmt.Println("Can you guess it?")
	// fmt.Println(target) // Omit this line; it displays the target number for debugging purposes.

	reader := bufio.NewReader(os.Stdin) // Create a reader to read input from the console.

	// Loop for a maximum of 10 guesses.
	for guesses := 0; guesses < 10; guesses++ {// Display the number of guesses remaining to the user.
		fmt.Println("You have", 10-guesses, "guesses left.")
		// Prompt the user to make a guess.
		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n') // Read the user's input from the console.
		if err != nil {
			log.Fatal(err) // Log any errors that occur while reading 
							//or converting input, and exit the program.
		}

		input = strings.TrimSpace(input) // Remove leading and trailing spaces from the input.
		guess, err := strconv.Atoi(input) // Convert the input string to an integer.
		if err != nil {
			log.Fatal(err) // Log any errors that occur while converting the input to an integer.
		}

		// Compare the user's guess with the target number and provide feedback.
		if guess < target {
			fmt.Println("Oops. Your guess was LOW.")
		} else if guess > target {
			fmt.Println("Oops. Your guess was HIGH")
		} else {
			// The user guessed the correct number, break out of the loop.
			fmt.Println("Good job! You guessed it!")
			break // Exit the loop when the correct guess is made.
		}
	}
}