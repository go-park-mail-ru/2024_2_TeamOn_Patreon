package models

const (
	Published  string = "PUBLISHED"  //  статус поста, который только что опубликован, проверку не прошел
	Complained string = "COMPLAINED" //  статус поста, на который пожаловались
	Allowed    string = "ALLOWED"    // статус поста, который успешно прошел проверку
	Blocked    string = "BLOCKED"    //  статус поста, который не прошел проверку
)

// Источник -> [тут](docs/api/moderation/Описание функционала для модератора.md)

// CheckStatus - проверяет корректно ли задан статус
func CheckStatus(status string) bool {
	switch status {
	case Published:
		return true
	case Complained:
		return true
	case Allowed:
		return true
	case Blocked:
		return true
	default:
		return false
	}
}
