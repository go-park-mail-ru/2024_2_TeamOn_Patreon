/*
 * PushART - СКАТ | API
 *
 * API для управления CSAT
 *
 * API version: 1.0.0
 */

package models

import (
	sModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/csat/service/models"
)

type Question struct {
	// Вопрос
	Question string `json:"question"`
	// ИД вопроса
	QuestionID string `json:"questionID"`
}

func MapServQuestionToControlQuestion(question sModels.Question) Question {
	return Question{
		Question:   question.Question,
		QuestionID: question.QuestionID,
	}
}
