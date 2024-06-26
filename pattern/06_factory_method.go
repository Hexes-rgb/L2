package pattern

import "fmt"

/*
Паттерн Фабричный метод применяется когда создание объектов в суперклассе предоставляет подклассам возможность изменять создаваемые объекты и
когда заранее неизвестно какой именно класс объекта нужно создавать.

Плюсы:
-Паттерн позволяет переместить код создания объекта из клиентского кода в отдельный метод или класс, что улучшает
читаемость и обеспечивает лучшую организацию кода.
-Фабричный метод делегирует создание объектов подклассам, позволяя легко добавлять новые типы объектов,
расширяя функциональность системы без изменения существующего кода.

Минусы:
-Паттерн может привести к увеличению количества классов в системе, что может усложнить её структуру.
-Для каждого нового типа объекта требуется создание нового подкласса, что может быть неудобным в некоторых ситуациях.

Примеры использования:
-Фреймворки для работы с базами данных
-Фреймворки для разработки игр
-Библиотеки для работы с элементами графического интерфейса
*/

type Product interface {
	Use() string
}

type ConcreteProduct struct{}

func (cp *ConcreteProduct) Use() string {
	return "Using ConcreteProduct"
}

type Creator interface {
	CreateProduct() Product
}

type ConcreteCreator struct{}

func (cc *ConcreteCreator) CreateProduct() Product {
	return &ConcreteProduct{}
}

func main() {
	creator := &ConcreteCreator{}
	product := creator.CreateProduct()
	fmt.Println(product.Use())
}
