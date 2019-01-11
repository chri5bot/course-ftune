package ftune

// Quiz model
type Quiz struct {
	questions []*Question
	grade     float64
}

// Question model
type Question struct {
	statement     string
	options       []*Option
	correctAnswer string
	chosenAnswer  string
	isCorrect     bool
}

// Option model
type Option struct {
	value string
}

// GradeQuiz grades a quiz
func (q *Quiz) GradeQuiz() {
	for _, qn := range q.questions {
		if qn.chosenAnswer == qn.correctAnswer {
			q.grade += float64(10 / len(q.questions))
		}
	}
}
