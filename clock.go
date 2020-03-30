package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {

	digitHolder := [10]placeholder{
		Zero,
		One,
		Two,
		Three,
		Four,
		Five,
		Six,
		Seven,
		Eight,
		Nine,
	}

	var displayHolder [8]placeholder

	for i := 0; ; i++ {
		currentTime := time.Now()
		var minute int = currentTime.Minute()
		seconds := currentTime.Second()
		hours := currentTime.Hour()

		//Display the hours at first as first two digit

		displayHolder[0] = digitHolder[hours/10]
		displayHolder[1] = digitHolder[hours%10]

		//display seperator digits after every two seconds
		if i%2 == 0 {
			displayHolder[2] = DoubleDot
			displayHolder[5] = DoubleDot
		} else{
			displayHolder[2]=EmptyDot
			displayHolder[5]=EmptyDot
		}

		//Display the minutes at index 3 and 4 of displayHolder

		displayHolder[3] = digitHolder[minute/10]
		displayHolder[4] = digitHolder[minute%10]

		//display the seconds at index 6 and 7

		displayHolder[6] = digitHolder[seconds/10]
		displayHolder[7] = digitHolder[seconds%10]

		//row 0,1,2,3,4 and column 0,2,3,4,5,6,7
		for row := range displayHolder[0] {
			for column := range displayHolder {
				fmt.Print(displayHolder[column][row], " ")
			}
			fmt.Println()
		}

		time.Sleep(1 * time.Second)
        
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
		

	}

}
