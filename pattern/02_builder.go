package pattern

import "fmt"

/*
Паттерн "Строитель" применяется когда процесс создания сложного объекта должен быть разделен на отдельные этапы или когда есть
необходмость создавать различные вариации объектов, не меняя код построения.

Плюсы:
-Разделение ответственности, разбиение построения на отдельные этапы
-Повторное использование кода построения

Минусы:
-Недостаточная гибкость. Если требуется добавить новые этапы построения или изменить порядок этапов,
это может потребовать изменения интерфейса Builder и всех его реализаций.
-Усложнение кода. Внедрение паттерна "Строитель" может привести к увеличению количества классов и интерфейсов, что может усложнить структуру программы.

Примеры использования:
-Построитель SQL запросов
-Построение сложных элементов графического интерфейса
*/

type Car struct {
	Model  string
	Color  string
	Option string
}

type CarBuilder interface {
	SetModel(model string)
	SetColor(color string)
	SetOption(option string)
	Build() *Car
}

type ConcreteCarBuilder struct {
	model  string
	color  string
	option string
}

func NewConcreteCarBuilder() *ConcreteCarBuilder {
	return &ConcreteCarBuilder{}
}

func (b *ConcreteCarBuilder) SetModel(model string) {
	b.model = model
}

func (b *ConcreteCarBuilder) SetColor(color string) {
	b.color = color
}

func (b *ConcreteCarBuilder) SetOption(option string) {
	b.option = option
}

func (b *ConcreteCarBuilder) Build() *Car {
	return &Car{
		Model:  b.model,
		Color:  b.color,
		Option: b.option,
	}
}

type Director struct {
	builder CarBuilder
}

func NewDirector(builder CarBuilder) *Director {
	return &Director{builder: builder}
}

func (d *Director) Construct(model, color, option string) *Car {
	d.builder.SetModel(model)
	d.builder.SetColor(color)
	d.builder.SetOption(option)
	return d.builder.Build()
}

func main() {
	builder := NewConcreteCarBuilder()
	director := NewDirector(builder)

	car := director.Construct("SUV", "black", "4WD")
	fmt.Printf("Car built: %+v\n", car)
}
