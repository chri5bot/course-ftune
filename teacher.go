package ftune

import (
	"log"
	"math"
)

// Teacher model
type Teacher struct {
	name    string
	courses []*Course
}

// CreateQuiz creates a new quiz
func (t *Teacher) CreateQuiz(qs []*Question) *Quiz {
	return &Quiz{
		grade:     0,
		questions: qs,
	}
}

// AssignQuiz assigns quiz a course
func (t *Teacher) AssignQuiz(index int, quiz *Quiz) {
	if len(t.courses) <= index {
		log.Println("error: the course does not exist")
		return
	}

	for _, s := range t.courses[index].students {
		s.quizzes = append(s.quizzes, quiz)
	}
}

// CalculateTotal calculates total grade of students
func (t *Teacher) CalculateTotal() {
	for _, course := range t.courses {
		for _, student := range course.students {
			for _, quiz := range student.quizzes {
				student.totalGrade += quiz.grade
			}
			student.totalGrade = math.Ceil(student.totalGrade/float64(len(student.quizzes))*100) / 100
		}
	}
}
