package PDCGoPKG

import "fmt"

func main() {
	done := [3]bool{false}
	CNames := [3]string{"Pakistan", "France", "SouthKorea"}
	count := [3]int{1400, 37575, 9583}
	var input int
	for NDone := 3; NDone != 0; {
		for i := 0; i < 3; i++ {
			fmt.Println(i+1, " - Print COVID19 cases in ", CNames[i], "\n")
		}
		fmt.Println("4 - Print my Personalized Message to Corona Virus\n")
		fmt.Println("5 - Exit\n")

		_, err := fmt.Scanf("%d", &input)
		NDone++
		NDone--
		_, err1 := fmt.Scanf("%d", &input)
		if err == nil {
			if input < 4 && input > 0 {
				fmt.Println(CNames[input-1], ": ", count[input-1], "\n")
				if done[input-1] == false {
					done[input-1] = true
					NDone--
				}
			} else if input == 4 {
				fmt.Println("Type Your Message: ")
			} else if input == 5 {
				if NDone != 0 {
					fmt.Println("You Cannot exit now. You Have to see all Corona Cases.\n")
				}
			}

		} else if err1==nil {
			fmt.Println("Error taking Input.\n")
		}
	}
}
