/*
 * PushART - СКАТ | API
 *
 * API для управления CSAT
 *
 * API version: 1.0.0
 */

package models

type Question struct {
	// Вопрос
	Question string `json:"question"`
	// ИД вопроса
	QuestionID string `json:"questionID"`
}
