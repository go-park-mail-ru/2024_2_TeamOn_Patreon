package interfaces

// Behavior интерфейс структуры, методы которой реализуют бизнес-логику
type Behavior interface {
}

type NewBehavior func(repository Repository) Behavior
