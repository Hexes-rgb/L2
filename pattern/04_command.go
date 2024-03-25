package pattern

import "fmt"

/*
Паттерн "Комманда" применяется когда необходимо ставить операции в очередь, выполняя их по расписанию или по сети, когда необходимо поддерживать отмену
операция, ведение журналов.

Плюсы:
-Клиентский код не зависит от конкретных классов получателей операций, что упрощает добавление новых команд и получателей.
-Команды могут быть легко комбинированы и применены в различных сценариях, что повышает гибкость и переиспользуемость кода.

Минусы:
-Создание объектов команд может потреблять больше памяти, особенно если в системе много различных команд.
-Введение дополнительных классов команд и получателей может привести к усложнению структуры программы.

Примеры использования:
-Система управления очередями задач
-Управление устройствами
*/

type Command interface {
	Execute()
}

type Receiver struct{}

func (r *Receiver) Action() {
	fmt.Println("Receiver: выполнение действия")
}

type ConcreteCommand struct {
	receiver *Receiver
}

func NewConcreteCommand(receiver *Receiver) *ConcreteCommand {
	return &ConcreteCommand{receiver: receiver}
}

func (c *ConcreteCommand) Execute() {
	fmt.Println("ConcreteCommand: выполнение команды")
	c.receiver.Action()
}

type Invoker struct {
	command Command
}

func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

func (i *Invoker) ExecuteCommand() {
	fmt.Println("Invoker: выполнение команды")
	i.command.Execute()
}

func main() {
	receiver := &Receiver{}
	command := NewConcreteCommand(receiver)

	invoker := &Invoker{}
	invoker.SetCommand(command)
	invoker.ExecuteCommand()
}
