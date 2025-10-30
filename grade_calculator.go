package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func averageGradeCalculator(grades map[string]int) float64 {
	sum := 0
	for _, grade := range grades {
		sum += grade
	}
	return float64(sum) / float64(len(grades))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter student name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	var subjectCount int
	for {
		fmt.Print("Enter number of subjects: ")
		countInput, _ := reader.ReadString('\n')
		countInput = strings.TrimSpace(countInput)
		count, err := strconv.Atoi(countInput)
		if err == nil && count > 0 {
			subjectCount = count
			break
		}
		fmt.Println("Please enter a valid positive number")
	}

	grades := make(map[string]int)
	for i := 0; i < subjectCount; i++ {
		fmt.Printf("Enter subject %d name: ", i+1)
		subject, _ := reader.ReadString('\n')
		subject = strings.TrimSpace(subject)

		var grade int
		for {
			fmt.Printf("Enter grade for %s: ", subject)
			gradeInput, _ := reader.ReadString('\n')
			gradeInput = strings.TrimSpace(gradeInput)
			g, err := strconv.Atoi(gradeInput)
			if err == nil && g >= 0 && g <= 100 {
				grade = g
				break
			}
			fmt.Println("Grade must be between 0 and 100")
		}
		grades[subject] = grade
	}

	fmt.Println("\nStudent:", name)
	for subject, grade := range grades {
		fmt.Println(subject+":", grade)
	}
	fmt.Printf("Average grade: %.2f\n", averageGradeCalculator(grades))
}
