package pattern

import "fmt"

/*
Паттерн "Фасад" применяется когда нужно разделить подсистемы на уровни доступа, упростить взаимодействие между ними.
Предоставление простого интерфейса к системе, который скрывает сложности реализации, так же является применением данного паттерна.

Плюсы:
-Высокая независимость клиента от деталей и особенностей реализации системы
-Упрощение использования системы, изоляция сложности

Минусы:
-Увеличение сложности фасада при добавлении новой функциональности
-Меньшая гибкость в настройке поведения системы

Примеры использования:
-Библиотека для работы с базой данных
-Фреймворк для разработки веб-приложений
-Система управления файлами
*/

type SubsystemA struct{}

func (s *SubsystemA) OperationA() {
	fmt.Println("Subsystem A: Operation...")
}

type SubsystemB struct{}

func (s *SubsystemB) OperationB() {
	fmt.Println("Subsystem B: Operation...")
}

type SubsystemC struct{}

func (s *SubsystemC) OperationC() {
	fmt.Println("Subsystem C: Operation...")
}

type SystemFacade struct {
	subsystemA *SubsystemA
	subsystemB *SubsystemB
	subsystemC *SubsystemC
}

func NewSystemFacade() *SystemFacade {
	return &SystemFacade{
		subsystemA: &SubsystemA{},
		subsystemB: &SubsystemB{},
		subsystemC: &SubsystemC{},
	}
}

func (f *SystemFacade) Operation() {
	f.subsystemA.OperationA()
	f.subsystemB.OperationB()
	f.subsystemC.OperationC()
}

func main() {
	facade := NewSystemFacade()
	facade.Operation()
}
