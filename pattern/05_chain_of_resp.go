package pattent

import "fmt"

/*
Паттерн "Цепочка вызовов" используется когда система должна обрабатывать запросы различными способами без явного указания обработчика и точно знать,
что запрос будет обработан одним из них.

Плюсы:
-Можно добавлять или удалять обработчики в цепочке динамически, изменяя поведение системы без изменения самой цепочки.
-Клиентский код не зависит от конкретных обработчиков и знает только о первом обработчике, что позволяет легко изменять цепочку обработчиков.

Минусы:
-Если ни один из обработчиков не обработает запрос, это может привести к необработанному запросу, что требует дополнительных мер для
обработки таких ситуаций.
-Неправильная настройка цепочки обработчиков может привести к зацикливанию обработки запросов.

Примеры использования:
-Цепочка обработки HTTP-запросов в веб-фреймворках
-Логирование событий
*/
type Handler interface {
	SetNext(handler Handler)
	Handle(request string)
}

type ConcreteHandler struct {
	nextHandler Handler
	name        string
}

func NewConcreteHandler(name string) *ConcreteHandler {
	return &ConcreteHandler{name: name}
}

func (h *ConcreteHandler) SetNext(handler Handler) {
	h.nextHandler = handler
}

func (h *ConcreteHandler) Handle(request string) {
	fmt.Printf("%s received the request: %s\n", h.name, request)
	if h.nextHandler != nil {
		h.nextHandler.Handle(request)
	}
}

func main() {
	handler1 := NewConcreteHandler("Handler 1")
	handler2 := NewConcreteHandler("Handler 2")
	handler3 := NewConcreteHandler("Handler 3")

	handler1.SetNext(handler2)
	handler2.SetNext(handler3)

	handler1.Handle("Request")
}
