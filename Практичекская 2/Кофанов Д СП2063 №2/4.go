package main

import "fmt"

func printVotes(votes []string) {
	ann := 0
	boris := 0
	viktor := 0

	for _, name := range votes {
		if name == "Артем" {
			ann++
		}
		if name == "Артур" {
			boris++
		}
		if name == "Даня" {
			viktor++
		}
	}

	total := float64(len(votes))
	fmt.Println("Анна:", ann, float64(ann)/total*100+'%')
	fmt.Println("Борис:", boris, float64(boris)/total*100+'%')
	fmt.Println("Виктор:", viktor, float64(viktor)/total*100+'%')
}

func main() {
	list := []string{"Даня", "Артур", "Даня", "Артем"}
	printVotes(list)
}
