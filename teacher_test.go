package ftune

import (
	"reflect"
	"testing"
)

func TestCreateQuiz(t *testing.T) {
	t.Run("Teachers can create multiple quizzes with many questions (each question is multiple choice)", func(t *testing.T) {

		te := &Teacher{
			name: "Kwan Lee",
		}

		testData := SetUpTestData()

		quizzes := make([]*Quiz, 0)

		firstQuiz := te.CreateQuiz(testData.firstQuestions)
		quizzes = append(quizzes, firstQuiz)

		secondQuiz := te.CreateQuiz(testData.secondQuestions)
		quizzes = append(quizzes, secondQuiz)

		want := len(quizzes)

		expected := 2

		if expected != want {
			t.Errorf("expected '%d' want '%d'", expected, want)
		}
	})

}

func TestAssingQuiz(t *testing.T) {
	t.Run("Teachers can assign quizzes to students", func(t *testing.T) {
		te := &Teacher{
			name: "Kwan Lee",
			courses: []*Course{
				{
					subject: "Mathematics",
					students: []*Student{
						{
							name: "Christian Torres",
						},
						{
							name: "Raul Hinojosa",
						},
					},
				},
			},
		}

		testData := SetUpTestData()

		te.AssignQuiz(0, testData.testQuiz)

		want := 0
		for _, cs := range te.courses {
			for _, st := range cs.students {
				if len(st.quizzes) > 0 {
					want++
				}
			}
		}

		expected := 2

		if expected != want {
			t.Errorf("expected '%d' want '%d'", expected, want)
		}

	})

}
func TestCalculateTotal(t *testing.T) {
	t.Run("For each teacher, they can calculate each student's total grade accumulated over a semester for their classes", func(t *testing.T) {

		testData := SetUpTestData()

		te := &Teacher{
			name: "Kwan Lee",
			courses: []*Course{
				testData.testCourse,
			},
		}

		te.CalculateTotal()

		var want []float64

		for _, course := range te.courses {
			for _, student := range course.students {
				want = append(want, student.totalGrade)
			}
		}

		expected := []float64{7.17, 7.33, 6.83}

		if reflect.DeepEqual(expected, want) {
			t.Errorf("expected '%.2f' want '%.2f'", expected, want)
		}
	})
}
