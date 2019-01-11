package ftune

import (
	"testing"
)

func TestAnswerQuiz(t *testing.T) {
	t.Run("Students solve/answer questions to complete the quiz, but they don't have to complete it at once. (Partial submissions can be made).", func(t *testing.T) {

		ss := []*Student{
			{
				name: "Raul Hinojosa",
			},
			{
				name: "Christian Torres",
			},
		}

		// assigns quiz for each student
		for _, s := range ss {
			testData := SetUpTestData()
			s.quizzes = append(s.quizzes, testData.testQuiz)
		}

		ss[0].AnswerQuestion(10, 10, 1)
		ss[0].AnswerQuestion(0, 10, 1)
		ss[0].AnswerQuestion(0, 0, 1)
		ss[0].AnswerQuestion(0, 1, 1)

		ss[1].AnswerQuestion(0, 0, 0)
		ss[1].AnswerQuestion(0, 0, 10)
		ss[1].AnswerQuestion(10, 10, 10)

		want := 0

		for _, s := range ss {
			for _, quiz := range s.quizzes {
				for _, q := range quiz.questions {
					// log.Println("Statement: ", q.statement)
					// log.Println("choosen answer", q.chosenAnswer)
					if q.chosenAnswer != "" {
						want++
					}
					// for _, o := range q.options {
					// log.Println("Option: ", o)
					// }
				}
			}

		}

		expected := 3

		if expected != want {
			t.Errorf("expected '%d' want '%d'", expected, want)
		}
	})
}
