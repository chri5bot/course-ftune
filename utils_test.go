package ftune

type TestData struct {
	firstQuestions  []*Question
	secondQuestions []*Question
	solvedQuiz      *Quiz
	testCourse      *Course
	testQuiz        *Quiz
}

func SetUpTestData() *TestData {
	firstQuestions := []*Question{
		{
			statement:     "how much is 1+1?",
			isCorrect:     false,
			correctAnswer: "2",
			chosenAnswer:  "",
			options: []*Option{
				{
					value: "2",
				},
				{
					value: "3",
				},
				{
					value: "4",
				},
			},
		},
		{
			statement:     "how much is 2-1?",
			isCorrect:     false,
			correctAnswer: "1",
			chosenAnswer:  "",
			options: []*Option{
				{
					value: "1",
				},
				{
					value: "4",
				},
			},
		},
	}

	secondQuestions := []*Question{
		{
			statement:     "how much is 1+1?",
			isCorrect:     false,
			correctAnswer: "2",
			chosenAnswer:  "",
			options: []*Option{
				{
					value: "2",
				},
				{
					value: "3",
				},
				{
					value: "4",
				},
			},
		},
		{
			statement:     "how much is 2-1?",
			isCorrect:     false,
			correctAnswer: "1",
			chosenAnswer:  "",
			options: []*Option{
				{
					value: "1",
				},
				{
					value: "4",
				},
			},
		},
		{
			statement:     "how much is 10-1?",
			isCorrect:     false,
			correctAnswer: "9",
			chosenAnswer:  "",
			options: []*Option{
				{
					value: "1",
				},
				{
					value: "4",
				},
				{
					value: "9",
				},
			},
		},
	}

	solvedQuiz := &Quiz{
		grade: 0,
		questions: []*Question{
			{
				statement:     "how much is 1*1?",
				isCorrect:     false,
				correctAnswer: "1",
				chosenAnswer:  "1",
				options: []*Option{
					{
						value: "1",
					},
					{
						value: "4",
					},
				},
			},
			{
				statement:     "how much is 2*1?",
				isCorrect:     false,
				correctAnswer: "2",
				chosenAnswer:  "0",
				options: []*Option{
					{
						value: "2",
					},
					{
						value: "4",
					},
					{
						value: "3",
					},
				},
			},
			{
				statement:     "how much is 3*3?",
				isCorrect:     false,
				correctAnswer: "9",
				chosenAnswer:  "9",
				options: []*Option{
					{
						value: "2",
					},
					{
						value: "4",
					},
					{
						value: "9",
					},
				},
			},
			{
				statement:     "how much is 4*3?",
				isCorrect:     false,
				correctAnswer: "12",
				chosenAnswer:  "12",
				options: []*Option{
					{
						value: "12",
					},
					{
						value: "4",
					},
					{
						value: "9",
					},
				},
			},
			{
				statement:     "how much is 5*3?",
				isCorrect:     false,
				correctAnswer: "15",
				chosenAnswer:  "15",
				options: []*Option{
					{
						value: "12",
					},
					{
						value: "15",
					},
					{
						value: "9",
					},
				},
			},
		},
	}

	testCourse := &Course{
		subject: "Mathematics",
		students: []*Student{
			{
				name: "Christian Torres",
				quizzes: []*Quiz{
					{
						grade: 8.0,
					},
					{
						grade: 4.0,
					},
					{
						grade: 9.0,
					},
					{
						grade: 10.0,
					},
					{
						grade: 5.0,
					},
					{
						grade: 7.0,
					},
				},
			},
			{
				name: "Raul Hinojosa",
				quizzes: []*Quiz{
					{
						grade: 3.0,
					},
					{
						grade: 10.0,
					},
					{
						grade: 10.0,
					},
					{
						grade: 10.0,
					},
					{
						grade: 9.0,
					},
					{
						grade: 2.0,
					},
				},
			},
			{
				name: "Gabriel Peña",
				quizzes: []*Quiz{
					{
						grade: 7.0,
					},
					{
						grade: 2.0,
					},
					{
						grade: 10.0,
					},
					{
						grade: 10.0,
					},
					{
						grade: 4.0,
					},
					{
						grade: 8.0,
					},
				},
			},
		},
	}

	testQuiz := &Quiz{
		grade:     0,
		questions: firstQuestions,
	}
	return &TestData{
		firstQuestions:  firstQuestions,
		secondQuestions: secondQuestions,
		solvedQuiz:      solvedQuiz,
		testCourse:      testCourse,
		testQuiz:        testQuiz,
	}
}
