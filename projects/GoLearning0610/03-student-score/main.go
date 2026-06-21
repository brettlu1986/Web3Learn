package main

import "fmt"

type Student struct {
	Name  string
	Score int
}

// go run ./03-student-score
func main() {

	students := []Student{
		{Name: "Alice", Score: 90},
		{Name: "Bob", Score: 78},
		{Name: "Cindy", Score: 95},
		{Name: "David", Score: 66},
	}

	avg := AverageScore(students)
	max := MaxScore(students)
	min := MinScore(students)

	fmt.Println("average:", avg)
	fmt.Println("max:", max)
	fmt.Println("min:", min)

	scoreMap := BuildScoreMap(students)

	fmt.Println("Bob score:", scoreMap["Bob"])
}

func AverageScore(students []Student) float64 {
	if len(students) == 0 {
		return 0
	}
	total := 0
	for _, s := range students {
		total += s.Score
	}

	return float64(total) / float64(len(students))
}

func MaxScore(students []Student) int {
	if len(students) == 0 {
		return 0
	}

	max := students[0].Score

	for _, s := range students {
		if s.Score > max {
			max = s.Score
		}
	}

	return max
}

func MinScore(students []Student) int {
	if len(students) == 0 {
		return 0
	}

	min := students[0].Score

	for _, s := range students {
		if s.Score < min {
			min = s.Score
		}
	}

	return min
}

func BuildScoreMap(students []Student) map[string]int {
	result := make(map[string]int)

	for _, s := range students {
		result[s.Name] = s.Score
	}
	return result
}
