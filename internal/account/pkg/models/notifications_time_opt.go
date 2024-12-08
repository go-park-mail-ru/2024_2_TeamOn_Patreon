package models

import "strconv"

type NotificationsTimeOpt struct {
	Time int
}

func (no *NotificationsTimeOpt) Validate() {
	no.validateTime()
}

func (no *NotificationsTimeOpt) validateTime() {
	if no.Time <= 0 {
		no.Time = 10
	}
}

func NewNotificationsTimeOpt(timeStr string) *NotificationsTimeOpt {
	// Установка значений по умолчанию
	time := 10

	// Преобразование `time` в целое число
	if timeStr != "" {
		if t, err := strconv.Atoi(timeStr); err == nil {
			time = t
		}
	}

	no := &NotificationsTimeOpt{Time: time}
	no.Validate()
	return no
}
