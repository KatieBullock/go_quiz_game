### Gophercise #1: Quiz Game

https://github.com/gophercises/quiz

The purpose of this program is to take a csv file in the format of "question,answer" and print each question to the user in the terminal. Users must input answers to each question within the time limit. If a user fails to answer all questions within the time limit, the unanswered questions will be marked as incorrect. Users receive feedback on the number of questions they answered correctly out of the total number of questions after they complete all problems or when the time runs out.

There are 3 flags that can be set: a csv file, a time limit (seconds), and a shuffle flag. They default to `problems.csv`, `30`, and `false` respectively.
