# Компонент 

## Общая структура компонентов

Компоненты имеют следующую общую структуру:

```
{
    id: number,
    type: string,
    path: string,
    data: {
        <field1>: any,
        ...
    },
    outputs: {
        <field1>: number,
        <field2>: number,
        ...
    }
}
```

- Поле id идентифицирует компонент.
- Поле type содержит тип компонента в виде строки. Список компонентов представлен [здесь](./).
- Поле path предназначен для хранения пути данных для [контекста](../context.md). Он будет использоваться в зависимости от компонента: по нему будут или записываться данные, или считываться для последующего использования.
- Поле data содержит данные, которые будут использоваться компонентом. Данные индивидуальны для каждого компонента.
- Поле outputs содержит выходы компонента - id следующего возможного компонента. Часть полей общие, а остальные индивидуальны.
Общие поля:
    + nextComponentId - id следующего компонента. Для некоторых комопнентов задаётся только во время выполнения.
    + idIfError - id следующего компонента, если произошла ошибка.

Для использования общих полей каждый компонент должен реализовать следующий интерфейс: 

```golang
type Outputs interface {
	GetNextComponentId() *int64
	GetIdIfError() *int64
}
```

Для использования данных и выходов компонент должен реализовать следующий интерфейс:

```golang
type Component interface {
	GetPath() string
	GetOutputs() Outputs
}
```

Данные интерфейсы будут использоваться в выполнении компонентов.

## Выполнение компонентов

Каждый компонент имеет свою логику выполнения. Выполнением компонентов занимается объект Executor.

### Виды компонентов

- Компонент действия
- Компонент управления
- Компонент ввода
- Компонент вывода

#### Компонент действия. Action component.

Компонент действия служит для преобразования данных.

Реализует следующий интерфейс:

```golang
type ActionComponent interface {
	Component

	Execute(ctx *Context) (*any, error)
}
```

Метод Execute принимает [контекст](../context.md), которой используется для получения новых данных.

#### Компонент управления. Control component.

Компонент управления используется для изменения потока выполнения.

Реализует следующий интерфейс:

```golang 
type ControlComponent interface {
	Component

	Execute(ctx *Context) error
}
```

В методе Execute происходит изменение поля nextComponentId с использованием [контекста](../context.md), данных и выходов компонента.

#### Компонент ввода

Компонент ввода используется для получения данных от пользователя.

Реализует следующий интерфейс: 

```golang
type InputComponent interface {
	Component

	Execute(ctx *Context, io io.IO) (*any, error)
}
```

Метод Execute получает данные от пользователя путем вызовов методов интерфейса ввода-вывода.


#### Компонент вывода

Компонент вывода используется для отправки данных пользователю.

Реализует следующий интерфейс: 

```golang
type OutputComponent interface {
	Component

	Execute(ctx *Context, io io.IO) error
}
```

Метод Execute выводит данные путем вызовов методов интерфейса ввода-вывода.
