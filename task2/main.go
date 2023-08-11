package main

import (
	"fmt"
	"math"
)

type Student struct {
	Name  string
	Score int
}

type Students []Student

func (s *Students) AddStudent(name string, score int) {
	student := Student{Name: name, Score: score}
	*s = append(*s, student)
}

func (s Students) AverageScore() float64 {
	totalScore := 0
	for _, student := range s {
		totalScore += student.Score
	}
	return float64(totalScore) / float64(len(s))
}

func (s Students) MinMaxScore() (min Student, max Student) {
	minScore := math.MaxInt32
	maxScore := math.MinInt32

	for _, student := range s {
		if student.Score < minScore {
			minScore = student.Score
			min = student
		}
		if student.Score > maxScore {
			maxScore = student.Score
			max = student
		}
	}

	return min, max
}

func main() {
	var students Students

	for i := 0; i < 5; i++ {
		var name string
		var score int

		fmt.Printf("Input %d Student's name: ", i+1)
		fmt.Scan(&name)

		fmt.Printf("Input %d Student's score: ", i+1)
		fmt.Scan(&score)

		students.AddStudent(name, score)
	}

	averageScore := students.AverageScore()
	minScoreStudent, maxScoreStudent := students.MinMaxScore()

	fmt.Printf("Average Score: %.2f\n", averageScore)
	fmt.Printf("Min score of student: %s (%d)\n", minScoreStudent.Name, minScoreStudent.Score)
	fmt.Printf("Max score of student: %s (%d)\n", maxScoreStudent.Name, maxScoreStudent.Score)
}
