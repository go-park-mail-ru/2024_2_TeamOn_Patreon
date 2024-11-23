package models

import (
	repModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/csat/repository/models"
)

// Service модель Question
type Question struct {
	// Вопрос
	Question string
	// ИД вопроса
	QuestionID string
}

func MapRepQuestionToServQuestion(question repModels.Question) Question {
	return Question{
		Question:   question.Question,
		QuestionID: question.QuestionID,
	}
}
