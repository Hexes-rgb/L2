package pattern

import "fmt"

/*
Паттерн "Стратегия" применяется когда есть несколько схожих классов, которые отличаются только поведением. Также используется когда есть класс с
множеством условных операторов, определяющих его поведение.

Плюсы:
-Позволяет легко добавлять новые стратегии или модифицировать существующие без изменения контекста.
-Каждая стратегия отвечает за свой алгоритм, что позволяет изолировать различные поведения от клиентов.
-Каждая стратегия отвечает за свой алгоритм, что позволяет изолировать различные поведения от клиентов.

Минусы:
-Паттерн может привести к увеличению числа классов в системе, что может усложнить ее структуру.
-Клиенты должны явно выбирать подходящую стратегию, что может увеличить сложность конфигурации.

Примеры использования:
-Сортировка
-Сжатие данных
-Различные алгоритмы поиска
*/

type Strategy interface {
	ExecuteOperation(int, int) int
}

type ConcreteStrategyAdd struct{}

func (s *ConcreteStrategyAdd) ExecuteOperation(num1, num2 int) int {
	return num1 + num2
}

type ConcreteStrategySubtract struct{}

func (s *ConcreteStrategySubtract) ExecuteOperation(num1, num2 int) int {
	return num1 - num2
}

type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy(num1, num2 int) int {
	if c.strategy == nil {
		return 0
	}
	return c.strategy.ExecuteOperation(num1, num2)
}

func main() {
	context := &Context{}

	context.SetStrategy(&ConcreteStrategyAdd{})
	result := context.ExecuteStrategy(10, 5)
	fmt.Println("10 + 5 =", result)

	context.SetStrategy(&ConcreteStrategySubtract{})
	result = context.ExecuteStrategy(10, 5)
	fmt.Println("10 - 5 =", result)
}
