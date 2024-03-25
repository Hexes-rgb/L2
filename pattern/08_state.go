package pattern

import "fmt"

/*
Паттерн "Состояние" применяется когда поведение объекта зависит от его состояния и может изменяться во времени выполнения

Плюсы:
-Позволяет избежать больших условных операторов, разбивая логику на отдельные классы состояний.
-Каждый класс состояния инкапсулирует поведение, связанное с конкретным состоянием, что делает код более чистым и поддерживаемым.
-Новые состояния могут быть легко добавлены, создавая новые классы состояний, без необходимости изменения кода контекста.

Минусы:
-Следить за переходами между состояниями может быть сложно, особенно если есть множество состояний и переходов между ними.
-Усложнение структуры программы, введение множества классов.

Примеры реализации:
-Состояние заказа на маркетплейсах
-Состояние персонажа в игре
-Система умного дома
*/

type Context struct {
	state State
}

func (c *Context) setState(state State) {
	c.state = state
}

func (c *Context) request() {
	c.state.handle()
}

type State interface {
	handle()
}

type ConcreteStateA struct{}

func (s *ConcreteStateA) handle() {
	fmt.Println("Handling request in ConcreteStateA")
}

type ConcreteStateB struct{}

func (s *ConcreteStateB) handle() {
	fmt.Println("Handling request in ConcreteStateB")
}

func main() {
	context := &Context{}

	context.setState(&ConcreteStateA{})

	context.request()

	context.setState(&ConcreteStateB{})

	context.request()
}
