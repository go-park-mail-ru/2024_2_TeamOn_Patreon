package validate

import "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/validate"

func Title(title string) (string, bool) {
	if len(title) < validate.MinLenSubTitle {
		return title, false
	}

	if len(title) > validate.MaxLenSubTitle {
		return title, false
	}
	title = validate.Sanitize(title)
	return title, true
}

func Description(description string) (string, bool) {
	if len(description) < validate.MinLenSubDescription {
		return description, false
	}
	if len(description) > validate.MaxLenSubDescription {
		return description, false
	}
	description = validate.Sanitize(description)
	return description, true
}

func Cost(cost int) bool {
	if cost < validate.MinSubCost {
		return false
	}

	if cost > validate.MaxSubCost {
		return false
	}
	return true
}

func Layer(layer int) bool {
	if layer < validate.MinSubLayer || layer > validate.MaxLayer {
		return false
	}
	return true
}
