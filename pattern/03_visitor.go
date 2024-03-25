package pattern

import "fmt"

/*
Паттер "Посетитель" применяется когда необходимо выполнить операцию над набором объектов разных классов с минимальными изменениями в самих классах и
когда они не должны знать о том, какая именно операция будет над ними выполняться.

Плюсы:
-Посетитель позволяет вынести операции над объектами в отдельные классы, что облегчает добавление новых операций без изменения классов объектов.
-Новые операции могут быть добавлены, создавая новые классы посетителей, не изменяя классы объектов.

Минусы:
-Посетитель нарушает инкапсуляцию объектов, так как требует открытия интерфейса объектов для доступа из посетителя.
-Паттерн может усложнить структуру программы, так как операции над объектами вынесены в отдельные классы.

Примеры использования:
-Валидация и анализ данных
-Генерация отчетов
-
*/

type Element interface {
	Accept(visitor Visitor)
}

type ConcreteElementA struct{}

func (e *ConcreteElementA) Accept(visitor Visitor) {
	visitor.VisitConcreteElementA(e)
}

type ConcreteElementB struct{}

func (e *ConcreteElementB) Accept(visitor Visitor) {
	visitor.VisitConcreteElementB(e)
}

type Visitor interface {
	VisitConcreteElementA(element *ConcreteElementA)
	VisitConcreteElementB(element *ConcreteElementB)
}

type ConcreteVisitor struct{}

func (v *ConcreteVisitor) VisitConcreteElementA(element *ConcreteElementA) {
	fmt.Println("Visit ConcreteElementA")
}

func (v *ConcreteVisitor) VisitConcreteElementB(element *ConcreteElementB) {
	fmt.Println("Visit ConcreteElementB")
}

type ObjectStructure struct {
	elements []Element
}

func (o *ObjectStructure) Attach(element Element) {
	o.elements = append(o.elements, element)
}

func (o *ObjectStructure) Accept(visitor Visitor) {
	for _, element := range o.elements {
		element.Accept(visitor)
	}
}

func main() {
	objectStructure := ObjectStructure{}

	objectStructure.Attach(&ConcreteElementA{})
	objectStructure.Attach(&ConcreteElementB{})

	visitor := &ConcreteVisitor{}
	objectStructure.Accept(visitor)
}
