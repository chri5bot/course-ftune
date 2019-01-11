package ftune

// Student model
type Student struct {
	name       string
	totalGrade float64
	quizzes    []*Quiz
}

// AnswerQuestion answers question
func (s *Student) AnswerQuestion(quiz int, question int, option int) {

	if len(s.quizzes) < quiz {
		// log.Println("error: the quiz does not exist")
		return
	}

	if len(s.quizzes[quiz].questions) < question {
		// log.Println("error: the question does not exist")
		return
	}

	if len(s.quizzes[quiz].questions[question].options) < option {
		// log.Println("error: the option does not exist")
		return
	}

	s.quizzes[quiz].questions[question].chosenAnswer = s.quizzes[quiz].questions[question].options[option].value
}
