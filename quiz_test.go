package ftune

import "testing"

func Test(t *testing.T) {
	t.Run("Quizzes need to get graded.", func(t *testing.T) {
		testData := SetUpTestData()

		testData.solvedQuiz.GradeQuiz()

		want := testData.solvedQuiz.grade

		expected := 8.0

		if expected != want {
			t.Errorf("expected '%.2f' want '%.2f'", expected, want)
		}
	})
}
