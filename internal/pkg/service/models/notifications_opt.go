package models

import "strconv"

const notReadStatus = "NOTREAD"

type NotificationsOpt struct {
	Offset int
	Limit  int
	Status string
}

func (no *NotificationsOpt) Validate() {
	no.validateOffset()
	no.validateStatus()
}

func (no *NotificationsOpt) validateOffset() {
	if no.Limit == 0 {
		no.Limit = 10
	}
}

func (no *NotificationsOpt) validateStatus() {
	if no.Status != "" && no.Status != notReadStatus {
		no.Status = ""
	}
}

func NewNotificationsOpt(offsetStr, limitStr, status string) *NotificationsOpt {
	// Установка значений по умолчанию
	offset := 0
	limit := 10

	// Преобразование `offset` и `limit` в целые числа
	if offsetStr != "" {
		if o, err := strconv.Atoi(offsetStr); err == nil {
			offset = o
		}
	}
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil {
			limit = l
		}
	}

	no := &NotificationsOpt{Offset: offset, Limit: limit, Status: status}
	no.Validate()
	return no
}
