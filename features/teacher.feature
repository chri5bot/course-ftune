Feature: The Teacher assigns a Quiz

    Scenario: the Teacher creates a Quiz
        When the Teacher groups questions and options, with their respective answers
        Then the Teacher can create the Quiz

    Scenario: the Teacher assigns a Quiz
        When the Teacher has the Quiz ready
        Then the Teacher can assign the Quiz


Feature: The Teacher calculates student's total grade

    Scenario: the Teacher calculates student's total grade
        When the semester is finished and the Students has been finished their quizzes
        Then the Teacher calculates total grade