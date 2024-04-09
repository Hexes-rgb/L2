Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:

```
<nil>
false
```

Несмотря на то, что базовая структура является равна `nil`, сам интерфейс не является `nil`. Интерфейс равен `nil`, только если и тип, и значение равны `nil`.
