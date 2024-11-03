package models

import "strconv"

type FeedOpt struct {
	Offset int
	Limit  int
}

func (fo *FeedOpt) Validate() {
	fo.validateOffset()
}

func (fo *FeedOpt) validateOffset() {
	if fo.Limit == 0 {
		fo.Limit = 10
	}
}

func NewFeedOpt(offsetStr, limitStr string) *FeedOpt {
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
	fo := &FeedOpt{Offset: offset, Limit: limit}
	fo.Validate()
	return fo
}
