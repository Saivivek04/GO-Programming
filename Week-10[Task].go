package main

import (
	"fmt"
	"strings"
)

func main() {
	myarray := []string{"CMRIT", "Cmrit", "CmRit", "MVSR", "CMRit", "IARE", "CBIT"}
	fmt.Println("The given array of strings is:", myarray)
	uniqueArray := removeDuplicate(myarray)
	fmt.Println("After Removing Duplicates:", uniqueArray)
}

func removeDuplicate(myarray []string) []string {
	visit := make(map[string]bool)
	result := []string{}
	for _, value := range myarray {
		lowerValue := strings.ToLower(value)
		if !visit[lowerValue] {
			visit[lowerValue] = true
			result = append(result, value)
		}
	}
	return result
}
